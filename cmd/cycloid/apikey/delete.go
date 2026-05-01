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

// NewDeleteCommand returns the cobra command holding
// the delete API key subcommand
func NewDeleteCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "delete [canonical...]",
		Aliases:           []string{"rm"},
		Args:              cyargs.RequireArgsOrFlag("canonical"),
		ValidArgsFunction: cyargs.CompleteAPIKeyCanonical,
		Short:             "delete an API key",
		Example: `# delete the API key 'my-key' in the org my-org
	cy --org my-org api-key delete my-key

	# delete multiple API keys at once
	cy --org my-org api-key delete key-one key-two

	# delete using the deprecated --canonical flag
	cy --org my-org api-key delete --canonical my-key`,
		RunE: remove,
	}

	// Keep --canonical for backward compatibility
	can := cyargs.AddAPIKeyCanonicalFlag(cmd)
	_ = cmd.Flags().MarkHidden(can)

	return cmd
}

// remove will send the DELETE request to the API in order to
// delete a generated token
func remove(cmd *cobra.Command, args []string) error {
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

	if output == "table" {
		output = "json"
	}
	p, err := factory.GetPrinter(output)
	if err != nil {
		return errors.Wrap(err, "unable to get printer")
	}

	deleted := make([]string, 0, len(args))
	for _, canonical := range args {
		_, err = m.DeleteAPIKey(org, canonical)
		if err != nil {
			return printer.SmartPrint(p, nil, err, "unable to delete API key "+canonical, printer.Options{}, cmd.OutOrStderr())
		}
		deleted = append(deleted, canonical)
	}
	return printer.SmartPrint(p, deleted, nil, "", printer.Options{}, cmd.OutOrStdout())
}
