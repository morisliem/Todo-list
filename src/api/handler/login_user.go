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
		json.NewDecoder(r.Body).Decode(&request)

		login := store.LoginRequest{
			Username: request["username"],
			Password: request["password"],
		}

		if validator.ValidateUsername(login.Username) != nil {
			res := validator.Response(validator.ValidateUsername(login.Username).Error())
			response.BadRequest(w, r, res)
			return
		}

		res, err := store.LoginUser(ctx, rdb, login)

		if err != nil {
			if err.Error() == validator.ErrorUserNotFound.Error() {
				response.NotFound(w, r, res)
				return

			} else if err.Error() == validator.ErrorWrongPassword.Error() {
				response.BadRequest(w, r, res)
				return

			} else {
				response.ServerError(w, r, validator.Response(err.Error()))
				return
			}
		}

		response.SuccessfullyOk(w, r, res)
	}
}
