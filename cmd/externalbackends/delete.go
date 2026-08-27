package externalbackends

import (
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/apiclient"
	"github.com/cycloidio/cycloid-cli/cmd/common"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
	"github.com/cycloidio/cycloid-cli/printer"
)

func NewDeleteCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete",
		Args:  cobra.NoArgs,
		Short: "delete an external backend configuration",
		Example: `
	# delete an existing external backend with ID 123
	cy --org my-org eb delete --id 123

	# delete it along with the Terraform states it stores, once the infrastructure is destroyed
	cy --org my-org eb delete --id 123 --delete-tfstate
`,
		RunE: del,
	}
	common.RequiredFlag(common.WithFlagID, cmd)
	cmd.Flags().
		Bool("delete-tfstate", false, "Also delete the Terraform states stored on the external backend. Only for external backends with the 'remote_tfstate' purpose.")

	return cmd
}

func del(cmd *cobra.Command, args []string) error {
	api := common.NewAPI()
	m := apiclient.NewAPIClient(api)

	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}
	id, err := cmd.Flags().GetUint32("id")
	if err != nil {
		return err
	}
	deleteTFState, err := cmd.Flags().GetBool("delete-tfstate")
	if err != nil {
		return err
	}

	opts := apiclient.DeleteExternalBackendOptions{DeleteTFState: deleteTFState}

	_, err = m.DeleteExternalBackend(org, id, opts)
	return cyout.PrintWithOptions(cmd, nil, err, "unable to delete external backend", printer.Options{})
}
