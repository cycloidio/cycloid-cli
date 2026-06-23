package teams

import "github.com/spf13/cobra"

func NewTeamMembersCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "members",
		Aliases: []string{"member"},
		Short:   "Manage team members.",
	}
	cmd.AddCommand(
		NewTeamMemberGetCommand(),
		NewTeamMemberListCommand(),
		NewTeamMemberAssignCommand(),
		NewTeamMemberUnAssignCommand(),
	)
	return cmd
}
