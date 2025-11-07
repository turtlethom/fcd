package cmd

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"
)

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

func NewStyles(config *Config) Styles {
	// Sets colors of the menu based on user configuration
	primary := lipgloss.Color(config.UserColors.Primary)
	secondary := lipgloss.Color(config.UserColors.Secondary)
	tertiary := lipgloss.Color(config.UserColors.Tertiary)

	NORMAL_PRESETS := lipgloss.NewStyle().
		BorderStyle(lipgloss.NormalBorder()).
		BorderLeft(true).
		BorderForeground(secondary)

	SELECTED_PRESETS := lipgloss.NewStyle().
		BorderStyle(lipgloss.NormalBorder()).
		BorderLeft(true).
		BorderForeground(primary)

	return Styles{
		Primary:   primary,
		Secondary: secondary,
		Tertiary:  tertiary,
		// Styling for FCD Main Menu Title
		MainTitle: lipgloss.NewStyle().
			Foreground(secondary).
			Background(primary).
			Bold(true).
			Padding(0, 1),
		// Styling for normal titles of items in list component
		NormalTitle: lipgloss.NewStyle().
			Inherit(NORMAL_PRESETS).
			Foreground(primary).
			Padding(0, 1),
		// Styling for normal descriptions of items in list component
		NormalDesc: lipgloss.NewStyle().
			Inherit(NORMAL_PRESETS).
			Foreground(secondary).
			Padding(0, 1),
		// Styling for title of selected item in list component
		SelectedTitle: lipgloss.NewStyle().
			Inherit(SELECTED_PRESETS).
			Foreground(secondary).
			Background(primary).
			Padding(0, 1).
			Bold(true),
		// Styling for description of selected item in list component
		SelectedDesc: lipgloss.NewStyle().
			Inherit(SELECTED_PRESETS).
			Foreground(secondary).
			Underline(true).
			Padding(0, 1),
		// Styling for "Filter:"
		StatusBar: lipgloss.NewStyle().
			Foreground(secondary).
			MarginBottom(1).
			Padding(0, 1),
		// Styling for status bar count
		StatusBarCount: lipgloss.NewStyle().
			Foreground(primary),
		// Styling for matching text to filter
		FilterMatch: lipgloss.NewStyle().
			Foreground(secondary).
			Underline(true),
		FilterPrompt: lipgloss.NewStyle().
			Foreground(primary).
			Bold(true),
		FilterCursor: lipgloss.NewStyle().
			Foreground(primary).
			Background(secondary),
	}
}	
// Responsible for rendering each item in the FCD Menu
func newDelegate(s Styles) list.DefaultDelegate {
	delegate := list.NewDefaultDelegate()
	// Normal configuration
	delegate.Styles.NormalTitle = s.NormalTitle
	delegate.Styles.NormalDesc = s.NormalDesc
	// Selected configuration
	delegate.Styles.SelectedTitle = s.SelectedTitle
	delegate.Styles.SelectedDesc = s.SelectedDesc
	// Filter configuration
	delegate.Styles.FilterMatch = s.FilterMatch

	return delegate
}

// Create the FCD Menu with configuration
func CreateBubblesList(items []list.Item, s Styles) list.Model {
	// Configured view of each item
	delegate := newDelegate(s)

	// Instantiating a Bubbles List
	l := list.New(items, delegate, 0, 10)

	// Clear out default list title styles
	l.Styles.Title = lipgloss.NewStyle()
	l.Styles.TitleBar = lipgloss.NewStyle()

	// Custom status bar styles
	l.Styles.StatusBar = s.StatusBar
	// Styles unfiltered items in filter menu
	l.Styles.StatusBarActiveFilter = s.StatusBarCount
	// Styles the right portion of items count
	l.Styles.StatusBarFilterCount = s.StatusBarCount
	// Manages Styles For Filter Menu
	l.FilterInput.PromptStyle = s.FilterPrompt 
	l.FilterInput.Cursor.Style = s.FilterCursor
	// Manages The Title Of The List
	l.Title = s.MainTitle.Render("FCD - Shortcut Menu")

	return l
}
