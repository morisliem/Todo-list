package store

import (
	"context"
	"todo-list/src/api/response"

	"github.com/go-redis/redis/v8"
)

type workFlowDetail []string

func AddWorkflow(ctx context.Context, db *redis.Client, username string, workflow string) (map[string]string, error) {
	checkUsername, err := db.HGetAll(ctx, username).Result()

	if len(checkUsername) == 0 {
		return map[string]string{}, &response.BadInputError{Message: response.ErrorUserNotFound.Error()}
	}

	if err != nil {
		return map[string]string{}, err
	}

	key := username + ":workflow"

	err = db.SAdd(ctx, key, workflow).Err()

	if err != nil {
		return map[string]string{}, response.ErrorFailedToAddWorkflow
	}

	res := response.Response(response.SuccessfullyAdded)
	return res, nil

}

func GetWorkflow(ctx context.Context, db *redis.Client, username string) (map[string]workFlowDetail, error) {
	result := map[string]workFlowDetail{}
	key := username + ":workflow"

	workflows, err := db.SMembers(ctx, key).Result()

	if err != nil {
		return map[string]workFlowDetail{}, err
	}

	if len(workflows) == 0 {
		return map[string]workFlowDetail{}, response.ErrorWorkflowNotFound
	}
	temp := []string{}
	temp = append(temp, workflows...)
	result["Workflows"] = temp
	return result, nil

}
