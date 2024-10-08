package memcache

import (
	"context"
	"errors"
)

type MemCache map[string]any

func (mc MemCache) GetAll(_ context.Context) ([]any, error) {
	all := []any{}
	for _, value := range mc {
		all = append(all, value)
	}
	return all, nil
}

func (mc MemCache) Get(_ context.Context, key string) (any, error) {
	value, ok := mc[key]
	if !ok {
		return "", errors.New("not found")
	}
	return value, nil
}

func (mc MemCache) Set(_ context.Context, key string, value any) error {
	mc[key] = value
	return nil
}

func New() *MemCache {
	return &MemCache{}
}
