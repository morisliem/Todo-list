package store

import (
	"context"
	"todo-list/src/api/response"

	"github.com/go-redis/redis/v8"
)

type Workflow struct {
	Workflows []string `json:"workflow"`
}

type workFlowDetail []string

func AddWorkflow(ctx context.Context, db *redis.Client, username string, workflow string) error {
	checkUsername, err := db.HGetAll(ctx, username).Result()

	if len(checkUsername) == 0 {
		return &response.BadInputError{Message: response.ErrorUserNotFound.Error()}
	}

	if err != nil {
		return err
	}

	key := username + ":workflow"
	err = db.SAdd(ctx, key, workflow).Err()

	if err != nil {
		return &response.DataStoreError{Message: response.ErrorFailedToAddWorkflow.Error()}
	}

	return nil

}

func GetWorkflow(ctx context.Context, db *redis.Client, username string) (Workflow, error) {
	workflow := Workflow{}
	key := username + ":workflow"

	workflows, err := db.SMembers(ctx, key).Result()

	if err != nil {
		return workflow, err
	}

	if len(workflows) == 0 {
		return workflow, nil
	}

	workflow.Workflows = append(workflow.Workflows, workflows...)

	return workflow, nil

}
