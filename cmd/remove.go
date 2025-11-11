package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/turtlethom/fcd/internal"
)


var removeCmd = &cobra.Command{
	Use:   "remove [label]",
	Short: "Remove a saved shortcut from your configuration",
	Long:  `Remove a saved shortcut from your configuration`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if err := handleRemove(config, args[0]); err != nil {
			fmt.Fprintln(os.Stderr, "fcd: Unable to remove shortcut:", err)
			os.Exit(1)
		}
		fmt.Fprintf(os.Stderr, "fcd: '%q' has been successfully removed\n", args[0])
	},
}
// handleRemove removes a Shortcut from Config based on the Shortcut.Label
//
// label - String representing user's input of a matching label to a Shortcut
func handleRemove(config *internal.Config, label string) error {
	// Convert commandArg to match shortcut label
	removeLabel := strings.ToUpper(strings.TrimSpace(label))
	// Create new shorcuts
	newShortcuts := make([]internal.Shortcut, 0, len(config.Shortcuts))
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
	return internal.SaveToConfig(config)
}
