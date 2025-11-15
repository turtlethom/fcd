/*
Copyright © 2025 Thomas James <wjamesthomas3@gmail.com>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// completionCmd represents the completion command for fcd
//
// This command creates the autocompletion scripts for fcd
var completionCmd = &cobra.Command{
	Use:       "completion [bash|zsh|fish|powershell]",
	Short:     "Generates fcd auto-completion scripts",
	ValidArgs: []string{"bash", "zsh", "fish"},
	Args: cobra.ExactArgs(1),
	Example: `
===============================================================
# BASH (per user)
    fcd completion bash > ~/.local/share/fcd/fcd.bash
    echo "source ~/.local/share/fcd/fcd.bash" >> ~/.bashrc
===============================================================
# ZSH (per user)
    mkdir -p ~/.zsh/completions
    fcd completion zsh > ~/.zsh/completions/_fcd
    echo 'fpath+=("$HOME/.zsh/completions")' >> ~/.zshrc
    echo 'autoload -Uz compinit && compinit' >> ~/.zshrc
===============================================================
# FISH (per user)
    fcd completion fish > ~/.config/fish/completions/fcd.fish
===============================================================
`,
	Run: func(cmd *cobra.Command, args []string) {
		switch args[0] {
		case "bash":
			cmd.Root().GenBashCompletion(os.Stdout)
		case "zsh":
			cmd.Root().GenZshCompletion(os.Stdout)
		case "fish":
			cmd.Root().GenFishCompletion(os.Stdout, true)
			// case "powershell":
			// 	cmd.Root().GenPowerShellCompletion(os.Stdout)
		}
	},
}
