package timer

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/MohakGupta2004/taskgo/internal/audio"
	"github.com/MohakGupta2004/taskgo/internal/ui"
)

// Timer handles the countdown logic and UI
type Timer struct {
	Duration time.Duration
	Title    string
	paused   bool
	stopChan chan struct{}
}

// New creates a new Timer
func New(duration time.Duration, title string) *Timer {
	return &Timer{
		Duration: duration,
		Title:    title,
		stopChan: make(chan struct{}),
	}
}

// Start begins the timer countdown
func (t *Timer) Start() {
	// Disable input buffering to read single keys
	disableInputBuffering()
	defer enableInputBuffering()

	targetTime := time.Now().Add(t.Duration)
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	// Channel to receive key presses
	keyChan := make(chan rune)
	go func() {
		reader := bufio.NewReader(os.Stdin)
		for {
			char, _, err := reader.ReadRune()
			if err == nil {
				keyChan <- char
			}
		}
	}()

	// Initial render
	t.render(t.Duration)

	remaining := t.Duration

	for {
		select {
		case <-t.stopChan:
			return
		case key := <-keyChan:
			if key == 'p' || key == 'P' {
				t.paused = !t.paused
				if t.paused {
					// Adjust targetTime when pausing so we don't lose time
					// But actually, when paused, we just stop updating remaining
					// When resuming, we need to recalculate targetTime
				} else {
					// Resuming: reset targetTime based on remaining
					targetTime = time.Now().Add(remaining)
				}
				t.render(remaining)
			} else if key == 'q' || key == 'Q' { // Optional: q to quit
				return
			}
		case <-ticker.C:
			if !t.paused {
				remaining = time.Until(targetTime)
				if remaining <= 0 {
					t.finish()
					return
				}
				t.render(remaining)
			} else {
				// If paused, we just update the UI to show paused state if needed
				// But we already updated on key press.
				// Maybe flash "PAUSED"?
				t.render(remaining)
			}
		}
	}
}

func (t *Timer) render(remaining time.Duration) {
	// Clear screen
	fmt.Print("\033[H\033[2J")

	minutes := int(remaining.Minutes())
	seconds := int(remaining.Seconds()) % 60
	timeStr := fmt.Sprintf("%02d:%02d", minutes, seconds)

	asciiArt := ui.RenderBigText(timeStr)
	banner := ui.RenderBanner()

	fmt.Println(banner)
	fmt.Println("")
	fmt.Println(ui.RenderTitle(t.Title))
	fmt.Println("")

	if t.paused {
		fmt.Println(ui.WarningStyle.Render(asciiArt))
		fmt.Println("")
		fmt.Println(ui.WarningStyle.Render("   [ PAUSED ]   "))
	} else {
		fmt.Println(ui.PrimaryStyle.Render(asciiArt))
		fmt.Println("")
	}

	fmt.Println("")
	fmt.Println(ui.SecondaryStyle.Render("Press 'p' to pause/resume, Ctrl+C to exit"))
}

func (t *Timer) finish() {
	fmt.Print("\033[H\033[2J") // Clear screen
	fmt.Println(ui.SuccessStyle.Render("🎉 Timer finished! 🎉"))
	fmt.Println("")
	audio.PlayMultipleBeeps(3)
}

// Helper functions for raw mode
func disableInputBuffering() {
	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
	exec.Command("stty", "-F", "/dev/tty", "-echo").Run()
}

func enableInputBuffering() {
	exec.Command("stty", "-F", "/dev/tty", "-cbreak").Run()
	exec.Command("stty", "-F", "/dev/tty", "echo").Run()
}
