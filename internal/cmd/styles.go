package cmd

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"
)

var (
	PRIMARY   = "#FF007F"
	SECONDARY = "#FFFFFF"

	selectedItemStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(PRIMARY)).
			BorderStyle(lipgloss.NormalBorder()).
			BorderLeft(true).
			BorderForeground(lipgloss.Color(PRIMARY)).
			Padding(0, 1)
	titleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(SECONDARY)).
			Background(lipgloss.Color(PRIMARY)).
			Bold(true).
			Padding(0, 1)
	normalItemStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(SECONDARY))
)

func handleDelegate() list.DefaultDelegate {
	delegate := list.NewDefaultDelegate()
	// Custom Styles
	delegate.Styles.SelectedTitle = selectedItemStyle
	delegate.Styles.SelectedDesc = selectedItemStyle
	delegate.Styles.NormalTitle = normalItemStyle
	delegate.Styles.NormalDesc = normalItemStyle
	return delegate
}

func CreateBubblesList(items []list.Item) list.Model {
	delegate := handleDelegate()
	l := list.New(items, delegate, 0, 10)

	// Clear out default list title styles
	l.Styles.Title = lipgloss.NewStyle()
	l.Styles.TitleBar = lipgloss.NewStyle()
	l.Title = titleStyle.Render("FCD - Shortcut Menu")
	return l
}
