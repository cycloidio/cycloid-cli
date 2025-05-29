package root

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var (
	supportedShells = []string{"bash", "fish", "zsh", "powershell"}
)

// NewCompletionCmd returns the cobra command that outputs shell completion code
func NewCompletionCmd() *cobra.Command {
	return &cobra.Command{
		Use: `completion (` + strings.Join(supportedShells, "|") + `)`,
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return fmt.Errorf("requires 1 arg, found %d", len(args))
			}
			return cobra.OnlyValidArgs(cmd, args)
		},
		ValidArgs: []string{"bash", "zsh"},
		Short:     "Output shell completion for the given shell (bash or zsh)",
		RunE:      completion,
		Example: `# generate completion for ZSH
cy completion zsh > _cy; cp _cy ~/.oh-my-zsh/functions; source ~/.zshrc

	# generate completion for Bash
	cy completion bash > cy; sudo cp cy /etc/bash_completion.d/; source /etc/profile
`,
	}
}

func completion(cmd *cobra.Command, args []string) error {
	switch args[0] {
	case "bash":
		return cmd.Root().GenBashCompletionV2(cmd.OutOrStdout(), true)
	case "fish":
		return cmd.Root().GenFishCompletion(cmd.OutOrStdout(), true)
	case "zsh":
		return cmd.Root().GenZshCompletion(cmd.OutOrStdout())
	case "powershell", "pwsh":
		return cmd.Root().GenPowerShellCompletionWithDesc(cmd.OutOrStdout())
	default:
		return fmt.Errorf("invalid shell argument '%s' valid options are: %v", args[0], supportedShells)
	}
}
