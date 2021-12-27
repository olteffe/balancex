//go:generate mockgen -source redis_repository.go -destination mock/redis_repository.go -package mock
package transaction

import "context"

// RedisRepository interface
type RedisRepository interface {
	GetRate(ctx context.Context, rate string) (float32, error)
}
