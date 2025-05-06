package api_key

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

// NewDeleteCommand returns the cobra command holding
// the delete API key subcommand
func NewDeleteCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete",
		Short: "delete an API key",
		Example: `
	# delete the API key 'my-key' in the org my-org
	cy api-key delete --org my-org --canonical my-key
`,
		RunE: remove,
	}

	WithFlagCanonical(cmd)
	cmd.MarkFlagRequired("canonical")
	return cmd
}

// remove will send the DELETE request to the API in order to
// delete a generated token
func remove(cmd *cobra.Command, args []string) error {
	org, err := cy_args.GetOrg(cmd)
	if err != nil {
		return fmt.Errorf("unable to get org flag: %w", err)
	}
	output, err := cmd.Flags().GetString("output")
	if err != nil {
		return fmt.Errorf("unable to get output flag: %w", err)
	}
	canonical, err := cmd.Flags().GetString("canonical")
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
	err = m.DeleteAPIKey(org, canonical)
	return printer.SmartPrint(p, nil, err, "unable to delete API key", printer.Options{}, cmd.OutOrStdout())
}
