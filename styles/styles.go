package styles

import "github.com/charmbracelet/lipgloss"

var (
	result = lipgloss.Color("245")
)

var terminalStyle = lipgloss.NewStyle().Foreground(result).Bold(true).Align(lipgloss.Center)
