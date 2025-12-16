/*
Copyright © 2025 Thomas James <wjamesthomas3@gmail.com>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/turtlethom/fcd/internal/data"
)

// printCmd is the cobra command for printing saved Shortcuts in fcd_config.json
var printCmd = &cobra.Command{
	Use:   "print",
	Short: "Print all your saved shortcuts",
	Long:  "Print all your saved shortcuts",
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		if USER_CONFIG == nil {
			fmt.Fprintln(os.Stderr, "config cannot be loaded")
			os.Exit(1)
		}
		handlePrint(USER_CONFIG)
	},
}

// handlePrint handles logic for printing all Shortcuts present based on fcd_config.json
func handlePrint(config *data.Config) {
	fmt.Fprintln(os.Stderr, "fcd: Printing all shortcut directories...")
	for _, sc := range config.Shortcuts {
		fmt.Fprintf(os.Stderr, "   %-20s %s\n", sc.Label, sc.Path)
	}
}
