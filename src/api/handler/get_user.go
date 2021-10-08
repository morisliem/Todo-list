package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"todo-list/src/api/validator"
	"todo-list/src/store"

	"github.com/go-chi/chi"
	"github.com/go-redis/redis/v8"
)

func GetUser(ctx context.Context, rdb *redis.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username := chi.URLParam(r, validator.URLUsername)

		if validator.ValidateUsername(username) != nil {

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(400)
			res := validator.Response(validator.ValidateUsername(username).Error())
			json.NewEncoder(w).Encode(res)
			return
		}

		res, err := store.GetUser(ctx, rdb, username)

		if err != nil {
			if err.Error() == validator.ErrorUserNotFound.Error() {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(404)
				json.NewEncoder(w).Encode(validator.Response(err.Error()))
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(500)
			json.NewEncoder(w).Encode(validator.Response(err.Error()))
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(res)
	}
}
