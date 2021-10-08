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

func AddTodo(ctx context.Context, rdb *redis.Client) http.HandlerFunc {
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

		if validator.ValidateTodoTitle(newTodo.Title) != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(400)
			res := validator.Response(validator.ValidateTodoTitle(newTodo.Title).Error())
			json.NewEncoder(w).Encode(res)
			return
		}

		if validator.ValidateTodoPriority(newTodo.Priority) != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(400)
			res := validator.Response(validator.ValidateTodoPriority(newTodo.Priority).Error())
			json.NewEncoder(w).Encode(res)
			return
		}

		if validator.ValidateTodoSeverity(newTodo.Severity) != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(400)
			res := validator.Response(validator.ValidateTodoSeverity(newTodo.Severity).Error())
			json.NewEncoder(w).Encode(res)
			return
		}

		if validator.ValidateTodoDeadline(newTodo.Deadline) != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(400)
			res := validator.Response(validator.ValidateTodoDeadline(newTodo.Deadline).Error())
			json.NewEncoder(w).Encode(res)
			return
		}

		if validator.ValidateTodoState(newTodo.State) != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(400)
			res := validator.Response(validator.ValidateTodoState(newTodo.State).Error())
			json.NewEncoder(w).Encode(res)
			return
		}

		res, err := store.AddTodo(ctx, rdb, username, newTodo)

		if err != nil {
			if err.Error() == validator.FailedToAddTodo {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(400)
				json.NewEncoder(w).Encode(res)
				return

			}
			if err.Error() == validator.FailedToUpdateUserTodo {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(400)
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
