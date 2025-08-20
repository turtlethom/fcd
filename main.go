package main

import (
	"log"
	"github.com/turtlethom/fcd/internal"
)

func main() {
	config, err := internal.HandleConfig()
	if err != nil {
		log.Fatal(err)
	}

	// for _, shortcut := range SHORTCUTS {
	// 	fmt.Printf("%s -> %s\n", shortcut.Label, shortcut.Path)
	// }
	internal.FlagInit(config)
}
