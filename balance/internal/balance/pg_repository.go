//go:generate mockgen -source pg_repository.go -destination mock/pg_repository.go -package mock
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
	FindUserID(ctx context.Context, userID uuid.UUID) error
}
