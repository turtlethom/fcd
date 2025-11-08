package cmd

import (
	"fmt"
	"os"
)

// HandleSubcommands handles the entirety of fcd functionality in the command line
//
// config - Pointer to Config struct that stores the user's current fcd_config.json data
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
	// fcd add
	case "add":
		{
			if err := HandleAdd(config, commandArg); err != nil {
				fmt.Fprintln(os.Stderr, "fcd: Error adding shortcut:", err)
				os.Exit(1)
			}
			fmt.Fprintf(os.Stderr, "fcd: Saved shortcut %q\n", commandArg)
		}
	// fcd remove
	case "remove":
		{
			if err := HandleRemove(config, commandArg); err != nil {
				fmt.Fprintln(os.Stderr, "fcd: Error removing shortcut:", err)
				os.Exit(1)
			}
			fmt.Fprintf(os.Stderr, "fcd: Successfully removed %q\n", commandArg)
		}
	// fcd branch
	case "branch":
		{
			// fmt.Fprintln(os.Stderr, "fcd: Successful branch: '", commandArg, "'")
			selectedBookmark := HandleBranch(config, commandArg)
			if selectedBookmark != "" {
				fmt.Println(selectedBookmark)
			}
		}
	// fcd clear
	case "clear":
		{
			if err := HandleClear(config); err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			} else {
				fmt.Fprintln(os.Stderr, "fcd: Successfully cleared all shortcuts")
			}
		}
	// fcd print
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
	// fcd setcolor
	case "setcolor":
		{
			HandleSetColor(config)
		}
	// fcd listcolors
	case "listcolors":
		{
			HandleListColors()
		}
	// fcd help
	case "help":
		{
			HandleHelp()
		}
	// fcd
	case "":
		{
			// If no shortcuts are present within fcd_config.json, exit
			if len(config.Shortcuts) == 0 {
				fmt.Fprintln(os.Stderr, "fcd: No existing shortcuts")
				os.Exit(1)
			}

			// Returns the path of a Shortcut selected within the bubbles/list component
			if len(os.Args) == 1 {
				selectedPath := HandleMenu(config)
				if selectedPath != "" {
					fmt.Println(selectedPath)
				}
			}
		}
	}

}
