package data

import (
	"context"
	"log"

	"github.com/n1cholasdunn/tasks_cli/auth"
	"google.golang.org/api/tasks/v1"
)

func CreateTask(ctx context.Context, tasklistID string, title string) (*tasks.Task, error) {
	srv, err := auth.NewTasksService(ctx, "credentials.json", tasks.TasksScope)
	if err != nil {
		log.Fatalf("Unable to retrieve tasks Client %v", err)
		return nil, err
	}

	newTask := &tasks.Task{
		Title: title,
	}

	task, err := srv.Tasks.Insert(tasklistID, newTask).Do()
	if err != nil {
		log.Printf("Unable to create task: %v", err)
		return nil, err
	}

	return task, nil
}
