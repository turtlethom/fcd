package internal

import (
	"flag"
	"fmt"
	"io"
	// "strings"
)

func printUsage() {
		fmt.Println("Usage:")
		fmt.Println("  fcd -a <PATH>        Add a shortcut")
		fmt.Println("  fcd -r <LABEL>       Remove a shortcut")
		fmt.Println("  fcd -b <LABEL>       Branch to a shortcut")
		fmt.Println("  fcd -c               Clear all shortcuts")
		fmt.Println("  fcd -p               Print all shortcuts")
}

func FlagInit(config *Config) {
	flag.Usage = printUsage
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
		// parts := strings.SplitN(*addFlag, ":", 2)
		// if len(parts) != 2 {
		// 	fmt.Println("Invalid format.")
		// }
		// config.Shortcuts = append(config.Shortcuts, Shortcut{
		//
		// })
		fmt.Println("Adding:", *addFlag)

	case *removeFlag != "":
		fmt.Println("Removing:", *removeFlag)

	case *branchFlag != "":
		fmt.Println("Branching:", *branchFlag)

	case *clearFlag:
		fmt.Println("Clearing all shortcuts...")

	case *printFlag:
		fmt.Println("Printing all shortcut directories...")

	case *helpFlag:
		printUsage()

	default:
		fmt.Println("Opening fcd menu...")
	}
}

