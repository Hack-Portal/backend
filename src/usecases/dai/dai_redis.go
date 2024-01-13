package dai

import (
	"context"
	"time"
)

// RedisDai はRedisに関するデータアクセスインターフェース
type RedisDai interface {
	Get(ctx context.Context, key string) ([]byte, bool, error)
	Set(ctx context.Context, key string, value []byte, deadline time.Duration) error
}

// Cache はキャッシュに関するデータアクセスインターフェース
type Cache[T any] interface {
	Reset(ctx context.Context, key string) error
	Get(ctx context.Context, key string, callback func(ctx context.Context) (T, error)) (T, error)
}
