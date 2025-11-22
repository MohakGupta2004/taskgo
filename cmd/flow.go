package cmd

import (
	"fmt"
	"os/exec"
	"runtime"
	"strings"
	"time"

	"github.com/MohakGupta2004/taskgo/internal/flow"
	"github.com/MohakGupta2004/taskgo/internal/timer"
	"github.com/MohakGupta2004/taskgo/internal/ui"
	"github.com/spf13/cobra"
)

var flowCmd = &cobra.Command{
	Use:   "flow",
	Short: "Manage focused work flows",
	Long:  `Create and run focused work flows with associated resources (links, apps).`,
}

var flowCreateCmd = &cobra.Command{
	Use:   "create [name]",
	Short: "Create a new flow",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		m, err := flow.NewManager()
		if err != nil {
			fmt.Println(ui.ErrorStyle.Render("Error initializing flow manager: " + err.Error()))
			return
		}

		if err := m.Create(name); err != nil {
			fmt.Println(ui.ErrorStyle.Render("Error creating flow: " + err.Error()))
			return
		}

		fmt.Println(ui.SuccessStyle.Render(fmt.Sprintf("Flow '%s' created successfully!", name)))
	},
}

var flowAddCmd = &cobra.Command{
	Use:   "add [name] [resource...]",
	Short: "Add resources to a flow",
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		resources := args[1:]

		m, err := flow.NewManager()
		if err != nil {
			fmt.Println(ui.ErrorStyle.Render("Error initializing flow manager: " + err.Error()))
			return
		}

		for _, res := range resources {
			if err := m.AddResource(name, res); err != nil {
				fmt.Println(ui.ErrorStyle.Render(fmt.Sprintf("Error adding resource '%s': %s", res, err.Error())))
				return
			}
		}

		fmt.Println(ui.SuccessStyle.Render(fmt.Sprintf("Added %d resources to flow '%s'", len(resources), name)))
	},
}

var zenMode bool

var flowRunCmd = &cobra.Command{
	Use:   "run [name]",
	Short: "Run a flow session",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		m, err := flow.NewManager()
		if err != nil {
			fmt.Println(ui.ErrorStyle.Render("Error initializing flow manager: " + err.Error()))
			return
		}

		f, err := m.Get(name)
		if err != nil {
			fmt.Println(ui.ErrorStyle.Render("Error loading flow: " + err.Error()))
			return
		}

		// Open resources
		fmt.Println(ui.RenderTitle(fmt.Sprintf("Starting Flow: %s", f.Name)))
		if zenMode {
			fmt.Println(ui.WarningStyle.Render("🧘 Entering Zen Mode..."))
		}
		fmt.Println("Opening resources...")

		openResources(f.Resources, zenMode)

		// Start a timer (indefinite or fixed? User didn't specify duration in run command, assuming indefinite or manual stop)
		// For now, let's just use a simple "Flow Session" timer that counts UP or just shows active state.
		// But our timer package is countdown. Let's use a default long duration like 4 hours for now, or update timer to support count up.
		// Given the constraints, let's start a 4h timer as a "session".
		// User said "taskgo flow run coding_flow", no duration.
		// Let's assume a standard work block or just open resources.
		// But user mentioned "I can break it and pause it".
		// Let's use a generic timer for now.

		// TODO: Maybe ask for duration or default to open-ended?
		// For this iteration, I'll default to 1 hour to show the UI, as `timer` requires duration.
		// Or I can modify `timer` to support "Stopwatch" mode.
		// Let's stick to a default duration for now to reuse `timer`.

		t := timer.New(4*time.Hour, fmt.Sprintf("Flow: %s", f.Name))
		t.Start()
	},
}

var flowListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all flows",
	Run: func(cmd *cobra.Command, args []string) {
		m, err := flow.NewManager()
		if err != nil {
			fmt.Println(ui.ErrorStyle.Render("Error initializing flow manager: " + err.Error()))
			return
		}

		flows := m.List()
		if len(flows) == 0 {
			fmt.Println("No flows found. Create one with 'taskgo flow create <name>'")
			return
		}

		fmt.Println(ui.RenderTitle("Available Flows"))
		for _, f := range flows {
			fmt.Println("- " + f)
		}
	},
}

func openResources(resources []string, zen bool) {
	if len(resources) == 0 {
		return
	}

	var urls []string
	var apps []string

	// Separate URLs and Apps
	for _, res := range resources {
		// Simple heuristic: if it contains "://" or starts with "www." or has a dot and no spaces, treat as URL
		// Otherwise treat as app
		if strings.Contains(res, "://") || strings.HasPrefix(res, "www.") || (strings.Contains(res, ".") && !strings.Contains(res, " ")) {
			urls = append(urls, res)
		} else {
			apps = append(apps, res)
		}
	}

	// Launch Apps
	for _, app := range apps {
		path, err := exec.LookPath(app)
		if err == nil {
			fmt.Printf("Launching app: %s...\n", app)
			exec.Command(path).Start()
		} else {
			fmt.Printf("Could not find app: %s\n", app)
		}
	}

	if len(urls) == 0 {
		return
	}

	if zen {
		// Try to detect default browser
		browser := ""
		if runtime.GOOS == "linux" {
			out, err := exec.Command("xdg-settings", "get", "default-web-browser").Output()
			if err == nil {
				desktopFile := strings.TrimSpace(string(out))
				// Remove .desktop suffix if present
				browser = strings.TrimSuffix(desktopFile, ".desktop")
			}
		}

		browsers := []string{}
		if browser != "" {
			browsers = append(browsers, browser)
		}
		// Add fallbacks
		browsers = append(browsers, "google-chrome", "chromium", "brave-browser", "firefox")

		for _, b := range browsers {
			if _, err := exec.LookPath(b); err == nil {
				args := []string{}
				if strings.Contains(b, "firefox") {
					args = append([]string{"--new-window"}, urls...)
				} else {
					// Chrome/Chromium/Brave/Edge
					args = append([]string{"--new-window", "--start-fullscreen"}, urls...)
				}

				fmt.Printf("Opening %d URLs in %s...\n", len(urls), b)
				err = exec.Command(b, args...).Start()
				if err == nil {
					return
				}
			}
		}
		fmt.Println("Could not find a supported browser for Zen Mode. Falling back to default.")
	}

	// Normal mode or fallback
	for _, url := range urls {
		fmt.Printf("Opening %s...\n", url)
		var err error
		switch runtime.GOOS {
		case "linux":
			err = exec.Command("xdg-open", url).Start()
		case "windows":
			err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
		case "darwin":
			err = exec.Command("open", url).Start()
		default:
			err = fmt.Errorf("unsupported platform")
		}
		if err != nil {
			fmt.Printf("Error opening %s: %v\n", url, err)
		}
	}
}

func init() {
	rootCmd.AddCommand(flowCmd)
	flowCmd.AddCommand(flowCreateCmd)
	flowCmd.AddCommand(flowAddCmd)
	flowCmd.AddCommand(flowRunCmd)
	flowCmd.AddCommand(flowListCmd)

	flowRunCmd.Flags().BoolVarP(&zenMode, "zen", "z", false, "Run in Zen Mode (Kiosk Mode)")
}
