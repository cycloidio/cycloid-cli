package stacks

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/internal"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/printer/factory"
)

var refFlag string

func NewGetCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "get",
		Short: "get a stack",
		Example: `
	# get a stack in 'my-org' using its ref
	cy --org my-org stacks get --ref my:stack-ref
`,
		RunE:    get,
		PreRunE: internal.CheckAPIAndCLIVersion,
	}
	common.RequiredPersistentFlag(common.WithFlagOrg, cmd)

	cmd.Flags().StringVar(&refFlag, "ref", "", "referential of the stack")
	cmd.MarkFlagRequired("ref")

	return cmd
}

func get(cmd *cobra.Command, args []string) error {
	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	org, err := cmd.Flags().GetString("org")
	if err != nil {
		return err
	}

	ref, err := cmd.Flags().GetString("ref")
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

	s, err := m.GetStack(org, ref)
	return printer.SmartPrint(p, s, err, "unable to get stack", printer.Options{}, cmd.OutOrStdout())
}


