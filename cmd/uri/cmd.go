package uri

import (
	"github.com/spf13/cobra"
)

func NewURICommands() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "uri",
		Short: "Use URI schemes with Cycloid",
	}

	cmd.AddCommand(
		NewGetCommand(),
		NewInterpolateCommand(),
	)

	return cmd
}
