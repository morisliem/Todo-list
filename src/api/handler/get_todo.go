package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"todo-list/src/api/response"
	"todo-list/src/api/validator"
	"todo-list/src/store"

	"github.com/go-chi/chi"
	"github.com/go-redis/redis/v8"
)

func GetTodos(ctx context.Context, rdb *redis.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		username := chi.URLParam(r, validator.URLUsername)
		_, err := store.GetUser(ctx, rdb, username)

		if err != nil {
			if err.Error() == validator.ErrorUserNotFound.Error() {
				response.NotFound(w, r, validator.Response(err.Error()))
				return

			}

			response.ServerError(w, r, validator.Response(err.Error()))
			return
		}

		res, err := store.GetTodos(ctx, rdb, username)

		if err != nil {
			response.ServerError(w, r, validator.Response(err.Error()))
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(res)
	}
}

func GetTodo(ctx context.Context, rdb *redis.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username := chi.URLParam(r, validator.URLUsername)
		todoId := chi.URLParam(r, validator.URLUTodoId)

		_, err := store.GetUser(ctx, rdb, username)

		if err != nil {
			if err.Error() == validator.ErrorUserNotFound.Error() {
				response.NotFound(w, r, validator.Response(err.Error()))
				return
			}

			response.ServerError(w, r, validator.Response(err.Error()))
			return
		}

		res, err := store.GetTodo(ctx, rdb, username, todoId)

		if err != nil {
			if err.Error() == validator.ErrorTodoNotFound.Error() {
				response.NotFound(w, r, validator.Response(err.Error()))
				return
			}

			response.ServerError(w, r, validator.Response(err.Error()))
			return
		}

		response.SuccessfullyOk(w, r, res)
	}
}
