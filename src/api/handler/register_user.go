package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"todo-list/src/api/response"
	"todo-list/src/api/validator"
	"todo-list/src/store"

	"github.com/go-redis/redis/v8"
	"github.com/rs/zerolog/log"
)

type AddUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Picture  string `json:"picture"`
}

func RegisterUser(ctx context.Context, rdb *redis.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := map[string]string{}
		err := json.NewDecoder(r.Body).Decode(&request)

		if err != nil {
			response.BadRequest(w, r, response.Response(response.ErrorFailedToDecode.Error()))

			log.Error().Err(err).Msg(response.ErrorFailedToDecode.Error())
			return
		}

		newUser := AddUserRequest{
			Username: request["username"],
			Password: request["password"],
			Name:     request["name"],
			Email:    request["email"],
			Picture:  request["picture"],
		}

		if validator.ValidateUsername(newUser.Username) != nil {
			response.BadRequest(w, r, response.Response(validator.ValidateUsername(newUser.Username).Error()))
			return
		}

		if validator.ValidateName(newUser.Name) != nil {
			response.BadRequest(w, r, response.Response(validator.ValidateName(newUser.Name).Error()))
			return
		}

		if validator.ValidatePassword(newUser.Password) != nil {
			response.BadRequest(w, r, response.Response(validator.ValidatePassword(newUser.Password).Error()))
			return
		}

		if validator.ValidateEmail(newUser.Email) != nil {
			response.BadRequest(w, r, response.Response(validator.ValidateEmail(newUser.Email).Error()))
			return
		}

		temp := store.User{
			Username: newUser.Username,
			Password: newUser.Password,
			Email:    newUser.Email,
			Picture:  newUser.Picture,
			Name:     newUser.Name,
		}

		err = store.AddUser(ctx, rdb, temp)

		switch err.(type) {
		case nil:
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(201)
			return
		case *response.DataStoreError:
			response.BadRequest(w, r, response.Response(err.Error()))
			log.Error().Err(err).Msg(err.Error())
			return
		default:
			response.ServerError(w, r)
			log.Error().Err(err).Msg(err.Error())
			return
		}
	}
}
