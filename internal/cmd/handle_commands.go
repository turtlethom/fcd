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

// MENU_COLORS represents the global colors that can be applied to bubbles/list component
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

// HandleAdd handles the addition of bookmarks, based on provided provided label or label/path pair; restricted to $HOME
//
// config - Pointer to a struct representing the current state of user's fcd_config.json
// commandArg - String representing user's input for a new Shortcut
func HandleAdd(config *Config, commandArg string) error {
	var shortcutLabel, shortcutPath string
	// If commandArg is equal to "label:shortcutPath", parse "shortcutLabel" and "shortcutPath"
	if strings.Contains(commandArg, ":") {
		parts := strings.SplitN(commandArg, ":", 2)
		shortcutLabel = parts[0]
		shortcutPath = parts[1]
	// If commandArg is equal to "shortcutPath", parse "shortcutLabel"
	} else {
		shortcutPath = commandArg
		shortcutLabel = filepath.Base(shortcutPath)
	}

	// Format "SHORTCUTLABEL" as "SHORTCUTLABEL", then add "SHORTCUTLABEL" and "shortcutPath"
	shortcutLabel = strings.ToUpper(strings.TrimSpace(shortcutLabel))
	if err := config.AddShortcut(shortcutLabel, shortcutPath); err != nil {
		return err
	}
	return SaveToConfig(config)
}

// HandleRemove removes a Shortcut from Config based on the Shortcut.Label
//
// commandArg - String representing user's input of a matching label to a Shortcut
func HandleRemove(config *Config, commandArg string) error {
	// Convert commandArg to match shortcut label
	removeLabel := strings.ToUpper(strings.TrimSpace(commandArg))
	// Create new shorcuts
	newShortcuts := make([]Shortcut, 0, len(config.Shortcuts))
	found := false
	for _, sc := range config.Shortcuts {
		// Exclude shortcut from being added back to new Shortcuts array
		if sc.Label == removeLabel {
			found = true
			continue
		}
		// Add each Shortcut to newShortcuts array
		newShortcuts = append(newShortcuts, sc)
	}
	// Check if label is invalid or Shortcut does not exist
	if !found {
		return fmt.Errorf("fcd: no shortcut found with label %q\n", removeLabel)
	}
	// Update new Shortcuts array and save data to fcd_config.json
	config.Shortcuts = newShortcuts
	return SaveToConfig(config)
}

// HandleBranch handles the 'branching' to a Shortcut based on the label provided
//
// config - Pointer to a struct representing the current state of user's fcd_config.json
// commandArg - String representing user's input of a corresponding label to branch to
func HandleBranch(config *Config, commandArg string) string {
	// Iterate and check LABEL exists in config.Shortcuts
	for _, sc := range config.Shortcuts {
		if sc.Label == strings.ToUpper(strings.TrimSpace(commandArg)) {
			return sc.Path
		}
	}
	fmt.Fprintf(os.Stderr, "fcd: Cannot find matching path for %q\n", commandArg)
	return ""
}

// handleClear handles clearing all saved Shortcuts from fcd_config.json
func HandleClear(config *Config) error {
	// Get confirmation from user on clearing Shortcuts (y/n)
	reader := bufio.NewReader(os.Stdin)
	fmt.Fprintf(os.Stderr, "fcd: Are you sure you want to clear all bookmarks? (y/n): ")

	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	switch strings.ToLower(input) {
	// Update the config.Shortcuts to empty array and save to fcd_config.json
	case "y":
		config.Shortcuts = []Shortcut{}
		return SaveToConfig(config)
	// Cancel operations if "n" or some other invalid answer
	case "n":
		fallthrough
	default:
		return fmt.Errorf("fcd: Canceled clearing bookmarks")
	}
}

// HandlePrint handles printing all shortcuts present within user's fcd_config.json
//
// config - Pointer to a struct representing the current state of user's fcd_config.json
func HandlePrint(config *Config) {
	for _, sc := range config.Shortcuts {
		fmt.Fprintf(os.Stderr, "%-10s %s\n", sc.Label, sc.Path)
	}
}

// HandleSetColors handles setting primary, secondary, and tertiary colors for 
// bubbles/list component
//
// config - Pointer to a struct representing the current state of user's fcd_config.json
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

// HandleListColors handles the printing of all available colors, as listed within MENU_COLORS
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

// HandleHelp handles printing the usage of fcd and all functionality present within the CLI
func HandleHelp() {
	fmt.Fprintln(os.Stderr, "fcd - fast change directory")
	fmt.Fprintln(os.Stderr, "  - Change into bookmarked directories using custom labels")
	fmt.Fprintln(os.Stderr, "  - Limited to directories within a user's home directory")
	fmt.Fprintln(os.Stderr)
	fmt.Fprintln(os.Stderr, "Usage:")
	fmt.Fprintln(os.Stderr, "  fcd                               Opens interactive menu for viewing saved shortcuts within user's configuration")
	fmt.Fprintln(os.Stderr, "  fcd add <PATH>                    Add a shortcut saved within user's configuration")
	fmt.Fprintln(os.Stderr, "  fcd remove <LABEL>                Remove a shortcut saved within user's configuration")
	fmt.Fprintln(os.Stderr, "  fcd branch <LABEL>                Changes directory to a shortcut saved within user's configuration")
	fmt.Fprintln(os.Stderr, "  fcd clear                         Clear all shortcuts saved within user's configuration")
	fmt.Fprintln(os.Stderr, "  fcd print                         Print all shortcuts saved within user's configuration")
	fmt.Fprintln(os.Stderr, "  fcd listcolors                    Prints all colors available with corresponding hex values")
	fmt.Fprintln(os.Stderr, "  fcd setcolor [-p] [-s] [t]        Prints all colors available with corresponding hex values")
}

// HandleMenu handles the initialization and customization of bubbles/list component
//
// config - Pointer to a struct representing the current state of user's fcd_config.json
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
