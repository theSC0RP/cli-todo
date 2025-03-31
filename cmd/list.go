package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/theSC0RP/cli-todo/db"
	"github.com/theSC0RP/cli-todo/my_table"
	"github.com/theSC0RP/cli-todo/todo"
)

var taskFilter string
var doneFilter string
var categoryFilter string
var priorityFilter int

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks",
	Long: `Lists all tasks from the to-do list with optional filters.

Available Filters:
    --filter-task, -t: Search by task content (substring search)
    --filter-done, -d: Filter by completion status (Y/N)
    --filter-priority, -p: Filter by priority (1-5)
		--filter-category", -c: Filter by category

Examples:
    cli-todo list --filter-task "buy"
    cli-todo list --filter-priority 5
    cli-todo list --filter-done Y
		cli-todo list --filter-category "personal"
    cli-todo list --filter-task "buy" --filter-priority 5 --filter-done false
	`,
	Run: func(cmd *cobra.Command, args []string) {
		sqlDB, err := db.ConnectDB()
		if err != nil {
			fmt.Println(connectionErrorMessage, err)
			return
		}

		exists, err := db.CheckIfTableExists(sqlDB, "todos")
		if err != nil {
			fmt.Println(tablCheckErrorMessage, err)
			return
		}

		if !exists {
			fmt.Println(noTodosMessage)
		}

		filterTaskFlagUsed := cmd.Flags().Lookup("filter-task").Changed
		filterCategoryFlagUsed := cmd.Flags().Lookup("filter-category").Changed
		filterDoneFlagUsed := cmd.Flags().Lookup("filter-done").Changed
		filterPriorityFlagUsed := cmd.Flags().Lookup("filter-priority").Changed

		conditionStrings := []string{}
		if filterTaskFlagUsed || filterCategoryFlagUsed || filterDoneFlagUsed || filterPriorityFlagUsed {

			if filterTaskFlagUsed {
				if taskFilter == "" {
					fmt.Println("Please enter key to search in tasks.")
					return
				} else {
					conditionStrings = append(conditionStrings, fmt.Sprintf("task LIKE '%%%s%%'", taskFilter))
				}
			}

			if filterCategoryFlagUsed {
				if categoryFilter == "" {
					fmt.Println("Please enter a valid non-empty category.")
					return
				} else {
					conditionStrings = append(conditionStrings, fmt.Sprintf("category = %%%s%%", categoryFilter))
				}
			}

			if filterDoneFlagUsed {
				doneFilter = strings.ToLower(doneFilter)
				if doneFilter == "" || (doneFilter != "y" && doneFilter != "n") {
					fmt.Println("Please press Y for yes and N for no to filter by completion status.")
					return
				} else {
					doneVal := 0
					if doneFilter == "y" {
						doneVal = 1
					}
					conditionStrings = append(conditionStrings, fmt.Sprintf("done = %d", doneVal))
				}
			}

			if filterPriorityFlagUsed {
				if priorityFilter <= 0 || priorityFilter > 5 {
					fmt.Println("Please enter a priority value between 1 (lowest) and 5 (highest).")
				} else {
					conditionStrings = append(conditionStrings, fmt.Sprintf("priority = %d", priorityFilter))
				}
			}
		}

		conditionString := ""
		if len(conditionStrings) > 0 {
			conditionString = fmt.Sprintf("WHERE %s", strings.Join(conditionStrings, " AND "))
		}

		orderString := `
ORDER BY 
	CASE WHEN todos.done = 1 THEN 2 ELSE 1 END, 
	CASE 	WHEN PRIORITY = 5 THEN 1 
				WHEN PRIORITY = 4 THEN 2 
				WHEN PRIORITY = 3 THEN 3 
				WHEN PRIORITY = 2 THEN 4 
	ELSE 5 END;`
		query := fmt.Sprintf(`SELECT * FROM todos %s %s`, conditionString, orderString)

		rows, err := sqlDB.Query(query)
		if err != nil {
			fmt.Printf("Could not execute select query: %v", err)
			return
		}
		defer rows.Close()

		var tasks []todo.Todo
		for rows.Next() {
			var task todo.Todo
			err := rows.Scan(&task.ID, &task.Task, &task.Priority, &task.Category, &task.Done)
			if err != nil {
				fmt.Printf("could not scan row: %v", err)
				return
			}
			tasks = append(tasks, task)
		}

		taskList := todo.TodoList{Todos: tasks}
		my_table.RenderTable(taskList)
	},
}

func init() {
	listCmd.Flags().StringVarP(&taskFilter, "filter-task", "t", "", "Filter by task substring")
	listCmd.Flags().StringVarP(&categoryFilter, "filter-category", "c", "", "Filter by category")
	listCmd.Flags().StringVarP(&doneFilter, "filter-done", "d", "", "Filter by completion status (Y/N)")
	listCmd.Flags().IntVarP(&priorityFilter, "filter-priority", "p", 0, "Filter by priority (1-5)")

	RootCmd.AddCommand(listCmd)
}
