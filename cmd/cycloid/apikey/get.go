package apikey

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/printer/factory"
)

// NewGetCommand returns the cobra command holding
// the get API key subcommand
func NewGetCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
		Args:  cobra.NoArgs,
		Short: "get API key",
		Example: `
	# get API key 'my-key' in the org my-org
	cy api-key get --org my-org --canonical my-key
`,
		RunE: get,
	}

	cyargs.AddAPIKeyCanonicalFlag(cmd)
	cmd.MarkFlagRequired("canonical")
	return cmd
}

// get will send the GET request to the API in order to
// get the generated token
func get(cmd *cobra.Command, args []string) error {
	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return fmt.Errorf("unable to get org flag: %w", err)
	}

	output, err := cyargs.GetOutput(cmd)
	if err != nil {
		return fmt.Errorf("unable to get output flag: %w", err)
	}

	canonical, err := cyargs.GetAPIKeyCanonical(cmd)
	if err != nil {
		return fmt.Errorf("unable to get canonical flag: %w", err)
	}

	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	// fetch the printer from the factory
	p, err := factory.GetPrinter(output)
	if err != nil {
		return errors.Wrap(err, "unable to get printer")
	}

	key, err := m.GetAPIKey(org, canonical)
	if err != nil {
		return printer.SmartPrint(p, nil, err, "unable to get API key", printer.Options{}, cmd.OutOrStderr())
	}

	return printer.SmartPrint(p, key, nil, "", printer.Options{}, cmd.OutOrStdout())
}
