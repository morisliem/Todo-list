package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"todo-list/src"

	"github.com/go-chi/chi"
	"github.com/go-redis/redis/v8"
)

func Get_user(ctx context.Context, rdb *redis.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username := chi.URLParam(r, "username")

		res, err := src.GetUser(ctx, rdb, username)

		if err != nil {
			w.WriteHeader(404)
			json.NewEncoder(w).Encode(res["Message"])
		}
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(res)
	}
}
