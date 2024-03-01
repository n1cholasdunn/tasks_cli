package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	title string
	notes string
	due   string
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "A brief description of your command",
	Long: `Create a new task with a title, optional notes, and an optional due date. For example:

tasksync create --title "Grocery shopping" --notes "Buy milk, eggs, and bread" --due "2024-03-01"`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("create called")
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	createCmd.Flags().StringVarP(&title, "title", "t", "", "A title for the task created")
	createCmd.Flags().StringVar(&title, "title", "", "A title for the task created")

	createCmd.Flags().StringVar(&notes, "notes", "", "Notes for the task created")
	createCmd.Flags().StringVarP(&notes, "notes", "n", "", "Notes for the task created")

	createCmd.Flags().StringVar(&due, "due", "", "Due date for the task (YYYY-MM-DD)")
	createCmd.Flags().StringVarP(&due, "due", "d", "", "Due date for the task (YYYY-MM-DD)")
}
