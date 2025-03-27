package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/theSC0RP/cli-todo/storage"
)

var done = &cobra.Command{
	Use:   "done [id]",
	Short: "Mark a completed task as done",
	Long:  "Mark a completed task as done",
	Run: func(cmd *cobra.Command, args []string) {
		id := args[0]
		tasks := storage.LoadTodos()

		if task, exists := tasks[id]; exists {
			task.Done = true

			tasks[id] = task
			storage.SaveTodos(tasks)
			fmt.Println("Congratulations! You completed the task: ", task.Task)
		} else {
			fmt.Println("Invalid task ID")
		}
	},
}

func init() {
	RootCmd.AddCommand(done)
}
