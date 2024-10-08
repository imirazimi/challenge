package contract

import "context"

//go:generate mockery --name Cache
type Cache interface {
	GetAll(ctx context.Context) ([]any, error)
	Get(ctx context.Context, key string) (any, error)
	Set(ctx context.Context, key string, value any) error
}
