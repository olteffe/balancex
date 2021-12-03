package balance

import (
	"context"
	"github.com/google/uuid"
	"github.com/olteffe/balancex/balance/internal/balance/models"
)

// PGRepository interface
type PGRepository interface {
	CreateBalance(ctx context.Context, balance *models.Balance) (string, error)
	GetBalance(ctx context.Context, balance *models.Balance) (*models.Balance, error)
	FindUserID(ctx context.Context, userID uuid.UUID) (*models.Balance, error)
}

// RedisRepository Redis balance repository interface
type RedisRepository interface {
	GetBalance(ctx context.Context, balance *models.Balance) (*models.Balance, error)
	ConvertBalance(ctx context.Context, balance *models.Balance) (*models.Balance, error)
}
