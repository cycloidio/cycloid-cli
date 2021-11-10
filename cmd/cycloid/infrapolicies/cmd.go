package infrapolicies

import (
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
)

func NewCommands() *cobra.Command {
	cmd := &cobra.Command{
		Use: "infrapolicy",
		Aliases: []string{
			"infra-policies",
			"infra-policy",
			"ip",
			"infrapolicies",
		},
		Short: "Manage the infraPolicies",
	}

	common.RequiredFlag(common.WithFlagOrg, cmd)

	cmd.AddCommand(NewValidateCommand(),
		NewUpdateCommand(),
		NewCreateCommand(),
		NewListCommand(),
		NewGetCommand(),
		NewDeleteCommand())
	return cmd
}
