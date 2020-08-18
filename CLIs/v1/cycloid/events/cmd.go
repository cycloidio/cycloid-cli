package events

import (
	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/common"
	"github.com/spf13/cobra"
)

func NewCommands() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "events",
		Short: "...",
		Long:  `........ . . .... .. .. ....`,
	}

	common.RequiredPersistentFlag(common.WithFlagOrg, cmd)

	cmd.AddCommand(NewSendCommand())

	return cmd
}
