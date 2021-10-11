package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"todo-list/src/api/response"
	"todo-list/src/api/validator"
	"todo-list/src/store"

	"github.com/go-redis/redis/v8"
)

func LoginUser(ctx context.Context, rdb *redis.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := map[string]string{}
		err := json.NewDecoder(r.Body).Decode(&request)

		if err != nil {
			response.BadRequest(w, r, response.Response(response.ErrorFailedToDecode.Error()))
			return
		}

		login := store.LoginRequest{
			Username: request["username"],
			Password: request["password"],
		}

		if validator.ValidateUsername(login.Username) != nil {
			response.BadRequest(w, r, response.Response(validator.ValidateUsername(login.Username).Error()))
			return
		}

		res, err := store.LoginUser(ctx, rdb, login)

		if err != nil {
			if err == response.ErrorUserNotFound {
				response.NotFound(w, r, res)
				return

			} else if err == response.ErrorWrongPassword {
				response.BadRequest(w, r, res)
				return

			} else {
				response.ServerError(w, r)
				return
			}
		}

		response.SuccessfullyOk(w, r, res)
	}
}
