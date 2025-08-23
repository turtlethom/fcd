package cmd

import (
	"flag"
	"fmt"
	"io"
	"os"
)

func HandleFlags(config *Config) {
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
			fmt.Fprintln(os.Stderr, "Error adding shortcut:", err)
			os.Exit(1)
		}
		fmt.Fprintln(os.Stderr, "Successfully added:", *addFlag)

	case *removeFlag != "":
		if err := HandleRemove(config, *removeFlag); err != nil {
			fmt.Fprintln(os.Stderr, "Error removing shortcut:", err)
			os.Exit(1)
		}
		fmt.Fprintln(os.Stderr, "Successfully removed:", *removeFlag)

	case *branchFlag != "":
		fmt.Fprintln(os.Stderr, "Successful branch:")

	case *clearFlag:
		if err := HandleClear(config); err != nil {
			fmt.Fprintln(os.Stderr, "Error clearing shortcuts:", err)
			os.Exit(1)
		}
		fmt.Fprintln(os.Stderr, "Successfully cleared all shortcuts")

	case *printFlag:
		HandlePrint(config)
		fmt.Fprintln(os.Stderr, "Printing all shortcut directories...")

	case *helpFlag:
		HandleHelp()

	default:
		selectedPath := HandleMenu(config)
		if selectedPath != "" {
			fmt.Println(selectedPath)
		}
	}
}

