package ui

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	PrimaryColor   = lipgloss.Color("#007BFF")
	SecondaryColor = lipgloss.Color("#6C757D")
	SuccessColor   = lipgloss.Color("#28A745")
	WarningColor   = lipgloss.Color("#FFC107")
	ErrorColor     = lipgloss.Color("#DC3545")
	LightGray      = lipgloss.Color("#F8F9FA")
	DarkGray       = lipgloss.Color("#343A40")
	OrangeColor    = lipgloss.Color("#FFA500")

	PrimaryStyle   = lipgloss.NewStyle().Foreground(PrimaryColor)
	SecondaryStyle = lipgloss.NewStyle().Foreground(SecondaryColor)
	SuccessStyle   = lipgloss.NewStyle().Foreground(SuccessColor)
	WarningStyle   = lipgloss.NewStyle().Foreground(WarningColor)
	ErrorStyle     = lipgloss.NewStyle().Foreground(ErrorColor)
	InfoStyle      = lipgloss.NewStyle().Foreground(PrimaryColor)

	TitleStyle = lipgloss.NewStyle().
			Foreground(PrimaryColor).
			Bold(true).
			Padding(1, 2).
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(PrimaryColor)

	TableHeaderStyle = lipgloss.NewStyle().
				Foreground(DarkGray).
				Background(LightGray).
				Bold(true).
				Padding(0, 1)

	TableCellStyle = lipgloss.NewStyle().
			Padding(0, 1)

	StatusTodoStyle = lipgloss.NewStyle().
			Foreground(OrangeColor)

	StatusInProgressStyle = lipgloss.NewStyle().
				Foreground(PrimaryColor)

	StatusCompletedStyle = lipgloss.NewStyle().
				Foreground(SuccessColor)

	PendingRowStyle = lipgloss.NewStyle().
			Foreground(OrangeColor)

	InProgressRowStyle = lipgloss.NewStyle().
				Foreground(PrimaryColor)

	CompletedRowStyle = lipgloss.NewStyle().
				Foreground(SuccessColor)

	TreeBranchStyle = lipgloss.NewStyle().
			Foreground(SecondaryColor).
			Bold(true)

	BannerStyle = lipgloss.NewStyle().
			Foreground(PrimaryColor).
			Bold(true).
			MarginBottom(1)
)

func RenderTitle(title string) string {
	return TitleStyle.Render(title)
}

func RenderBanner() string {
	banner := `
  _______        _      _____       
 |__   __|      | |    / ____|      
    | | __ _ ___| | __| |  __  ___  
    | |/ _` + "`" + ` / __| |/ /| | |_ |/ _ \ 
    | | (_| \__ \   < | |__| | (_) |
    |_|\__,_|___/_|\_\ \_____|\___/ 
`
	return BannerStyle.Render(banner)
}
