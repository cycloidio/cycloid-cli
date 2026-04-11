package manager

import (
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
	"github.com/cycloidio/cycloid-cli/printer"
)

func NewAcceptCommand() *cobra.Command {
	return &cobra.Command{
		Use:               "accept <id-or-name>",
		Args:              cobra.ExactArgs(1),
		ValidArgsFunction: cyargs.CompletePluginManagerID,
		Short:             "[beta] Accept a plugin manager invitation",
		Example: `
  cy beta plugin manager accept 1
  cy beta plugin manager accept my-manager
`,
		RunE: acceptPluginManager,
	}
}

func acceptPluginManager(cmd *cobra.Command, args []string) error {
	return updateInviteStatus(cmd, args, "accepted")
}

func updateInviteStatus(cmd *cobra.Command, args []string, status string) error {
	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}

	id, err := cyargs.ResolvePluginManagerID(org, args[0], m)
	if err != nil {
		return err
	}

	result, _, err := m.UpdatePluginManager(org, id, status)
	return cyout.PrintWithOptions(cmd, result, err, "unable to update plugin manager", printer.Options{})
}
