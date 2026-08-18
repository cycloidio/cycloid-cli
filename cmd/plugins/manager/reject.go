package manager

import (
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/internal/cyargs"
)

func NewRejectCommand() *cobra.Command {
	return &cobra.Command{
		Use:               "reject <id-or-name>",
		Args:              cobra.ExactArgs(1),
		ValidArgsFunction: cyargs.CompletePluginManagerID,
		Short:             "Reject a plugin manager invitation",
		Example: `
  cy plugin manager reject 1
  cy plugin manager reject my-manager
`,
		RunE: rejectPluginManager,
	}
}

func rejectPluginManager(cmd *cobra.Command, args []string) error {
	return updateInviteStatus(cmd, args, "rejected")
}
