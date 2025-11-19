package cmd

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/MohakGupta2004/taskgo/internal/task"
	"github.com/MohakGupta2004/taskgo/internal/ui"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(ui.RenderBanner())
		tasks, err := taskManager.List()
		if err != nil {
			fmt.Println(ui.ErrorStyle.Render("Error listing tasks: " + err.Error()))
			return
		}

		if len(tasks) == 0 {
			fmt.Println(ui.WarningStyle.Render("No tasks found."))
			return
		}

		fmt.Println(ui.RenderTitle("Task List"))

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"ID", "Title", "Status", "Created At", "Completed At"})
		table.SetBorder(false)
		table.SetHeaderLine(false)
		table.SetRowLine(false)
		table.SetCenterSeparator("")
		table.SetColumnSeparator("")
		table.SetRowSeparator("")
		table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
		table.SetAlignment(tablewriter.ALIGN_LEFT)

		for _, t := range tasks {
			status := string(t.Status)
			switch t.Status {
			case task.StatusTodo:
				status = ui.StatusTodoStyle.Render(status)
			case task.StatusInProgress:
				status = ui.StatusInProgressStyle.Render(status)
			case task.StatusCompleted:
				status = ui.StatusCompletedStyle.Render(status)
			}

			completedAt := ""
			if t.CompletedAt != nil {
				completedAt = t.CompletedAt.Format(time.RFC822)
			}

			table.Append([]string{
				strconv.Itoa(t.ID),
				t.Title,
				status,
				t.CreatedAt.Format(time.RFC822),
				completedAt,
			})
		}
		table.Render()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
