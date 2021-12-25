package transaction

import (
	"context"

	"github.com/olteffe/balancex/balance/internal/transaction/models"
	"github.com/olteffe/balancex/balance/pkg/utils"
)

// PGRepository interface
type PGRepository interface {
	CreateTransaction(ctx context.Context, transaction *models.Transaction) (string, error)
	GetTransactions(ctx context.Context, balance *models.Transaction) (*models.TransactionList, error)
	FindUsersID(ctx context.Context, transaction *models.Transaction) (bool, error)
	FindUserID(ctx context.Context, transaction *utils.TransactionsRequest) (bool, error)
}

// RedisRepository interface
type RedisRepository interface {
	ConvertTransaction(ctx context.Context, transaction *models.Transaction) (*models.Transaction, error)
	ConvertTransactions(ctx context.Context, transaction *utils.TransactionsRequest) (*models.Transaction, error)
}
