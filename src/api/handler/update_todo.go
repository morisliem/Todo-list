package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
	"todo-list/src/api/response"
	"todo-list/src/api/validator"
	"todo-list/src/store"

	"github.com/go-chi/chi"
	"github.com/go-redis/redis/v8"
)

func UpdateTodo(ctx context.Context, rdb *redis.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username := chi.URLParam(r, response.URLUsername)
		todoId := chi.URLParam(r, response.URLUTodoId)
		request := map[string]string{}

		err := json.NewDecoder(r.Body).Decode(&request)

		if err != nil {
			response.BadRequest(w, r, response.Response(response.ErrorFailedToDecode.Error()))
			return
		}

		if validator.ValidateUsername(username) != nil {
			response.BadRequest(w, r, response.Response(validator.ValidateUsername(username).Error()))
			return
		}

		if validator.ValidateTodoId(todoId) != nil {
			response.BadRequest(w, r, response.Response(validator.ValidateUsername(todoId).Error()))
			return
		}

		_, err = store.GetUser(ctx, rdb, username)

		if err != nil {
			if err == response.ErrorUserNotFound {
				response.NotFound(w, r, response.Response(err.Error()))
				return
			}
			response.ServerError(w, r)
			return
		}

		newTodo := store.Todo{
			Title:       request["title"],
			Description: request["description"],
			Label:       request["label"],
			Deadline:    request["deadline"],
			Severity:    request["severity"],
			Priority:    request["priority"],
			State:       request["state"],
			Created_at:  time.Now(),
		}

		res, err := store.UpdateTodo(ctx, rdb, username, todoId, newTodo)

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
