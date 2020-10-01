package events

import (
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/spf13/cobra"
)

func NewCommands() *cobra.Command {
	var cmd = &cobra.Command{
		Use: "event",
		Aliases: []string{
			"events",
			"e",
		},
		Short: "Manage the events",
	}

	common.RequiredPersistentFlag(common.WithFlagOrg, cmd)

	cmd.AddCommand(NewSendCommand())

	return cmd
}
