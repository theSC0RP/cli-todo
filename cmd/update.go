package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/theSC0RP/cli-todo/storage"
)

var updatedTask string

var updateCmd = &cobra.Command{
	Use: "update [id]",
	Run: func(cmd *cobra.Command, args []string) {
		var id = args[0]

		// Load existing todos
		todos := storage.LoadTodos()

		// Check if the task exists
		if task, exists := todos[id]; exists {
			// Update task description if flag is provided
			if updatedTask != "" {
				task.Task = updatedTask
				todos[id] = task
				storage.SaveTodos(todos)
				fmt.Printf("Task %s updated to: \"%s\"\n", id, updatedTask)
			} else {
				fmt.Println("No update provided. Use -t to specify a new task description.")
			}
		} else {
			fmt.Println("Task not found")
		}
	},
}

func init() {
	updateCmd.Flags().StringVarP(&updatedTask, "task", "t", "", "Updated task title")
	RootCmd.AddCommand(updateCmd)
}
