package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

func HandleHelp() {
	fmt.Fprintln(os.Stderr, "Usage:")
	fmt.Fprintln(os.Stderr, "  fcd                  Search saved shortcuts")
	fmt.Fprintln(os.Stderr, "  fcd -a <PATH>        Add a shortcut")
	fmt.Fprintln(os.Stderr, "  fcd -r <LABEL>       Remove a shortcut")
	fmt.Fprintln(os.Stderr, "  fcd -b <LABEL>       Branch to a shortcut")
	fmt.Fprintln(os.Stderr, "  fcd -c               Clear all shortcuts")
	fmt.Fprintln(os.Stderr, "  fcd -p               Print all shortcuts")
}

func HandleAdd(config *Config, addFlag *string) error {
	var label, path string
	// If arg value is "LABEL:PATH"
	if strings.Contains(*addFlag, ":") {
		parts := strings.SplitN(*addFlag, ":", 2)
		label = parts[0]
		path = parts[1]
		// If only "PATH" is provided
	} else {
		path = *addFlag
		label = filepath.Base(path)
	}

	label = strings.ToUpper(strings.TrimSpace(label))
	if err := config.AddShortcut(label, path); err != nil {
		return err
	}
	return Save(config)
}

// Removes file based on provided label
func HandleRemove(config *Config, removeFlag string) error {
	removeLabel := strings.ToUpper(strings.TrimSpace(removeFlag))
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
		return fmt.Errorf("no shortcut found with label %q", removeLabel)
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
	model := Model{
		Choices:  config.Shortcuts,
		Cursor:   0,
		Selected: -1,
	}

	program := tea.NewProgram(model, tea.WithOutput(os.Stderr))
	finalModel, err := program.StartReturningModel()
	if err != nil {
		log.Fatal(err)
	}

	m := finalModel.(Model)

	if m.Selected >= 0 {
		return m.Choices[m.Selected].Path // return path without printing
	}
	return ""
}
