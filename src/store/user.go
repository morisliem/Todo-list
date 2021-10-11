package store

import (
	"context"
	"strings"
	"time"
	"todo-list/src/api/response"

	"github.com/go-redis/redis/v8"
)

type User struct {
	Username   string    `json:"username"`
	Password   string    `json:"password"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Picture    string    `json:"picture"`
	Created_at time.Time `json:"created_at"`
	Deleted_at time.Time `json:"deleted_at"`
	Todo_lists []string  `json:"todo_list"`
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

func AddUser(ctx context.Context, db *redis.Client, usr User) error {

	err := db.HMSet(
		ctx, usr.Username,
		HmapKeyUserName, usr.Name,
		HmapKeyUserPassword, usr.Password,
		HmapKeyUserEmail, usr.Email,
		HmapKeyUserPicture, usr.Picture,
		HmapKeyUserCreatedAt, time.Now(),
		HmapKeyUserDeletedAt, usr.Deleted_at).Err()

	if err != nil {
		return &response.DataStoreError{Message: err.Error()}
	}

	return nil
}

func GetUser(ctx context.Context, db *redis.Client, usr string) (User, error) {
	value, err := db.HGetAll(ctx, usr).Result()
	user := User{}

	if err != nil {
		return user, err
	}

	if len(value) == 0 {
		return user, response.ErrorUserNotFound
	}

	for key, val := range value {
		if key == "created_at" {
			tmp, err := time.Parse("2006-01-02T15:04:05Z07:00", val)
			if err != nil {
				return user, err
			}
			user.Created_at = tmp
		}
		if key == "email" {
			user.Email = val
		}
		if key == "name" {
			user.Name = val
		}
		if key == "password" {
			user.Password = val
		}
		if key == "picture" {
			user.Picture = val
		}
		if key == "todosId" {
			tmp := strings.Split(val, " ")
			for i, v := range tmp {
				if i%2 == 0 {
					user.Todo_lists = append(user.Todo_lists, v)
				}
			}
		}
	}
	user.Username = usr

	return user, nil
}

func LogoutUser(ctx context.Context, db *redis.Client, usr string) (map[string]string, error) {

	res := response.Response(response.SuccessfullyLogout)
	return res, nil
}
