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

type getWorkflowResponse struct {
	Workflow []string `json:"workflows"`
}

func GetWorkflow(ctx context.Context, rdb *redis.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username := chi.URLParam(r, response.URLUsername)

		if validator.ValidateUsername(username) != nil {
			response.BadRequest(w, r, response.Response(validator.ValidateUsername(username).Error()))
			return
		}

		res, err := store.GetWorkflow(ctx, rdb, username)

		getWorkflowResponse := getWorkflowResponse{
			Workflow: res.Workflows,
		}

		switch err.(type) {
		case nil:
			response.SuccessfullyOk(w, r)
			json.NewEncoder(w).Encode(getWorkflowResponse)
			return
		case *response.BadInputError:
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
