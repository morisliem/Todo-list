package src

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/go-redis/redis/v8"
)

func EstablishedConnection(ctx context.Context) (*redis.Client, error) {
	redisPort, _ := strconv.Atoi(os.Getenv("REDIS_PORT"))
	redisDB, _ := strconv.Atoi(os.Getenv("REDIS_DB"))
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", os.Getenv("REDIS_HOST"), redisPort),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       redisDB,
	})

	_, err := client.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}

	return client, nil
}
