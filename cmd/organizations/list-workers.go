package organizations

import (
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/apiclient"
	"github.com/cycloidio/cycloid-cli/cmd/common"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
	"github.com/cycloidio/cycloid-cli/printer"
)

func NewListWorkersCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-workers",
		Args:  cobra.NoArgs,
		Short: "list the organization workers",
		RunE:  listWorkers,
	}

	return cmd
}

func listWorkers(cmd *cobra.Command, args []string) error {
	api := common.NewAPI()
	m := apiclient.NewAPIClient(api)

	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}

	ws, _, err := m.ListOrganizationWorkers(org)
	return cyout.PrintWithOptions(cmd, ws, err, "unable to list organization workers", printer.Options{})
}
