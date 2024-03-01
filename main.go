/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"context"
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

	data.CreateTask(ctx, selectedTaskListId, "Tifa Task")

	selectedTask, err := forms.SelectTask(ctx, selectedTaskListId)
	if err != nil {
		log.Fatalf("Error selecting task: %v", err)
	}
	log.Printf("Selected task: %s", selectedTask)
	//	forms.Form()
	//
	// cmd.Execute()
}
