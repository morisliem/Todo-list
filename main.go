package main

import (
	"context"
	"net/http"
	"todo-list/src"
	"todo-list/src/api/handler"

	"github.com/go-chi/chi"
)

var (
	ctx    = context.Background()
	rdb    = src.EstablishedConnection()
	router = chi.NewRouter()
)

func main() {

	router.Post("/register", handler.RegisterUser(ctx, rdb))
	router.Post("/login", handler.LoginUser(ctx, rdb))
	router.Post("/users/{username}/logout", handler.LogoutUser(ctx, rdb))
	router.Post("/users/{username}/workflow", handler.AddWorkflow(ctx, rdb))
	router.Post("/users/{username}/todos", handler.AddTodo(ctx, rdb))
	router.Post("/users/{username}", handler.AddPicture(ctx, rdb))
	router.Put("/users/{username}/todos/{todoId}", handler.UpdateTodo(ctx, rdb))
	router.Patch("/users/{username}/todos/{todoId}", handler.UpdateTodoState(ctx, rdb))
	router.Get("/users/{username}/todos", handler.GetTodos(ctx, rdb))
	router.Get("/users/{username}/todos/{todoId}", handler.GetTodo(ctx, rdb))
	router.Get("/users/{username}", handler.GetUser(ctx, rdb))
	router.Get("/users/{username}/workflows", handler.GetWorkflow(ctx, rdb))
	router.Delete("/users/{username}/todos/{todoId}", handler.DeleteTodo(ctx, rdb))

	http.ListenAndServe(":8080", router)

}
