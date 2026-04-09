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
		Use:               "get [canonical...]",
		Args:              cyargs.RequireArgsOrFlag("canonical"),
		ValidArgsFunction: cyargs.CompleteAPIKeyCanonical,
		Short:             "get API key",
		Example: `
	# get API key 'my-key' in the org my-org
	cy --org my-org api-key get my-key

	# get multiple API keys
	cy --org my-org api-key get key-one key-two

	# get using the deprecated --canonical flag
	cy --org my-org api-key get --canonical my-key
`,
		RunE: get,
	}

	// Keep --canonical for backward compatibility
	can := cyargs.AddAPIKeyCanonicalFlag(cmd)
	_ = cmd.Flags().MarkDeprecated(can, "pass the canonical as a positional argument instead")

	return cmd
}

// get will send the GET request to the API in order to
// get the generated token
func get(cmd *cobra.Command, args []string) error {
	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return fmt.Errorf("unable to get org flag: %w", err)
	}

	// Support --canonical for backward compat; positional args take precedence
	if len(args) == 0 {
		can, err := cyargs.GetAPIKeyCanonical(cmd)
		if err != nil {
			return fmt.Errorf("unable to get canonical flag: %w", err)
		}
		args = []string{can}
	}

	output, err := cyargs.GetOutput(cmd)
	if err != nil {
		return fmt.Errorf("unable to get output flag: %w", err)
	}

	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	p, err := factory.GetPrinter(output)
	if err != nil {
		return errors.Wrap(err, "unable to get printer")
	}

	if len(args) == 1 {
		key, _, err := m.GetAPIKey(org, args[0])
		if err != nil {
			return printer.SmartPrint(p, nil, err, "unable to get API key", printer.Options{}, cmd.OutOrStderr())
		}
		return printer.SmartPrint(p, key, nil, "", printer.Options{}, cmd.OutOrStdout())
	}

	results := make([]interface{}, 0, len(args))
	for _, canonical := range args {
		key, _, err := m.GetAPIKey(org, canonical)
		if err != nil {
			return printer.SmartPrint(p, nil, err, "unable to get API key "+canonical, printer.Options{}, cmd.OutOrStderr())
		}
		results = append(results, key)
	}
	return printer.SmartPrint(p, results, nil, "", printer.Options{}, cmd.OutOrStdout())
}
