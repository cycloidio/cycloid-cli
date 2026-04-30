package apikey

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/client/models"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
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
	_ = cmd.Flags().MarkHidden(can)
	cyout.RegisterModel(cmd, models.APIKey{})

	return cmd
}

// get will send the GET request to the API in order to
// get the generated token
func get(cmd *cobra.Command, args []string) error {
	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return fmt.Errorf("unable to get org flag: %w", err)
	}

	if can, err := cyargs.GetAPIKeyCanonical(cmd); err != nil {
		return fmt.Errorf("unable to get canonical flag: %w", err)
	} else if can != "" {
		found := false
		for _, a := range args {
			if a == can {
				found = true
				break
			}
		}
		if !found {
			args = append(args, can)
		}
	}

	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	if len(args) == 1 {
		key, _, err := m.GetAPIKey(org, args[0])
		return cyout.PrintWithOptions(cmd, key, err, "unable to get API key", apiKeyTableOptions)
	}

	results := make([]*models.APIKey, 0, len(args))
	for _, canonical := range args {
		key, _, err := m.GetAPIKey(org, canonical)
		if err != nil {
			return cyout.PrintWithOptions(cmd, nil, err, "unable to get API key "+canonical, apiKeyTableOptions)
		}
		results = append(results, key)
	}
	return cyout.PrintWithOptions(cmd, results, nil, "", apiKeyTableOptions)
}
