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

		err, fileName := SavePicture(w, r, username)

		switch err.(type) {
		case *response.DataStoreError:
			response.BadRequest(w, r, response.Response(err.Error()))
			log.Error().Err(err).Msg(err.Error())
			return
		}

		err = store.AddUserPicture(ctx, rdb, username, fileName)

		switch err.(type) {
		case nil:
			response.SuccessfullyCreated(w, r)
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

func SavePicture(w http.ResponseWriter, r *http.Request, usr string) (error, string) {
	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		response.BadRequest(w, r, response.Response(response.ErrorToParsePicture.Error()))
		log.Error().Err(err).Msg(response.ErrorToParsePicture.Error())
		return &response.DataStoreError{Message: response.ErrorToParsePicture.Error()}, ""
	}

	file, handler, err := r.FormFile("picture")
	if err != nil {
		response.BadRequest(w, r, response.Response(response.ErrorToRetrieveFile.Error()))
		log.Error().Err(err).Msg(response.ErrorToRetrieveFile.Error())

		return &response.DataStoreError{Message: response.ErrorToRetrieveFile.Error()}, ""
	}
	defer file.Close()

	filename := usr + fmt.Sprintf("-%v", handler.Filename)
	fmt.Println(filename)
	tempFile, err := ioutil.TempFile("./img/user_picture", "upload-*.png")

	if err != nil {
		response.BadRequest(w, r, response.Response(response.ErrorToCreateTempFile.Error()))
		log.Error().Err(err).Msg(response.ErrorToCreateTempFile.Error())

		return &response.DataStoreError{Message: response.ErrorToCreateTempFile.Error()}, ""
	}
	defer tempFile.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		response.BadRequest(w, r, response.Response(response.ErrorToSaveFile.Error()))
		log.Error().Err(err).Msg(response.ErrorToSaveFile.Error())
		return &response.DataStoreError{Message: response.ErrorToSaveFile.Error()}, ""
	}
	tempFile.Write(fileBytes)

	return nil, handler.Filename
}
