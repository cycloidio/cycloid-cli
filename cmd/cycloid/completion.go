package root

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// NewCompletionCmd returns the cobra command that outputs shell completion code
func NewCompletionCmd() *cobra.Command {
	return &cobra.Command{
		Use: "completion (zsh|bash)",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return fmt.Errorf("requires 1 arg, found %d", len(args))
			}
			return cobra.OnlyValidArgs(cmd, args)
		},
		ValidArgs: []string{"bash", "zsh"},
		Short:     "Output shell completion for the given shell (bash or zsh)",
		RunE:      completion,
		Example: `
	# generate completion for ZSH
	cy completion zsh > _cy; cp _cy ~/.oh-my-zsh/functions; source ~/.zshrc

	# generate completion for Bash
	cy completion bash > cy; sudo cp cy /etc/bash_completion.d/; source /etc/profile
`,
	}
}

func completion(cmd *cobra.Command, args []string) error {
	switch args[0] {
	case "bash":
		return cmd.Root().GenBashCompletion(os.Stdout)
	case "zsh":
		return cmd.Root().GenZshCompletion(os.Stdout)
	default:
		return nil
	}
}
