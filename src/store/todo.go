package store

import (
	"context"
	"fmt"
	"strings"
	"time"
	"todo-list/src/api/validator"

	"github.com/go-redis/redis/v8"
	uuid "github.com/satori/go.uuid"
)

type todoDetail map[string]string

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

// Missing validation and have to clean the code
func GetTodos(ctx context.Context, db *redis.Client, usr string) (map[string]todoDetail, error) {
	listOfTodosTitle, err := db.HMGet(ctx, usr, HmapKeyUserTodos).Result()
	result := map[string]todoDetail{}
	count := 0

	if err != nil {
		return map[string]todoDetail{}, err
	}

	todo := strings.Split(fmt.Sprintf("%v", listOfTodosTitle[0]), " ")

	for i, v := range todo {
		if i%2 == 0 {
			count++
			key := usr + ":todo:" + fmt.Sprintf("%v", v)
			td, _ := db.HGetAll(ctx, key).Result()
			index := "Todo " + fmt.Sprintf("%v", count)
			tmp := map[string]string{}
			for key, val := range td {
				tmp[key] = val
			}
			result[index] = tmp
		}
	}

	return result, nil
}
