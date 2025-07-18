package credentials

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/printer/factory"
)

func NewListCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:     "list",
		Args:    cobra.NoArgs,
		Short:   "list the credentials",
		Example: `cy --org my-org credentials list -o json`,
		RunE:    list,
	}

	cyargs.AddCredentialTypeFlag(cmd)
	return cmd
}

func list(cmd *cobra.Command, args []string) error {
	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}
	credT, err := cyargs.GetCredentialType(cmd)
	if err != nil {
		return err
	}
	output, err := cyargs.GetOutput(cmd)
	if err != nil {
		return errors.Wrap(err, "unable to get output flag")
	}

	// fetch the printer from the factory
	p, err := factory.GetPrinter(output)
	if err != nil {
		return errors.Wrap(err, "unable to get printer")
	}

	creds, err := m.ListCredentials(org, credT)
	if err != nil {
		return printer.SmartPrint(p, nil, err, "unable to list credential", printer.Options{}, cmd.OutOrStderr())
	}

	return printer.SmartPrint(p, creds, nil, "", printer.Options{}, cmd.OutOrStdout())
}
