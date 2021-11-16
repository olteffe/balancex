package server

import (
	"context"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-redis/redis/v8"
	grpcRecovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpcCtxTags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpcPrometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/opentracing/opentracing-go"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"

	"github.com/olteffe/balancex/balance/config"
	balanceGrpc "github.com/olteffe/balancex/balance/internal/balance/delivery/grpc"
	balService "github.com/olteffe/balancex/balance/internal/balance/proto"
	"github.com/olteffe/balancex/balance/internal/balance/repository"
	"github.com/olteffe/balancex/balance/internal/balance/service"
	"github.com/olteffe/balancex/balance/internal/interceptors"
	"github.com/olteffe/balancex/balance/pkg/logger"
	"github.com/olteffe/balancex/balance/pkg/metrics"
)

// Server struct
type Server struct {
	logger  logger.Logger
	cfg     *config.Config
	pgxPool *pgxpool.Pool
	redis   *redis.Client
	tracer  opentracing.Tracer
}

// NewServer Server constructor
func NewServer(logger logger.Logger, cfg *config.Config, pgxPool *pgxpool.Pool, redis *redis.Client,
	tracer opentracing.Tracer) *Server {
	return &Server{cfg: cfg, logger: logger, pgxPool: pgxPool, redis: redis, tracer: tracer}
}

// Run server
func (s *Server) Run() error {
	metric, err := metrics.CreateMetrics(s.cfg.Metrics.URL, s.cfg.Metrics.ServiceName)
	if err != nil {
		s.logger.Errorf("CreateMetrics Error: %s", err)
	}
	s.logger.Info(
		"Metrics available URL: %s, ServiceName: %s",
		s.cfg.Metrics.URL,
		s.cfg.Metrics.ServiceName,
	)

	im := interceptors.NewInterceptorManager(s.logger, s.cfg, metric)
	balanceRepository := repository.NewBalanceRepository(s.pgxPool)
	balanceService := service.NewBalanceService(balanceRepository, s.logger, s.cfg)

	ctx, cancel := context.WithCancel(context.Background())
	router := echo.New()
	router.GET("/metrics", echo.WrapHandler(promhttp.Handler()))

	go func() {
		if err := router.Start(s.cfg.Metrics.URL); err != nil {
			s.logger.Errorf("router.Start metrics: %v", err)
			cancel()
		}
	}()

	l, err := net.Listen("tcp", s.cfg.Server.Port)
	if err != nil {
		return err
	}
	defer l.Close()

	server := grpc.NewServer(grpc.KeepaliveParams(keepalive.ServerParameters{
		MaxConnectionIdle: s.cfg.Server.MaxConnectionIdle * time.Minute,
		Timeout:           s.cfg.Server.Timeout * time.Second,
		MaxConnectionAge:  s.cfg.Server.MaxConnectionAge * time.Minute,
		Time:              s.cfg.Server.Timeout * time.Minute,
	}),
		grpc.UnaryInterceptor(im.Logger),
		grpc.ChainUnaryInterceptor(
			grpcCtxTags.UnaryServerInterceptor(),
			grpcPrometheus.UnaryServerInterceptor,
			grpcRecovery.UnaryServerInterceptor(),
		),
	)

	balanceGrpcMicroservice := balanceGrpc.NewBalanceMicroservice(balanceService, s.logger, s.cfg)
	balService.RegisterBalanceServiceServer(server, balanceGrpcMicroservice)
	grpcPrometheus.Register(server)
	s.logger.Info("Emails Service initialized")

	if s.cfg.Server.Mode != "Production" {
		reflection.Register(server)
	}

	go func() {
		s.logger.Infof("Server is listening on port: %v", s.cfg.Server.Port)
		s.logger.Fatal(server.Serve(l))
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	select {
	case v := <-quit:
		s.logger.Errorf("signal.Notify: %v", v)
	case done := <-ctx.Done():
		s.logger.Errorf("ctx.Done: %v", done)
	}

	if err := router.Shutdown(ctx); err != nil {
		s.logger.Errorf("Metrics router.Shutdown: %v", err)
	}
	server.GracefulStop()
	s.logger.Info("Server Exited Properly")

	return nil
}
