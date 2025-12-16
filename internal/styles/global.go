package styles

import "github.com/charmbracelet/lipgloss"

// Styles holds all the styles for the bubbles/list component
type Styles struct {
	Primary   lipgloss.Color
	Secondary lipgloss.Color
	Tertiary  lipgloss.Color

	NormalTitle    lipgloss.Style
	NormalDesc     lipgloss.Style
	SelectedTitle  lipgloss.Style
	SelectedDesc   lipgloss.Style
	MainTitle      lipgloss.Style
	StatusBar      lipgloss.Style
	StatusBarCount lipgloss.Style
	FilterMatch    lipgloss.Style
	FilterPrompt   lipgloss.Style
	FilterCursor   lipgloss.Style
}
