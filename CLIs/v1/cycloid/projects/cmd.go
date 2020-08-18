package projects

import (
	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/common"
	"github.com/spf13/cobra"
)

func NewCommands() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "project",
		Short: "...",
		Long:  `........ . . .... .. .. ....`,
	}
	common.RequiredPersistentFlag(common.WithFlagOrg, cmd)

	cmd.AddCommand(NewUpdateCommand(),
		NewDeleteCommand(),
		NewCreateCommand(),
		NewListCommand(),
		NewGetCommand())

	return cmd
}
