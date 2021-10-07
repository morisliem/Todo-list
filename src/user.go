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

	if err != nil {
		res := Response(FailedToAddUser)
		return res, err
	}

	res := Response(SuccessfullyAdded)
	return res, nil
}

func GetUser(ctx context.Context, db *redis.Client, usr string) (map[string]string, error) {
	value, _ := db.HGetAll(ctx, usr).Result()

	result := map[string]string{}
	if len(value) == 0 {
		res := Response(UserNotFoundError().Error())
		return res, nil
	}

	for key, val := range value {
		result[key] = val
	}

	return result, nil
}

func LoginUser(ctx context.Context, db *redis.Client, usr LoginRequest) (map[string]string, error) {
	password, _ := db.HGet(ctx, usr.Username, HmapKeyUserPassword).Result()

	if len(password) == 0 {
		res := Response(UserNotFoundError().Error())
		return res, nil
	}

	if password != usr.Password {
		res := Response(WrongPassword().Error())
		return res, nil
	}

	res := Response(SuccessfullyLogin)
	return res, nil
}

func LogoutUser(ctx context.Context, db *redis.Client, usr string) (map[string]string, error) {

	res := Response(SuccessfullyLogout)
	return res, nil
}
