package handler

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"todo-list/src/api/response"
	"todo-list/src/store"

	"github.com/go-chi/chi"
	"github.com/go-redis/redis/v8"
	"github.com/rs/zerolog/log"
)

func AddPicture(ctx context.Context, rdb *redis.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username := chi.URLParam(r, response.URLUsername)
		_, err := store.GetUser(ctx, rdb, username)

		switch err.(type) {
		case *response.NotFoundError:
			response.NotFound(w, r, response.Response(err.Error()))
			log.Error().Err(err).Msg(err.Error())
			return
		}

		fileName, err := SavePicture(w, r, username)

		switch err.(type) {
		case *response.DataStoreError:
			response.BadRequest(w, r, response.Response(err.Error()))
			log.Error().Err(err).Msg(err.Error())
			return
		}

		err = store.AddUserPicture(ctx, rdb, username, fileName)

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

func SavePicture(w http.ResponseWriter, r *http.Request, usr string) (string, error) {
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		response.BadRequest(w, r, response.Response(response.ErrorToParsePicture.Error()))
		log.Error().Err(err).Msg(response.ErrorToParsePicture.Error())
		return "", &response.DataStoreError{Message: response.ErrorToParsePicture.Error()}
	}

	file, handler, err := r.FormFile("picture")
	if err != nil {
		response.BadRequest(w, r, response.Response(response.ErrorToRetrieveFile.Error()))
		log.Error().Err(err).Msg(response.ErrorToRetrieveFile.Error())

		return "", &response.DataStoreError{Message: response.ErrorToRetrieveFile.Error()}
	}
	defer file.Close()

	filename := usr + fmt.Sprintf("-%v", handler.Filename)
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		response.BadRequest(w, r, response.Response(response.ErrorToSaveFile.Error()))
		log.Error().Err(err).Msg(response.ErrorToSaveFile.Error())
		return "", &response.DataStoreError{Message: response.ErrorToSaveFile.Error()}
	}

	err = ioutil.WriteFile("./img/user_picture/"+filename, fileBytes, 0777)
	if err != nil {
		response.BadRequest(w, r, response.Response(response.ErrorToSaveFile.Error()))
		log.Error().Err(err).Msg(response.ErrorToSaveFile.Error())
		return "", &response.DataStoreError{Message: response.ErrorToSaveFile.Error()}
	}

	return filename, nil
}
