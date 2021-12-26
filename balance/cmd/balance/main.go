package main

import (
	"io"
	"log"
	"os"

	"github.com/opentracing/opentracing-go"

	"github.com/olteffe/balancex/balance/config"
	"github.com/olteffe/balancex/balance/internal/server"
	"github.com/olteffe/balancex/balance/pkg/jaeger"
	"github.com/olteffe/balancex/balance/pkg/logger"
	"github.com/olteffe/balancex/balance/pkg/postgres"
	"github.com/olteffe/balancex/balance/pkg/redis"
)

func main() {
	log.Println("Starting server")

	// config
	configPath := config.GetConfigPath(os.Getenv("config"))
	cfg, err := config.GetConfig(configPath)
	if err != nil {
		log.Fatalf("Loading config: %v", err)
	}

	// logger
	appLogger := logger.NewApiLogger(cfg)
	appLogger.InitLogger()
	appLogger.Infof(
		"AppVersion: %s, LogLevel: %s, Mode: %s, SSL: %v",
		cfg.Server.AppVersion,
		cfg.Logger.Level,
		cfg.Server.Mode,
		cfg.Server.SSL,
	)
	appLogger.Infof("Success parsed config: %#v", cfg.Server.AppVersion)

	// jaeger
	tracer, closer, err := jaeger.InitJaeger(cfg)
	if err != nil {
		appLogger.Fatal("cannot create tracer", err)
	}
	appLogger.Info("Jaeger connected")

	// opentracing
	opentracing.SetGlobalTracer(tracer)
	defer func(closer io.Closer) {
		err := closer.Close()
		if err != nil {
			appLogger.Warn("Jaeger closer")
		}
	}(closer)
	appLogger.Info("Opentracing connected")

	// postgres
	pgxPool, err := postgres.NewPgxConn(cfg)
	if err != nil {
		appLogger.Fatalf("NewPgxConn: %+v", err)
	}
	appLogger.Infof("PostgreSQL connected: %+v", pgxPool.Stat().TotalConns())

	// redis
	redisClient := redis.NewRedisClient(cfg)
	appLogger.Info("Redis connected")

	s := server.NewServer(appLogger, cfg, pgxPool, redisClient, tracer)
	appLogger.Fatal(s.Run())
}
