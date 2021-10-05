package main

import (
	"context"
	"net/http"
	"todo-list/src"
	"todo-list/src/handler"

	"github.com/go-chi/chi"
)

var ctx = context.Background()

func main() {
	rdb := src.EstablishedConnection()
	router := chi.NewRouter()

	router.Post("/register", handler.Register_user(ctx, rdb))
	router.Post("/login", handler.Login_user(ctx, rdb))
	router.Post("/{username}/logout", handler.Logout_user(ctx, rdb))
	router.Get("/{username}", handler.Get_user(ctx, rdb))
	router.Post("/{username}/workflow", handler.Add_workflow(ctx, rdb))
	router.Get("/{username}/workflow", handler.Get_workflow(ctx, rdb))

	http.ListenAndServe(":8080", router)

}
