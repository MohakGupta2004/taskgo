package cmd

import (
	"fmt"
	"strings"

	"github.com/MohakGupta2004/taskgo/internal/config"
	"github.com/MohakGupta2004/taskgo/internal/ui"
	"github.com/spf13/cobra"
)

var checkoutCmd = &cobra.Command{
	Use:   "checkout [group]",
	Short: "Switch the active task group",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		groupName := args[0]
		if strings.EqualFold(groupName, "default") || strings.EqualFold(groupName, "general") {
			groupName = ""
		}

		ctx, err := config.LoadContext()
		if err != nil {
			fmt.Println(ui.ErrorStyle.Render("Error loading context: " + err.Error()))
			return
		}

		ctx.CurrentGroup = groupName
		if err := config.SaveContext(ctx); err != nil {
			fmt.Println(ui.ErrorStyle.Render("Error saving context: " + err.Error()))
			return
		}

		displayGroup := groupName
		if displayGroup == "" {
			displayGroup = "General"
		}
		fmt.Println(ui.SuccessStyle.Render(fmt.Sprintf("Switched to group '%s'", displayGroup)))
	},
}

func init() {
	rootCmd.AddCommand(checkoutCmd)
}
