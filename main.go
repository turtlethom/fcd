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
	cmd.HandleFlags(config)
}
