package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/theSC0RP/cli-todo/db"
	"github.com/theSC0RP/cli-todo/todo"
)

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
		if todoPriority < 1 || todoPriority > 5 {
			fmt.Println("Error: Todo priority should be between 1-5 (both inclusive).")
			return
		}

		task := todo.Todo{
			Task:     args[0],
			Done:     false,
			Priority: todoPriority,
			Category: todoCategory,
		}

		sqlDB, err := db.ConnectDB()
		if err != nil {
			fmt.Println(connectionErrorMessage, err)
			return
		}
		defer db.CloseConnection(sqlDB)

		err = db.CreateTableIfNotExists(sqlDB, "todos", todo.TodoColumns)
		if err != nil {
			fmt.Println("Table could not be created: ", err)
			return
		}

		query := "INSERT INTO todos (task, done, priority, category) VALUES (?, ?, ?, ?);"
		result, err := sqlDB.Exec(query, task.Task, task.Done, task.Priority, task.Category)
		if err != nil {
			fmt.Printf("Failed to add task: %v\n", err)
			return
		}

		id, err := result.LastInsertId()
		if err != nil {
			fmt.Printf("Failed to fetch last inserted Iid: %v\n", err)
			return
		}

		fmt.Printf("Task added: [%d] %s (Priority: %d, Category: %s)\n", id, args[0], todoPriority, todoCategory)

		defer db.CloseConnection(sqlDB)
	},
}

func init() {
	addCmd.Flags().IntVarP(&todoPriority, "priority", "p", 3, "Priority of the task: \n\t5-highest\n\t4-high\n\t3-medium (default)\n\t2-low\n\t1-lowest\n\n")
	addCmd.Flags().StringVarP(&todoCategory, "category", "c", "", "Category of the task")

	RootCmd.AddCommand(addCmd)
}
