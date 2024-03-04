package forms

import (
	"context"
	"log"

	"github.com/charmbracelet/huh"
	"github.com/n1cholasdunn/tasks_cli/auth"
	"github.com/n1cholasdunn/tasks_cli/data"
	"google.golang.org/api/tasks/v1"
)

func MutliselectTaskList(ctx context.Context, taskListId string) ([]string, error) {
	var selectedTasks []string
	var options []huh.Option[string]

	srv, err := auth.NewTasksService(ctx, "credentials.json", tasks.TasksScope)
	if err != nil {
		log.Fatalf("Error initializing Google Tasks service: %v", err)
	}

	taskList, err := data.FetchTasks(srv, taskListId)
	if err != nil {
		log.Fatalf("Error fetching task list: %v", err)
	}
	itemCount := len(taskList)

	for _, task := range taskList {
		option := huh.NewOption(task.Title, task.Id)
		options = append(options, option)
	}

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewMultiSelect[string]().
				Title("What task do you want to modify?").Options(options...).Height(itemCount * 2).
				Value(&selectedTasks),
		))

	err = form.Run()
	if err != nil {
		return nil, err
	}
	return selectedTasks, nil
}
