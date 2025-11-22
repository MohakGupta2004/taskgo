package audio

import (
	"encoding/base64"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

// Base64-encoded notification sound (simple beep WAV file)
// This is a 880Hz beep for 0.25 seconds (8kHz sample rate)
const notificationSoundBase64 = `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`

// PlayNotificationSound plays a notification sound when the Pomodoro timer ends.
// It tries multiple methods in order of preference:
// 1. paplay (PulseAudio)
// 2. aplay (ALSA)
// 3. ffplay (FFmpeg)
// 4. Terminal bell (fallback)
func PlayNotificationSound() {
	// Decode the base64 sound data
	soundData, err := base64.StdEncoding.DecodeString(notificationSoundBase64)
	if err != nil {
		// Fallback to terminal bell
		fmt.Print("\a")
		return
	}

	// Create a temporary file for the sound
	tmpDir := os.TempDir()
	soundFile := filepath.Join(tmpDir, "taskgo-notification.wav")

	// Write the sound data to the temp file
	if err := os.WriteFile(soundFile, soundData, 0644); err != nil {
		// Fallback to terminal bell
		fmt.Print("\a")
		return
	}
	defer os.Remove(soundFile) // Clean up after playing

	// Try different audio players in order of preference
	players := []struct {
		cmd  string
		args []string
	}{
		{"paplay", []string{soundFile}},
		{"aplay", []string{soundFile}},
		{"ffplay", []string{"-nodisp", "-autoexit", "-hide_banner", "-loglevel", "quiet", soundFile}},
	}

	played := false
	for _, player := range players {
		if _, err := exec.LookPath(player.cmd); err == nil {
			// Player is available, try to use it
			cmd := exec.Command(player.cmd, player.args...)
			if err := cmd.Run(); err == nil {
				played = true
				break
			}
		}
	}

	// If no player worked, use terminal bell as final fallback
	if !played {
		fmt.Print("\a")
	}
}

// PlayMultipleBeeps plays multiple notification sounds for emphasis
func PlayMultipleBeeps(count int) {
	for i := 0; i < count; i++ {
		PlayNotificationSound()
		if i < count-1 {
			// Small delay between beeps
			exec.Command("sleep", "0.3").Run()
		}
	}
}
