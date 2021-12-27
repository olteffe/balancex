package service

import (
	"context"
	"fmt"

	"github.com/opentracing/opentracing-go" //todo -> opentelemetry

	"github.com/olteffe/balancex/balance/internal/balance"
	"github.com/olteffe/balancex/balance/internal/balance/models"
	"github.com/olteffe/balancex/balance/pkg/logger"
)

// BalanceService Balance service
type balanceService struct {
	balanceRepo balance.PGRepository
	redisRepo   balance.RedisRepository
	logger      logger.Logger
}

const (
	mainCurrency = "RUB"
)

// NewBalanceService Balance service constructor
func NewBalanceService(balanceRepo balance.PGRepository, redisRepo balance.RedisRepository,
	logger logger.Logger) *balanceService {
	return &balanceService{balanceRepo: balanceRepo, redisRepo: redisRepo, logger: logger}
}

// CreateBalance Create user balance
func (b *balanceService) CreateBalance(ctx context.Context, balance *models.Balance) (string, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "balanceService.CreateBalance")
	defer span.Finish()

	if err := b.balanceRepo.FindUserID(ctx, balance.UserID); err != nil {
		return "", fmt.Errorf("CreateBalance.FindUserID: %w", err)
	}

	convert, err := b.convertRate(ctx, balance)
	if err != nil {
		return "", fmt.Errorf("CreateBalance.convertRate: %w", err)
	}
	return b.balanceRepo.CreateBalance(ctx, convert)
}

// GetBalance get user balance
func (b *balanceService) GetBalance(ctx context.Context, balance *models.Balance) (*models.Balance, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "balanceService.GetBalance")
	defer span.Finish()

	convert, err := b.convertRate(ctx, balance)
	if err != nil {
		return nil, fmt.Errorf("GetBalance.convertRate: %w", err)
	}
	return b.balanceRepo.GetBalance(ctx, convert)
}

// convertRate converts currency
func (b *balanceService) convertRate(ctx context.Context, balance *models.Balance) (*models.Balance, error) {
	if balance.Currency == mainCurrency {
		return balance, nil
	}
	rate, err := b.redisRepo.GetRate(ctx, balance.Currency)
	if err != nil {
		return nil, fmt.Errorf("convertRate: %w", err)
	}
	balance.Amount = uint64(rate * float32(balance.Amount)) // problems with law
	return balance, nil
}
