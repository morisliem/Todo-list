package handler

import (
	"context"
	"encoding/json"
	"net/http"
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
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(400)
			res := validator.Response(validator.FailedToDecode)
			json.NewEncoder(w).Encode(res)
			return
		}

		username := chi.URLParam(r, validator.URLUsername)
		newWorkflow := request["workflow"]

		if validator.ValidateUsername(username) != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(400)
			res := validator.Response(validator.ValidateUsername(username).Error())
			json.NewEncoder(w).Encode(res)
			return
		}

		if validator.ValidateWorkflow(newWorkflow) != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(400)
			res := validator.Response(validator.ValidateWorkflow(newWorkflow).Error())
			json.NewEncoder(w).Encode(res)
			return
		}

		res, err := store.AddWorkflow(ctx, rdb, username, newWorkflow)

		if err != nil {
			if err.Error() == validator.FailedToAddWorkflow {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(404)
				json.NewEncoder(w).Encode(res)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(500)
			json.NewEncoder(w).Encode(validator.Response(err.Error()))
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(201)
		json.NewEncoder(w).Encode(res)
	}
}
