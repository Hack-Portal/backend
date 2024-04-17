package redis

import (
	"context"
	"errors"
	"log"
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

	connectTimeout  time.Duration
	conenctWaitTime time.Duration
	connectAttempts int

	client *redis.Client
}

type RedisClient interface {
	Connect(db int) (*redis.Client, error)
	Close()
}

func New(
	address, password string,
	connectTimeout, connectWaitTime time.Duration,
	connectAttempts int,
) RedisClient {
	return &redisClient{
		address:         address,
		password:        password,
		connectTimeout:  connectTimeout,
		conenctWaitTime: connectWaitTime,
		connectAttempts: connectAttempts,
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
			log.Println("failed to connect redis: ", err)
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
