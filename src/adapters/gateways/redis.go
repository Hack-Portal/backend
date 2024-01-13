package gateways

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/Hack-Portal/backend/src/usecases/dai"
	"github.com/redis/go-redis/v9"
	"golang.org/x/sync/singleflight"
)

type redisGateway struct {
	db *redis.Client
}

// NewRedisGateway はredisGatewayのインスタンスを生成する
func NewRedisGateway(db *redis.Client) dai.RedisDai {
	return &redisGateway{
		db: db,
	}
}

// Get はキャッシュを取得する
func (r *redisGateway) Get(ctx context.Context, key string) ([]byte, bool, error) {
	bytes, err := r.db.Get(ctx, key).Bytes()
	// Cache not found
	if err == redis.Nil {
		return nil, false, nil
	}

	if err != nil {
		return nil, false, err
	}

	return bytes, true, nil
}

// Set はキャッシュを設定する
func (r *redisGateway) Set(ctx context.Context, key string, value []byte, deadline time.Duration) error {
	return r.db.Set(ctx, key, value, deadline).Err()
}

type cache[T any] struct {
	db         dai.RedisDai
	expiration time.Duration
	sfg        *singleflight.Group
}

// NewCache はcacheのインスタンスを生成する
func NewCache[T any](db *redis.Client, expiration time.Duration) dai.Cache[T] {
	return &cache[T]{
		db:         NewRedisGateway(db),
		expiration: expiration,
		sfg:        &singleflight.Group{},
	}
}

// Reset はキャッシュをリセットする
func (h *cache[T]) Reset(ctx context.Context, key string) error {
	return h.db.Set(ctx, key, nil, 0)
}

// Get はキャッシュを取得する
func (h *cache[T]) Get(ctx context.Context, key string, callback func(ctx context.Context) (T, error)) (T, error) {
	a, err, _ := h.sfg.Do(key, func() (interface{}, error) {
		bytes, exists, err := h.db.Get(ctx, key)
		if err != nil {
			return nil, err
		}

		if exists {
			return bytes, nil
		}

		t, err := callback(ctx)
		if err != nil {
			return nil, err
		}

		bytes, err = json.Marshal(t)
		if err != nil {
			return nil, err
		}

		err = h.db.Set(ctx, key, bytes, h.expiration)
		if err != nil {
			return nil, err
		}

		return bytes, nil
	})

	var t T
	if err != nil {
		return t, err
	}

	bytes, ok := a.([]byte)
	if !ok {
		return t, fmt.Errorf("failed to get from cache: invalid type %T", a)
	}

	err = json.Unmarshal(bytes, &t)
	if err != nil {
		return t, err
	}
	return t, nil
}
