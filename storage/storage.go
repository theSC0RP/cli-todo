package storage

import (
	"encoding/json"
	"os"

	"github.com/theSC0RP/cli-todo/todo"
)

var todo_file = ".cli-todos"

func SaveTodos(todos map[string]todo.Todo) {
	file, _ := json.MarshalIndent(todos, "", "    ")
	os.WriteFile(todo_file, file, 0644)
}

func LoadTodos() map[string]todo.Todo {
	var todos map[string]todo.Todo
	file, err := os.ReadFile(todo_file)

	if err == nil {
		json.Unmarshal(file, &todos)
	} else {
		todos = make(map[string]todo.Todo)
	}

	return todos
}
