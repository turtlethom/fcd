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

	if *addFlag != "" {
		fmt.Println("Adding: ", *addFlag)
	} else {
		fmt.Println("Syntax: fcd -a <PATH> or <LABEL:PATH>")
	}

	if *removeFlag != "" {
		fmt.Println("Removing: ", *removeFlag)
	} else {
		fmt.Println("Syntax: fcd -r <LABEL>")
	}

	if *branchFlag != "" {
		fmt.Println("Branching: ", *branchFlag)
	} else {
		fmt.Println("Syntax: fcd -b <LABEL>")
	}

	if *clearFlag {
		fmt.Println("Clearing all shortcuts...")
	} else {
		fmt.Println("Something went wrong?")
	}

	if *printFlag {
		fmt.Println("Printing all shortcut directories...")
	} else {
		fmt.Println("Something went wrong?")
	}
}
