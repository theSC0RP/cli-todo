package cmd

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/theSC0RP/cli-todo/db"
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
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Error: Please enter a valid [id]")
			return
		}

		sqlDB, err := db.ConnectDB()
		if err != nil {
			fmt.Println("Error connecting to the db")
			return
		}

		exists, err := db.CheckIfTableExists(sqlDB, "todos")
		if err != nil {
			fmt.Println(err)
			return
		}

		if !exists {
			fmt.Println("No todos present, please add a todo using \"add\" command: Check \"cli-todo add -help\"")
		}

		tx, err := sqlDB.Begin()
		if err != nil {
			fmt.Println("Could not start a transaction:", err)
			return
		}
		var task string

		// Fetch the entire Todo record before updating
		selectQuery := "SELECT task FROM todos WHERE id = ?"
		err = tx.QueryRow(selectQuery, id).Scan(&task)
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

		deleteQuery := "DELETE FROM todos WHERE id = ?;"
		_, err = tx.Exec(deleteQuery, id)
		if err != nil {
			tx.Rollback()
			fmt.Println("Error deleting todo:", err)
			return
		}

		if err = tx.Commit(); err != nil {
			fmt.Println("Error committing transaction:", err)
			return
		}

		fmt.Printf("\nTodo \"%s\" deleted successfully.\n", task)
	},
}

func init() {
	RootCmd.AddCommand(deleteCmd)
}
