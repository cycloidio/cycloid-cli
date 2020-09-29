package stacks

import (
	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/common"
	"github.com/spf13/cobra"
)

func NewCommands() *cobra.Command {
	var cmd = &cobra.Command{
		Use: "stack",
		Aliases: []string{
			"s",
			"stacks",
		},
		Short: "Manage the stacks",
	}
	common.RequiredPersistentFlag(common.WithFlagOrg, cmd)

	cmd.AddCommand(NewListCommand(),
		NewGetCommand())

	return cmd
}
