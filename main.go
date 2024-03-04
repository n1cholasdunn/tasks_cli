package main

import (
	"context"
	"fmt"
	"log"

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
		title, notes, dueDate, err := forms.AddTaskForm()
		if err != nil {
			log.Fatalf("Unable to complete add task form: %v", err)
		}

		data.CreateTask(ctx, selectedTaskListId, title, notes, dueDate)

	case "update":
		// slice of task ids to
		// var taskQueue []string
		selectedTaskId, err := forms.SelectTask(ctx, selectedTaskListId)
		if err != nil {
			log.Fatalf("Error selecting task: %v", err)
		}
		editedFields, err := forms.EditTaskForm()
		if err != nil {
			log.Fatalf("Error getting edited fields from form: %v", err)
		}

		editConfirmed, err := forms.ConfirmEditForm()
		if err != nil {
			log.Fatalf("Error confirming edit: %v", err)
		}
		if editConfirmed {
			data.PatchTask(ctx, selectedTaskListId, selectedTaskId, editedFields["title"], editedFields["notes"], editedFields["due"])
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

	case "mark":
		tasksToMark, err := forms.MutliselectTaskList(ctx, selectedTaskListId)
		if err != nil {
			log.Fatalf("Error selecting task: %v", err)
		}
		for _, taskId := range tasksToMark {
			data.MarkTask(ctx, selectedTaskListId, taskId)
			fmt.Printf("Task %s marked\n", taskId)
		}
	}
	// cmd.Execute()
}
