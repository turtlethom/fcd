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
		Use:   "set",
		Example: "fcd colors set --primary crimson --secondary white --tertiary black",
		Short: "Sets the colors for the fcd menu",
		Long:  "Sets the colors for the fcd menu",
		Run: func(cmd *cobra.Command, args []string) {
			handleSetColor(config)
		},
	}
)

func handleSetColor(config *internal.Config) {
	// If --primary color
	if primaryColor != "" {
		hex, err := validateColorChoice(primaryColor)
		if err != nil {
			fmt.Fprintln(os.Stderr, "primary color not set: ", err)
		} else {
			fmt.Fprintf(os.Stderr, "fcd: setting primary color to [%v - %v]\n", primaryColor, hex)
			// config.UserColors.Primary = hex
		}
	}
	// If --secondary color
	if secondaryColor != "" {
		hex, err := validateColorChoice(secondaryColor)
		if err != nil {
			fmt.Fprintln(os.Stderr, "secondary color not set: ", err)
		} else {
			fmt.Fprintf(os.Stderr, "fcd: setting secondary color to [%v - %v]\n", secondaryColor, hex)
			// config.UserColors.Secondary = hex
		}
	}
	// If --tertiary color
	if tertiaryColor != "" {
		hex, err := validateColorChoice(tertiaryColor)
		if err != nil {
			fmt.Fprintln(os.Stderr, "tertiary color not set: ", err)
		} else {
			fmt.Fprintf(os.Stderr, "fcd: setting tertiary color to [%v - %v]\n", tertiaryColor, hex)
			// config.UserColors.Tertiary = hex
		}
	}

	internal.SaveToConfig(config)
}

func validateColorChoice(color string) (string, error) {
		if hex, ok := MENU_COLORS[color]; ok {
			return hex, nil
		}
	return "", fmt.Errorf("[%v - N/A]", color)
}

func init() {
	colorsSetCmd.Flags().StringVar(&primaryColor, "primary", "", "Set the primary color for fcd menu")
	colorsSetCmd.Flags().StringVar(&secondaryColor, "secondary", "", "Set the primary color for fcd menu")
	colorsSetCmd.Flags().StringVar(&tertiaryColor, "tertiary", "", "Set the primary color for fcd menu")
}
