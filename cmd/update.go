package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/MohakGupta2004/taskgo/internal/task"
	"github.com/MohakGupta2004/taskgo/internal/ui"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update [id] [status]",
	Short: "Update task status (todo, in-progress, completed)",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println(ui.ErrorStyle.Render("Invalid ID"))
			return
		}

		statusStr := strings.ToLower(args[1])
		var status task.TaskStatus
		switch statusStr {
		case "todo":
			status = task.StatusTodo
		case "in-progress":
			status = task.StatusInProgress
		case "completed":
			status = task.StatusCompleted
		default:
			fmt.Println(ui.ErrorStyle.Render("Invalid status. Use: todo, in-progress, completed"))
			return
		}

		if err := taskManager.Update(id, status); err != nil {
			fmt.Println(ui.ErrorStyle.Render("Error updating task: " + err.Error()))
			return
		}
		fmt.Println(ui.SuccessStyle.Render("Task updated successfully!"))
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
