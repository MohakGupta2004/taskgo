package cmd

import (
	"fmt"
	"strings"

	"github.com/MohakGupta2004/taskgo/internal/ui"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [task]",
	Short: "Add a new task",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		title := strings.Join(args, " ")
		if err := taskManager.Add(title); err != nil {
			fmt.Println(ui.ErrorStyle.Render("Error adding task: " + err.Error()))
			return
		}
		fmt.Println(ui.SuccessStyle.Render("Task added successfully!"))
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
