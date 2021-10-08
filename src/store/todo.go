package store

import (
	"context"
	"fmt"
	"time"
	"todo-list/src/api/validator"

	"github.com/go-redis/redis/v8"
	uuid "github.com/satori/go.uuid"
)

type Todo struct {
	Id          string
	Title       string `json:"title"`
	Description string `json:"description"`
	Label       string `json:"label"`
	Deadline    string `json:"deadline"`
	Severity    string `json:"severity"`
	Priority    string `json:"priority"`
	State       string `json:"state"`
	Created_at  time.Time
	Deleted_at  time.Time
}

const (
	HmapKeyTodoId          = "id"
	HmapKeyTodoTitle       = "title"
	HmapKeyTodoDescription = "description"
	HmapKeyTodoLabel       = "label"
	HmapKeyTodoDeadline    = "deadline"
	HmapKeyTodoSevertiy    = "severity"
	HmapKeyTodoPriority    = "priority"
	HmapKeyTodoState       = "state"
	HmapKeyTodoCreatedAt   = "created_at"
	HmapKeyTodoDeletedAt   = "deleted_at"
)

func AddTodo(ctx context.Context, db *redis.Client, usr string, td Todo) (map[string]string, error) {
	todoId := uuid.NewV4().String()
	HmapKey := usr + ":todo:" + string(todoId)
	var todos string

	err := db.HMSet(ctx, HmapKey,
		HmapKeyTodoTitle, td.Title,
		HmapKeyTodoDescription, td.Description,
		HmapKeyTodoLabel, td.Label,
		HmapKeyTodoDeadline, td.Deadline,
		HmapKeyTodoSevertiy, td.Severity,
		HmapKeyTodoPriority, td.Priority,
		HmapKeyTodoState, td.State,
		HmapKeyTodoCreatedAt, td.Created_at,
		HmapKeyTodoDeletedAt, td.Deleted_at).Err()

	if err != nil {
		res := validator.Response(validator.FailedToAddTodo)
		return res, err
	}

	todoListFromUserHash, err := db.HMGet(ctx, usr, HmapKeyUserTodos).Result()

	if err != nil {
		res := validator.Response(validator.FailedToUpdateUserTodo)
		return res, err
	}

	if todoListFromUserHash[0] == nil {
		todoListFromUserHash = todoListFromUserHash[1:]
	}

	todoListFromUserHash = append(todoListFromUserHash, todoId)

	for _, v := range todoListFromUserHash {
		todos += fmt.Sprintf("%v ", v)
	}

	err = db.HMSet(ctx, usr, HmapKeyUserTodos, todos).Err()

	if err != nil {
		res := validator.Response(validator.FailedToUpdateUserTodo)
		return res, err
	}

	res := validator.Response(validator.SuccessfullyAdded)
	return res, nil
}
