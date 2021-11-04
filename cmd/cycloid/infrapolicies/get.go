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

// NewGetCommand returns the cobra command
// to get a infrapolicy
func NewGetCommand() *cobra.Command {

	var cmd = &cobra.Command{
		Use:   "get",
		Short: "get a infrapolicy",
		Example: `
	# get a infrapolicy my_policy
	cy --org my-org ip get \
	   --cannonical my_policy 
		`,
		RunE:    get,
		PreRunE: internal.CheckAPIAndCLIVersion,
	}
	common.RequiredFlag(WithFlagCannonical, cmd)

	return cmd

}

func get(cmd *cobra.Command, args []string) error {
	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	org, err := cmd.Flags().GetString("org")
	if err != nil {
		return err
	}

	cannonical, err := cmd.Flags().GetString("cannonical")
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

	res, err := m.GetInfraPolicy(org, cannonical)
	return printer.SmartPrint(p, res, err, "unable to get infrapolicy", printer.Options{}, cmd.OutOrStdout())

}
