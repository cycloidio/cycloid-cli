package teams

import "github.com/spf13/cobra"

func NewTeamsCommands() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "teams",
		Aliases: []string{"team"},
		Short:   "Commands to manage teams",
	}
	cmd.AddCommand(
		NewTeamMembersCommand(),
		NewCreateTeamCommand(),
		NewUpdateTeamCommand(),
		NewGetTeamCommand(),
		NewDeleteTeamCommand(),
		NewListTeamCommand(),
	)
	return cmd
}
