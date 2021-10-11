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
			response.BadRequest(w, r, response.Response(response.ErrorFailedToDecode.Error()))
			return
		}

		username := chi.URLParam(r, response.URLUsername)
		newWorkflow := request["workflow"]

		if validator.ValidateUsername(username) != nil {
			res := response.Response(validator.ValidateUsername(username).Error())
			response.BadRequest(w, r, res)
			return
		}

		if validator.ValidateWorkflow(newWorkflow) != nil {
			res := response.Response(validator.ValidateWorkflow(newWorkflow).Error())
			response.BadRequest(w, r, res)
			return
		}

		res, err := store.AddWorkflow(ctx, rdb, username, newWorkflow)
		switch err.(type) {
		case nil:
			response.SuccessfullyCreated(w, r, res)
		case *response.BadInputError:
			response.BadRequest(w, r, response.Response(err.Error()))
			return
		case *response.DataStoreError:
			response.NotFound(w, r, res)
			return
		case *response.ServerInternalError:
			response.NotFound(w, r, res)
			return
		default:
			response.NotFound(w, r, res)
			return
		}
	}
}
