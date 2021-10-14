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

type updateTodoRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Label       string `json:"label"`
	Deadline    string `json:"deadline"`
	Severity    string `json:"severity"`
	Priority    string `json:"priority"`
	State       string `json:"state"`
}

func UpdateTodo(ctx context.Context, rdb *redis.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username := chi.URLParam(r, response.URLUsername)
		todoId := chi.URLParam(r, response.URLUTodoId)
		var request updateTodoRequest

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

		switch err.(type) {
		case *response.NotFoundError:
			response.NotFound(w, r, response.Response(err.Error()))
			log.Error().Err(err).Msg(err.Error())
			return
		}

		newTodo := store.Todo{
			Title:       request.Title,
			Description: request.Description,
			Label:       request.Label,
			Deadline:    request.Deadline,
			Severity:    request.Severity,
			Priority:    request.Priority,
			State:       request.State,
		}

		err = store.UpdateTodo(ctx, rdb, username, todoId, newTodo)

		switch err.(type) {
		case nil:
			response.SuccessfullyCreated(w, r)
			return
		case *response.NotFoundError:
			response.NotFound(w, r, response.Response(err.Error()))
			log.Error().Err(err).Msg(err.Error())
			return
		case *response.DataStoreError:
			response.NotFound(w, r, response.Response(err.Error()))
			log.Error().Err(err).Msg(err.Error())
			return
		default:
			response.ServerError(w, r)
			log.Error().Err(err).Msg(err.Error())
			return

		}
	}
}
