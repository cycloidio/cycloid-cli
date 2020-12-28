package members

import (
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
)

func NewCommands() *cobra.Command {
	var (
		example = `
	# Manage members of my-org organization
	cy --org my-org members [invite|list|get|update|delete]
	`
		short = "Manage members from the organization"
		long  = short
	)

	var cmd = &cobra.Command{
		Use:     "members",
		Example: example,
		Short:   short,
		Long:    long,
	}
	common.RequiredPersistentFlag(common.WithFlagOrg, cmd)

	cmd.AddCommand(NewUpdateCommand(),
		NewDeleteCommand(),
		NewListCommand(),
		NewGetCommand(),
		NewInviteCommand(),
	  NewListInvitesCommand())

	return cmd
}
