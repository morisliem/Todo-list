package store

import (
	"context"
	"time"
	"todo-list/src/api/validator"

	"github.com/go-redis/redis/v8"
)

type User struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Picture    string `json:"picture"`
	Created_at time.Time
	Deleted_at time.Time
	Todo_lists string
}

const (
	HmapKeyUserName      = "name"
	HmapKeyUserPassword  = "password"
	HmapKeyUserEmail     = "email"
	HmapKeyUserPicture   = "picture"
	HmapKeyUserCreatedAt = "created_at"
	HmapKeyUserDeletedAt = "deleted_at"
	HmapKeyUserTodos     = "todosId"
)

func AddUser(ctx context.Context, db *redis.Client, usr User) (map[string]string, error) {
	err := db.HMSet(
		ctx, usr.Username,
		HmapKeyUserName, usr.Name,
		HmapKeyUserPassword, usr.Password,
		HmapKeyUserEmail, usr.Email,
		HmapKeyUserPicture, usr.Picture,
		HmapKeyUserCreatedAt, usr.Created_at,
		HmapKeyUserDeletedAt, usr.Deleted_at).Err()

	if err != nil {
		res := validator.Response(validator.FailedToAddUser)
		return res, err
	}

	res := validator.Response(validator.SuccessfullyAdded)
	return res, nil
}

func GetUser(ctx context.Context, db *redis.Client, usr string) (map[string]string, error) {
	value, err := db.HGetAll(ctx, usr).Result()
	result := map[string]string{}

	if err != nil {
		return result, err
	}

	if len(value) == 0 {
		return result, validator.ErrorUserNotFound
	}

	for key, val := range value {
		result[key] = val
	}

	return result, nil
}

func LogoutUser(ctx context.Context, db *redis.Client, usr string) (map[string]string, error) {

	res := validator.Response(validator.SuccessfullyLogout)
	return res, nil
}
