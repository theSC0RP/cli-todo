package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/theSC0RP/cli-todo/storage"
)

var deleteCmd = &cobra.Command{
	Use:   "delete [id]",
	Short: "Delete a task by its ID",
	Long: `Deletes a task from the to-do list using the provided ID.

If the specified task ID does not exist, no changes will be made.

Usage examples:
    cli-todo delete 3
    cli-todo delete 7`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id := args[0]
		tasks := storage.LoadTodos()

		// Check if the task exists before deleting
		if _, exists := tasks[id]; !exists {
			fmt.Printf("Task with ID %s not found.\n", id)
			return
		}

		delete(tasks, id)
		storage.SaveTodos(tasks)

		fmt.Printf("Task %s deleted successfully.\n", id)
	},
}

func init() {
	RootCmd.AddCommand(deleteCmd)
}
