package cmd

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"
)

var (
	// Colors For FCD Menu
	// TODO: Allow configurable colors for menu
	PRIMARY   = lipgloss.Color("#FF007F")
	SECONDARY = lipgloss.Color("#FFFFFF")

	// Styles Global To All Normal Styles
	NORMAL_PRESETS = lipgloss.NewStyle().
			BorderStyle(lipgloss.NormalBorder()).
			BorderLeft(true).
			BorderForeground(SECONDARY)
	// Styles Global To All Selected Styles
	SELECTED_PRESETS = lipgloss.NewStyle().
			BorderStyle(lipgloss.NormalBorder()).
			BorderLeft(true).
			BorderForeground(PRIMARY)

	// Styling For FCD Menu Title
	mainTitleStyle = lipgloss.NewStyle().
			Foreground(SECONDARY).
			Background(PRIMARY).

			Bold(true).
			Padding(0, 1)

	// Styling For Unselected/Normal Title
	normalTitleStyle = lipgloss.NewStyle().
			Inherit(NORMAL_PRESETS).

			Foreground(PRIMARY).
			Padding(0, 1)

	// Styling For Unselected/Normal Descriptions
	normalDescStyle = lipgloss.NewStyle().
			Inherit(NORMAL_PRESETS).

			Foreground(SECONDARY).
			Padding(0, 1)

	// Styling For Selected Title
	selectedTitleStyle = lipgloss.NewStyle().
			Inherit(SELECTED_PRESETS).

			Foreground(SECONDARY).
			Background(PRIMARY).
			Padding(0, 1).
			Bold(true)

	// Styling For Selected Description
	selectedDescStyle = lipgloss.NewStyle().
			Inherit(SELECTED_PRESETS).

			Foreground(SECONDARY).
			Underline(true).
			Padding(0, 1)
)

// Responsible for rendering each item in the FCD Menu
func handleDelegate() list.DefaultDelegate {
	delegate := list.NewDefaultDelegate()
	// Normal configuration
	delegate.Styles.NormalTitle = normalTitleStyle
	delegate.Styles.NormalDesc = normalDescStyle
	// Selected configuration
	delegate.Styles.SelectedTitle = selectedTitleStyle
	delegate.Styles.SelectedDesc = selectedDescStyle
	// Filter configuration
	delegate.Styles.FilterMatch = lipgloss.NewStyle().
		Foreground(SECONDARY).
		Underline(true)

	return delegate
}

// Create the FCD Menu with configuration
func CreateBubblesList(items []list.Item) list.Model {
	// Configured view of each item
	delegate := handleDelegate()

	// Instantiating a Bubbles List
	l := list.New(items, delegate, 0, 10)

	// Clear out default list title styles
	l.Styles.Title = lipgloss.NewStyle()
	l.Styles.TitleBar = lipgloss.NewStyle()

	// Custom status bar styles
	l.Styles.StatusBar = lipgloss.NewStyle().
		Foreground(SECONDARY).
		MarginBottom(1).
		Padding(0, 1)
	// Styles unfiltered items in filter menu
	l.Styles.StatusBarActiveFilter = lipgloss.NewStyle().
		Foreground(SECONDARY)
	// Styles the right portion of items count
	l.Styles.StatusBarFilterCount = lipgloss.NewStyle().
		Foreground(PRIMARY)

	// Manages Styles For Filter Menu
	l.FilterInput.PromptStyle = lipgloss.NewStyle().
		Foreground(PRIMARY).
		Bold(true)
	l.FilterInput.Cursor.Style = lipgloss.NewStyle().
		Foreground(PRIMARY).
		Background(SECONDARY)

	// Manages The Title Of The List
	l.Title = mainTitleStyle.Render("FCD - Shortcut Menu")

	return l
}
