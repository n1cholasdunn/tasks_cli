package data

import (
	"context"
	"log"

	"github.com/n1cholasdunn/tasks_cli/auth"
	"google.golang.org/api/tasks/v1"
)

func PatchTask(ctx context.Context, tasklistID string, taskId string, title string, notes string, dueDate string) (*tasks.Task, error) {
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

	task.Title = title
	if notes != "" {
		task.Notes = notes
	}
	if dueDate != "" {
		task.Due = dueDate
	}

	_, err = srv.Tasks.Patch(tasklistID, taskId, task).Do()
	if err != nil {
		log.Printf("Unable to update task: %v", err)
		return nil, err
	}

	return task, nil
}
