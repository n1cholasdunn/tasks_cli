package data

import (
	"log"

	"google.golang.org/api/tasks/v1"
)

func FetchTaskLists(srv *tasks.Service) ([]*tasks.TaskList, error) {
	r, err := srv.Tasklists.List().MaxResults(10).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve task lists: %v", err)
	}
	return r.Items, err
}

func FetchTasks(srv *tasks.Service, taskListId string) ([]*tasks.Task, error) {
	tasks, err := srv.Tasks.List(taskListId).Do()
	if err != nil {
		log.Printf("Unable to retrieve tasks from list %s: %v", taskListId, err)
	}
	return tasks.Items, err
}
