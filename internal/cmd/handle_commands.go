package cmd

import (
	"bufio"
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

// Adds bookmarks based on provided LABEL or LABEL:PATH; restricted to $HOME
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

// Removes file based on provided LABEL
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

// 'Branches' to the path that matches the provided LABEL
func HandleBranch(config *Config, commandArg string) string {
	for _, sc := range config.Shortcuts {
		if sc.Label == strings.ToUpper(strings.TrimSpace(commandArg)) {
			return sc.Path
		}
	}
	fmt.Fprintf(os.Stderr, "fcd: Cannot find matching path for %q\n", commandArg)
	return ""
}

func HandleClear(config *Config) error {
	reader := bufio.NewReader(os.Stdin)

	fmt.Fprintf(os.Stderr, "fcd: Are you sure you want to clear all bookmarks? (y/n): ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	switch strings.ToLower(input) {
	case "y":
		config.Shortcuts = []Shortcut{}
		return Save(config)
	case "n":
		fmt.Fprintf(os.Stderr, "fcd: Canceled clearing bookmarks\n")
		return nil
	default:
		fmt.Fprintf(os.Stderr, "fcd: Defaulting to no - operation ceased \n")
		return fmt.Errorf("fcd: Something went wrong")
	}
}

func HandlePrint(config *Config) {
	for _, sc := range config.Shortcuts {
		fmt.Fprintf(os.Stderr, "%-10s %s\n", sc.Label, sc.Path)
	}
}

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

func HandleMenu(config *Config) string {
	// Ensures ANSI colors will be rendered for both stdout and stderr
	lipgloss.SetColorProfile(termenv.TrueColor)
	// Creating list items based on bookmarks
	items := make([]list.Item, len(config.Shortcuts))
	for i, sc := range config.Shortcuts {
		items[i] = sc
	}
	// Initialize styles
	styles := NewStyles(config)
	bubblesList := CreateBubblesList(items, styles)
	// Create Model
	model := Model{
		List: bubblesList,
	}
	// Start bubbletea for FCD menu
	program := tea.NewProgram(
		model,
		tea.WithOutput(os.Stderr),
		tea.WithAltScreen(),
	)
	finalModel, err := program.Run()
	if err != nil {
		log.Fatal(err)
	}
	// Get result of user choice and return it at the end
	m := finalModel.(Model)
	if sel, ok := m.List.SelectedItem().(Shortcut); ok {
		return sel.Path
	}
	return ""
}
