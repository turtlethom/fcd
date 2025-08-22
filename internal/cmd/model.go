package cmd

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Core state of Bubble Tea program
type Model struct {
	Cursor   int        // Current item highlighted
	Choices  []Shortcut // Menu options
	Selected int        // Selected option
}

var (
	normalBox = lipgloss.NewStyle().
			Border(lipgloss.NormalBorder())
	cursorBox = normalBox.Copy().
			BorderForeground(lipgloss.Color("205")).
			Bold(true)
	selectedBox = normalBox.Copy().
			BorderForeground(lipgloss.Color("229")).
			Bold(true).
			Background(lipgloss.Color("57"))
)

// Initializes Bubble Tea, can return initial commands
func (m Model) Init() tea.Cmd {
	return nil
}

// Receives current model and a message, returns new model and a command
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
			if m.Cursor < len(m.Choices)-1 {
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
		symbol := ") " // default unselected
		if m.Cursor == i {
			symbol = "*) " // hovered
		}
		line := fmt.Sprintf("%s%s", symbol, choice.Label)
		// Show path only if hovered
		if m.Cursor == i {
			pathStyle := lipgloss.NewStyle().
				Foreground(lipgloss.Color("240"))
			line += fmt.Sprintf(" → %s", pathStyle.Render(choice.Path))
		}
		s += line + "\n"
	}
	s += "\n(j/k) to move, Enter to select, q to quit.\n"
	return s
}
