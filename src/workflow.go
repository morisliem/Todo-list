package src

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

func AddWorkflow(ctx context.Context, db *redis.Client, username string, workflow string) (map[string]string, error) {
	result := map[string]string{}
	checkUsername, _ := db.HGetAll(ctx, username).Result()

	if len(checkUsername) == 0 {
		result["Message"] = "Cannot add workflow, user not found"
		return result, nil
	}

	key := username + ":workflow"

	err := db.SAdd(ctx, key, workflow).Err()

	if err != nil {
		result["Message"] = "Failed to add workflow"
		return result, err
	}

	result["Message"] = "Successfully added"
	return result, nil

}

func GetWorkflow(ctx context.Context, db *redis.Client, username string) (map[string]string, error) {
	result := map[string]string{}
	key := username + ":workflow"

	workflows, _ := db.SMembers(ctx, key).Result()

	if len(workflows) == 0 {
		result["Message"] = "Workflow not found"
		return result, nil
	}

	result["Workflows"] = fmt.Sprint(workflows)
	return result, nil

}
