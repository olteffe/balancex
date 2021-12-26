//go:generate mockgen -source service.go -destination mock/service.go -package mock
package balance

import (
	"context"
	"github.com/olteffe/balancex/balance/internal/balance/models"
)

// Service interface
type Service interface {
	CreateBalance(ctx context.Context, balance *models.Balance) (string, error)
	GetBalance(ctx context.Context, balance *models.Balance) (*models.Balance, error)
}
