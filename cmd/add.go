package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/turtlethom/fcd/internal"
)

// addCmd is the cobra command for branching to a Shortcut directory saved within fcd_config.json
var addCmd = &cobra.Command{
	Use:   "add [label|label:path]",
	Short: "Add a directory path to be saved as a shortcut",
	Long:  `Add a directory path to be saved as a shortcut`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if err := handleAdd(config, args[0]); err != nil {
			fmt.Fprintln(os.Stderr, "fcd: Error adding shortcut:", err)
			os.Exit(1)
		}
		fmt.Fprintf(os.Stderr, "fcd: %q successfully saved as a shortcut\n", args[0])
	},
}

func handleAdd(config *internal.Config, labelPathArg string) error {
	var shortcutLabel, shortcutPath string
	// If labelPathArg = "label:path", parse "label" and "path"
	if strings.Contains(labelPathArg, ":") {
		parts := strings.SplitN(labelPathArg, ":", 2)
		shortcutLabel = parts[0]
		shortcutPath = parts[1]
	// If labelPathArg = "path", parse "label" from base of path
	} else {
		shortcutPath = labelPathArg
		shortcutLabel = filepath.Base(shortcutPath)
	}
	// Format "label:path" -> "LABEL:path"
	shortcutLabel = strings.ToUpper(strings.TrimSpace(shortcutLabel))
	if err := config.AddShortcut(shortcutLabel, shortcutPath); err != nil {
		return err
	}

	return internal.SaveToConfig(config)

}
