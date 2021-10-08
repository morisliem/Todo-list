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

// Missing validation and have to clean the code
func GetTodos(ctx context.Context, rdb *redis.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		username := chi.URLParam(r, validator.URLUsername)
		res, err := store.GetTodos(ctx, rdb, username)

		if err != nil {
			w.WriteHeader(500)
			json.NewEncoder(w).Encode(validator.Response(err.Error()))
			return
		}

		w.WriteHeader(200)
		json.NewEncoder(w).Encode(res)
	}
}
