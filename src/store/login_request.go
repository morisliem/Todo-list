package store

import (
	"context"
	"todo-list/src/api/response"

	"github.com/go-redis/redis/v8"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func LoginUser(ctx context.Context, db *redis.Client, usr LoginRequest) error {
	password, err := db.HGet(ctx, usr.Username, HmapKeyUserPassword).Result()

	if len(password) == 0 {
		return &response.BadInputError{Message: response.ErrorUserNotFound.Error()}
	}

	if password != usr.Password {
		return &response.BadInputError{Message: response.ErrorWrongPassword.Error()}
	}

	if err != nil {
		return err
	}

	return nil
}
