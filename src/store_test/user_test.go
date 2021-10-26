package store_test

import (
	"fmt"
	"testing"
	"time"
	"todo-list/src/store"

	"github.com/alicebob/miniredis/v2"
)

func TestUserStore(t *testing.T) {
	redis, err := miniredis.Run()
	if err != nil {
		fmt.Println(err)
	}
	defer redis.Close()

	newUser := store.User{
		Username:   "moris",
		Password:   "moris123",
		Name:       "moris",
		Email:      "moris",
		Picture:    "moris.png",
		Created_at: time.Now(),
	}

	redis.HSet("moris",
		"username", newUser.Name,
		"password", newUser.Password,
		"name", newUser.Name,
		"email", newUser.Email,
		"picture", newUser.Picture,
		"created_at", fmt.Sprintf("%v", newUser.Created_at))

	res := redis.HGet("moris", "username")
	if res != newUser.Username {
		t.Errorf("Failed to get the correct value of username")
	}

	res = redis.HGet("moris", "password")
	if res != newUser.Password {
		t.Errorf("Failed to get the correct value of password")
	}

	res = redis.HGet("moris", "name")
	if res != newUser.Name {
		t.Errorf("Failed to get the correct value of name")
	}

	res = redis.HGet("moris", "email")
	if res != newUser.Email {
		t.Errorf("Failed to get the correct value of email")
	}

	res = redis.HGet("moris", "picture")
	if res != newUser.Picture {
		t.Errorf("Failed to get the correct value of picture")
	}

	res = redis.HGet("moris", "created_at")
	if res != fmt.Sprintf("%v", newUser.Created_at) {
		t.Errorf("Failed to get the correct value of created_at")
	}

}
