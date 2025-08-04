package stacks

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/printer/factory"
)

func NewGetCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:     "get",
		Args:    cobra.NoArgs,
		Short:   "get information on a stack",
		Example: `cy --org my-org stacks get --ref my:stack-ref`,
		RunE:    get,
	}

	cmd.MarkFlagRequired(cyargs.AddStackRefFlag(cmd))
	return cmd
}

func get(cmd *cobra.Command, args []string) error {
	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}

	ref, err := cyargs.GetStackRef(cmd)
	if err != nil {
		return err
	}

	output, err := cyargs.GetOutput(cmd)
	if err != nil {
		return errors.Wrap(err, "unable to get output flag")
	}

	// Initialize middleware after all arguments are collected
	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	// fetch the printer from the factory
	p, err := factory.GetPrinter(output)
	if err != nil {
		return errors.Wrap(err, "unable to get printer")
	}

	s, err := m.GetStack(org, ref)
	if err != nil {
		return printer.SmartPrint(p, nil, err, "failed to get stack from API", printer.Options{}, cmd.OutOrStderr())
	}

	return printer.SmartPrint(p, s, nil, "", printer.Options{}, cmd.OutOrStdout())
}
