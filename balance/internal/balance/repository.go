package balance

import (
	"context"
	"github.com/olteffe/balancex/balance/internal/balance/models"
)

// PGRepository interface
type PGRepository interface {
	CreateBalance(ctx context.Context, balance *models.Balance) (*models.Balance, error)
	GetBalance(ctx context.Context, balance *models.Balance) (*models.Balance, error)
}

// RedisRepository Redis balance repository interface
type RedisRepository interface {
	GetBalance(ctx context.Context, balance *models.Balance) (*models.Balance, error)
}
