package roles

import (
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
)

func NewCommands() *cobra.Command {
	var (
		example = `
	# Manage roles of my-org organization
	cy --org my-org roles [list|get|delete]
	`
		short = "Manage roles from the organization"
		long  = short
	)

	var cmd = &cobra.Command{
		Use:     "roles",
		Example: example,
		Short:   short,
		Long:    long,
	}
	common.RequiredPersistentFlag(common.WithFlagOrg, cmd)

	cmd.AddCommand(NewDeleteCommand(),
		NewListCommand(),
		NewGetCommand())

	return cmd
}
