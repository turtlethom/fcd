package internal

import (
	"flag"
	"fmt"
)

func FlagInit() {
	addFlag := flag.String("a", "", "Add a shortcut to a directory within $HOME")
	removeFlag := flag.String("r", "", "Remove a shortcut from your saved shortcuts based on a LABEL")
	branchFlag := flag.String("b", "", "Branch to a shortcut directory based on LABEL")
	clearFlag := flag.Bool("c", false, "Clear all saved shortcuts")
	printFlag := flag.Bool("p", false, "Print all saved shortcuts")

	// Parse command line args
	flag.Parse()

	// Handle flags
	switch {
	case *addFlag != "":
		fmt.Println("Adding:", *addFlag)

	case *removeFlag != "":
		fmt.Println("Removing:", *removeFlag)

	case *branchFlag != "":
		fmt.Println("Branching:", *branchFlag)

	case *clearFlag:
		fmt.Println("Clearing all shortcuts...")

	case *printFlag:
		fmt.Println("Printing all shortcut directories...")

	default:
		fmt.Println("Usage:")
		fmt.Println("  fcd -a <PATH>        Add a shortcut")
		fmt.Println("  fcd -r <LABEL>       Remove a shortcut")
		fmt.Println("  fcd -b <LABEL>       Branch to a shortcut")
		fmt.Println("  fcd -c               Clear all shortcuts")
		fmt.Println("  fcd -p               Print all shortcuts")
	}
}
