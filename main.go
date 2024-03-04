package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/n1cholasdunn/tasks_cli/data"
	"github.com/n1cholasdunn/tasks_cli/forms"
)

func main() {
	ctx := context.Background()

	selectedTaskListId, err := forms.SelectTaskList(ctx)
	if err != nil {
		log.Fatalf("Error selecting task list: %v", err)
	}

	operation, err := forms.SelectCRUD()
	if err != nil {
		log.Fatalf("Error selecting operation: %v", err)
	}

	switch operation {
	case "add":
		fmt.Println("Enter the task title here:")

		reader := bufio.NewReader(os.Stdin)

		title, err := reader.ReadString('\n')
		title = strings.TrimSpace(title)
		if err != nil {
			log.Fatalf("Unable to read title: %v", err)
		}

		data.CreateTask(ctx, selectedTaskListId, title)

	case "update":
		selectedTaskId, err := forms.SelectTask(ctx, selectedTaskListId)
		if err != nil {
			log.Fatalf("Error selecting task: %v", err)
		}
		fmt.Println("Enter the new task title here:")
		reader := bufio.NewReader(os.Stdin)
		title, err := reader.ReadString('\n')
		title = strings.TrimSpace(title)
		if err != nil {
			log.Fatalf("Unable to read title: %v", err)
		}

		editConfirmed, err := forms.ConfirmEditForm()
		if err != nil {
			log.Fatalf("Error confirming edit: %v", err)
		}
		if editConfirmed {
			data.PatchTask(ctx, selectedTaskListId, selectedTaskId, title)
			fmt.Printf("Task %s updated\n", selectedTaskId)
		}

	case "delete":
		selectedTaskId, err := forms.SelectTask(ctx, selectedTaskListId)
		if err != nil {
			log.Fatalf("Error selecting task: %v", err)
		}
		deleteConfirmed, err := forms.ConfirmDeleteForm()
		if err != nil {
			log.Fatalf("Error confirming delete: %v", err)
		}
		if deleteConfirmed {
			data.DeleteTask(ctx, selectedTaskListId, selectedTaskId)
			fmt.Printf("Task %s deleted\n", selectedTaskId)
		}
	case "list":
		_, err := forms.SelectTask(ctx, selectedTaskListId)
		if err != nil {
			log.Fatalf("Error selecting task: %v", err)
		}

	}
	// cmd.Execute()
}
