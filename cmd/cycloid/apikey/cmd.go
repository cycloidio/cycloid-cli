package apikey

import (
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
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
		Hidden:  true,
	}
	common.RequiredPersistentFlag(common.WithFlagOrg, cmd)

	cmd.AddCommand(
		NewCreateCommand(),
		NewListCommand(),
	)
	return cmd
}
