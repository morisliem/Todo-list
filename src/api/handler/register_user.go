package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
	"todo-list/src/api/response"
	"todo-list/src/api/validator"
	"todo-list/src/store"

	"github.com/go-redis/redis/v8"
)

func RegisterUser(ctx context.Context, rdb *redis.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := map[string]string{}
		err := json.NewDecoder(r.Body).Decode(&request)

		if err != nil {
			res := validator.Response(validator.FailedToDecode)
			response.BadRequest(w, r, res)
			return
		}

		newUser := store.User{
			Username:   request["username"],
			Password:   request["password"],
			Name:       request["name"],
			Email:      request["email"],
			Picture:    request["picture"],
			Created_at: time.Now(),
		}

		if validator.ValidateUsername(newUser.Username) != nil {
			res := validator.Response(validator.ValidateUsername(newUser.Username).Error())
			response.BadRequest(w, r, res)
			return
		}

		if validator.ValidateName(newUser.Name) != nil {
			res := validator.Response(validator.ValidateName(newUser.Name).Error())
			response.BadRequest(w, r, res)
			return
		}

		if validator.ValidatePassword(newUser.Password) != nil {
			res := validator.Response(validator.ValidatePassword(newUser.Password).Error())
			response.BadRequest(w, r, res)
			return
		}

		if validator.ValidateEmail(newUser.Email) != nil {
			res := validator.Response(validator.ValidateEmail(newUser.Email).Error())
			response.BadRequest(w, r, res)
			return
		}

		res, err := store.AddUser(ctx, rdb, newUser)

		if err != nil {
			if err.Error() == validator.FailedToAddUser {
				response.BadRequest(w, r, res)
				return
			}

			response.ServerError(w, r, validator.Response(err.Error()))
			return

		}

		response.SuccessfullyCreated(w, r, res)
	}
}
