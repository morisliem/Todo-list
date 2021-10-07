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

		username := chi.URLParam(r, src.URLUsername)

		if src.ValidateUsername(username) != nil {
			w.WriteHeader(400)
			res := src.Response(src.ValidateUsername(username).Error())
			json.NewEncoder(w).Encode(res)
			return
		}

		res, err := src.LogoutUser(ctx, rdb, username)

		if err != nil {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(res)
			return
		}

		w.WriteHeader(200)
		json.NewEncoder(w).Encode(res)
	}
}
