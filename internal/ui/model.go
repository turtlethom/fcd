package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"fmt"
)

// Core state of Bubble Tea program
type Model struct {
	Cursor 	 int			// Current item highlighted
	Choices  []string // Menu options
	Selected int			// Selected option
}

// Initializes Bubble Tea, can return initial commands
func (m Model) Init() tea.Cmd {
	return nil
}

// Receives current model and a message, returns new model and a commands
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	// Handles keyboard input
	case tea.KeyMsg:
		switch msg.String() {
		case "k":
			if m.Cursor > 0 {
				m.Cursor--
			}
		case "j":
			if m.Cursor < len(m.Choices) - 1 {
				m.Cursor++
			}
		case "enter":
			m.Selected = m.Cursor
			return m, tea.Quit
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
	case tea.MouseMsg:
	}
	// Default - return the unchanged model with no command
	return m, nil
}

// Returns string that should be rendered to the terminal
// Controls how UI looks based on model's state
func (m Model) View() string {
	s := "Main Menu\n\n"
	for i, choice := range m.Choices {
		cursor := " "
		if m.Cursor == i {
			cursor = ">"
		}
		line := fmt.Sprintf("%s %s", cursor, choice)
		if m.Selected == i {
			line = fmt.Sprintln(" (selected)", line)
		}
		s += line + "\n"
	}
	s += "\n(j/k) to move, Enter to select, q to quit.\n"
	return s
}
