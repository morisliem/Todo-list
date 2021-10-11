package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"todo-list/src/api/response"
	"todo-list/src/store"

	"github.com/go-chi/chi"
	"github.com/go-redis/redis/v8"
)

func GetTodos(ctx context.Context, rdb *redis.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		username := chi.URLParam(r, response.URLUsername)
		_, err := store.GetUser(ctx, rdb, username)

		if err != nil {
			if err == response.ErrorUserNotFound {
				response.NotFound(w, r, response.Response(err.Error()))
				return

			}

			response.ServerError(w, r)
			return
		}

		res, err := store.GetTodos(ctx, rdb, username)

		if err != nil {
			response.ServerError(w, r)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(res)
	}
}

func GetTodo(ctx context.Context, rdb *redis.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username := chi.URLParam(r, response.URLUsername)
		todoId := chi.URLParam(r, response.URLUTodoId)

		_, err := store.GetUser(ctx, rdb, username)

		if err != nil {
			if err == response.ErrorUserNotFound {
				response.NotFound(w, r, response.Response(err.Error()))
				return
			}

			response.ServerError(w, r)
			return
		}

		res, err := store.GetTodo(ctx, rdb, username, todoId)

		if err != nil {
			if err == response.ErrorTodoNotFound {
				response.NotFound(w, r, response.Response(err.Error()))
				return
			}

			response.ServerError(w, r)
			return
		}

		response.SuccessfullyOk(w, r, res)
	}
}
