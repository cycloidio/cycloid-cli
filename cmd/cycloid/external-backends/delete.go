package externalBackends

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/internal"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/printer/factory"
)

func NewDeleteCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "delete",
		Short: "delete an external backend configuration",
		Example: `
	# delete an existing external backend with ID 123
	cy --org my-org eb delete --id 123
`,
		RunE:    del,
		PreRunE: internal.CheckAPIAndCLIVersion,
	}
	common.RequiredFlag(common.WithFlagID, cmd)
	common.RequiredPersistentFlag(common.WithFlagOrg, cmd)

	return cmd
}

func del(cmd *cobra.Command, args []string) error {
	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	org, err := cmd.Flags().GetString("org")
	if err != nil {
		return err
	}
	id, err := cmd.Flags().GetUint32("id")
	if err != nil {
		return err
	}

	output, err := cmd.Flags().GetString("output")
	if err != nil {
		return errors.Wrap(err, "unable to get output flag")
	}
	// fetch the printer from the factory
	p, err := factory.GetPrinter(output)
	if err != nil {
		return errors.Wrap(err, "unable to get printer")
	}

	err = m.DeleteExternalBackend(org, id)
	return printer.SmartPrint(p, nil, err, "unable to delete external backend", printer.Options{}, cmd.OutOrStdout())
}
