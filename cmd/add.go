package cmd

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/theSC0RP/cli-todo/storage"
	"github.com/theSC0RP/cli-todo/todo"
)

func getFirstMissingNumber(tasks map[string]todo.Todo) string {
	var ids []int

	for id := range tasks {
		intID, err := strconv.Atoi(id)
		if err == nil {
			ids = append(ids, intID)
		}
	}

	sort.Ints(ids) // Sort IDs in ascending order

	missingNum := 1
	for _, num := range ids {
		if num == missingNum {
			missingNum++
		} else {
			break
		}
	}
	return strconv.Itoa(missingNum)
}

var todoPriority int
var todoCategory string

var addCmd = &cobra.Command{
	Use:   "add [task]",
	Short: "Add a new task",
	Long: `Adds a new task to the to-do list.

Each task is assigned a unique ID. The task must have a description, 
and you can optionally specify a priority and a category.

Priority levels:
    5 - Highest
    4 - High
    3 - Medium (default)
    2 - Low
    1 - Lowest

Usage examples:
    cli-todo add "Buy groceries"
    cli-todo add "Finish project report" -p 5 -c "Work"`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		tasks := storage.LoadTodos()

		id := getFirstMissingNumber(tasks)

		if todoPriority < 1 || todoPriority > 5 {
			fmt.Println("Error: Todo priority should be between 1-5 (both inclusive).")
			return
		}

		// Add the new task
		tasks[id] = todo.Todo{
			ID:       id,
			Task:     args[0],
			Done:     false,
			Priority: todoPriority,
			Category: todoCategory,
		}

		// Save tasks to storage
		storage.SaveTodos(tasks)

		fmt.Printf("Task added: [%s] %s (Priority: %d, Category: %s)\n", id, args[0], todoPriority, todoCategory)
	},
}

func init() {
	addCmd.Flags().IntVarP(&todoPriority, "priority", "p", 3, "Priority of the task: \n\t5-highest\n\t4-high\n\t3-medium (default)\n\t2-low\n\t1-lowest\n\n")
	addCmd.Flags().StringVarP(&todoCategory, "category", "c", "", "Category of the task")

	RootCmd.AddCommand(addCmd)
}
