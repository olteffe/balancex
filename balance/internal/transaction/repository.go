package transaction

import (
	"context"

	"github.com/olteffe/balancex/balance/internal/transaction/models"
)

// PGRepository interface
type PGRepository interface {
	CreateTransaction(ctx context.Context, transaction *models.Transaction) (string, error)
	GetTransactions(ctx context.Context, balance *models.Transaction) (*models.TransactionList, error)
	FindUsersID(ctx context.Context, transaction *models.Transaction) (bool, error)
	FindUserID(ctx context.Context, transaction *models.TransactionsRequest) (bool, error)
}

// RedisRepository interface
type RedisRepository interface {
	ConvertTransaction(ctx context.Context, transaction *models.Transaction) (*models.Transaction, error)
	ConvertTransactions(ctx context.Context, transaction *models.TransactionsRequest) (*models.Transaction, error)
}
