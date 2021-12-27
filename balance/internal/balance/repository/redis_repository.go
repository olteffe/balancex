package repository

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/olteffe/balancex/balance/pkg/grpc_errors"
	"github.com/opentracing/opentracing-go"
	"strconv"

	"github.com/olteffe/balancex/balance/pkg/logger"
)

// balanceRedisRepo redis repository
type balanceRedisRepo struct {
	redisClient *redis.Client
	logger      logger.Logger
}

// NewBalanceRedisRepo redis repository constructor
func NewBalanceRedisRepo(redisClient *redis.Client, logger logger.Logger) *balanceRedisRepo {
	return &balanceRedisRepo{redisClient: redisClient, logger: logger}
}

// GetRate Get exchange rate
func (b *balanceRedisRepo) GetRate(ctx context.Context, rate string) (float32, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "RedisRepository.GetRate")
	defer span.Finish()

	coefficient, err := b.redisClient.Get(ctx, rate).Result()
	switch {
	case err == redis.Nil:
		return 0, grpc_errors.ErrConvertCurrency
	case err != nil:
		return 0, grpc_errors.ErrDB
	case coefficient == "":
		return 0, grpc_errors.ErrConvertCurrency
	}

	float, err := strconv.ParseFloat(coefficient, 32)
	if err != nil {
		return 0, grpc_errors.ErrConvertCurrency
	}
	return float32(float), nil
}
