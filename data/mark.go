package data

import (
	"context"
	"fmt"
	"log"

	"github.com/n1cholasdunn/tasks_cli/auth"
	"github.com/n1cholasdunn/tasks_cli/helpers"
	"google.golang.org/api/tasks/v1"
)

func MarkTask(ctx context.Context, tasklistID string, taskId string) (*tasks.Task, error) {
	srv, err := auth.NewTasksService(ctx, "credentials.json", tasks.TasksScope)
	if err != nil {
		log.Fatalf("Unable to retrieve tasks Client %v", err)
		return nil, err
	}

	task, err := srv.Tasks.Get(tasklistID, taskId).Do()
	if err != nil {
		log.Printf("Unable to get task: %v", err)
		return nil, err
	}
	completedDate := helpers.GetTodayAsRFC3339()
	fmt.Println(completedDate)
	task.Completed = &completedDate
	task.Status = "completed"

	_, err = srv.Tasks.Patch(tasklistID, taskId, task).Do()
	if err != nil {
		log.Printf("Unable to mark task: %v", err)
		return nil, err
	}

	return task, nil
}
