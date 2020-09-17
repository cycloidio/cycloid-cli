package externalBackends

import (
	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/common"
	"github.com/spf13/cobra"
)

func NewCommands() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "external-backends",
		Short: "...",
		Long:  `........ . . .... .. .. ....`,
	}
	common.RequiredPersistentFlag(common.WithFlagOrg, cmd)

	cmd.AddCommand(NewGetCommand(),
		NewDeleteCommand(),
		NewListCommand(),
		NewCreateCommand())

	return cmd
}
