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
	"github.com/rs/zerolog/log"
)

type AddTodoRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Label       string `json:"label"`
	Deadline    string `json:"deadline"`
	Severity    string `json:"severity"`
	Priority    string `json:"priority"`
	State       string `json:"state"`
	Created_at  time.Time
}

func AddTodo(ctx context.Context, rdb *redis.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request AddTodoRequest
		err := json.NewDecoder(r.Body).Decode(&request)

		if err != nil {
			response.BadRequest(w, r, response.Response(response.ErrorFailedToDecode.Error()))
			return
		}

		username := chi.URLParam(r, response.URLUsername)
		_, err = store.GetUser(ctx, rdb, username)

		if err != nil {
			if err == response.ErrorUserNotFound {
				response.NotFound(w, r, response.Response(err.Error()))
				return
			}

			response.ServerError(w, r)
			return
		}

		if request.validateRequest(w, r) {
			newTodo := store.Todo{
				Title:       request.Title,
				Description: request.Description,
				Label:       request.Label,
				Deadline:    request.Deadline,
				Severity:    request.Severity,
				Priority:    request.Priority,
				State:       request.State,
				Created_at:  time.Now(),
			}

			err = store.AddTodo(ctx, rdb, username, newTodo)

			switch err.(type) {
			case nil:
				response.SuccessfullyCreated(w, r)
				return
			case *response.BadInputError:
				response.BadRequest(w, r, response.Response(err.Error()))
				log.Error().Err(err).Msg(err.Error())
				return
			case *response.DataStoreError:
				response.BadRequest(w, r, response.Response(err.Error()))
				log.Error().Err(err).Msg(err.Error())
				return
			default:
				response.ServerError(w, r)
				log.Error().Err(err).Msg(err.Error())
				return

			}
		}
	}
}

func (td *AddTodoRequest) validateRequest(w http.ResponseWriter, r *http.Request) bool {
	if validator.ValidateTodoTitle(td.Title) != nil {
		res := response.Response(validator.ValidateTodoTitle(td.Title).Error())
		response.BadRequest(w, r, res)
		return false
	}

	if validator.ValidateTodoPriority(td.Priority) != nil {
		res := response.Response(validator.ValidateTodoPriority(td.Priority).Error())
		response.BadRequest(w, r, res)
		return false
	}

	if validator.ValidateTodoSeverity(td.Severity) != nil {
		res := response.Response(validator.ValidateTodoSeverity(td.Severity).Error())
		response.BadRequest(w, r, res)
		return false
	}

	if validator.ValidateTodoDeadline(td.Deadline) != nil {
		res := response.Response(validator.ValidateTodoDeadline(td.Deadline).Error())
		response.BadRequest(w, r, res)
		return false
	}

	if validator.ValidateTodoState(td.State) != nil {
		res := response.Response(validator.ValidateTodoState(td.State).Error())
		response.BadRequest(w, r, res)
		return false
	}
	return true
}
