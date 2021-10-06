package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
	"todo-list/src"

	"github.com/go-redis/redis/v8"
)

func RegisterUserHandler(ctx context.Context, rdb *redis.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := map[string]string{}
		err := json.NewDecoder(r.Body).Decode(&request)

		if err != nil {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode("Failed to decode")
			return
		}

		newUser := src.User{
			Username:   request["username"],
			Password:   request["password"],
			Name:       request["name"],
			Email:      request["email"],
			Picture:    request["picture"],
			Created_at: time.Now(),
		}

		if src.ValidateUsername(newUser.Username) != nil {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(src.ValidateUsername(request["username"]).Error())
			return
		}

		if src.ValidateName(newUser.Name) != nil {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(src.ValidateName(request["name"]).Error())
			return
		}

		if src.ValidatePassword(newUser.Password) != nil {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(src.ValidatePassword(request["password"]).Error())
			return
		}

		if src.ValidateEmail(newUser.Email) != nil {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(src.ValidateEmail(request["email"]).Error())
			return
		}

		res, err := src.AddUser(ctx, rdb, newUser)

		if err != nil {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(res)
			return
		}

		w.WriteHeader(201)
		json.NewEncoder(w).Encode(res)

	}
}
