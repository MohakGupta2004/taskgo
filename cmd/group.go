package cmd

import (
	"fmt"
	"time"

	"github.com/MohakGupta2004/taskgo/internal/config"
	"github.com/MohakGupta2004/taskgo/internal/ui"
	"github.com/spf13/cobra"
)

var groupCmd = &cobra.Command{
	Use:   "group [name]",
	Short: "Manage task groups",
	Long:  `Show current group, list all groups, or create a new one.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			var groupName string
			var validity string

			// Check for positional validity: taskgo group 2h work
			if len(args) >= 2 {
				if _, err := time.ParseDuration(args[0]); err == nil {
					validity = args[0]
					groupName = args[1]
				} else {
					groupName = args[0]
				}
			} else {
				groupName = args[0]
			}

			// Check flag if positional not found
			if validity == "" {
				validity, _ = cmd.Flags().GetString("validity")
			}

			// Save to context if validity is provided
			if validity != "" {
				ctx, err := config.LoadContext()
				if err != nil {
					fmt.Println(ui.ErrorStyle.Render("Error loading context: " + err.Error()))
					return
				}

				if ctx.GroupValidity == nil {
					ctx.GroupValidity = make(map[string]string)
				}
				ctx.GroupValidity[groupName] = validity

				if err := config.SaveContext(ctx); err != nil {
					fmt.Println(ui.ErrorStyle.Render("Error saving context: " + err.Error()))
					return
				}
				fmt.Println(ui.SuccessStyle.Render(fmt.Sprintf("Group '%s' validity set to %s.", groupName, validity)))
			} else {
				fmt.Println(ui.SuccessStyle.Render(fmt.Sprintf("Group '%s' is ready to use.", groupName)))
			}
			return
		}

		// Show current group
		ctx, err := config.LoadContext()
		if err != nil {
			fmt.Println(ui.ErrorStyle.Render("Error loading context: " + err.Error()))
			return
		}

		currentGroup := ctx.CurrentGroup
		if currentGroup == "" {
			currentGroup = "General"
		}
		fmt.Println(ui.InfoStyle.Render(fmt.Sprintf("Current group: %s", currentGroup)))
	},
}

var groupListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all task groups",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := taskManager.List()
		if err != nil {
			fmt.Println(ui.ErrorStyle.Render("Error listing groups: " + err.Error()))
			return
		}

		groups := make(map[string]bool)
		for _, t := range tasks {
			groupName := t.Group
			if groupName == "" {
				groupName = "General"
			}
			groups[groupName] = true
		}

		if len(groups) == 0 {
			fmt.Println(ui.WarningStyle.Render("No groups found."))
			return
		}

		fmt.Println(ui.TitleStyle.Render("Task Groups"))
		for g := range groups {
			fmt.Println(ui.SecondaryStyle.Render("- " + g))
		}
	},
}

func init() {
	groupCmd.Flags().StringP("validity", "v", "", "Default validity duration for the group")
	groupCmd.AddCommand(groupListCmd)
	rootCmd.AddCommand(groupCmd)
}
