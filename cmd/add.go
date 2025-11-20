package cmd

import (
	"fmt"
	"strings"
	"time"

	"github.com/MohakGupta2004/taskgo/internal/config"
	"github.com/MohakGupta2004/taskgo/internal/ui"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [task]",
	Short: "Add a new task",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		group, _ := cmd.Flags().GetString("group")
		validity, _ := cmd.Flags().GetString("validity")

		if group == "" {
			ctx, err := config.LoadContext()
			if err == nil && ctx.CurrentGroup != "" {
				group = ctx.CurrentGroup
			}
		}

		title := strings.Join(args, " ")

		// Check if the first argument is a validity duration (only if flag not set)
		if validity == "" && len(args) > 1 {
			// Try parsing the first argument as duration
			if _, err := time.ParseDuration(args[0]); err == nil {
				validity = args[0]
				title = strings.Join(args[1:], " ")
			}
		}

		err := taskManager.Add(title, group, validity)
		if err != nil {
			fmt.Println(ui.ErrorStyle.Render("Error adding task: " + err.Error()))
			return
		}
		fmt.Println(ui.SuccessStyle.Render("Task added successfully!"))
	},
}

func init() {
	addCmd.Flags().StringP("group", "g", "", "Group for the task")
	addCmd.Flags().StringP("validity", "v", "", "Validity duration (e.g. 1h, 30m)")
	rootCmd.AddCommand(addCmd)
}
