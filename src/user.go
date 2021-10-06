package src

import (
	"context"
	"time"

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
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

const (
	HmapKeyUserName      = "name"
	HmapKeyUserPassword  = "password"
	HmapKeyUserEmail     = "email"
	HmapKeyUserPicture   = "picture"
	HmapKeyUserCreatedAt = "created_at"
	HmapKeyUserDeletedAt = "deleted_at"
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

	result := map[string]string{}
	if err != nil {
		result["Message"] = "Failed to add user"
		return result, err
	}

	result["Message"] = SuccessfullyAdded()
	return result, nil
}

func GetUser(ctx context.Context, db *redis.Client, usr string) (map[string]string, error) {
	value, _ := db.HGetAll(ctx, usr).Result()

	result := map[string]string{}
	if len(value) == 0 {
		result["Message"] = UserNotFoundError().Error()
		return result, nil
	}

	for key, val := range value {
		result[key] = val
	}

	return result, nil
}

func LoginUser(ctx context.Context, db *redis.Client, usr LoginRequest) (map[string]string, error) {
	password, _ := db.HGet(ctx, usr.Username, HmapKeyUserPassword).Result()
	result := map[string]string{}

	if len(password) == 0 {
		result["Message"] = UserNotFoundError().Error()
		return result, nil
	}

	if password != usr.Password {
		result["Message"] = WrongPassword().Error()
		return result, nil
	}

	result["Message"] = SuccessfullyLogin()
	return result, nil
}

func LogoutUser(ctx context.Context, db *redis.Client, usr string) (map[string]string, error) {
	result := map[string]string{}

	result["Message"] = SuccessfullyLogout()
	return result, nil
}
