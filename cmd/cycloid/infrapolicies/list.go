package infrapolicies

import (
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/internal"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/printer/factory"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// NewListCommand returns the cobra command
// to list the infrapolicies in a organization
func NewListCommand() *cobra.Command {

	var cmd = &cobra.Command{
		Use:   "list",
		Short: "list infrapolicies",
		Example: `
	# list all infrapolicies in an organization
	cy --org my-org ip list -output=table|json|yaml

	# parse the result using
	cy --org my-org ip list -output=json | jq --color-output .
		`,
		RunE:    list,
		PreRunE: internal.CheckAPIAndCLIVersion,
	}

	return cmd

}

func list(cmd *cobra.Command, args []string) error {
	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	org, err := common.GetOrg(cmd)
	if err != nil {
		return err
	}

	//to allow to specify the output flag as specified in cmd/cycloid.go
	output, err := cmd.Flags().GetString("output")
	if err != nil {
		return errors.Wrap(err, "unable to get output flag")
	}

	// fetch the printer from the factory
	p, err := factory.GetPrinter(output)
	if err != nil {
		return errors.Wrap(err, "unable to get printer")
	}

	res, err := m.ListInfraPolicies(org)
	return printer.SmartPrint(p, res, err, "unable to list infrapolicies", printer.Options{}, cmd.OutOrStdout())

}
