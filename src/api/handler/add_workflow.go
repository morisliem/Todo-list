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

type AddWorkflowRequest struct {
	Workflow string `json:"workflow"`
}

func AddWorkflow(ctx context.Context, rdb *redis.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request AddWorkflowRequest
		err := json.NewDecoder(r.Body).Decode(&request)

		if err != nil {
			response.BadRequest(w, r, response.Response(response.ErrorFailedToDecode.Error()))
			return
		}

		username := chi.URLParam(r, response.URLUsername)
		newWorkflow := request.Workflow

		if validator.ValidateUsername(username) != nil {
			res := response.Response(validator.ValidateUsername(username).Error())
			response.BadRequest(w, r, res)
			return
		}

		if validator.ValidateWorkflow(request.Workflow) != nil {
			res := response.Response(validator.ValidateWorkflow(newWorkflow).Error())
			response.BadRequest(w, r, res)
			return
		}

		err = store.AddWorkflow(ctx, rdb, username, newWorkflow)

		switch err.(type) {
		case nil:
			response.SuccessfullyCreated(w, r)
			return

		case *response.BadInputError:
			response.NotFound(w, r, response.Response(err.Error()))
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
