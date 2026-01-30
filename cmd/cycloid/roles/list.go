package roles

import (
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/printer/factory"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func NewListCommand() *cobra.Command {
	var (
		example = `cy --org my-org roles list`
		short   = "Get the list of the current organization roles"
		long    = short
	)

	var cmd = &cobra.Command{
		Use:     "list",
		Args:    cobra.NoArgs,
		Example: example,
		Short:   short,
		Long:    long,
		RunE:    listRoles,
	}

	return cmd
}

func listRoles(cmd *cobra.Command, args []string) error {
	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	output, err := cyargs.GetOutput(cmd)
	if err != nil {
		return err
	}

	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}

	p, err := factory.GetPrinter(output)
	if err != nil {
		return errors.Wrap(err, "unable to get printer")
	}

	roles, err := m.ListRoles(org)
	if err != nil {
		return printer.SmartPrint(p, nil, err, "unable to list roles", printer.Options{}, cmd.OutOrStderr())
	}

	return printer.SmartPrint(p, roles, nil, "unable to list roles", printer.Options{}, cmd.OutOrStdout())
}
