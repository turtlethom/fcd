package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/termenv"
)

var (
	selectedItemStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FF007F")). // white text
			BorderStyle(lipgloss.NormalBorder()).
			BorderLeft(true).
			BorderForeground(lipgloss.Color("#FF007F")).
			Padding(0, 1)
	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Padding(0, 1).
			Foreground(lipgloss.Color("#FFFFFF")).
			Background(lipgloss.Color("#FF007F"))
	normalItemStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#AAAAAA")) // gray text
)

func HandleHelp() {
	fmt.Fprintln(os.Stderr, "fcd - fast change directory")
	fmt.Fprintln(os.Stderr, "  - Change into bookmarked directories using custom labels")
	fmt.Fprintln(os.Stderr, "  - Limited to directories within a user's home directory")
	fmt.Fprintln(os.Stderr)
	fmt.Fprintln(os.Stderr, "Usage:")
	fmt.Fprintln(os.Stderr, "  fcd                   Search saved shortcuts")
	fmt.Fprintln(os.Stderr, "  fcd add <PATH>        Add a shortcut")
	fmt.Fprintln(os.Stderr, "  fcd remove <LABEL>    Remove a shortcut")
	fmt.Fprintln(os.Stderr, "  fcd branch <LABEL>    Branch to a shortcut")
	fmt.Fprintln(os.Stderr, "  fcd clear             Clear all shortcuts")
	fmt.Fprintln(os.Stderr, "  fcd print             Print all shortcuts")
}

func HandleAdd(config *Config, commandArg string) error {
	var label, path string
	// If arg value has "LABEL:PATH"
	if strings.Contains(commandArg, ":") {
		parts := strings.SplitN(commandArg, ":", 2)
		label = parts[0]
		path = parts[1]
		// If only "PATH" is provided
	} else {
		path = commandArg
		label = filepath.Base(path)
	}

	label = strings.ToUpper(strings.TrimSpace(label))
	if err := config.AddShortcut(label, path); err != nil {
		return err
	}
	return Save(config)
}

// Removes file based on provided label
func HandleRemove(config *Config, commandArg string) error {
	removeLabel := strings.ToUpper(strings.TrimSpace(commandArg))
	newShortcuts := make([]Shortcut, 0, len(config.Shortcuts))
	found := false
	for _, sc := range config.Shortcuts {
		if sc.Label == removeLabel {
			found = true
			continue
		}
		newShortcuts = append(newShortcuts, sc)
	}
	if !found {
		return fmt.Errorf("fcd: no shortcut found with label %q", removeLabel)
	}
	config.Shortcuts = newShortcuts
	return Save(config)
}

func HandleClear(config *Config) error {
	config.Shortcuts = []Shortcut{}
	return Save(config)
}

func HandlePrint(config *Config) {
	for _, sc := range config.Shortcuts {
		fmt.Fprintf(os.Stderr, "%-10s %s\n", sc.Label, sc.Path)
	}
}

func HandleMenu(config *Config) string {
	lipgloss.SetColorProfile(termenv.TrueColor)
	items := make([]list.Item, len(config.Shortcuts))
	for i, sc := range config.Shortcuts {
		items[i] = sc
	}

	delegate := list.NewDefaultDelegate()
	// Custom Styles
	delegate.Styles.SelectedTitle = selectedItemStyle
	delegate.Styles.SelectedDesc = selectedItemStyle
	delegate.Styles.NormalTitle = normalItemStyle
	delegate.Styles.NormalDesc = normalItemStyle

	l := list.New(items, delegate, 0, 10)

	// Clear out default list title styles
	l.Styles.Title = lipgloss.NewStyle()
	l.Styles.TitleBar = lipgloss.NewStyle()
	l.Title = titleStyle.Render("FCD - Shortcut Menu")

	model := Model{
		List: l,
	}

	program := tea.NewProgram(
		model,
		tea.WithOutput(os.Stderr),
		tea.WithAltScreen(),
	)

	finalModel, err := program.Run()
	if err != nil {
		log.Fatal(err)
	}

	m := finalModel.(Model)
	if sel, ok := m.List.SelectedItem().(Shortcut); ok {
		return sel.Path
	}
	return ""
}
