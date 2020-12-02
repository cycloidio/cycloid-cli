package apikey

import (
	"fmt"
	"os"

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
			return list(org, output)
		},
	}
	return cmd
}

// list will send the GET request to the API in order to
// list the generated tokens
func list(org, output string) error {
	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	keys, err := m.ListAPIKey(org)
	if err != nil {
		return fmt.Errorf("unable to list API keys: %w", err)
	}

	// fetch the printer from the factory
	p, err := factory.GetPrinter(output)
	if err != nil {
		return fmt.Errorf("unable to get printer: %w", err)
	}

	// print the result on the standard output
	if err := p.Print(keys, printer.Options{}, os.Stdout); err != nil {
		return fmt.Errorf("unable to print result: %w", err)
	}
	return nil
}
