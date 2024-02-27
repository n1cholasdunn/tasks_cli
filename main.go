package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/n1cholasdunn/tasks_cli/auth"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/api/tasks/v1"
)

func main() {
	ctx := context.Background()
	b, err := os.ReadFile("credentials.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, tasks.TasksReadonlyScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	client := auth.GetClient(config)

	srv, err := tasks.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to retrieve tasks Client %v", err)
	}

	r, err := srv.Tasklists.List().MaxResults(10).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve task lists. %v", err)
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
}
