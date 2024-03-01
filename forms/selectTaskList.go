package forms

import (
	"context"
	"fmt"
	"log"

	"github.com/charmbracelet/huh"
	"github.com/n1cholasdunn/tasks_cli/auth"
	"github.com/n1cholasdunn/tasks_cli/data"
	"google.golang.org/api/tasks/v1"
)

func SelectTaskList(ctx context.Context) (string, error) {
	var selectedTaskList string
	var options []huh.Option[string]

	srv, err := auth.NewTasksService(ctx, "credentials.json", tasks.TasksScope)
	if err != nil {
		log.Fatalf("Error initializing Google Tasks service: %v", err)
	}

	taskLists, err := data.FetchTaskLists(srv)
	if err != nil {
		log.Fatalf("Error fetching task lists: %v", err)
	}

	for _, taskList := range taskLists {
		option := huh.NewOption(taskList.Title, taskList.Id)
		options = append(options, option)
	}

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("What task list do you want to modify?").
				Options(options...).
				Value(&selectedTaskList),
		)).WithTheme(huh.ThemeCatppuccin())

	err = form.Run()
	if err != nil {
		return "", fmt.Errorf("error running form: %v", err)
	}
	return selectedTaskList, nil
}
