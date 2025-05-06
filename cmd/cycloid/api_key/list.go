package apikey

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cy_args"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/printer/factory"
)

// NewListCommand returns the cobra command holding
// the list API key subcommand
func NewListCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "list API keys",
		Example: `# list API keys in the org my-org
cy api-key list --org my-org`,
		RunE: list,
	}
	return cmd
}

// list the generated tokens
func list(cmd *cobra.Command, args []string) error {
	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	org, err := cy_args.GetOrg(cmd)
	if err != nil {
		return err
	}

	output, err := cmd.Flags().GetString("output")
	if err != nil {
		return fmt.Errorf("unable to get output flag: %w", err)
	}

	// fetch the printer from the factory
	p, err := factory.GetPrinter(output)
	if err != nil {
		return errors.Wrap(err, "unable to get printer")
	}

	keys, err := m.ListAPIKeys(org)
	return printer.SmartPrint(p, keys, err, "unable to list API keys", printer.Options{}, cmd.OutOrStdout())
}
