package main

import (
	// tea "github.com/charmbracelet/bubbletea"
	// "github.com/turtlethom/fcd/ui"
	// "log"
	// "fmt"
	"github.com/turtlethom/fcd/utils"
	"github.com/turtlethom/fcd/internal"
)

var FCD_FILE string = ".fcd.txt"
var LABELS_PATHS map[string]string = utils.ParseLabelsPaths(FCD_FILE)

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
	// for k, v := range LABELS_PATHS {
	//    fmt.Printf("%s -> %s\n", k, v)
	// }
	internal.FlagInit()
	
}
