package apikey

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/printer/factory"
)

// NewListCommand returns the cobra command holding
// the list API key subcommand
func NewListCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Args:  cobra.NoArgs,
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

	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}

	output, err := cyargs.GetOutput(cmd)
	if err != nil {
		return err
	}

	// fetch the printer from the factory
	p, err := factory.GetPrinter(output)
	if err != nil {
		return errors.Wrap(err, "unable to get printer")
	}

	keys, err := m.ListAPIKeys(org)
	if err != nil {
		return printer.SmartPrint(p, nil, err, "unable to list API keys", printer.Options{}, cmd.OutOrStderr())
	}
	return printer.SmartPrint(p, keys, nil, "", printer.Options{}, cmd.OutOrStdout())
}
