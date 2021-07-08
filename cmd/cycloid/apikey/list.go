package apikey

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/printer/factory"
)

// NewListCommand returns the cobra command holding
// the list API key subcommand
func NewListCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "list API keys",
		Example: `
	# list API keys in the org my-org
	cy api-key list --org my-org
`,
		RunE: func(cmd *cobra.Command, args []string) error {
			org, err := cmd.Flags().GetString("org")
			if err != nil {
				return fmt.Errorf("unable to get org flag: %w", err)
			}
			output, err := cmd.Flags().GetString("output")
			if err != nil {
				return fmt.Errorf("unable to get output flag: %w", err)
			}
			return list(cmd, org, output)
		},
	}
	return cmd
}

// list will send the GET request to the API in order to
// list the generated tokens
func list(cmd *cobra.Command, org, output string) error {
	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	// fetch the printer from the factory
	p, err := factory.GetPrinter(output)
	if err != nil {
		return errors.Wrap(err, "unable to get printer")
	}
	keys, err := m.ListAPIKey(org)
	return printer.SmartPrint(p, keys, err, "unable to list API keys", printer.Options{}, cmd.OutOrStdout())
}
