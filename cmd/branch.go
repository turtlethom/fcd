/*
Copyright © 2025 Thomas James <wjamesthomas3@gmail.com>
*/
package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"github.com/turtlethom/fcd/internal"
)

// branchCmd is the cobra command for branching to a Shortcut directory saved within fcd_config.json
var branchCmd = &cobra.Command{
	Use:   "branch [label]",
	Short: "Jump to one of your saved shortcut directories",
	Long:  "Jump to one of your saved shortcut directories",
	Run: func(cmd *cobra.Command, args []string) {
		selectedShortcut := handleBranch(config, args[0])
		if selectedShortcut != "" {
			fmt.Println(selectedShortcut)
		}
	},
}

// handleBranch handles the logic for branching to a Shortcut
func handleBranch(config *internal.Config, label string) string {
	// Iterate and check LABEL exists in config.Shortcuts
	for _, sc := range config.Shortcuts {
		if sc.Label == strings.ToUpper(strings.TrimSpace(label)) {
			return sc.Path
		}
	}
	fmt.Fprintf(os.Stderr, "fcd: Cannot find matching path for %q\n", label)
	return ""
}
