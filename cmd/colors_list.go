package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var colorsListCmd = &cobra.Command{
	Use: "list",
	Short: "Lists the colors available for the fcd menu",
	Long: "Lists the colors available for the fcd menu",
	Run: func(cmd *cobra.Command, args []string) {
		handleListColor()
	},
}

func handleListColor() {
		fmt.Fprintln(os.Stderr, "fcd: Printing all supported colors")
	for name, hex := range MENU_COLORS {
		if name == "none" {
			fmt.Fprintf(os.Stderr, "|%10s --> %s\n", name, "DEFAULT|")
			continue
		}
		fmt.Fprintf(os.Stderr, "|%10s --> %s|\n", name, hex)
	}
}
