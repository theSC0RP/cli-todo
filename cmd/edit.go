package cmd

import (
	"database/sql"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/theSC0RP/cli-todo/db"
	"github.com/theSC0RP/cli-todo/todo"
)

var editedTask string
var editedCategory string
var editedPriority int

var updateCmd = &cobra.Command{
	Use:   "edit [id]",
	Short: "Edit an existing task",
	Long:  "",
	Args:  cobra.ExactArgs(1), // Ensures exactly one argument (ID) is provided
	Run: func(cmd *cobra.Command, args []string) {
		id := args[0]

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

		sqlDB, err := db.ConnectDB()
		if err != nil {
			fmt.Print(connectionErrorMessage, err)
			return
		}
		defer db.CloseConnection(sqlDB)

		exists, err := db.CheckIfTableExists(sqlDB, "todos")
		if err != nil {
			fmt.Print(tablCheckErrorMessage, err)
			return
		}

		if !exists {
			fmt.Print(noTodosMessage)
			return
		}

		tx, err := sqlDB.Begin()
		if err != nil {
			fmt.Println("Could not start a transaction:", err)
			return
		}

		var todoItem todo.Todo

		// Fetch the entire Todo record before updating
		selectQuery := "SELECT id, task, done, priority, category FROM todos WHERE id = ?"
		err = tx.QueryRow(selectQuery, id).Scan(&todoItem.ID, &todoItem.Task, &todoItem.Done, &todoItem.Priority, &todoItem.Category)
		if err != nil {
			if err == sql.ErrNoRows {
				tx.Rollback()
				fmt.Println("Todo with the given ID does not exist.", err)
				return
			}
			tx.Rollback()
			fmt.Println("Error fetching the todo:", err)
			return
		}

		// Apply changes
		if editedPriority != 0 {
			todoItem.Priority = editedPriority
		}
		if editedTask != "" {
			todoItem.Task = editedTask
		}
		if editedCategory != "" {
			todoItem.Category = editedCategory
		}

		updateQuery := "UPDATE todos SET task = ?, priority = ?, category = ? WHERE id = ?;"
		_, err = tx.Exec(updateQuery, todoItem.Task, todoItem.Priority, todoItem.Category, id)
		if err != nil {
			tx.Rollback()
			fmt.Println("Error updating todo:", err)
			return
		}

		if err = tx.Commit(); err != nil {
			fmt.Println("Error committing transaction:", err)
			return
		}

		fmt.Printf("Todo updated successfully.\n")
	},
}

func init() {
	updateCmd.Flags().IntVarP(&editedPriority, "priority", "p", 0, "Changed priority (1-lowest to 5-highest)")
	updateCmd.Flags().StringVarP(&editedTask, "task", "t", "", "Changed task description")
	updateCmd.Flags().StringVarP(&editedCategory, "category", "c", "", "Changed category")

	RootCmd.AddCommand(updateCmd)
}
