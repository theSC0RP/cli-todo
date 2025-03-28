package todo

type Todo struct {
	ID       string `json:"id"`
	Task     string `json:"task"`
	Done     bool   `json:"done"`
	Priority int    `json:"priority"`
	Category string `json:"category"`
}

var TodoFields = []string{"ID", "Task", "Priority", "Category", "Complete"}
