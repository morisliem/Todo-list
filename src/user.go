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

func AddUser(ctx context.Context, db *redis.Client, usr User) (map[string]string, error) {
	err := db.HMSet(
		ctx, usr.Username,
		"name", usr.Name,
		"password", usr.Password,
		"email", usr.Email,
		"picture", usr.Picture,
		"created_at", usr.Created_at,
		"deleted_at", usr.Deleted_at).Err()

	result := map[string]string{}
	if err != nil {
		result["Message"] = "Failed to add"
		return result, err
	}

	result["Message"] = "Successfully added"
	return result, nil
}

func GetUser(ctx context.Context, db *redis.Client, usr string) (map[string]string, error) {
	value, err := db.HGetAll(ctx, usr).Result()
	if err != nil {
		panic(err)
	}

	result := map[string]string{}
	if len(value) == 0 {
		result["Message"] = "User not found"
		return result, nil
	}

	for key, val := range value {
		result[key] = val
	}

	return result, nil
}

func LoginUser(ctx context.Context, db *redis.Client, usr User) (map[string]string, error) {
	password, err := db.HGet(ctx, usr.Username, "password").Result()
	result := map[string]string{}

	if err != nil {
		result["Message"] = "User not found"
		return result, err
	}

	if password != usr.Password {
		result["Message"] = "Invalid password"
		return result, nil
	}

	result["Message"] = "Logged in successfully"
	return result, nil
}

func LogoutUser(ctx context.Context, db *redis.Client, usr string) (map[string]string, error) {
	result := map[string]string{}

	result["Message"] = "Logged out successfully"
	return result, nil
}
