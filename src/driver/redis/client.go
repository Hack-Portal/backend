package redis

import (
	"context"
	"errors"
	"time"

	"github.com/redis/go-redis/v9"
)

var (
	ErrRedisConnectionFailed = errors.New("redis connection failed")
)

const (
	DefaultConnectTimeout = 5
	DefaultConnectWait    = 5
	DefaultConnectAttempt = 3
)

type redisClient struct {
	address  string
	password string

	connectTimeout  int
	connectAttempts int
	conenctWaitTime int

	client *redis.Client
}

type RedisClient interface {
	Connect(db int) (*redis.Client, error)
	Close()
}

func New(address, password string, connectTimeout, connectAttempts, connectWaitTime *int) RedisClient {
	var (
		_connectTimeout  int = DefaultConnectTimeout
		_connectAttempts int = DefaultConnectAttempt
		_connectWaitTime int = DefaultConnectWait
	)

	if connectTimeout != nil {
		_connectTimeout = *connectTimeout
	}

	if connectAttempts != nil {
		_connectAttempts = *connectAttempts
	}

	if connectWaitTime != nil {
		_connectWaitTime = *connectWaitTime
	}

	return &redisClient{
		address:         address,
		password:        password,
		connectTimeout:  _connectTimeout,
		connectAttempts: _connectAttempts,
		conenctWaitTime: _connectWaitTime,
	}
}

func (rc *redisClient) Connect(db int) (*redis.Client, error) {
	rc.client = redis.NewClient(&redis.Options{
		Addr:        rc.address,
		Password:    rc.password, // no password set
		DB:          db,          // use default DB
		DialTimeout: time.Duration(rc.connectTimeout) * time.Second,
	})

	sleep := func() {
		time.Sleep(time.Duration(rc.conenctWaitTime) * time.Second)
	}

	for i := 0; i < rc.connectAttempts; i++ {
		_, err := rc.client.Ping(context.Background()).Result()
		if err != nil {
			sleep()
			continue
		}

		return rc.client, nil
	}

	return nil, ErrRedisConnectionFailed
}

func (rc *redisClient) Close() {
	rc.client.Close()
}
