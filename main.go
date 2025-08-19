package main

import (
	// tea "github.com/charmbracelet/bubbletea"
	// "github.com/turtlethom/fcd/ui"
	// "log"
	"github.com/turtlethom/fcd/utils"
	
)

var FCD_FILE string = ".fcd.txt"

func main() {
	// model := ui.Model{
	// 	Choices: []string{"First", "Second", "Third", "Quit"},
	// 	Cursor: 0,
	// 	Selected: -1,
	// }
	// program := tea.NewProgram(model)
	// if err := program.Start(); err != nil {
	// 	log.Fatal(err)
	// }
	utils.ReadFile(FCD_FILE)
	
}
