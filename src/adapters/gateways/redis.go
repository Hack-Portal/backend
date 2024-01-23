package gateways

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/Hack-Portal/backend/src/usecases/dai"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/redis/go-redis/v9"
	"golang.org/x/sync/singleflight"
)

type RedisGateway struct {
	db *redis.Client
}

func NewRedisGateway(db *redis.Client) dai.RedisDai {
	return &RedisGateway{
		db: db,
	}
}

func (r *RedisGateway) Get(ctx context.Context, key string) ([]byte, bool, error) {
	defer newrelic.FromContext(ctx).StartSegment("Get-gateway").End()
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

func (r *RedisGateway) Set(ctx context.Context, key string, value []byte, deadline time.Duration) error {
	defer newrelic.FromContext(ctx).StartSegment("Set-gateway").End()
	return r.db.Set(ctx, key, value, deadline).Err()
}

type Cache[T any] struct {
	db         dai.RedisDai
	expiration time.Duration
	sfg        *singleflight.Group
}

func NewCache[T any](db *redis.Client, expiration time.Duration) dai.Cache[T] {
	return &Cache[T]{
		db:         NewRedisGateway(db),
		expiration: expiration,
		sfg:        &singleflight.Group{},
	}
}

func (h *Cache[T]) Reset(ctx context.Context, key string) error {
	defer newrelic.FromContext(ctx).StartSegment("Reset-gateway").End()
	return h.db.Set(ctx, key, nil, 0)
}

func (h *Cache[T]) Get(ctx context.Context, key string, callback func(ctx context.Context) (T, error)) (T, error) {
	defer newrelic.FromContext(ctx).StartSegment("Get-gateway").End()
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
