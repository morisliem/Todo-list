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

type Todo struct {
	Id          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Label       string    `json:"label"`
	Deadline    string    `json:"deadline"`
	Severity    string    `json:"severity"`
	Priority    string    `json:"priority"`
	State       string    `json:"state"`
	Created_at  time.Time `json:"created_at"`
	Deleted_at  time.Time `json:"deleted_at"`
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
)

func AddTodo(ctx context.Context, db *redis.Client, usr string, td Todo) error {
	todoId := uuid.NewV4().String()
	HmapKey := usr + ":todo:" + string(todoId)

	workflow, err := GetWorkflow(ctx, db, usr)

	if len(workflow.Workflows) == 0 {
		return &response.BadInputError{Message: response.ErrorEmptyWorkflow.Error()}
	}
	if err != nil {
		return err
	}

	isWorklowExist := isWorkflowExist(workflow.Workflows, td.State)
	if !isWorklowExist {
		return &response.BadInputError{Message: response.ErrorWorkflowNotExist.Error()}
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
		return &response.DataStoreError{Message: response.ErrorFailedToAddTodo.Error()}
	}

	todoListFromUserHash, err := db.HMGet(ctx, usr, HmapKeyUserTodos).Result()

	if err != nil {
		return err
	}

	todos := updateTodoListInUserHash(todoListFromUserHash, todoId)

	splitTodos := strings.Split(todos, " ")
	var updatedTodos string

	for i, v := range splitTodos {
		if i%2 == 0 {
			updatedTodos += v
		} else {
			updatedTodos += " "
		}
	}

	err = db.HMSet(ctx, usr, HmapKeyUserTodos, updatedTodos).Err()

	if err != nil {
		return &response.DataStoreError{Message: response.ErrorFailedToUpdateUserTodo.Error()}
	}

	return nil
}

func GetTodos(ctx context.Context, db *redis.Client, usr string) ([]Todo, error) {
	listOfTodosTitle, err := db.HGet(ctx, usr, HmapKeyUserTodos).Result()
	result := []Todo{}

	if err != nil {
		return result, err
	}

	if len(listOfTodosTitle) == 0 {
		return result, &response.NotFoundError{Message: response.ErrorTodoNotFound.Error()}
	}

	todo := strings.Split(listOfTodosTitle, " ")
	for _, v := range todo {
		if v != "" {
			key := usr + ":todo:" + fmt.Sprintf("%v", v)
			td, err := db.HGetAll(ctx, key).Result()
			tempTodo := Todo{}
			if err != nil {
				return result, &response.NotFoundError{Message: response.ErrorTodoNotFound.Error()}
			}

			for key, val := range td {
				if key == HmapKeyTodoCreatedAt {
					tmp, err := time.Parse(time.RFC3339, val)
					if err != nil {
						return result, err
					}
					tempTodo.Created_at = tmp
				}
				if key == HmapKeyTodoTitle {
					tempTodo.Title = val
				}
				if key == HmapKeyTodoDescription {
					tempTodo.Description = val
				}
				if key == HmapKeyTodoLabel {
					tempTodo.Label = val
				}
				if key == HmapKeyTodoDeadline {
					tempTodo.Deadline = val
				}
				if key == HmapKeyTodoSeverity {
					tempTodo.Severity = val
				}
				if key == HmapKeyTodoPriority {
					tempTodo.Priority = val
				}
				if key == HmapKeyTodoState {
					tempTodo.State = val
				}
			}
			tempTodo.Id = v
			result = append(result, tempTodo)
		}
	}

	return result, nil

}

func GetTodo(ctx context.Context, db *redis.Client, usr string, todoId string) (Todo, error) {
	key := usr + ":todo:" + fmt.Sprintf("%v", todoId)
	result := Todo{}

	todo, err := db.HGetAll(ctx, key).Result()

	if len(todo) == 0 {
		return result, &response.NotFoundError{Message: response.ErrorTodoNotFound.Error()}
	}

	if err != nil {
		return result, err
	}

	for _, v := range todo {
		if v == HmapKeyTodoTitle {
			result.Title = v
		}
		if v == HmapKeyTodoDescription {
			result.Description = v
		}
		if v == HmapKeyTodoLabel {
			result.Label = v
		}
		if v == HmapKeyTodoDeadline {
			result.Deadline = v
		}
		if v == HmapKeyTodoSeverity {
			result.Severity = v
		}
		if v == HmapKeyTodoPriority {
			result.Priority = v
		}
		if v == HmapKeyTodoState {
			result.State = v
		}
		if v == HmapKeyTodoCreatedAt {
			tmp, err := time.Parse(time.RFC3339, v)
			if err != nil {
				return result, err
			}
			result.Created_at = tmp
		}
	}
	result.Id = todoId

	return result, nil
}

func UpdateTodo(ctx context.Context, db *redis.Client, usr string, todoId string, td Todo) error {
	key := usr + ":todo:" + fmt.Sprintf("%v", todoId)
	tmp := map[string]string{}

	todo, err := db.HGetAll(ctx, key).Result()

	if len(todo) == 0 {
		return &response.NotFoundError{Message: response.ErrorTodoNotFound.Error()}
	}

	if err != nil {
		return err
	}

	dataToUpdate := dataToChange(tmp, td)

	for i, k := range dataToUpdate {
		err := db.HMSet(ctx, key, i, k).Err()

		if err != nil {
			return &response.DataStoreError{Message: err.Error()}
		}
	}

	return nil
}

func RemoveTodo(ctx context.Context, db *redis.Client, usr string, todoId string) error {
	key := usr + ":todo:" + todoId

	isUserExist, err := db.HMGet(ctx, usr, HmapKeyUserName).Result()

	if isUserExist[0] == nil {
		return &response.NotFoundError{Message: response.ErrorUserNotFound.Error()}
	}

	if err != nil {
		return err
	}

	isTodoExist, err := db.HMGet(ctx, key, HmapKeyTodoTitle).Result()

	if isTodoExist[0] == nil {
		return &response.NotFoundError{Message: response.ErrorTodoNotFound.Error()}
	}

	if err != nil {
		return err
	}

	err = db.Del(ctx, key).Err()

	if err != nil {
		return err
	}

	todoFromUserHash, err := db.HGet(ctx, usr, HmapKeyUserTodos).Result()
	if len(todoFromUserHash) == 0 {
		return &response.DataStoreError{Message: response.ErrorFailedToUpdateUserTodo.Error()}
	}

	if err != nil {
		return err
	}

	tmp := strings.Split(todoFromUserHash, " ")
	var todos string

	if len(tmp) == 2 && tmp[0] == todoId {
		err = db.HDel(ctx, usr, HmapKeyUserTodos).Err()
		if err != nil {
			return err
		}
	} else {
		for i, v := range tmp {
			if i%2 == 0 {
				if v == todoId {
					todos += ""
				} else {
					todos += fmt.Sprintf("%v ", v)
				}
			}
		}
		err = db.HMSet(ctx, usr, HmapKeyUserTodos, todos).Err()

		if err != nil {
			return &response.DataStoreError{Message: response.ErrorFailedToUpdateUserTodo.Error()}
		}
	}

	return nil
}

func UpdateTodoState(ctx context.Context, db *redis.Client, usr string, todoId string, newState string) error {
	key := usr + ":todo:" + todoId

	todo, err := db.HGetAll(ctx, key).Result()

	if len(todo) == 0 {
		return &response.NotFoundError{Message: response.ErrorTodoNotFound.Error()}
	}
	if err != nil {
		return err
	}

	workflow, err := GetWorkflow(ctx, db, usr)
	if err != nil {
		return err
	}
	if len(workflow.Workflows) == 0 {
		return &response.DataStoreError{Message: response.ErrorEmptyWorkflow.Error()}
	}

	isWorklowExist := false
	for _, v := range workflow.Workflows {
		if strings.EqualFold(v, newState) {
			isWorklowExist = true
			break
		}
	}

	if !isWorklowExist {
		return &response.NotFoundError{Message: response.ErrorWorkflowNotExist.Error()}
	}

	err = db.HMSet(ctx, key, HmapKeyTodoState, newState).Err()

	if err != nil {
		return err
	}

	return nil
}

func isWorkflowExist(workFlowList workFlowDetail, workFlow string) bool {
	isExist := false
	for _, v := range workFlowList {
		if strings.EqualFold(v, workFlow) {
			isExist = true
			break
		}
	}
	return isExist
}

func updateTodoListInUserHash(todoList []interface{}, todoId string) string {
	var todos string

	if todoList[0] == nil {
		todoList = todoList[1:]
	}

	todoList = append(todoList, todoId)

	for _, v := range todoList {
		todos += fmt.Sprintf("%v ", v)
	}

	return todos
}

func dataToChange(dataToUpdate map[string]string, td Todo) map[string]string {
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

	return dataToUpdate
}
