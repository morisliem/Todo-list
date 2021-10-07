package main

import (
	"context"
	"net/http"
	"todo-list/src"
	"todo-list/src/handler"

	"github.com/go-chi/chi"
)

var (
	ctx    = context.Background()
	rdb    = src.EstablishedConnection()
	router = chi.NewRouter()
)

func main() {

	router.Post("/register", handler.RegisterUserHandler(ctx, rdb))
	router.Post("/login", handler.LoginUserHandler(ctx, rdb))
	router.Post("users/{username}/logout", handler.LogoutUserHandler(ctx, rdb))
	router.Get("users/{username}", handler.GetUserHandler(ctx, rdb))
	router.Post("users/{username}/workflow", handler.AddWorkflowHandler(ctx, rdb))
	router.Get("users/{username}/workflows", handler.GetWorkflowHandler(ctx, rdb))

	http.ListenAndServe(":8080", router)

}
