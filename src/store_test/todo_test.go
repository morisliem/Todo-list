package store_test

import (
	"fmt"
	"testing"
	"time"
	"todo-list/src/store"

	"github.com/alicebob/miniredis/v2"
	uuid "github.com/satori/go.uuid"
)

func TestTodoStore(t *testing.T) {
	redis, err := miniredis.Run()
	if err != nil {
		fmt.Println(err)
	}
	defer redis.Close()

	todoId := uuid.NewV4().String()

	newTodo := store.Todo{
		Id:          todoId,
		Title:       "Make unit test",
		Description: "Making unit test for todo list app",
		Label:       "Exercise",
		Deadline:    "10-10-2021",
		Severity:    "High",
		Priority:    "High",
		State:       "Ongoing",
		Created_at:  time.Now(),
	}

	key := "moris:todo:" + todoId

	redis.HSet(key,
		"id", newTodo.Id,
		"title", newTodo.Title,
		"description", newTodo.Description,
		"label", newTodo.Label,
		"deadline", newTodo.Deadline,
		"severity", newTodo.Severity,
		"priority", newTodo.Priority,
		"state", newTodo.State,
		"created_at", fmt.Sprintf("%v", newTodo.Created_at))

	res := redis.HGet(key, "id")
	if res != newTodo.Id {
		t.Errorf("Failed to get the correct value of id")
	}

	res = redis.HGet(key, "title")
	if res != newTodo.Title {
		t.Errorf("Failed to get the correct value of title")
	}

	res = redis.HGet(key, "description")
	if res != newTodo.Description {
		t.Errorf("Failed to get the correct value of description")
	}

	res = redis.HGet(key, "label")
	if res != newTodo.Label {
		t.Errorf("Failed to get the correct value of label")
	}

	res = redis.HGet(key, "deadline")
	if res != newTodo.Deadline {
		t.Errorf("Failed to get the correct value of deadline")
	}

	res = redis.HGet(key, "severity")
	if res != newTodo.Severity {
		t.Errorf("Failed to get the correct value of severity")
	}

	res = redis.HGet(key, "priority")
	if res != newTodo.Priority {
		t.Errorf("Failed to get the correct value of priority")
	}

	res = redis.HGet(key, "state")
	if res != newTodo.State {
		t.Errorf("Failed to get the correct value of state")
	}

	res = redis.HGet(key, "created_at")
	if res != fmt.Sprintf("%v", newTodo.Created_at) {
		t.Errorf("Failed to get the correct value of created_at")
	}

}
