package handler

import (
	"context"
	"encoding/json"
	"net/http"
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
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(400)
			res := validator.Response(validator.ValidateUsername(login.Username).Error())
			json.NewEncoder(w).Encode(res)
			return
		}

		res, err := store.LoginUser(ctx, rdb, login)

		if err != nil {
			if err.Error() == validator.ErrorUserNotFound.Error() {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(404)
				json.NewEncoder(w).Encode(res)
				return

			} else if err.Error() == validator.ErrorWrongPassword.Error() {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(404)
				json.NewEncoder(w).Encode(res)
				return

			} else {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(500)
				json.NewEncoder(w).Encode(validator.Response(err.Error()))
				return
			}
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(res)

	}
}
