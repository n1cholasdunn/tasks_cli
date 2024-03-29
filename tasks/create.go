package tasks

import (
	"context"
	"log"
	"net/http"

	"google.golang.org/api/option"
	"google.golang.org/api/tasks/v1"
)

func CreateTask(ctx context.Context, client *http.Client, tasklistID string, title string) (*tasks.Task, error) {
	srv, err := tasks.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to retrieve tasks Client %v", err)
		return nil, err
	}

	newTask := &tasks.Task{
		Title: title,
	}

	// Create the task in the specified task list
	task, err := srv.Tasks.Insert(tasklistID, newTask).Do()
	if err != nil {
		log.Printf("Unable to create task: %v", err)
		return nil, err
	}

	return task, nil
}
