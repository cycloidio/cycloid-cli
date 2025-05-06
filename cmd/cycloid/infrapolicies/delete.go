package infrapolicies

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

// NewDeleteCommand returns the cobra command
// to delete a infrapolicy
func NewDeleteCommand() *cobra.Command {

	var cmd = &cobra.Command{
		Use:   "delete",
		Short: "delete a infrapolicy",
		Example: `
	# create a infrapolicy my_policy
	cy --org my-org ip delete \
	   --canonical my_policy 
		`,
		RunE:    delete,
		PreRunE: internal.CheckAPIAndCLIVersion,
	}
	common.RequiredFlag(WithFlagcanonical, cmd)

	return cmd

}

func delete(cmd *cobra.Command, args []string) error {
	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	org, err := cy_args.GetOrg(cmd)
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

	err = m.DeleteInfraPolicy(org, canonical)
	return printer.SmartPrint(p, nil, err, "unable to delete infrapolicy", printer.Options{}, cmd.OutOrStdout())

}
