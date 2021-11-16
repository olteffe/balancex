package service

import (
	"github.com/olteffe/balancex/balance/config"
	"github.com/olteffe/balancex/balance/internal/balance"
	"github.com/olteffe/balancex/balance/pkg/logger"
)

// BalanceService Balance service
type BalanceService struct {
	balanceRepo balance.BalanceRepository
	logger      logger.Logger
	cfg         *config.Config
}

// NewBalanceService Balance service constructor
func NewBalanceService(balanceRepo balance.BalanceRepository, logger logger.Logger, cfg *config.Config) *BalanceService {
	return &BalanceService{balanceRepo: balanceRepo, logger: logger, cfg: cfg}
}
