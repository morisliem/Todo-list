package src

import "github.com/go-redis/redis/v8"

func EstablishedConnection() *redis.Client {
	redis_DB := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	return redis_DB
}
