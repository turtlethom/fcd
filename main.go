/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"log"

	"github.com/turtlethom/fcd/cmd"
	"github.com/turtlethom/fcd/internal"
)

func main() {
	config, err := internal.HandleConfig()
	if err != nil {
		log.Fatal(err)
	}
	cmd.SetConfig(config)
	cmd.Execute()
}
