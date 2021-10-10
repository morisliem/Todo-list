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

func AddTodo(ctx context.Context, rdb *redis.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := map[string]string{}
		err := json.NewDecoder(r.Body).Decode(&request)

		if err != nil {
			res := validator.Response(validator.FailedToDecode)
			response.BadRequest(w, r, res)
			return
		}

		username := chi.URLParam(r, validator.URLUsername)
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

		if validator.ValidateTodoTitle(newTodo.Title) != nil {
			res := validator.Response(validator.ValidateTodoTitle(newTodo.Title).Error())
			response.BadRequest(w, r, res)
			return
		}

		if validator.ValidateTodoPriority(newTodo.Priority) != nil {
			res := validator.Response(validator.ValidateTodoPriority(newTodo.Priority).Error())
			response.BadRequest(w, r, res)
			return
		}

		if validator.ValidateTodoSeverity(newTodo.Severity) != nil {
			res := validator.Response(validator.ValidateTodoSeverity(newTodo.Severity).Error())
			response.BadRequest(w, r, res)
			return
		}

		if validator.ValidateTodoDeadline(newTodo.Deadline) != nil {
			res := validator.Response(validator.ValidateTodoDeadline(newTodo.Deadline).Error())
			response.BadRequest(w, r, res)
			return
		}

		if validator.ValidateTodoState(newTodo.State) != nil {
			res := validator.Response(validator.ValidateTodoState(newTodo.State).Error())
			response.BadRequest(w, r, res)
			return
		}

		res, err := store.AddTodo(ctx, rdb, username, newTodo)

		if err != nil {
			if err.Error() == validator.FailedToAddTodo {
				response.BadRequest(w, r, res)
				return

			}
			if err.Error() == validator.FailedToUpdateUserTodo {
				response.BadRequest(w, r, res)
				return
			}

			response.BadRequest(w, r, validator.Response(err.Error()))
			return
		}
		response.SuccessfullyCreated(w, r, res)
	}
}
