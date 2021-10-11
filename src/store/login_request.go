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

func LoginUser(ctx context.Context, db *redis.Client, usr LoginRequest) (map[string]string, error) {
	password, err := db.HGet(ctx, usr.Username, HmapKeyUserPassword).Result()

	if len(password) == 0 {
		res := response.Response(response.ErrorUserNotFound.Error())
		return res, nil
	}

	if password != usr.Password {
		res := response.Response(response.ErrorWrongPassword.Error())
		return res, nil
	}

	if err != nil {
		return map[string]string{}, err
	}

	res := response.Response(response.SuccessfullyLogin)
	return res, nil
}
