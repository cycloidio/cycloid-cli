package organizations

import (
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/organizations/licence"
)

func NewCommands() *cobra.Command {
	cmd := &cobra.Command{
		Use: "organization",
		Aliases: []string{
			"org",
			"organizations",
			"o",
		},
		Short: "Manage the organizations",
	}
	cmd.AddCommand(
		NewCreateCommand(),
		NewUpdateCommand(),
		NewListCommand(),
		NewListWorkersCommand(),
		NewDeleteCommand(),
		NewListChildrensCommand(),
		NewCreateChildCommand(),
		NewGetCommand(),
		NewSubscriptionCommands(),
		licence.NewCommands(),
	)
	return cmd
}
