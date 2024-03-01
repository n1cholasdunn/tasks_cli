/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/n1cholasdunn/tasks_cli/auth"
	"github.com/n1cholasdunn/tasks_cli/data"
	"github.com/n1cholasdunn/tasks_cli/forms"
	"google.golang.org/api/tasks/v1"
)

func main() {
	ctx := context.Background()

	srv, err := auth.NewTasksService(ctx, "credentials.json", tasks.TasksReadonlyScope)
	if err != nil {
		log.Fatalf("Error initializing Google Tasks service: %v", err)
	}

	taskLists, err := data.FetchTaskLists(srv)
	if err != nil {
		log.Fatalf("Error fetching task lists: %v", err)
	}
	for _, taskList := range taskLists {
		fmt.Printf("%s (%s)\n", taskList.Title, taskList.Id)
		tasks, err := data.FetchTasks(srv, taskList.Id)
		if err != nil {
			log.Printf("Error fetching tasks from list %s: %v", taskList.Title, err)
			continue
		}
		for _, task := range tasks {
			fmt.Printf("- %s\n", task.Title)
		}
	}
	/*
		r, err := srv.Tasklists.List().MaxResults(10).Do()
		if err != nil {
			log.Fatalf("Unable to retrieve task lists: %v", err)
		}
		fmt.Println("Task Lists:")
		if len(r.Items) > 0 {
			for _, taskList := range r.Items {
				fmt.Printf("%s (%s)\n", taskList.Title, taskList.Id)

				// fetch tasks within current task list
				tasks, err := srv.Tasks.List(taskList.Id).Do()
				if err != nil {
					log.Printf("Unable to retrieve tasks from list %s: %v", taskList.Title, err)
					continue
				}

				fmt.Println("Tasks:")
				if len(tasks.Items) > 0 {
					for _, task := range tasks.Items {
						fmt.Printf("- %s\n", task.Title)
					}
				} else {
					fmt.Println("- No tasks found.")
				}
			}
		} else {
			fmt.Print("No task lists found.")
		}
	*/
	forms.Form()
	// cmd.Execute()
}
