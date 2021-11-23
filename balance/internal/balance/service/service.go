package service

import (
	"context"

	"github.com/opentracing/opentracing-go" //todo -> opentelemetry

	"github.com/olteffe/balancex/balance/internal/balance"
	"github.com/olteffe/balancex/balance/internal/balance/models"
	"github.com/olteffe/balancex/balance/pkg/grpc_errors"
	"github.com/olteffe/balancex/balance/pkg/logger"
)

// BalanceService Balance service
type balanceService struct {
	balanceRepo balance.PGRepository
	redisRepo   balance.RedisRepository
	logger      logger.Logger
}

// NewBalanceService Balance service constructor
func NewBalanceService(balanceRepo balance.PGRepository, redisRepo balance.RedisRepository,
	logger logger.Logger) *balanceService {
	return &balanceService{balanceRepo: balanceRepo, redisRepo: redisRepo, logger: logger}
}

// CreateBalance Create user balance
func (b *balanceService) CreateBalance(ctx context.Context, balance *models.Balance) (*models.Balance, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "balanceService.CreateBalance")
	defer span.Finish()

	existsUser, err := b.balanceRepo.FindUserID(ctx, balance.UserID)
	if existsUser != nil || err == nil {
		return nil, grpc_errors.ErrUserExists
	}

	convert, err := b.redisRepo.ConvertBalance(ctx, balance)
	if err != nil {
		return nil, grpc_errors.ErrConvertCurrency
	}
	return b.balanceRepo.CreateBalance(ctx, convert)
}

// GetBalance get user balance
func (b *balanceService) GetBalance(ctx context.Context, balance *models.Balance) (*models.Balance, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "balanceService.GetBalance")
	defer span.Finish()

	convert, err := b.redisRepo.ConvertBalance(ctx, balance)
	if err != nil {
		return nil, grpc_errors.ErrConvertCurrency
	}
	return b.balanceRepo.GetBalance(ctx, convert)
}
