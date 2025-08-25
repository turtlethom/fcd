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
			fmt.Fprintln(os.Stderr, "fcd: Error adding shortcut:", err)
			os.Exit(1)
		}
		fmt.Fprintln(os.Stderr, "fcd: Saved shortcut '", *addFlag, "'")

	case *removeFlag != "":
		if err := HandleRemove(config, *removeFlag); err != nil {
			fmt.Fprintln(os.Stderr, "fcd: Error removing shortcut:", err)
			os.Exit(1)
		}
		fmt.Fprintln(os.Stderr, "fcd: Successfully removed '", *removeFlag, "'")

	// TODO: Allow to change/branch into a selected directory based on LABEL
	case *branchFlag != "":
		fmt.Fprintln(os.Stderr, "fcd: Successful branch: '", *branchFlag, "'")

	case *clearFlag:
		if err := HandleClear(config); err != nil {
			fmt.Fprintln(os.Stderr, "fcd: Error clearing shortcuts:", err)
			os.Exit(1)
		}
		fmt.Fprintln(os.Stderr, "fcd: Successfully cleared all shortcuts")

	// Print all saved shortcuts from user
	case *printFlag:
		// If no added shortcuts, exit
		if len(config.Shortcuts) <= 0 {
			fmt.Fprintln(os.Stderr, "fcd: No existing shortcuts")
			os.Exit(1)
		}
		fmt.Fprintln(os.Stderr, "fcd: Printing all shortcut directories...")
		HandlePrint(config)

	// Fcd help page
	case *helpFlag:
		HandleHelp()

	default:
		// If no added shortcuts, exit
		if len(config.Shortcuts) <= 0 {
			fmt.Fprintln(os.Stderr, "fcd: No existing shortcuts")
			os.Exit(1)
		}
		// Grabbing path for handling flags
		selectedPath := HandleMenu(config)
		if selectedPath != "" {
			fmt.Println(selectedPath)
		}
	}
}

