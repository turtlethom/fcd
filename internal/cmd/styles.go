package cmd

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"
)

var (
	PRIMARY   = "#FF007F"
	SECONDARY = "#FFFFFF"
	DIMMED    = "#808080"

	// Styling For FCD Menu Title
	mainTitleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(SECONDARY)).
			Background(lipgloss.Color(PRIMARY)).
			Bold(true).
			Padding(0, 1)
	// Styling For Unselected/Normal Items
	normalItemStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(SECONDARY))
	// normalTitleStyle = lipgloss.NewStyle().
	// 		Foreground(lipgloss.Color(SECONDARY))
	// Styling For Currently Selected Item
	selectedItemStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(PRIMARY)).
			BorderStyle(lipgloss.NormalBorder()).
			BorderLeft(true).
			BorderForeground(lipgloss.Color(PRIMARY)).
			Padding(0, 1)
)

func handleDelegate() list.DefaultDelegate {
	delegate := list.NewDefaultDelegate()
	// Normal configuration
	delegate.Styles.NormalTitle = normalItemStyle
	delegate.Styles.NormalDesc = normalItemStyle
	// Selected configuration
	delegate.Styles.SelectedTitle = selectedItemStyle
	delegate.Styles.SelectedDesc = selectedItemStyle
	
	// Filter configuration
	delegate.Styles.DimmedTitle = lipgloss.NewStyle().Foreground(lipgloss.Color(PRIMARY))
	delegate.Styles.DimmedDesc = lipgloss.NewStyle().Foreground(lipgloss.Color(SECONDARY))
	delegate.Styles.FilterMatch = lipgloss.NewStyle().Foreground(lipgloss.Color(PRIMARY)).Underline(true)

	return delegate
}

func CreateBubblesList(items []list.Item) list.Model {
	delegate := handleDelegate()
	l := list.New(items, delegate, 0, 10)

	// Clear out default list title styles
	l.Styles.Title = lipgloss.NewStyle()
	l.Styles.TitleBar = lipgloss.NewStyle()

	// Ensure the status bar doesn't override styles
	l.Styles.StatusBar = lipgloss.NewStyle().
		Foreground(lipgloss.Color(DIMMED))
	l.Styles.StatusBarActiveFilter = lipgloss.NewStyle().
		Foreground(lipgloss.Color(DIMMED))
	l.Styles.StatusBarFilterCount = lipgloss.NewStyle().
		Foreground(lipgloss.Color(DIMMED))

	// Manages Styles For Filter Menu
	l.FilterInput.PromptStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color(PRIMARY)).
		Bold(true)
	l.FilterInput.Cursor.Style = lipgloss.NewStyle().
		Foreground(lipgloss.Color(PRIMARY)).
		Background(lipgloss.Color(SECONDARY))

	// Manages The Title Of The List
	l.Title = mainTitleStyle.Render("FCD - Shortcut Menu")

	return l
}
