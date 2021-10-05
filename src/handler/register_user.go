package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
	"todo-list/src"

	"github.com/go-redis/redis/v8"
)

func Register_user(ctx context.Context, rdb *redis.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := map[string]string{}
		json.NewDecoder(r.Body).Decode(&request)

		newUser := src.User{
			Username:   request["username"],
			Password:   request["password"],
			Name:       request["name"],
			Email:      request["email"],
			Picture:    request["picture"],
			Created_at: time.Now(),
		}

		res, err := src.AddUser(ctx, rdb, newUser)

		if err != nil {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(res)
		}
		w.WriteHeader(201)
		json.NewEncoder(w).Encode(res)
	}
}
