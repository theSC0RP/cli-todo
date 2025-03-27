package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/theSC0RP/cli-todo/storage"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks",
	Long:  "List all tasks",
	Run: func(cmd *cobra.Command, args []string) {
		tasks := storage.LoadTodos()

		for _, task := range tasks {
			status := "[ ]"
			if task.Done {
				status = "[x]"
			}
			fmt.Printf("\n%s | %s | %s\n", task.ID, task.Task, status)
		}
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
