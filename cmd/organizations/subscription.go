package organizations

import "github.com/spf13/cobra"

func NewSubscriptionCommands() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "subscription",
		Aliases: []string{"subscriptions"},
		Short:   "Manage organization's subscriptions",
		Args:    cobra.NoArgs,
	}
	cmd.AddCommand(
		NewCreateOrUpdateSubscriptionCommand(),
	)

	return cmd
}
