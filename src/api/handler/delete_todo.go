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
	"github.com/rs/zerolog/log"
)

func DeleteTodo(ctx context.Context, rdb *redis.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username := chi.URLParam(r, response.URLUsername)
		todoId := chi.URLParam(r, response.URLUTodoId)

		if validator.ValidateUsername(username) != nil {
			response.BadRequest(w, r, response.Response(validator.ValidateUsername(username).Error()))
			log.Error().Err(validator.ValidateUsername(username)).Msg(validator.ValidateUsername(username).Error())
			return
		}

		if validator.ValidateTodoId(todoId) != nil {
			response.BadRequest(w, r, response.Response(validator.ValidateTodoId(todoId).Error()))
			log.Error().Err(validator.ValidateTodoId(todoId)).Msg(validator.ValidateTodoId(todoId).Error())
			return
		}

		err := store.RemoveTodo(ctx, rdb, username, todoId)
		switch err.(type) {
		case nil:
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(200)
			return
		case *response.BadInputError:
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(response.Response(err.Error()))
			log.Error().Err(err).Msg(err.Error())
			return
		default:
			response.ServerError(w, r)
			log.Error().Err(err).Msg(err.Error())
		}
	}
}
