package repository

import (
	"context"
	"strconv"

	"github.com/go-redis/redis/v8"
	"github.com/opentracing/opentracing-go"

	"github.com/olteffe/balancex/balance/pkg/grpc_errors"
	"github.com/olteffe/balancex/balance/pkg/logger"
)

// transactionRedisRepo redis repository
type transactionRedisRepo struct {
	redisClient *redis.Client
	logger      logger.Logger
}

// NewTransactionRedisRepo redis repository constructor
func NewTransactionRedisRepo(redisClient *redis.Client, logger logger.Logger) *transactionRedisRepo {
	return &transactionRedisRepo{redisClient: redisClient, logger: logger}
}

// GetRate Get exchange rate
func (t *transactionRedisRepo) GetRate(ctx context.Context, rate string) (float32, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "RedisRepository.GetRate")
	defer span.Finish()

	coefficient, err := t.redisClient.Get(ctx, rate).Result()
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
