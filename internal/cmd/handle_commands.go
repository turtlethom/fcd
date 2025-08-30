package cmd

import (
	"fmt"
	"os"
)

func HandleSubcommands(config *Config) {
	/* os.Args = {"fcd", command, commandArg} */
	userArguments := os.Args
	var command string
	var commandArg string
	// If no added shortcuts, exit
	if len(config.Shortcuts) == 0 {
		fmt.Fprintln(os.Stderr, "fcd: No existing shortcuts")
		os.Exit(1)
	}

	// RETURNING OUTPUT FROM FCD MENU
	if len(userArguments) == 1 {
		selectedPath := HandleMenu(config)
		if selectedPath != "" {
			fmt.Println(selectedPath)
		}
	}

	// No arg command initialization
	if len(userArguments) == 2 {
		command = userArguments[1]
		commandArg = ""
	}

	// Single arg command initialization
	if len(userArguments) == 3 {
		command = userArguments[1]
		commandArg = userArguments[2]
	}

	switch command {
	case "add":
		{
			if err := HandleAdd(config, commandArg); err != nil {
				fmt.Fprintln(os.Stderr, "fcd: Error adding shortcut:", err)
				os.Exit(1)
			}
			fmt.Fprintln(os.Stderr, "fcd: Saved shortcut '", commandArg, "'")
		}

	case "remove":
		{
			if err := HandleRemove(config, commandArg); err != nil {
				fmt.Fprintln(os.Stderr, "fcd: Error removing shortcut:", err)
				os.Exit(1)
			}
			fmt.Fprintln(os.Stderr, "fcd: Successfully removed '", commandArg, "'")
		}
	case "branch":
		{
			fmt.Fprintln(os.Stderr, "fcd: Successful branch: '", commandArg, "'")
		}
	case "clear":
		{
			if err := HandleClear(config); err != nil {
				fmt.Fprintln(os.Stderr, "fcd: Error clearing shortcuts:", err)
				os.Exit(1)
			}
			fmt.Fprintln(os.Stderr, "fcd: Successfully cleared all shortcuts")
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
	case "help":
		{
			HandleHelp()
		}
	}
}
