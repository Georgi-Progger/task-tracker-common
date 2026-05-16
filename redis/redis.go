package redis

import (
	"context"
	"os"

	"github.com/Georgi-Progger/task-tracker-common/logger"
	"github.com/redis/go-redis/v9"
)

func NewRedisClient(ctx context.Context, dsn string, logger logger.Logger) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr: dsn,
		DB:   0,
	})

	ping, err := client.Ping(ctx).Result()
	if err != nil {
		logger.Error(err, "Failed to ping redis")
		os.Exit(1)
	}
	logger.Info(ping)

	return client, nil
}
