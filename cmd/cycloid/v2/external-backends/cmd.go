package externalBackends

import (
	"github.com/spf13/cobra"
	pEb "github.com/cycloidio/youdeploy-cli/cmd/cycloid/v1/external-backends"

)

func NewCommands() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "external-backends",
		Short: "...",
		Long:  `........ . . .... .. .. ....`,
	}

	cmd.AddCommand(pEb.NewUpdateCommand(),
		pEb.NewGetCommand(),
		pEb.NewDeleteCommand(),
		NewListCommand(),
		pEb.NewCreateCommand())

	return cmd
}
