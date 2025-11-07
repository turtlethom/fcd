package cmd

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/turtlethom/fcd/internal/utils"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/termenv"
)

/*
TODO: Move menu colors somewhere else
Make err functions to avoid repetitive printf to os.Stderr
*/
var MENU_COLORS = map[string]string{
	"crimson":   "#DC143C",
	"coral":     "#F08080",
	"pink":      "#FF007F",
	"red":       "#E06C75",
	"orange":    "#D19A66",
	"tangerine": "#FF9500",
	"amber":     "#FFBF00",
	"gold":      "#DAA520",
	"yellow":    "#E5C078",
	"green":     "#98C379",
	"emerald":   "#50C878",
	"aqua":      "#56B6C2",
	"navy":      "#000080",
	"amethyst":  "#9966CC",
	"indigo":    "#4B0082",
	"violet":    "#C678DD",
	"gray":      "#5C6370",
	"white":     "#FFFFFF",
	"black":     "#000000",
	"brown":     "#A0522D",
	"none":      "",
}

// Adds bookmarks based on provided LABEL or LABEL:PATH; restricted to $HOME
func HandleAdd(config *Config, commandArg string) error {
	var label, path string
	// If arg value has "LABEL:PATH"
	if strings.Contains(commandArg, ":") {
		parts := strings.SplitN(commandArg, ":", 2)
		label = parts[0]
		path = parts[1]
		// If only "PATH" is provided
	} else {
		path = commandArg
		label = filepath.Base(path)
	}

	label = strings.ToUpper(strings.TrimSpace(label))
	if err := config.AddShortcut(label, path); err != nil {
		return err
	}
	return SaveToConfig(config)
}

// Removes file based on provided LABEL
func HandleRemove(config *Config, commandArg string) error {
	// Convert commandArg to match shortcut label
	removeLabel := strings.ToUpper(strings.TrimSpace(commandArg))
	// Create new shorcuts
	newShortcuts := make([]Shortcut, 0, len(config.Shortcuts))
	found := false
	for _, sc := range config.Shortcuts {
		// Exclude shortcut from being added back to newShorcuts array
		if sc.Label == removeLabel {
			found = true
			continue
		}
		// Add shortcut to newShortcuts array
		newShortcuts = append(newShortcuts, sc)
	}
	if !found {
		return fmt.Errorf("fcd: no shortcut found with label %q\n", removeLabel)
	}
	config.Shortcuts = newShortcuts
	return SaveToConfig(config)
}

// 'Branches' to the path that matches the provided LABEL
func HandleBranch(config *Config, commandArg string) string {
	for _, sc := range config.Shortcuts {
		if sc.Label == strings.ToUpper(strings.TrimSpace(commandArg)) {
			return sc.Path
		}
	}
	fmt.Fprintf(os.Stderr, "fcd: Cannot find matching path for %q\n", commandArg)
	return ""
}

// Clear all bookmarks from fcd_config.json
func HandleClear(config *Config) error {
	// Get confirmation from user (y/n)
	reader := bufio.NewReader(os.Stdin)
	fmt.Fprintf(os.Stderr, "fcd: Are you sure you want to clear all bookmarks? (y/n): ")

	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	switch strings.ToLower(input) {
	// Assign the shortcuts to empty array and save
	case "y":
		config.Shortcuts = []Shortcut{}
		return SaveToConfig(config)
	case "n":
		fallthrough
	default:
		// return fmt.Errorf("fcd: Defaulting to no - operation ceased")
		return fmt.Errorf("fcd: Canceled clearing bookmarks")
	}
}

func HandlePrint(config *Config) {
	for _, sc := range config.Shortcuts {
		fmt.Fprintf(os.Stderr, "%-10s %s\n", sc.Label, sc.Path)
	}
}

func HandleSetColor(config *Config) error {
	// Create setcolor command
	setColorCmd := flag.NewFlagSet("setcolor", flag.ExitOnError)
	// Attach flags to command
	primaryFlag := setColorCmd.String("p", "", "primary color of fcd menu")
	secondaryFlag := setColorCmd.String("s", "", "secondary color of fcd menu")
	tertiaryFlag := setColorCmd.String("t", "", "tertiary color of fcd menu")
	setColorCmd.Parse(os.Args[2:])
	// Checking if no arguments provided to flags
	if *primaryFlag == "" && *secondaryFlag == "" && *tertiaryFlag == "" {
		fmt.Fprintf(os.Stderr, "fcd: No color arguments provided\n")
		fmt.Fprintf(os.Stderr, "fcd: Use -p, -s, or -t\n")
	}
	// Checking if primary color flag has valid argument
	if *primaryFlag != "" {
		hexCode, err := utils.SelectColor(*primaryFlag)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fcd: %s\n", err)
			os.Exit(1)
		}
		if hexCode == "" {
			fmt.Fprintf(os.Stderr, "fcd: Set primary to %s\n", "none")
		} else {
			fmt.Fprintf(os.Stderr, "fcd: Set primary to %s\n", hexCode)
		}
		config.UserColors.Primary = hexCode
	}
	// Checking if secondary color flag has valid argument
	if *secondaryFlag != "" {
		hexCode, err := utils.SelectColor(*secondaryFlag)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fcd: %v\n", err)
			os.Exit(1)
		}
		if hexCode == "" {
			fmt.Fprintf(os.Stderr, "fcd: Set secondary to %s\n", "none")
		} else {
			fmt.Fprintf(os.Stderr, "fcd: Set secondary to %s\n", hexCode)
		}
		config.UserColors.Secondary = hexCode
	}
	// Checking if tertiary color flag has valid argument
	if *tertiaryFlag != "" {
		hexCode, err := utils.SelectColor(*tertiaryFlag)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fcd: %v\n", err)
			os.Exit(1)
		}
		if hexCode == "" {
			fmt.Fprintf(os.Stderr, "fcd: Set tertiary to %s\n", "none")
		} else {
			fmt.Fprintf(os.Stderr, "fcd: Set tertiary to %s\n", hexCode)
		}
		config.UserColors.Tertiary = hexCode
	}

	return SaveToConfig(config)
}

func HandleListColors() {
	fmt.Fprintln(os.Stderr, "fcd: Printing all supported colors")
	for name, hex := range MENU_COLORS {
		if name == "none" {
			fmt.Fprintf(os.Stderr, "|%10s --> %s\n", name, "DEFAULT|")
			continue
		}
		fmt.Fprintf(os.Stderr, "|%10s --> %s|\n", name, hex)
	}
}

func HandleHelp() {
	fmt.Fprintln(os.Stderr, "fcd - fast change directory")
	fmt.Fprintln(os.Stderr, "  - Change into bookmarked directories using custom labels")
	fmt.Fprintln(os.Stderr, "  - Limited to directories within a user's home directory")
	fmt.Fprintln(os.Stderr)
	fmt.Fprintln(os.Stderr, "Usage:")
	fmt.Fprintln(os.Stderr, "  fcd                   Search saved shortcuts")
	fmt.Fprintln(os.Stderr, "  fcd add <PATH>        Add a shortcut")
	fmt.Fprintln(os.Stderr, "  fcd remove <LABEL>    Remove a shortcut")
	fmt.Fprintln(os.Stderr, "  fcd branch <LABEL>    Branch to a shortcut")
	fmt.Fprintln(os.Stderr, "  fcd clear             Clear all shortcuts")
	fmt.Fprintln(os.Stderr, "  fcd print             Print all shortcuts")
}

func HandleMenu(config *Config) string {
	// Ensures ANSI colors will be rendered for both stdout and stderr
	lipgloss.SetColorProfile(termenv.TrueColor)
	// Creating list items based on bookmarks
	items := make([]list.Item, len(config.Shortcuts))
	for i, sc := range config.Shortcuts {
		items[i] = sc
	}
	// Initialize styles
	styles := NewStyles(config)
	bubblesList := CreateBubblesList(items, styles)
	// Create Model
	model := Model{
		List: bubblesList,
	}
	// Start bubbletea for FCD menu
	program := tea.NewProgram(
		model,
		tea.WithOutput(os.Stderr),
		tea.WithAltScreen(),
	)
	finalModel, err := program.Run()
	if err != nil {
		log.Fatal(err)
	}
	// Get result of user choice and return it at the end
	m := finalModel.(Model)
	if sel, ok := m.List.SelectedItem().(Shortcut); ok {
		return sel.Path
	}
	return ""
}
