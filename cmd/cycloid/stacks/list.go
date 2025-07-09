package stacks

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/internal"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/printer/factory"
)

func NewListCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use: "list",
		Aliases: []string{
			"ls",
		},
		Args:    cobra.NoArgs,
		Short:   "list the stacks",
		Example: `cy --org my-org stack list`,
		RunE:    list,
		PreRunE: internal.CheckAPIAndCLIVersion,
	}

	return cmd
}

func list(cmd *cobra.Command, args []string) error {
	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}

	output, err := cyargs.GetOutput(cmd)
	if err != nil {
		return err
	}

	// fetch the printer from the factory
	p, err := factory.GetPrinter(output)
	if err != nil {
		return errors.Wrap(err, "unable to get printer")
	}

	stacks, err := m.ListStacks(org)
	if err != nil {
		return printer.SmartPrint(p, nil, err, "failed to list stacks from API", printer.Options{}, cmd.OutOrStdout())
	}

	return printer.SmartPrint(p, stacks, nil, "", printer.Options{}, cmd.OutOrStdout())
}
