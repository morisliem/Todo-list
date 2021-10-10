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
		username := chi.URLParam(r, validator.URLUsername)
		todoId := chi.URLParam(r, validator.URLUTodoId)
		request := map[string]string{}

		err := json.NewDecoder(r.Body).Decode(&request)

		if err != nil {
			res := validator.Response(validator.FailedToDecode)
			response.BadRequest(w, r, res)
			return
		}

		if validator.ValidateUsername(username) != nil {
			res := validator.Response(validator.ValidateUsername(username).Error())
			response.BadRequest(w, r, res)
			return
		}

		if validator.ValidateTodoId(todoId) != nil {
			res := validator.Response(validator.ValidateUsername(todoId).Error())
			response.BadRequest(w, r, res)
			return
		}

		_, err = store.GetUser(ctx, rdb, username)

		if err != nil {
			if err.Error() == validator.ErrorUserNotFound.Error() {
				response.NotFound(w, r, validator.Response(err.Error()))
				return
			}
			response.ServerError(w, r, validator.Response(err.Error()))
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
