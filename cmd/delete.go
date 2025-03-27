package cmd

import (
	"github.com/spf13/cobra"
	"github.com/theSC0RP/cli-todo/storage"
)

var deleteCmd = &cobra.Command{
	Use:   "delete [id]",
	Short: "Delete the task having the provided id",
	Long:  "Delete the task having the provided id",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var id = args[0]
		var tasks = storage.LoadTodos()

		delete(tasks, id)

		storage.SaveTodos(tasks)
	},
}

func init() {
	RootCmd.AddCommand(deleteCmd)
}
