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
)

func AddWorkflow(ctx context.Context, rdb *redis.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := map[string]string{}
		err := json.NewDecoder(r.Body).Decode(&request)

		if err != nil {
			res := validator.Response(validator.FailedToDecode)
			response.BadRequest(w, r, res)
			return
		}

		username := chi.URLParam(r, validator.URLUsername)
		newWorkflow := request["workflow"]

		if validator.ValidateUsername(username) != nil {
			res := validator.Response(validator.ValidateUsername(username).Error())
			response.BadRequest(w, r, res)
			return
		}

		if validator.ValidateWorkflow(newWorkflow) != nil {
			res := validator.Response(validator.ValidateWorkflow(newWorkflow).Error())
			response.BadRequest(w, r, res)
			return
		}

		res, err := store.AddWorkflow(ctx, rdb, username, newWorkflow)

		if err != nil {
			if err.Error() == validator.FailedToAddWorkflow {
				response.NotFound(w, r, res)
				return
			}
			response.ServerError(w, r, validator.Response(err.Error()))
			return
		}

		response.SuccessfullyCreated(w, r, res)
	}
}
