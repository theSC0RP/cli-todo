package todo

type Todo struct {
	ID       string `db:"id"`
	Task     string `db:"task"`
	Done     bool   `db:"done"`
	Priority int    `db:"priority"`
	Category string `db:"category"`
}

var TodoFields = []string{"ID", "Task", "Priority", "Category", "Complete"}

var TodoColumns = map[string]string{
	"id":       "INTEGER PRIMARY KEY AUTOINCREMENT",
	"task":     "TEXT",
	"priority": "INTEGER",
	"category": "TEXT",
	"done":     "BOOLEAN",
}
