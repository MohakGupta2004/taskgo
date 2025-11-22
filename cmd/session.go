package cmd

import (
	"fmt"
	"time"

	"github.com/MohakGupta2004/taskgo/internal/timer"
	"github.com/MohakGupta2004/taskgo/internal/ui"
	"github.com/spf13/cobra"
)

var sessionCmd = &cobra.Command{
	Use:   "session [duration]",
	Short: "Start a pomodoro session (alternating work/break)",
	Long:  `Start a pomodoro session that alternates between 25m work and 5m break intervals until the specified duration is reached. Default duration is 2h.`,
	Args:  cobra.MaximumNArgs(1),
	Run:   runSession,
}

func runSession(cmd *cobra.Command, args []string) {
	totalDuration := 2 * time.Hour

	if len(args) > 0 {
		d, err := time.ParseDuration(args[0])
		if err != nil {
			fmt.Println(ui.ErrorStyle.Render("Invalid time format. Use format like 2h, 1h30m"))
			return
		}
		totalDuration = d
	}

	startTime := time.Now()
	endTime := startTime.Add(totalDuration)

	workDuration := 25 * time.Minute
	breakDuration := 5 * time.Minute

	for {
		if time.Now().After(endTime) {
			break
		}

		// Work Session
		remaining := time.Until(endTime)
		currentWorkDuration := workDuration
		if remaining < workDuration {
			currentWorkDuration = remaining
		}

		if currentWorkDuration > 0 {
			t := timer.New(currentWorkDuration, "Session: Work")
			t.Start()
		}

		if time.Now().After(endTime) {
			break
		}

		// Break Session
		remaining = time.Until(endTime)
		currentBreakDuration := breakDuration
		if remaining < breakDuration {
			currentBreakDuration = remaining
		}

		if currentBreakDuration > 0 {
			t := timer.New(currentBreakDuration, "Session: Break")
			t.Start()
		}
	}

	fmt.Println(ui.SuccessStyle.Render("🎉 Session completed! 🎉"))
}

func init() {
	rootCmd.AddCommand(sessionCmd)
}
