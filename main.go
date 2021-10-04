package main

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
	"todo-list/src"

	"github.com/go-chi/chi"
)

var ctx = context.Background()

func main() {
	rdb := src.EstablishedConnection()
	router := chi.NewRouter()

	router.Post("/register", func(w http.ResponseWriter, r *http.Request) {
		request := map[string]string{}
		json.NewDecoder(r.Body).Decode(&request)

		newUser := src.User{
			Username:   request["username"],
			Password:   request["password"],
			Name:       request["name"],
			Email:      request["email"],
			Picture:    request["picture"],
			Created_at: time.Now(),
		}

		res, err := src.AddUser(ctx, rdb, newUser)

		if err != nil {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(res)
		}
		w.WriteHeader(201)
		json.NewEncoder(w).Encode(res)
	})

	router.Post("/login", func(w http.ResponseWriter, r *http.Request) {
		request := map[string]string{}
		json.NewDecoder(r.Body).Decode(&request)

		login := src.User{
			Username: request["username"],
			Password: request["password"],
		}

		res, err := src.LoginUser(ctx, rdb, login)

		if err != nil {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(res)
		}

		w.WriteHeader(200)
		json.NewEncoder(w).Encode(res)

	})

	router.Post("/{username}/logout", func(w http.ResponseWriter, r *http.Request) {
		username := chi.URLParam(r, "username")
		res, err := src.LogoutUser(ctx, rdb, username)

		if err != nil {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(res)
		}
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(res)
	})

	router.Get("/{username}", func(w http.ResponseWriter, r *http.Request) {
		username := chi.URLParam(r, "username")

		res, err := src.GetUser(ctx, rdb, username)

		if err != nil {
			w.WriteHeader(404)
			json.NewEncoder(w).Encode(res["Message"])
		}
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(res)
	})

	http.ListenAndServe(":8080", router)

}
