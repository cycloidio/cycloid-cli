package infrapolicies

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

// NewGetCommand returns the cobra command
// to get a infrapolicy
func NewGetCommand() *cobra.Command {

	var cmd = &cobra.Command{
		Use:   "get",
		Args:  cobra.NoArgs,
		Short: "get a infrapolicy",
		Example: `
	# get a infrapolicy my_policy
	cy --org my-org ip get \
	   --canonical my_policy 
		`,
		RunE:    get,
		PreRunE: internal.CheckAPIAndCLIVersion,
	}
	common.RequiredFlag(WithFlagcanonical, cmd)

	return cmd

}

func get(cmd *cobra.Command, args []string) error {
	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}

	canonical, err := cmd.Flags().GetString("canonical")
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

	res, err := m.GetInfraPolicy(org, canonical)
	return printer.SmartPrint(p, res, err, "unable to get infrapolicy", printer.Options{}, cmd.OutOrStdout())

}
