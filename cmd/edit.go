package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/theSC0RP/cli-todo/storage"
)

var editedTask string

var updateCmd = &cobra.Command{
	Use: "edit [id]",
	Run: func(cmd *cobra.Command, args []string) {
		var id = args[0]

		// Load existing todos
		todos := storage.LoadTodos()

		// Check if the task exists
		if task, exists := todos[id]; exists {
			// Edit task description if flag is provided
			if editedTask != "" {
				task.Task = editedTask
				todos[id] = task
				storage.SaveTodos(todos)
				fmt.Printf("Task %s changed to: \"%s\"\n", id, editedTask)
			} else {
				fmt.Println("No change provided. Use -t to specify the edited task.")
			}
		} else {
			fmt.Println("Task not found")
		}
	},
}

func init() {
	updateCmd.Flags().StringVarP(&editedTask, "task", "t", "", "edited task")
	RootCmd.AddCommand(updateCmd)
}
