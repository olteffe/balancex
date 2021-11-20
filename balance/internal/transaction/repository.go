package transaction

import (
	"context"

	"github.com/olteffe/balancex/balance/internal/transaction/models"
)

// PGRepository interface
type PGRepository interface {
	CreateTransaction(ctx context.Context, transaction *models.Transaction) (*models.Transaction, error)
	GetTransactions(ctx context.Context, balance *models.Transaction) (*models.TransactionList, error)
}
