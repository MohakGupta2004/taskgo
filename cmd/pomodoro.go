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

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		remaining := time.Until(targetTime)
		if remaining <= 0 {
			fmt.Print("\r")
			fmt.Println(ui.SuccessStyle.Render("Pomodoro finished! Take a break.   "))
			fmt.Print("\a")
			break
		}

		fmt.Printf("\r%s: %02d:%02d", ui.PrimaryStyle.Render("Time Remaining"), int(remaining.Minutes()), int(remaining.Seconds())%60)
		<-ticker.C
	}
}

func init() {
	rootCmd.AddCommand(pomodoroCmd)
}
