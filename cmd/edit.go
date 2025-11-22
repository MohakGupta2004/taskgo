package cmd

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/MohakGupta2004/taskgo/internal/config"
	"github.com/MohakGupta2004/taskgo/internal/ui"
	"github.com/spf13/cobra"
)

var editCmd = &cobra.Command{
	Use:   "edit [id] [new title]",
	Short: "Edit task title or validity",
	Long: `Edit a task's title or validity.

Examples:
  taskgo edit 1 "New task title"           # Edit task title
  taskgo edit 1 --validity 2h              # Edit task validity
  taskgo edit 1 --validity none            # Remove task validity
  taskgo edit --group work --validity 4h   # Edit group validity`,
	Args: cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		groupFlag, _ := cmd.Flags().GetString("group")
		validityFlag, _ := cmd.Flags().GetString("validity")

		// Edit group validity
		if groupFlag != "" {
			if validityFlag == "" {
				fmt.Println(ui.ErrorStyle.Render("Please specify --validity when editing a group"))
				return
			}

			// First, validate the duration format if not "none"
			if validityFlag != "none" {
				if _, err := time.ParseDuration(validityFlag); err != nil {
					fmt.Println(ui.ErrorStyle.Render("Invalid duration format"))
					return
				}
			}

			// Update all existing tasks in the group
			if err := taskManager.UpdateGroupValidity(groupFlag, validityFlag); err != nil {
				fmt.Println(ui.ErrorStyle.Render("Error updating group tasks: " + err.Error()))
				return
			}

			// Update the group validity in context for future tasks
			ctx, err := config.LoadContext()
			if err != nil {
				fmt.Println(ui.ErrorStyle.Render("Error loading context: " + err.Error()))
				return
			}

			if ctx.GroupValidity == nil {
				ctx.GroupValidity = make(map[string]string)
			}

			if validityFlag == "none" {
				delete(ctx.GroupValidity, groupFlag)
				fmt.Println(ui.SuccessStyle.Render(fmt.Sprintf("Removed validity for group '%s' and all its tasks", groupFlag)))
			} else {
				ctx.GroupValidity[groupFlag] = validityFlag
				fmt.Println(ui.SuccessStyle.Render(fmt.Sprintf("Group '%s' validity set to %s for all existing and future tasks", groupFlag, validityFlag)))
			}

			if err := config.SaveContext(ctx); err != nil {
				fmt.Println(ui.ErrorStyle.Render("Error saving context: " + err.Error()))
				return
			}
			return
		}

		// Edit task
		if len(args) < 1 {
			fmt.Println(ui.ErrorStyle.Render("Please specify a task ID or use --group flag"))
			return
		}

		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println(ui.ErrorStyle.Render("Invalid task ID"))
			return
		}

		// Edit task validity
		if validityFlag != "" {
			if err := taskManager.UpdateValidity(id, validityFlag); err != nil {
				fmt.Println(ui.ErrorStyle.Render("Error updating task validity: " + err.Error()))
				return
			}
			if validityFlag == "none" {
				fmt.Println(ui.SuccessStyle.Render("Task validity removed successfully!"))
			} else {
				fmt.Println(ui.SuccessStyle.Render(fmt.Sprintf("Task validity updated to %s!", validityFlag)))
			}
			return
		}

		// Edit task title
		if len(args) < 2 {
			fmt.Println(ui.ErrorStyle.Render("Please specify a new title or use --validity flag"))
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
	editCmd.Flags().StringP("validity", "v", "", "Set or update validity duration (use 'none' to remove)")
	editCmd.Flags().StringP("group", "g", "", "Edit group validity instead of task")
	rootCmd.AddCommand(editCmd)
}
