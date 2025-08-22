package cmd

import (
	"flag"
	"fmt"
	"log"
	"io"
)

func FlagInit(config *Config) {
	flag.Usage = HandleHelp
	addFlag := flag.String("a", "", "Add a shortcut to a directory within $HOME")
	removeFlag := flag.String("r", "", "Remove a shortcut from your saved shortcuts based on a LABEL")
	branchFlag := flag.String("b", "", "Branch to a shortcut directory based on LABEL")
	clearFlag := flag.Bool("c", false, "Clear all saved shortcuts")
	printFlag := flag.Bool("p", false, "Print all saved shortcuts")
	helpFlag := flag.Bool("h", false, "Display usage for 'fcd' command")

	// Change how flag handles errors
	flag.CommandLine.SetOutput(io.Discard) // mute built-in error messages

	// Parse command line args
	flag.Parse()

	// Handle flags
	switch {
	case *addFlag != "":
		if err := HandleAdd(config, addFlag); err != nil {
			log.Fatal(err)
		}
		fmt.Println("Adding:", *addFlag)

	case *removeFlag != "":
		if err := HandleRemove(config, *removeFlag); err != nil {
			log.Fatal(err)
		}
		fmt.Println("Removing:", *removeFlag)

	case *branchFlag != "":
		fmt.Println("Branching:", *branchFlag)

	case *clearFlag:
		if err := HandleClear(config); err != nil {
			log.Fatal(err)
		}
		fmt.Println("Clearing all shortcuts...")

	case *printFlag:
		HandlePrint(config)
		fmt.Println("Printing all shortcut directories...")

	case *helpFlag:
		HandleHelp()

	default:
		fmt.Println("Opening fcd menu...")
		HandleMenu(config)
	}
}

