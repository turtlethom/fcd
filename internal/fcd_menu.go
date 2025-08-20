package internal

import (
	"log"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/turtlethom/fcd/internal/ui"
)

func InitMenu() {
	model := ui.Model{
		Choices: []string{"First", "Second", "Third", "Quit"},
		Cursor: 0,
		Selected: -1,
	}
	program := tea.NewProgram(model)
	if err := program.Start(); err != nil {
		log.Fatal(err)
	}
}
