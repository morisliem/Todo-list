package handler

import (
	"context"
	"net/http"
	"todo-list/src/api/response"
	"todo-list/src/api/validator"
	"todo-list/src/store"

	"github.com/go-chi/chi"
	"github.com/go-redis/redis/v8"
)

func LogoutUser(ctx context.Context, rdb *redis.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		username := chi.URLParam(r, validator.URLUsername)

		if validator.ValidateUsername(username) != nil {
			res := validator.Response(validator.ValidateUsername(username).Error())
			response.BadRequest(w, r, res)
			return
		}

		res, err := store.LogoutUser(ctx, rdb, username)

		if err != nil {
			response.BadRequest(w, r, res)
			return
		}

		response.SuccessfullyOk(w, r, res)
	}
}
