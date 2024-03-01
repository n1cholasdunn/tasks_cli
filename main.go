/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"context"
	"log"

	"github.com/n1cholasdunn/tasks_cli/forms"
)

func main() {
	ctx := context.Background()

	selectTaskId, err := forms.SelectTaskList(ctx)
	if err != nil {
		log.Fatalf("Error selecting task list: %v", err)
	}
	selectedTask, err := forms.SelectTask(ctx, selectTaskId)
	if err != nil {
		log.Fatalf("Error selecting task: %v", err)
	}
	log.Printf("Selected task: %s", selectedTask)
	//	forms.Form()
	//
	// cmd.Execute()
}
