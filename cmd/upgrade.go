package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/MohakGupta2004/taskgo/internal/ui"
	"github.com/spf13/cobra"
)

var upgradeCmd = &cobra.Command{
	Use:   "upgrade",
	Short: "Upgrade taskgo to the latest version",
	Long:  `Pull the latest changes from the repository and rebuild the executable.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(ui.InfoStyle.Render("Checking for updates..."))

		// 1. git pull
		pullCmd := exec.Command("git", "pull")
		pullCmd.Stdout = os.Stdout
		pullCmd.Stderr = os.Stderr
		if err := pullCmd.Run(); err != nil {
			fmt.Println(ui.ErrorStyle.Render("Failed to pull latest changes: " + err.Error()))
			return
		}

		fmt.Println(ui.InfoStyle.Render("Building latest version..."))

		// 2. go build
		// We assume we are in the project root or can build from here.
		buildCmd := exec.Command("go", "build", "-o", "taskgo", "main.go")
		buildCmd.Stdout = os.Stdout
		buildCmd.Stderr = os.Stderr
		if err := buildCmd.Run(); err != nil {
			fmt.Println(ui.ErrorStyle.Render("Failed to build: " + err.Error()))
			return
		}

		fmt.Println(ui.SuccessStyle.Render("Successfully upgraded taskgo!"))
	},
}

func init() {
	rootCmd.AddCommand(upgradeCmd)
}
