package transaction

import (
	"context"

	tr "github.com/olteffe/balancex/balance/internal/transaction/models"
)

// Service interface
type Service interface {
	CreateTransaction(ctx context.Context, transaction *tr.Transaction) (string, error)
	GetTransactions(ctx context.Context, transaction *tr.TransactionsRequest) (*tr.TransactionList, error)
}
