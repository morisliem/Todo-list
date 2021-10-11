package store

import (
	"context"
	"fmt"
	"strings"
	"time"
	"todo-list/src/api/response"

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
	Update_at   time.Time
}

type AddTodoResponse struct {
}

const (
	HmapKeyTodoId          = "id"
	HmapKeyTodoTitle       = "title"
	HmapKeyTodoDescription = "description"
	HmapKeyTodoLabel       = "label"
	HmapKeyTodoDeadline    = "deadline"
	HmapKeyTodoSeverity    = "severity"
	HmapKeyTodoPriority    = "priority"
	HmapKeyTodoState       = "state"
	HmapKeyTodoCreatedAt   = "created_at"
	HmapKeyTodoDeletedAt   = "deleted_at"
	HmapKeyTodoUpdatedAt   = "updated_at"
)

func AddTodo(ctx context.Context, db *redis.Client, usr string, td Todo) (map[string]string, error) {
	todoId := uuid.NewV4().String()
	HmapKey := usr + ":todo:" + string(todoId)
	var todos string

	workflow, err := GetWorkflow(ctx, db, usr)
	if err != nil {
		return map[string]string{}, err
	}
	if len(workflow) == 0 {
		return map[string]string{}, &response.DataStoreError{Message: response.ErrorEmptyWorkflow.Error()}
	}

	isWorklowExist := false
	for _, v := range workflow["Workflows"] {
		if strings.EqualFold(v, td.State) {
			isWorklowExist = true
			break
		}
	}

	if !isWorklowExist {
		return map[string]string{}, &response.BadInputError{Message: response.ErrorWorkflowNotExist.Error()}
	}

	err = db.HMSet(ctx, HmapKey,
		HmapKeyTodoTitle, td.Title,
		HmapKeyTodoDescription, td.Description,
		HmapKeyTodoLabel, td.Label,
		HmapKeyTodoDeadline, td.Deadline,
		HmapKeyTodoSeverity, td.Severity,
		HmapKeyTodoPriority, td.Priority,
		HmapKeyTodoState, td.State,
		HmapKeyTodoCreatedAt, td.Created_at,
		HmapKeyTodoDeletedAt, td.Deleted_at).Err()

	if err != nil {
		return map[string]string{}, &response.DataStoreError{Message: response.ErrorFailedToAddTodo.Error()}
	}

	todoListFromUserHash, err := db.HMGet(ctx, usr, HmapKeyUserTodos).Result()

	if err != nil {
		return map[string]string{}, &response.DataStoreError{Message: response.ErrorFailedToUpdateUserTodo.Error()}
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
		return map[string]string{}, &response.DataStoreError{Message: response.ErrorFailedToUpdateUserTodo.Error()}
	}

	res := response.Response(response.SuccessfullyAdded)
	return res, nil
}

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
			tmp := map[string]string{}
			key := usr + ":todo:" + fmt.Sprintf("%v", v)
			index := "Todo " + fmt.Sprintf("%v", count)

			td, err := db.HGetAll(ctx, key).Result()

			if err != nil {
				return map[string]todoDetail{}, response.ErrorTodoNotFound
			}

			for key, val := range td {
				tmp[key] = val
			}
			result[index] = tmp
		}
	}

	return result, nil
}

func GetTodo(ctx context.Context, db *redis.Client, usr string, todoId string) (map[string]string, error) {
	key := usr + ":todo:" + fmt.Sprintf("%v", todoId)

	todo, err := db.HGetAll(ctx, key).Result()

	if len(todo) == 0 {
		return map[string]string{}, response.ErrorTodoNotFound
	}

	if err != nil {
		return map[string]string{}, err
	}

	return todo, nil
}

func UpdateTodo(ctx context.Context, db *redis.Client, usr string, todoId string, td Todo) (map[string]string, error) {
	key := usr + ":todo:" + fmt.Sprintf("%v", todoId)
	dataToUpdate := map[string]string{}

	todo, err := db.HGetAll(ctx, key).Result()

	if len(todo) == 0 {
		return map[string]string{}, response.ErrorTodoNotFound
	}

	if err != nil {
		return map[string]string{}, err
	}

	if len(strings.TrimSpace(td.Title)) != 0 {
		dataToUpdate[HmapKeyTodoTitle] = td.Title
	}

	if len(strings.TrimSpace(td.Description)) != 0 {
		dataToUpdate[HmapKeyTodoDescription] = td.Description
	}

	if len(strings.TrimSpace(td.Label)) != 0 {
		dataToUpdate[HmapKeyTodoLabel] = td.Label
	}

	if len(strings.TrimSpace(td.Deadline)) != 0 {
		dataToUpdate[HmapKeyTodoDeadline] = td.Deadline
	}

	if len(strings.TrimSpace(td.Priority)) != 0 {
		dataToUpdate[HmapKeyTodoPriority] = td.Priority
	}

	if len(strings.TrimSpace(td.Severity)) != 0 {
		dataToUpdate[HmapKeyTodoSeverity] = td.Severity
	}

	dataToUpdate[HmapKeyTodoUpdatedAt] = fmt.Sprintf("%v", time.Now())

	for i, k := range dataToUpdate {
		err := db.HMSet(ctx, key, i, k).Err()

		if err != nil {
			return map[string]string{}, err
		}
	}

	return response.Response(response.SuccessfullyUpdated), nil
}
