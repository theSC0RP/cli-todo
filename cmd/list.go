package cmd

import (
	"github.com/spf13/cobra"
	"github.com/theSC0RP/cli-todo/my_table"
	"github.com/theSC0RP/cli-todo/storage"
	"github.com/theSC0RP/cli-todo/todo"
	"github.com/theSC0RP/cli-todo/utils"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks",
	Long:  "List all tasks",
	Run: func(cmd *cobra.Command, args []string) {
		tasks := storage.LoadTodos()

		todos := utils.MapValues(tasks)
		todoList := todo.TodoList{Todos: todos}

		my_table.RenderTable(todoList)
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
