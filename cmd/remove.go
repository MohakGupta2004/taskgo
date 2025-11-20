package cmd

import (
	"fmt"
	"strconv"

	"github.com/MohakGupta2004/taskgo/internal/config"
	"github.com/MohakGupta2004/taskgo/internal/ui"
	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "remove [id]",
	Short: "Remove a task",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if args[0] == "*" || args[0] == "all" {
			ctx, err := config.LoadContext()
			group := "General"
			if err == nil && ctx.CurrentGroup != "" {
				group = ctx.CurrentGroup
			}

			if err := taskManager.RemoveByGroup(group); err != nil {
				fmt.Println(ui.ErrorStyle.Render("Error removing tasks: " + err.Error()))
				return
			}
			fmt.Println(ui.SuccessStyle.Render(fmt.Sprintf("Removed all tasks in group '%s'", group)))
			return
		}

		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println(ui.ErrorStyle.Render("Invalid ID"))
			return
		}

		if err := taskManager.Remove(id); err != nil {
			fmt.Println(ui.ErrorStyle.Render("Error removing task: " + err.Error()))
			return
		}
		fmt.Println(ui.SuccessStyle.Render("Task removed successfully!"))
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}
