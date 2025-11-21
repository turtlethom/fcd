/*
Copyright © 2025 Thomas James <wjamesthomas3@gmail.com>
*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/termenv"
	"github.com/spf13/cobra"

	"github.com/turtlethom/fcd/internal"
)

var (
	// Config represents the global state of the user's configuration for fcd
	config *internal.Config

	// rootCmd represents the base command when called without any subcommands
	rootCmd = &cobra.Command{
		Use:   "fcd",
		Short: "fcd changes to bookmarked directories",
		Long: `fcd (fast change directory) is a CLI tool for navigating to directories saved by the user.

		[WARNING]: Valid shortcuts limited to the user's home directory
		All user configuration is stored within 'fcd_config.json'`,
		Run: func(cmd *cobra.Command, args []string) {
			if config == nil || len(config.Shortcuts) == 0 {
				fmt.Fprintln(os.Stderr, "fcd: shortcuts empty, use 'fcd add' to add a new shortcut")
				cmd.Help()
				os.Exit(1)
			}
			// Main output for fcd
			selectedPath := handleRoot(config)
			if selectedPath != "" {
				fmt.Println(selectedPath)
			}
		},
	}
)

// HandleRoot handles the initialization and customization of bubbles/list component for fcd
// This handles the core logic of fcd
//
// config - Pointer to a struct representing the current state of user's fcd_config.json
func handleRoot(config *internal.Config) string {
	// Ensures ANSI colors will be rendered for both stdout and stderr
	lipgloss.SetColorProfile(termenv.TrueColor)
	// Creating list items based on bookmarks
	items := make([]list.Item, len(config.Shortcuts))
	for i, sc := range config.Shortcuts {
		items[i] = sc
	}
	// Initialize styles
	styles := internal.NewStyles(config)
	bubblesList := internal.CreateBubblesList(items, styles)
	// Create Model
	model := internal.Model{
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
	m := finalModel.(internal.Model)
	if sel, ok := m.List.SelectedItem().(internal.Shortcut); ok {
		return sel.Path
	}
	return ""
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(branchCmd)
	rootCmd.AddCommand(clearCmd)
	rootCmd.AddCommand(colorsCmd)
	rootCmd.AddCommand(completionCmd)
	rootCmd.AddCommand(removeCmd)
	rootCmd.AddCommand(printCmd)
}

// SetConfig assigns the shared config (called from main)
//
// This function essentially ensures the chain of subcommands have access to the Config instance
func SetConfig(localConfig *internal.Config) {
	config = localConfig
}
