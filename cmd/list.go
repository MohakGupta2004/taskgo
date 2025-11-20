package cmd

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/MohakGupta2004/taskgo/internal/task"
	"github.com/MohakGupta2004/taskgo/internal/ui"
	"github.com/charmbracelet/lipgloss"
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

		// Group tasks
		groupedTasks := make(map[string][]task.Task)
		var groups []string
		for _, t := range tasks {
			groupName := t.Group
			if groupName == "" {
				groupName = "General"
			}
			if _, exists := groupedTasks[groupName]; !exists {
				groups = append(groups, groupName)
			}
			groupedTasks[groupName] = append(groupedTasks[groupName], t)
		}

		// Iterate over groups
		for _, group := range groups {
			// Render Tree Branch / Group Header
			fmt.Println(ui.TreeBranchStyle.Render("├── " + group))

			table := tablewriter.NewWriter(os.Stdout)
			table.SetHeader([]string{"ID", "Title", "Status", "Created At", "Completed At"})
			table.SetBorder(true)
			table.SetHeaderLine(true)
			table.SetRowLine(false)
			table.SetCenterSeparator("|")
			table.SetColumnSeparator("|")
			table.SetRowSeparator("-")
			table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
			table.SetAlignment(tablewriter.ALIGN_LEFT)
			table.SetAutoWrapText(false)
			table.SetReflowDuringAutoWrap(false)

			for _, t := range groupedTasks[group] {
				statusStr := string(t.Status)
				if t.Status == task.StatusTodo {
					statusStr = "pending"
				}

				titleStr := t.Title

				// Apply styling based on status
				switch t.Status {
				case task.StatusTodo:
					statusStr = ui.StatusTodoStyle.Render(statusStr)
					titleStr = ui.PendingRowStyle.Render(titleStr)
				case task.StatusInProgress:
					statusStr = ui.StatusInProgressStyle.Render(statusStr)
					titleStr = ui.InProgressRowStyle.Render(titleStr)
				case task.StatusCompleted:
					statusStr = ui.StatusCompletedStyle.Render(statusStr)
					titleStr = ui.CompletedRowStyle.Render(titleStr)
				}

				// Wrap title if it's too long
				titleStr = lipgloss.NewStyle().Width(40).Render(titleStr)

				completedAt := ""
				if t.CompletedAt != nil {
					completedAt = t.CompletedAt.Format("02 Jan 06 15:04 MST")
				}

				validUntil := ""
				if t.ValidUntil != nil {
					remaining := time.Until(*t.ValidUntil).Round(time.Minute)
					if remaining > 0 {
						validUntil = remaining.String()
					} else {
						validUntil = "Expired"
					}
				}

				// Apply orange color to all columns for pending tasks
				row := []string{
					strconv.Itoa(t.ID),
					titleStr,
					statusStr,
					t.CreatedAt.Format("02 Jan 06 15:04 MST"),
					completedAt,
					validUntil,
				}

				if t.Status == task.StatusTodo {
					for i, cell := range row {
						// Title and Status are already styled
						if i != 1 && i != 2 {
							row[i] = lipgloss.NewStyle().Foreground(ui.OrangeColor).Render(cell)
						}
					}
				}

				table.Append(row)
			}
			table.Render()
			fmt.Println("│") // Spacer between groups
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
