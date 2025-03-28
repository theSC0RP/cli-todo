package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/theSC0RP/cli-todo/storage"
)

var doneCmd = &cobra.Command{
	Use:   "done [id]",
	Short: "Mark a task as completed",
	Long: `Marks a task as completed in the to-do list using the provided ID.

The task's "Done" status will be updated to true.

Usage examples:
    cli-todo done 3
    cli-todo done 7`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id := args[0]
		tasks := storage.LoadTodos()

		// Check if the task exists
		task, exists := tasks[id]
		if !exists {
			fmt.Printf("Error: Task with ID %s not found.\n", id)
			return
		}

		// Mark task as done
		if task.Done {
			fmt.Printf("Task %s is already marked as completed.\n", id)
			return
		}

		task.Done = true
		tasks[id] = task
		storage.SaveTodos(tasks)

		fmt.Printf("✅ Task %s marked as completed: %s\n", id, task.Task)
	},
}

func init() {
	RootCmd.AddCommand(doneCmd)
}
