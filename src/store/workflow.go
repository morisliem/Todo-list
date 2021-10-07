package store

import (
	"context"
	"fmt"
	"todo-list/src"

	"github.com/go-redis/redis/v8"
)

func AddWorkflow(ctx context.Context, db *redis.Client, username string, workflow string) (map[string]string, error) {
	checkUsername, _ := db.HGetAll(ctx, username).Result()

	if len(checkUsername) == 0 {
		temp := src.FailedToAddWorkflow + "," + src.UserNotFoundError().Error()
		res := src.Response(temp)
		return res, nil
	}

	key := username + ":workflow"

	err := db.SAdd(ctx, key, workflow).Err()

	if err != nil {
		res := src.Response(src.FailedToAddWorkflow)
		return res, err
	}

	res := src.Response(src.SuccessfullyAdded)
	return res, nil

}

func GetWorkflow(ctx context.Context, db *redis.Client, username string) (map[string]string, error) {
	result := map[string]string{}
	key := username + ":workflow"

	workflows, _ := db.SMembers(ctx, key).Result()

	if len(workflows) == 0 {
		res := src.Response(src.WorkflowNotFoundError().Error())
		return res, nil
	}

	result["Workflows"] = fmt.Sprint(workflows)
	return result, nil

}
