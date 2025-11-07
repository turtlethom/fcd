package cmd

import (
	"fmt"
	"os"
)

func HandleSubcommands(config *Config) {
	/* os.Args = {"fcd", command, commandArg} */
	var command string = ""
	var commandArg string = ""

	// No arg command initialization
	if len(os.Args) == 2 {
		command = os.Args[1]
	}

	// Multi arg command initialization
	if len(os.Args) >= 3 {
		command = os.Args[1]
		commandArg = os.Args[2]
	}

	switch command {
	case "add":
		{
			if err := HandleAdd(config, commandArg); err != nil {
				fmt.Fprintln(os.Stderr, "fcd: Error adding shortcut:", err)
				os.Exit(1)
			}
			fmt.Fprintf(os.Stderr, "fcd: Saved shortcut %q\n", commandArg)
		}

	case "remove":
		{
			if err := HandleRemove(config, commandArg); err != nil {
				fmt.Fprintln(os.Stderr, "fcd: Error removing shortcut:", err)
				os.Exit(1)
			}
			fmt.Fprintf(os.Stderr, "fcd: Successfully removed %q\n", commandArg)
		}
	case "branch":
		{
			// fmt.Fprintln(os.Stderr, "fcd: Successful branch: '", commandArg, "'")
			selectedBookmark := HandleBranch(config, commandArg)
			if selectedBookmark != "" {
				fmt.Println(selectedBookmark)
			}
		}
	case "clear":
		{
			if err := HandleClear(config); err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			} else {
				fmt.Fprintln(os.Stderr, "fcd: Successfully cleared all shortcuts")
			}
		}
	case "print":
		{
			// If no added shortcuts, exit
			if len(config.Shortcuts) <= 0 {
				fmt.Fprintln(os.Stderr, "fcd: No existing shortcuts")
				os.Exit(1)
			}
			fmt.Fprintln(os.Stderr, "fcd: Printing all shortcut directories...")
			HandlePrint(config)
		}
	case "setcolor":
		{
			HandleSetColor(config)
		}
	case "listcolors":
		{
			HandleListColors()
		}
	case "help":
		{
			HandleHelp()
		}
	case "":
		{
			// If no added shortcuts, exit
			if len(config.Shortcuts) == 0 {
				fmt.Fprintln(os.Stderr, "fcd: No existing shortcuts")
				os.Exit(1)
			}

			// RETURNING OUTPUT FROM FCD MENU
			if len(os.Args) == 1 {
				selectedPath := HandleMenu(config)
				if selectedPath != "" {
					fmt.Println(selectedPath)
				}
			}
		}
	}

}
