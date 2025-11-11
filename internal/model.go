package internal

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var docStyle = lipgloss.NewStyle().Margin(2, 1)

// Core state of Bubble Tea program
type Model struct {
	List     list.Model // Menu options
	Selected int        // Selected option
}

// Init Initializes Bubble Tea Model, can return initial commands
func (m Model) Init() tea.Cmd {
	return nil
}

// Update receives current model and a message, returns new model and a command
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	// Handles keyboard input
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
				m.Selected = m.List.Index()
			return m, tea.Quit
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	// Handles window sizing
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.List.SetSize(msg.Width-h, msg.Height-v)
	}

	m.List, cmd = m.List.Update(msg)
	return m, cmd
}

// View controls how UI looks based on model's state
//
// Returns a string that should be rendered to the terminal
func (m Model) View() string {
	return docStyle.Render(m.List.View())
}
