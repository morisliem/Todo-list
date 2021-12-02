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

		err = request.validateRequest(w, r)
		if err != nil {
			response.BadRequest(w, r, response.Response(err.Error()))
			log.Error().Err(err).Msg(err.Error())
			return
		}

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
			response.SuccessfullyOk(w, r)
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

func (req *AddUserRequest) validateRequest(w http.ResponseWriter, r *http.Request) error {
	if err := validator.ValidateUsername(req.Username); err != nil {
		return err
	}

	if err := validator.ValidateName(req.Name); err != nil {
		return err
	}

	if err := validator.ValidatePassword(req.Password); err != nil {
		return err
	}

	if err := validator.ValidateEmail(req.Email); err != nil {
		return err
	}

	return nil
}
