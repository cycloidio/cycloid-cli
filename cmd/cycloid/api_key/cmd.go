package api_key

import (
	"github.com/spf13/cobra"
)

var (
	example = `
	# Manage API keys of my-org organization
	cy --org my-org api-key [create|list|get|delete]
`
	short = "Manage organization API keys"
)

func NewCommands() *cobra.Command {
	cmd := &cobra.Command{
		Use: "api-key",
		Aliases: []string{
			"api-keys",
			"ak",
		},
		Example: example,
		Short:   short,
	}

	cmd.AddCommand(
		NewDeleteCommand(),
		NewGetCommand(),
		NewListCommand(),
	)
	return cmd
}
