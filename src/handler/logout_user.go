package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"todo-list/src"

	"github.com/go-chi/chi"
	"github.com/go-redis/redis/v8"
)

func LogoutUserHandler(ctx context.Context, rdb *redis.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username := chi.URLParam(r, "username")
		res, err := src.LogoutUser(ctx, rdb, username)

		if err != nil {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(res)
		}
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(res)
	}
}
