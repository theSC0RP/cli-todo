package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/theSC0RP/cli-todo/storage"
)

var editedTask string
var editedCategory string
var editedPriority int

var updateCmd = &cobra.Command{
	Use:   "edit [id]",
	Short: "Edit an existing task",
	Long: `Edits an existing task in the to-do list.

You must provide the task ID to edit. You can modify the task description, 
priority, and category. If a field is not specified, it remains unchanged.

Priority levels:
    5 - Highest
    4 - High
    3 - Medium
    2 - Low
    1 - Lowest

Usage examples:
    cli-todo edit 3 -t "Buy fresh vegetables"
    cli-todo edit 7 -p 4 -c "Personal"
    cli-todo edit 5 -t "Submit assignment" -p 5 -c "College"`,
	Args: cobra.ExactArgs(1), // Ensures exactly one argument (ID) is provided
	Run: func(cmd *cobra.Command, args []string) {
		id := args[0]

		// Load existing todos
		todos := storage.LoadTodos()

		// Check if task exists
		task, exists := todos[id]
		if !exists {
			fmt.Println("Task not found. Use 'list' to view existing tasks.")
			return
		}

		// Ensure at least one change is provided
		if editedTask == "" && editedPriority == 0 && editedCategory == "" {
			fmt.Println("No change provided.")
			fmt.Println("\tUse -t to specify the edited task.")
			fmt.Println("\tUse -p to specify new priority.")
			fmt.Println("\tUse -c to specify new category.")
			return
		}

		// Validate priority range
		if editedPriority != 0 && (editedPriority < 1 || editedPriority > 5) {
			fmt.Println("Priority must be between 1 (lowest) and 5 (highest).")
			return
		}

		// Apply changes
		if editedPriority != 0 {
			task.Priority = editedPriority
		}
		if editedTask != "" {
			task.Task = editedTask
		}
		if editedCategory != "" {
			task.Category = editedCategory
		}

		// Save the updated task
		todos[id] = task
		storage.SaveTodos(todos)
		fmt.Printf("Task %s updated successfully.\n", id)
	},
}

func init() {
	updateCmd.Flags().IntVarP(&editedPriority, "priority", "p", 0, "Changed priority (1-lowest to 5-highest)")
	updateCmd.Flags().StringVarP(&editedTask, "task", "t", "", "Changed task description")
	updateCmd.Flags().StringVarP(&editedCategory, "category", "c", "", "Changed category")

	RootCmd.AddCommand(updateCmd)
}
