package data

import (
	"context"
	"log"

	"github.com/n1cholasdunn/tasks_cli/auth"
	"google.golang.org/api/tasks/v1"
)

func DeleteTask(ctx context.Context, tasklistID string, taskId string) error {
	srv, err := auth.NewTasksService(ctx, "credentials.json", tasks.TasksScope)
	if err != nil {
		log.Fatalf("Unable to retrieve tasks Client %v", err)
		return err
	}

	err = srv.Tasks.Delete(tasklistID, taskId).Do()
	if err != nil {
		log.Printf("Unable to create task: %v", err)
		return err
	}

	return err
}
