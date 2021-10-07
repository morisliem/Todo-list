package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"todo-list/src"

	"github.com/go-chi/chi"
	"github.com/go-redis/redis/v8"
)

func AddWorkflowHandler(ctx context.Context, rdb *redis.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := map[string]string{}
		err := json.NewDecoder(r.Body).Decode(&request)

		if err != nil {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(src.FailedToDecode)
			return
		}

		username := chi.URLParam(r, src.URLUsername)
		newWorkflow := request["workflow"]

		if src.ValidateUsername(username) != nil {
			w.WriteHeader(400)
			res := src.Response(src.ValidateUsername(username).Error())
			json.NewEncoder(w).Encode(res)
			return
		}

		if src.ValidateWorkflow(newWorkflow) != nil {
			w.WriteHeader(400)
			res := src.Response(src.ValidateWorkflow(newWorkflow).Error())
			json.NewEncoder(w).Encode(res)
			return
		}

		res, err := src.AddWorkflow(ctx, rdb, username, newWorkflow)

		if err != nil {
			w.WriteHeader(404)
			json.NewEncoder(w).Encode(res)
			return
		}

		w.WriteHeader(201)
		json.NewEncoder(w).Encode(res)
	}
}
