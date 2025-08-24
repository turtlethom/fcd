package cmd

import (
	"github.com/charmbracelet/lipgloss"
	tea "github.com/charmbracelet/bubbletea"
)

// Core state of Bubble Tea program
type Model struct {
	Cursor   int        // Current item highlighted
	Choices  []Shortcut // Menu options
	Selected int        // Selected option
	Styles   Styles			// Styles rendered to Stderr
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
	s := m.RenderTitle()
	s += m.RenderChoices() + "\n"
	s += m.RenderHelp()
	return s
}
