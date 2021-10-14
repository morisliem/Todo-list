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
		return user, &response.NotFoundError{Message: response.ErrorUserNotFound.Error()}
	}

	for key, val := range value {
		if key == HmapKeyUserCreatedAt {
			tmp, err := time.Parse(time.RFC3339, val)
			if err != nil {
				return user, err
			}
			user.Created_at = tmp
		}
		if key == HmapKeyUserEmail {
			user.Email = val
		}
		if key == HmapKeyUserName {
			user.Name = val
		}
		if key == HmapKeyUserPassword {
			user.Password = val
		}
		if key == HmapKeyUserPicture {
			user.Picture = val
		}
		if key == HmapKeyUserTodos {
			tmp := strings.Split(val, " ")
			for _, v := range tmp {
				if v != "" {
					user.Todo_lists = append(user.Todo_lists, v)
				}
			}
		}
	}
	user.Username = usr

	return user, nil
}

func LogoutUser(ctx context.Context, db *redis.Client, usr string) error {
	isUserExist, err := db.HGet(ctx, usr, HmapKeyUserName).Result()

	if len(isUserExist) == 0 {
		return &response.NotFoundError{Message: response.ErrorUserNotFound.Error()}
	}

	if err != nil {
		return err
	}

	return nil
}

func AddUserPicture(ctx context.Context, db *redis.Client, usr string, pict string) error {
	err := db.HMSet(ctx, usr, HmapKeyUserPicture, pict).Err()

	if err != nil {
		return &response.BadInputError{Message: response.ErrorFailedToAddPict.Error()}
	}

	return nil
}

// func uploadPict() {

// }
