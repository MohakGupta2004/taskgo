package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/MohakGupta2004/taskgo/internal/ui"
	"github.com/spf13/cobra"
)

var editCmd = &cobra.Command{
	Use:   "edit [id] [new title]",
	Short: "Edit task title",
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println(ui.ErrorStyle.Render("Invalid ID"))
			return
		}

		newTitle := strings.Join(args[1:], " ")
		if err := taskManager.UpdateTitle(id, newTitle); err != nil {
			fmt.Println(ui.ErrorStyle.Render("Error editing task: " + err.Error()))
			return
		}
		fmt.Println(ui.SuccessStyle.Render("Task updated successfully!"))
	},
}

func init() {
	rootCmd.AddCommand(editCmd)
}
