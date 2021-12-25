package transaction

import (
	"context"

	tr "github.com/olteffe/balancex/balance/internal/transaction/models"
	"github.com/olteffe/balancex/balance/pkg/utils"
)

// Service interface
type Service interface {
	CreateTransaction(ctx context.Context, transaction *tr.Transaction) (string, error)
	GetTransactions(ctx context.Context, transaction *utils.TransactionsRequest) (*tr.TransactionList, error)
}
