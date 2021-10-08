package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
	"todo-list/src/api/validator"
	"todo-list/src/store"

	"github.com/go-chi/chi"
	"github.com/go-redis/redis/v8"
)

// Missing validation (half done)
func UpdateTodo(ctx context.Context, rdb *redis.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username := chi.URLParam(r, validator.URLUsername)
		todoId := chi.URLParam(r, validator.URLUTodoId)
		request := map[string]string{}

		err := json.NewDecoder(r.Body).Decode(&request)

		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(400)
			res := validator.Response(validator.FailedToDecode)
			json.NewEncoder(w).Encode(res)
			return
		}

		_, err = store.GetUser(ctx, rdb, username)

		if err != nil {
			if err.Error() == validator.ErrorUserNotFound.Error() {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(404)
				json.NewEncoder(w).Encode(validator.Response(err.Error()))
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(500)
			json.NewEncoder(w).Encode(validator.Response(err.Error()))
			return
		}

		newTodo := store.Todo{
			Title:       request["title"],
			Description: request["description"],
			Label:       request["label"],
			Deadline:    request["deadline"],
			Severity:    request["severity"],
			Priority:    request["priority"],
			State:       request["state"],
			Created_at:  time.Now(),
		}

		res, err := store.UpdateTodo(ctx, rdb, username, todoId, newTodo)

		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(500)
			json.NewEncoder(w).Encode(validator.Response(err.Error()))
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(res)
	}
}
