package store_test

import (
	"fmt"
	"testing"

	"github.com/alicebob/miniredis/v2"
)

func TestWorkflowStore(t *testing.T) {
	redis, err := miniredis.Run()
	if err != nil {
		fmt.Println(err)
	}
	defer redis.Close()

	workflows := []string{"finish"}
	workflows = append(workflows, "start")

	redis.SAdd("workflow", workflows...)

	res, err := redis.SMembers("workflow")

	if err != nil {
		fmt.Println(err)
	}

	if len(res) != 2 {
		t.Errorf("Failed to get the data")
	}

	if res[0] != workflows[0] {
		t.Errorf("Different value")

	}
	if res[1] != workflows[1] {
		t.Errorf("Different value")

	}
}
