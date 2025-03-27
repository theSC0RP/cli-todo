package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/theSC0RP/cli-todo/storage"
	"github.com/theSC0RP/cli-todo/todo"
)

func getFirstMissingNumber(nums []int) string {
	missingNum := 1
	numsSet := make(map[int]struct{})
	for _, n := range nums {
		numsSet[int(n)] = struct{}{}
	}

	for {
		if _, exists := numsSet[missingNum]; exists {
			missingNum++
		} else {
			return strconv.Itoa(missingNum)
		}
	}

}

func getIds(tasks map[string]todo.Todo, ids *[]int) {
	for id, _ := range tasks {
		intID, _ := strconv.Atoi(id)
		*ids = append(*ids, intID)
	}
}

var addCmd = &cobra.Command{
	Use:   "add [task]",
	Short: "Add a new task",
	Long:  "Add a new task",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		tasks := storage.LoadTodos()

		var ids []int
		getIds(tasks, &ids)

		fmt.Println(ids)
		id := getFirstMissingNumber(ids)

		fmt.Println("new ID: ", id)
		tasks[id] = todo.Todo{ID: id, Task: args[0], Done: false}
		fmt.Println(tasks)

		storage.SaveTodos(tasks)
		fmt.Println("Added :", args[0])
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
