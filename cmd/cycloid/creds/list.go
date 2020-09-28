package creds

import (
	"os"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/common"
	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/youdeploy-cli/printer"
	"github.com/cycloidio/youdeploy-cli/printer/factory"
)

func NewListCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "list",
		Short: "list the credentials",
		Example: `
	# list the credentials with the org 'my-org' in JSON format
	cy --org my-org credentials list -o json
`,
		RunE: list,
	}

	WithFlagType(cmd)
	return cmd
}

func list(cmd *cobra.Command, args []string) error {
	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	org, err := cmd.Flags().GetString("org")
	if err != nil {
		return err
	}
	credT, err := cmd.Flags().GetString("type")
	if err != nil {
		return err
	}
	output, err := cmd.Flags().GetString("output")
	if err != nil {
		return errors.Wrap(err, "unable to get output flag")
	}

	creds, err := m.ListCredentials(org, credT)
	if err != nil {
		return errors.Wrap(err, "unable to list credentials")
	}

	// fetch the printer from the factory
	p, err := factory.GetPrinter(output)
	if err != nil {
		return errors.Wrap(err, "unable to get printer")
	}

	// print the result on the standard output
	if err := p.Print(creds, printer.Options{}, os.Stdout); err != nil {
		return errors.Wrap(err, "unable to print result")
	}

	return nil
}
