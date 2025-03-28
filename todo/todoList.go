package todo

import (
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/theSC0RP/cli-todo/my_table"
)

var _ my_table.TableRowProvider = (*TodoList)(nil)

type TodoList struct {
	Todos []Todo
}

func (tl TodoList) ToTableRows() []table.Row {
	var rows []table.Row
	for _, t := range tl.Todos {
		isDone := "✅"
		if !t.Done {
			isDone = "❌"
		}
		rows = append(rows, table.Row{t.ID, t.Task, Priorities.ToStr[t.Priority], t.Category, isDone})
	}
	return rows
}

func (tl TodoList) ToTableHeader() table.Row {
	row := make(table.Row, len(TodoFields))
	for i, field := range TodoFields {
		row[i] = field
	}
	return row
}
