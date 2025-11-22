package cmd

import (
	"fmt"
	"strconv"
	"time"

	"github.com/MohakGupta2004/taskgo/internal/timer"
	"github.com/MohakGupta2004/taskgo/internal/ui"
	"github.com/spf13/cobra"
)

var pomodoroCmd = &cobra.Command{
	Use:   "pomodoro [duration]",
	Short: "Start a pomodoro timer (default 25m)",
	Long:  `Start a pomodoro timer. You can specify the duration using Go's time format (e.g., 25m, 1h, 1h30m). Default is 25m.`,
	Args:  cobra.MaximumNArgs(1),
	Run:   runPomodoro,
}

func runPomodoro(cmd *cobra.Command, args []string) {
	duration := 25 * time.Minute

	if len(args) > 0 {
		input := args[0]
		// Try parsing as duration (e.g., "25m", "1h")
		d, err := time.ParseDuration(input)
		if err == nil {
			duration = d
		} else {
			// Fallback: try parsing as integer minutes for backward compatibility
			if m, err := strconv.Atoi(input); err == nil {
				duration = time.Duration(m) * time.Minute
			} else {
				fmt.Println(ui.ErrorStyle.Render("Invalid time format. Use format like 25m, 1h, 1h30m"))
				return
			}
		}
	}

	t := timer.New(duration, fmt.Sprintf("Pomodoro (%s)", duration.String()))
	t.Start()
}

func init() {
	rootCmd.AddCommand(pomodoroCmd)
}
