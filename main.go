package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"todo-list/src"

	"github.com/go-chi/chi"
)

func main() {
	usr := src.NewUser()

	user1 := src.User{Username: "moris", Password: "hello123", ID: "100"}

	usr.AddUser(user1)

	fmt.Println("Starting server...")

	router := chi.NewRouter()

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(usr)
	})

	router.Get("/{username}/hciwhdco", func(w http.ResponseWriter, r *http.Request) {
		test := chi.URLParam(r, "username")
		fmt.Println(test)
		fmt.Println(r)
		json.NewEncoder(w).Encode(usr)
	})

	router.Post("/login", func(w http.ResponseWriter, r *http.Request) {
		request := map[string]string{}
		loginStatus := false
		var specificUser src.User

		json.NewDecoder(r.Body).Decode(&request)

		username := request["username"]
		password := request["password"]

		for i := range usr.Users {
			request_username := usr.Users[i].Username
			request_password := usr.Users[i].Password

			if request_username == username && request_password == password {
				specificUser = usr.Users[i]
				loginStatus = true
				break
			}
		}

		if loginStatus {
			usr.AddUser(src.User{
				ID:         specificUser.ID,
				Username:   specificUser.Username,
				Password:   specificUser.Password,
				Name:       specificUser.Name,
				Email:      specificUser.Email,
				Picture:    specificUser.Picture,
				IsLoggedIn: loginStatus,
			})

			w.Write([]byte("Logged in successfully"))
		} else {
			w.Write([]byte("Invalid username or password"))
		}
	})

	//

	router.Post("/post", func(w http.ResponseWriter, r *http.Request) {
		request := map[string]string{}

		json.NewDecoder(r.Body).Decode(&request)

		usr.AddUser(src.User{
			ID:       request["ID"],
			Username: request["username"],
			Password: request["password"],
			Name:     request["name"],
			Email:    request["email"],
			Picture:  request["picture"],
		})

		w.Write([]byte("Successfully added"))
	})

	log.Fatal(http.ListenAndServe(":8080", router))

}
