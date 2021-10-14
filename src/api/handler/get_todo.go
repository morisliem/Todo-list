package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
	"todo-list/src/api/response"
	"todo-list/src/store"

	"github.com/go-chi/chi"
	"github.com/go-redis/redis/v8"
	"github.com/rs/zerolog/log"
)

type getTodoResponse struct {
	Id          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Label       string    `json:"label"`
	Deadline    string    `json:"deadline"`
	Severity    string    `json:"severity"`
	Priority    string    `json:"priority"`
	State       string    `json:"state"`
	Created_at  time.Time `json:"created_at"`
	Deleted_at  time.Time `json:"deleted_at"`
}

type listOfTodo struct {
	Todos []getTodoResponse `json:"todos"`
}

func GetTodos(ctx context.Context, rdb *redis.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		username := chi.URLParam(r, response.URLUsername)
		_, err := store.GetUser(ctx, rdb, username)

		switch err.(type) {
		case *response.NotFoundError:
			response.NotFound(w, r, response.Response(err.Error()))
			log.Error().Err(err).Msg(err.Error())
			return
		}

		res, err := store.GetTodos(ctx, rdb, username)
		var todoTemp getTodoResponse
		var tmp []getTodoResponse
		var todoResponse listOfTodo

		for _, v := range res {
			todoTemp.Id = v.Id
			todoTemp.Title = v.Title
			todoTemp.Description = v.Description
			todoTemp.Label = v.Label
			todoTemp.Deadline = v.Deadline
			todoTemp.Severity = v.Severity
			todoTemp.Priority = v.Priority
			todoTemp.State = v.State
			todoTemp.Created_at = v.Created_at
			todoTemp.Deleted_at = v.Deleted_at
			tmp = append(tmp, todoTemp)
		}
		todoResponse.Todos = tmp

		switch err.(type) {
		case nil:
			response.SuccessfullyOk(w, r)
			json.NewEncoder(w).Encode(todoResponse)
			return
		case *response.NotFoundError:
			response.NotFound(w, r, response.Response(err.Error()))
			log.Error().Err(err).Msg(err.Error())
			return
		default:
			response.ServerError(w, r)
			log.Error().Err(err).Msg(err.Error())
			return

		}
	}
}

func GetTodo(ctx context.Context, rdb *redis.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username := chi.URLParam(r, response.URLUsername)
		todoId := chi.URLParam(r, response.URLUTodoId)

		_, err := store.GetUser(ctx, rdb, username)

		switch err.(type) {
		case *response.NotFoundError:
			response.NotFound(w, r, response.Response(err.Error()))
			log.Error().Err(err).Msg(err.Error())
			return
		}

		res, err := store.GetTodo(ctx, rdb, username, todoId)

		getTodoRes := getTodoResponse{
			Id:          res.Id,
			Title:       res.Title,
			Description: res.Description,
			Label:       res.Label,
			Deadline:    res.Deadline,
			Severity:    res.Severity,
			Priority:    res.Priority,
			State:       res.State,
			Created_at:  res.Created_at,
		}

		switch err.(type) {
		case nil:
			response.SuccessfullyOk(w, r)
			json.NewEncoder(w).Encode(getTodoRes)
			return

		case *response.NotFoundError:
			response.NotFound(w, r, response.Response(err.Error()))
			log.Error().Err(err).Msg(err.Error())
			return

		default:
			response.ServerError(w, r)
			log.Error().Err(err).Msg(err.Error())
			return
		}
	}
}
