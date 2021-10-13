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

type loginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func LoginUser(ctx context.Context, rdb *redis.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request loginRequest
		err := json.NewDecoder(r.Body).Decode(&request)

		if err != nil {
			response.BadRequest(w, r, response.Response(response.ErrorFailedToDecode.Error()))
			log.Error().Err(err).Msg(response.ErrorFailedToDecode.Error())
			return
		}

		login := store.LoginRequest{
			Username: request.Username,
			Password: request.Password,
		}

		if validator.ValidateUsername(login.Username) != nil {
			response.BadRequest(w, r, response.Response(validator.ValidateUsername(login.Username).Error()))
			log.Error().Err(validator.ValidateUsername(login.Username)).Msg(validator.ValidateUsername(login.Username).Error())
			return
		}

		err = store.LoginUser(ctx, rdb, login)

		switch err.(type) {
		case nil:
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(200)
			return
		case *response.BadInputError:
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
