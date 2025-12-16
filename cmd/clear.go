/*
Copyright © 2025 Thomas James <wjamesthomas3@gmail.com>
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"github.com/turtlethom/fcd/internal/data"
)

var clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Clear all saved shortcuts",
	Long:  "Clear all saved shortcuts",
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		if err := handleClear(USER_CONFIG); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		} else {
			fmt.Fprintln(os.Stderr, "fcd: Successfully cleared all shortcuts")
		}
	},
}

// handleClear handles clearing all saved Shortcut from fcd_config.json
func handleClear(config *data.Config) error {
	// Get confirmation from user on clearing Shortcuts (y/n)
	reader := bufio.NewReader(os.Stdin)
	fmt.Fprintf(os.Stderr, "fcd: Are you sure you want to clear all bookmarks? (y/n): ")

	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	switch strings.ToLower(input) {
	// Update the config.Shortcuts to empty array and save to fcd_config.json
	case "y":
		config.Shortcuts = []data.Shortcut{}
		return data.SaveToConfig(config)
	// Cancel operations if "n" or some other invalid answer
	case "n":
		fallthrough
	default:
		return fmt.Errorf("fcd: Canceled clearing bookmarks")
	}
}
