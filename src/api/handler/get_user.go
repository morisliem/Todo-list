package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
	"todo-list/src/api/response"
	"todo-list/src/api/validator"
	"todo-list/src/store"

	"github.com/go-chi/chi"
	"github.com/go-redis/redis/v8"
)

type GetUserResponse struct {
	Username   string    `json:"username"`
	Password   string    `json:"password"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Picture    string    `json:"picture"`
	Created_at time.Time `json:"created_at"`
	Deleted_at time.Time `json:"deleted_at"`
	Todo_lists []string  `json:"todo_list"`
}

func GetUser(ctx context.Context, rdb *redis.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username := chi.URLParam(r, response.URLUsername)

		if validator.ValidateUsername(username) != nil {
			response.BadRequest(w, r, response.Response(validator.ValidateUsername(username).Error()))
			return
		}

		res, err := store.GetUser(ctx, rdb, username)

		getUserResponse := GetUserResponse{
			Username:   res.Username,
			Password:   res.Password,
			Name:       res.Name,
			Email:      res.Email,
			Created_at: res.Created_at,
			Deleted_at: res.Deleted_at,
			Picture:    res.Picture,
			Todo_lists: res.Todo_lists,
		}

		if err != nil {
			if err == response.ErrorUserNotFound {
				response.NotFound(w, r, response.Response(err.Error()))
				return
			}

			response.ServerError(w, r)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(getUserResponse)
	}
}
