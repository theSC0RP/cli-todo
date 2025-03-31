package cmd

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/theSC0RP/cli-todo/db"
	"github.com/theSC0RP/cli-todo/todo"
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
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Error: Please enter a valid [id]")
			return
		}

		task, err := changeTaskStatus(true, id)
		if err != nil {
			return
		}

		fmt.Printf("✅ Task \"%d: %s\" marked as completed\n", id, task)
	},
}

var undoneCmd = &cobra.Command{
	Use:   "undone [id]",
	Short: "Mark a task as not complete",
	Long: `Marks a task as not completed in the to-do list using the provided ID.

The task's "Done" status will be updated to false.

Usage examples:
    cli-todo undone 3
    cli-todo undone 7`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Error: Please enter a valid [id]")
			return
		}

		task, err := changeTaskStatus(false, id)
		if err != nil {
			return
		}

		fmt.Printf("Task \"%d: %s\" marked as incomplete\n", id, task)
	},
}

func changeTaskStatus(status bool, id int) (string, error) {
	sqlDB, err := db.ConnectDB()
	if err != nil {
		fmt.Print(connectionErrorMessage, err)
		return "", err
	}

	exists, err := db.CheckIfTableExists(sqlDB, "todos")
	if err != nil {
		fmt.Print(tablCheckErrorMessage, err)
		return "", err
	}

	if !exists {
		fmt.Print(noTodosMessage)
		return "", err
	}

	tx, err := sqlDB.Begin()
	if err != nil {
		fmt.Println("Could not start a transaction:", err)
		return "", err
	}

	var todoItem todo.Todo

	// Fetch the entire Todo record before updating
	selectQuery := "SELECT id, task, done, priority, category FROM todos WHERE id = ?"
	err = tx.QueryRow(selectQuery, id).Scan(&todoItem.ID, &todoItem.Task, &todoItem.Done, &todoItem.Priority, &todoItem.Category)
	if err != nil {
		if err == sql.ErrNoRows {
			tx.Rollback()
			fmt.Println("Todo with the given ID does not exist.", err)
			return "", err
		}
		tx.Rollback()
		fmt.Println("Error fetching the todo:", err)
		return "", err
	}

	if todoItem.Done == status {
		statusString := "incomplete"
		if status {
			statusString = "complete"
		}

		fmt.Printf("Task is already marked as %s!", statusString)
		tx.Rollback()
		return "", fmt.Errorf("Task is already marked as %s!", statusString)
	}

	updateQuery := "UPDATE todos SET done = ? WHERE id = ?"
	_, err = tx.Exec(updateQuery, status, id)
	if err != nil {
		tx.Rollback()
		fmt.Println("Error marking todo as done:", err)
		return "", err
	}

	if err = tx.Commit(); err != nil {
		fmt.Println("Error committing transaction:", err)
		return "", err
	}

	return todoItem.Task, nil
}

func init() {
	RootCmd.AddCommand(doneCmd)
	RootCmd.AddCommand(undoneCmd)
}
