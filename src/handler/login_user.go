package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"todo-list/src"

	"github.com/go-redis/redis/v8"
)

func LoginUserHandler(ctx context.Context, rdb *redis.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := map[string]string{}
		json.NewDecoder(r.Body).Decode(&request)

		login := src.LoginRequest{
			Username: request["username"],
			Password: request["password"],
		}

		if src.ValidateUsername(login.Username) != nil {
			w.WriteHeader(400)
			res := src.Response(src.ValidateUsername(login.Username).Error())
			json.NewEncoder(w).Encode(res)
			return
		}

		res, err := src.LoginUser(ctx, rdb, login)

		if err != nil {
			w.WriteHeader(404)
			json.NewEncoder(w).Encode(res)
			return
		}

		w.WriteHeader(200)
		json.NewEncoder(w).Encode(res)

	}
}
