package cmd

import (
	"fmt"
	"time"

	"github.com/MohakGupta2004/taskgo/internal/ui"
	"github.com/spf13/cobra"
)

var pomodoroCmd = &cobra.Command{
	Use:   "pomodoro [minutes]",
	Short: "Start a pomodoro timer (default 25 minutes)",
	Args:  cobra.MaximumNArgs(1),
	Run:   runPomodoro,
}

func runPomodoro(cmd *cobra.Command, args []string) {
	minutes := 25
	if len(args) > 0 {
		var err error
		_, err = fmt.Sscanf(args[0], "%d", &minutes)
		if err != nil {
			fmt.Println(ui.ErrorStyle.Render("Invalid duration"))
			return
		}
	}

	duration := time.Duration(minutes) * time.Minute
	targetTime := time.Now().Add(duration)

	fmt.Println(ui.RenderTitle(fmt.Sprintf("Pomodoro started for %d minutes", minutes)))

	// Clear screen once
	fmt.Print("\033[H\033[2J")

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	// Initial render
	renderTimer(minutes, 0)

	for {
		remaining := time.Until(targetTime)
		if remaining <= 0 {
			fmt.Print("\033[H\033[2J") // Clear screen
			fmt.Println(ui.SuccessStyle.Render("Pomodoro finished! Take a break."))
			fmt.Print("\a")
			break
		}

		// Clear screen and move cursor to top-left
		fmt.Print("\033[H\033[2J")
		renderTimer(int(remaining.Minutes()), int(remaining.Seconds())%60)

		<-ticker.C
	}
}

func renderTimer(min, sec int) {
	timeStr := fmt.Sprintf("%02d:%02d", min, sec)
	asciiArt := ui.RenderBigText(timeStr)

	banner := ui.RenderBanner()

	fmt.Println(banner)
	fmt.Println("")
	fmt.Println(ui.PrimaryStyle.Render(asciiArt))
	fmt.Println("")
	fmt.Println(ui.SecondaryStyle.Render("Press Ctrl+C to stop"))
}

func init() {
	rootCmd.AddCommand(pomodoroCmd)
}
