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
		var request AddUserRequest
		err := json.NewDecoder(r.Body).Decode(&request)

		if err != nil {
			response.BadRequest(w, r, response.Response(response.ErrorFailedToDecode.Error()))
			log.Error().Err(err).Msg(response.ErrorFailedToDecode.Error())
			return
		}

		request.validateRequest(w, r)

		temp := store.User{
			Username: request.Username,
			Password: request.Password,
			Email:    request.Email,
			Picture:  request.Picture,
			Name:     request.Name,
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

func (req *AddUserRequest) validateRequest(w http.ResponseWriter, r *http.Request) {
	if validator.ValidateUsername(req.Username) != nil {
		response.BadRequest(w, r, response.Response(validator.ValidateUsername(req.Username).Error()))
		return
	}

	if validator.ValidateName(req.Name) != nil {
		response.BadRequest(w, r, response.Response(validator.ValidateName(req.Name).Error()))
		return
	}

	if validator.ValidatePassword(req.Password) != nil {
		response.BadRequest(w, r, response.Response(validator.ValidatePassword(req.Password).Error()))
		return
	}

	if validator.ValidateEmail(req.Email) != nil {
		response.BadRequest(w, r, response.Response(validator.ValidateEmail(req.Email).Error()))
		return
	}
}
