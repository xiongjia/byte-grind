package util

import (
	"context"
	"log/slog"
	"os"
	"time"

	redis "github.com/redis/go-redis/v9"
)

// RedisMgr connects to Redis (address can be overridden via REDIS_ADDR env var),
// writes k1=v1 and reads it back. Returns error on failure.
func RedisMgr() error {
	addr := os.Getenv("REDIS_ADDR")
	if addr == "" {
		addr = "192.168.71.180:6379"
	}
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "",
		DB:       0,
	})

	ctx := context.TODO()
	if err := rdb.Set(ctx, "k1", "v1", 0).Err(); err != nil {
		return err
	}

	if val, err := rdb.Get(ctx, "k1").Result(); err == nil {
		slog.Debug("v", slog.String("v", val))
	}

	result, err := rdb.SetNX(ctx, "key1", 1, 10*time.Second).Result()
	if err != nil {
		return err
	}
	slog.Debug("result", slog.Bool("result", result))

	return nil
}
