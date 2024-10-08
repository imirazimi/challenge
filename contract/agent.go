package contract

import "context"

//go:generate mockery --name Agent
type Agent interface {
	JSON(ctx context.Context, data []string, format string) (string, error)
}
