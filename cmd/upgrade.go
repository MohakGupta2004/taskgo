package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/MohakGupta2004/taskgo/internal/ui"
	"github.com/spf13/cobra"
)

var upgradeCmd = &cobra.Command{
	Use:   "upgrade",
	Short: "Upgrade taskgo to the latest version",
	Long:  `Clones the latest repository to a temporary directory and runs the installation script.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(ui.InfoStyle.Render("Initiating upgrade..."))

		// 1. Create temporary directory
		tempDir, err := os.MkdirTemp("", "taskgo-upgrade")
		if err != nil {
			fmt.Println(ui.ErrorStyle.Render("Failed to create temp directory: " + err.Error()))
			return
		}
		defer os.RemoveAll(tempDir) // Cleanup

		// 2. Clone repository
		fmt.Println(ui.InfoStyle.Render("Cloning latest repository..."))
		cloneCmd := exec.Command("git", "clone", "https://github.com/MohakGupta2004/taskgo.git", tempDir)
		cloneCmd.Stdout = os.Stdout
		cloneCmd.Stderr = os.Stderr
		if err := cloneCmd.Run(); err != nil {
			fmt.Println(ui.ErrorStyle.Render("Failed to clone repository: " + err.Error()))
			return
		}

		// 3. Run install.sh
		fmt.Println(ui.InfoStyle.Render("Running installation script..."))
		installScript := filepath.Join(tempDir, "install.sh")

		// Make executable just in case
		exec.Command("chmod", "+x", installScript).Run()

		installCmd := exec.Command(installScript)
		installCmd.Dir = tempDir
		installCmd.Stdout = os.Stdout
		installCmd.Stderr = os.Stderr

		// We might need sudo for the install script's final move step,
		// but the script handles sudo internally for the mv command if run by user.
		// However, if the user runs `taskgo upgrade`, they might not be root.
		// The install.sh uses `sudo mv`. This will prompt for password if needed.

		if err := installCmd.Run(); err != nil {
			fmt.Println(ui.ErrorStyle.Render("Installation failed: " + err.Error()))
			return
		}

		fmt.Println(ui.SuccessStyle.Render("Upgrade complete!"))
	},
}

func init() {
	rootCmd.AddCommand(upgradeCmd)
}
