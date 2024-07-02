package infrapolicies

import (
	"github.com/spf13/cobra"
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

	cmd.AddCommand(NewValidateCommand(),
		NewUpdateCommand(),
		NewCreateCommand(),
		NewListCommand(),
		NewGetCommand(),
		NewDeleteCommand())

	return cmd
}
