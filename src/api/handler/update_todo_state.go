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

type NewTodoState struct {
	NewState string `json:"new_state"`
}

func UpdateTodoState(ctx context.Context, rdb *redis.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		result := map[string]string{}
		err := json.NewDecoder(r.Body).Decode(&result)

		if err != nil {
			response.BadRequest(w, r, response.Response(response.ErrorFailedToDecode.Error()))
			log.Error().Err(err).Msg(response.ErrorFailedToDecode.Error())
			return
		}

		newState := NewTodoState{
			NewState: result["new_state"],
		}

		username := chi.URLParam(r, response.URLUsername)
		todoId := chi.URLParam(r, response.URLUTodoId)

		if validator.ValidateUsername(username) != nil {
			response.BadRequest(w, r, response.Response(validator.ValidateUsername(username).Error()))
			log.Error().Err(validator.ValidateUsername(username)).Msg(validator.ValidateUsername(username).Error())
			return
		}

		if validator.ValidateTodoId(todoId) != nil {
			response.BadRequest(w, r, response.Response(validator.ValidateTodoId(todoId).Error()))
			log.Error().Err(validator.ValidateTodoId(todoId)).Msg(validator.ValidateUsername(username).Error())
			return
		}

		err = store.UpdateTodoState(ctx, rdb, username, todoId, newState.NewState)

		switch err.(type) {
		case nil:
			response.SuccessfullyOk(w, r)
			return

		case *response.DataStoreError:
			response.BadRequest(w, r, response.Response(err.Error()))
			log.Error().Err(err).Msg(err.Error())
			return

		case *response.NotFoundError:
			response.NotFound(w, r, response.Response(err.Error()))
			log.Error().Err(err).Msg(err.Error())
			return

		default:
			response.ServerError(w, r)
			log.Error().Err(err).Msg(err.Error())
		}
	}
}
