package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"todo-list/src"

	"github.com/go-chi/chi"
	"github.com/go-redis/redis/v8"
)

func Add_workflow(ctx context.Context, rdb *redis.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := map[string]string{}
		json.NewDecoder(r.Body).Decode(&request)

		username := chi.URLParam(r, "username")
		newWorkflow := request["workflow"]

		res, err := src.AddWorkflow(ctx, rdb, username, newWorkflow)

		if err != nil {
			w.WriteHeader(404)
			json.NewEncoder(w).Encode(res)
		}

		w.WriteHeader(201)
		json.NewEncoder(w).Encode(res)
	}
}
