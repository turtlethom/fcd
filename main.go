package main

import (
	"log"
	"github.com/turtlethom/fcd/internal/cmd"
)

func main() {
	config, err := cmd.HandleConfig()
	if err != nil {
		log.Fatal(err)
	}

	// for _, shortcut := range SHORTCUTS {
	// 	fmt.Printf("%s -> %s\n", shortcut.Label, shortcut.Path)
	// }
	cmd.FlagInit(config)
}
