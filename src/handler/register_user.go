package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
	"todo-list/src"
	"todo-list/src/store"

	"github.com/go-redis/redis/v8"
)

func RegisterUser(ctx context.Context, rdb *redis.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := map[string]string{}
		err := json.NewDecoder(r.Body).Decode(&request)

		if err != nil {
			w.WriteHeader(400)
			res := src.Response(src.FailedToDecode)
			json.NewEncoder(w).Encode(res)
			return
		}

		newUser := store.User{
			Username:   request["username"],
			Password:   request["password"],
			Name:       request["name"],
			Email:      request["email"],
			Picture:    request["picture"],
			Created_at: time.Now(),
		}

		if src.ValidateUsername(newUser.Username) != nil {
			w.WriteHeader(400)
			res := src.Response(src.ValidateUsername(newUser.Username).Error())
			json.NewEncoder(w).Encode(res)
			return
		}

		if src.ValidateName(newUser.Name) != nil {
			w.WriteHeader(400)
			res := src.Response(src.ValidateName(newUser.Name).Error())
			json.NewEncoder(w).Encode(res)
			return
		}

		if src.ValidatePassword(newUser.Password) != nil {
			w.WriteHeader(400)
			res := src.Response(src.ValidatePassword(newUser.Password).Error())
			json.NewEncoder(w).Encode(res)
			return
		}

		if src.ValidateEmail(newUser.Email) != nil {
			w.WriteHeader(400)
			res := src.Response(src.ValidateEmail(newUser.Email).Error())
			json.NewEncoder(w).Encode(res)
			return
		}

		res, err := store.AddUser(ctx, rdb, newUser)

		if err != nil {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(res)
			return
		}

		w.WriteHeader(201)
		json.NewEncoder(w).Encode(res)

	}
}
