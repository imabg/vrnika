package redis

import (
	"context"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
)

type Redis struct {
	Client *redis.Client
}

var ctx = context.Background()

// ConnectToRedis Create new instance of Redis
func ConnectToRedis() *Redis {
	return &Redis{
		Client: redis.NewClient(&redis.Options{
			Addr:     os.Getenv("REDIS_URL"),
			Username: os.Getenv("REDIS_USERNAME"),
			Password: os.Getenv("REDIS_PWD"),
		}),
	}
}

// Ping check redis status
func (r *Redis) Ping() {
	if err := r.Client.Ping(ctx); err != nil {
		panic(err)
	}
}

// Set sets value in the store
func (r *Redis) Set(key string, value string, expireIn time.Duration) error {
	err := r.Client.Set(ctx, key, value, expireIn).Err()
	if err != nil {
		return err
	}
	return nil
}

// Retrieve gets value from the store
func (r *Redis) Retrieve(key string) (string, error) {
	val, err := r.Client.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}
