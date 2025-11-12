/*
Copyright © 2025 Thomas James <wjamesthomas3@gmail.com>
*/
package cmd

import (
	"github.com/spf13/cobra"
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
	"none":      "#------",
}

var colorsCmd = &cobra.Command{
	Use:   "colors",
	Short: "Manage fcd color themes and styles",
	Long: `Manage fcd color themes and styles.

	Use 'fcd colors set' to customize theme colors.
	Use 'fcd colors list' to list all available colors`,
}

func init() {
	colorsCmd.AddCommand(colorsSetCmd)
	colorsCmd.AddCommand(colorsListCmd)
}
