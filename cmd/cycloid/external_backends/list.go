package external_backends

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/internal"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cy_args"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/printer/factory"
)

var (
	example = `
	# List all the external backends within my-org organization in JSON output format
	cy --org my-org external-backends list --output=json

	# List all the external backends within my-org organization in YAML output format
	cy --org my-org external-backends list --output=yaml
`
	short = "Get the list of organization external backends"
	long  = short
)

func NewListCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:     "list",
		Example: example,
		Short:   short,
		Long:    long,
		RunE:    list,
		PreRunE: internal.CheckAPIAndCLIVersion,
	}

	return cmd
}

func list(cmd *cobra.Command, args []string) error {
	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	org, err := cy_args.GetOrg(cmd)
	if err != nil {
		return err
	}
	output, err := cmd.Flags().GetString("output")
	if err != nil {
		return err
	}

	// fetch the printer from the factory
	p, err := factory.GetPrinter(output)
	if err != nil {
		return errors.Wrap(err, "unable to get printer")
	}

	ebs, err := m.ListExternalBackends(org)
	return printer.SmartPrint(p, ebs, err, "unable to list external backends", printer.Options{}, cmd.OutOrStdout())
}
