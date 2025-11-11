package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/turtlethom/fcd/internal"
)

var (
	primaryColor   string
	secondaryColor string
	tertiaryColor  string

	colorsSetCmd = &cobra.Command{
		Use:     "set",
		Example: "fcd colors set --primary crimson --secondary white --tertiary black",
		Short:   "Sets the colors for the fcd menu",
		Long:    "Sets the colors for the fcd menu",
		Run: func(cmd *cobra.Command, args []string) {
			if len(primaryColor)+len(secondaryColor)+len(tertiaryColor) == 0 {
				cmd.Help()
			}
			handleSetColor(config)
		},
	}
)

func handleSetColor(config *internal.Config) {
	// If --primary color
	if primaryColor != "" {
		hex, valid := validateColorChoice("primary", primaryColor)
		if valid {
			config.UserColors.Primary = hex
		}
	}
	// If --secondary color
	if secondaryColor != "" {
		hex, valid := validateColorChoice("secondary", secondaryColor)
		if valid {
			config.UserColors.Secondary = hex
		}
	}
	// If --tertiary color
	if tertiaryColor != "" {
		hex, valid := validateColorChoice("tertiary", tertiaryColor)
		if valid {
			config.UserColors.Tertiary = hex
		}
	}

	internal.SaveToConfig(config)
}

func validateColorChoice(colorFor string, colorName string) (string, bool) {
	if hex, ok := MENU_COLORS[colorName]; ok {
		// If color is valid and notify color exists and will be set
		fmt.Fprintf(os.Stderr, "fcd: setting %v color to [%v - %v]\n", colorFor, colorName, hex)
		return hex, true
	}
	// Else color is invalid, notify that color does not exist and will not be set
	fmt.Fprintf(os.Stderr, "Error: %v color not set, '%v' does not exist\n", colorFor, colorName)
	return "", false
}

func init() {
	colorsSetCmd.Flags().StringVar(&primaryColor, "primary", "", "Set the primary color for fcd menu")
	colorsSetCmd.Flags().StringVar(&secondaryColor, "secondary", "", "Set the primary color for fcd menu")
	colorsSetCmd.Flags().StringVar(&tertiaryColor, "tertiary", "", "Set the primary color for fcd menu")
}
