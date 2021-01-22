package terracost

import (
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
)

func NewCommands() *cobra.Command {
	cmd := &cobra.Command{
		Use: "terracost",
	}

	common.RequiredFlag(common.WithFlagOrg, cmd)

	cmd.AddCommand(NewEstimateCommand())
	return cmd
}
