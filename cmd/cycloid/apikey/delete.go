package apikey

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
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
		RunE: func(cmd *cobra.Command, args []string) error {
			org, err := cmd.Flags().GetString("org")
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
			return remove(org, canonical, output)
		},
	}

	WithFlagCanonical(cmd)
	cmd.MarkFlagRequired("canonical")
	return cmd
}

// remove will send the DELETE request to the API in order to
// delete a generated token
func remove(org, canonical, output string) error {
	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	if err := m.DeleteAPIKey(org, canonical); err != nil {
		return fmt.Errorf("unable to delete API key: %w", err)
	}
	return nil
}
