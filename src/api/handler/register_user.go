package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
	"todo-list/src/api/validator"
	"todo-list/src/store"

	"github.com/go-redis/redis/v8"
)

func RegisterUser(ctx context.Context, rdb *redis.Client) http.HandlerFunc {
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

		newUser := store.User{
			Username:   request["username"],
			Password:   request["password"],
			Name:       request["name"],
			Email:      request["email"],
			Picture:    request["picture"],
			Created_at: time.Now(),
		}

		if validator.ValidateUsername(newUser.Username) != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(400)
			res := validator.Response(validator.ValidateUsername(newUser.Username).Error())
			json.NewEncoder(w).Encode(res)
			return
		}

		if validator.ValidateName(newUser.Name) != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(400)
			res := validator.Response(validator.ValidateName(newUser.Name).Error())
			json.NewEncoder(w).Encode(res)
			return
		}

		if validator.ValidatePassword(newUser.Password) != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(400)
			res := validator.Response(validator.ValidatePassword(newUser.Password).Error())
			json.NewEncoder(w).Encode(res)
			return
		}

		if validator.ValidateEmail(newUser.Email) != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(400)
			res := validator.Response(validator.ValidateEmail(newUser.Email).Error())
			json.NewEncoder(w).Encode(res)
			return
		}

		res, err := store.AddUser(ctx, rdb, newUser)

		if err != nil {
			if err.Error() == validator.FailedToAddUser {
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
