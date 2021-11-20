package transaction

import (
	"context"

	bl "github.com/olteffe/balancex/balance/internal/balance/models"
	tr "github.com/olteffe/balancex/balance/internal/transaction/models"
)

// Service interface
type Service interface {
	CreateTransaction(ctx context.Context, transaction *tr.Transaction) (*bl.Balance, error)
	GetTransactions(ctx context.Context, balance *bl.Balance) (*tr.TransactionList, error)
}
