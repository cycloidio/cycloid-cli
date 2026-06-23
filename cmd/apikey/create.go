package apikey

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/apiclient"
	"github.com/cycloidio/cycloid-cli/cmd/common"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/gen/models"
)

// NewCreateCommand returns the cobra command holding
// the create API key subcommand
func NewCreateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create",
		Args:  cobra.NoArgs,
		Short: "create an API key",
		Example: `# Create an admin API key
cy --org myOrg api-key create --canonical "admin-api-key" --owner "Admin" --rules '[{"action": "organization:**", "effect": "allow", "resources": []}]'

# Recreate an existing API key with new rules
cy --org myOrg api-key create --canonical "admin-api-key" --rules '[...]' --recreate`,
		RunE: create,
	}

	cyargs.AddAPIKeyNameFlag(cmd)
	cyargs.AddAPIKeyCanonicalFlag(cmd)
	cyargs.AddOwnerFlag(cmd)
	cyargs.AddAPIKeyRulesFlag(cmd)
	cyargs.AddDescriptionFlag(cmd)
	cmd.Flags().Bool("recreate", false, "delete the API key if it already exists and create a new one with the provided settings")

	return cmd
}

// create the generated tokens
func create(cmd *cobra.Command, args []string) error {
	api := common.NewAPI()
	m := apiclient.NewMiddleware(api)

	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}

	apiKey, _ := cyargs.GetAPIKeyCanonical(cmd)
	apiKeyName, _ := cyargs.GetAPIKeyName(cmd)
	if apiKey == "" && apiKeyName == "" {
		return fmt.Errorf("missing either --canonical or --name flags")
	}

	if apiKey == "" {
		apiKey = common.GenerateCanonical(apiKeyName)
	}

	if apiKeyName == "" {
		apiKeyName = apiKey
	}

	recreate, err := cmd.Flags().GetBool("recreate")
	if err != nil {
		return err
	}

	owner, err := cyargs.GetOwner(cmd)
	if err != nil {
		return err
	}

	rules, err := cyargs.GetAPIKeyRules(cmd)
	if err != nil {
		return err
	}

	description, err := cyargs.GetAPIKeyDescription(cmd)
	if err != nil {
		return err
	}

	var rulesModel []*models.NewRule
	err = json.Unmarshal([]byte(rules), &rulesModel)
	if err != nil {
		return fmt.Errorf("failed to read rules argument as JSON: %w", err)
	}

	_, _, getErr := m.GetAPIKey(org, apiKey)
	if getErr == nil {
		if !recreate {
			return fmt.Errorf("API key %q already exists; use --recreate to delete and recreate it", apiKey)
		}
		_, err := m.DeleteAPIKey(org, apiKey)
		if err != nil {
			return cyout.PrintWithOptions(cmd, nil, err, "unable to delete existing API key before recreation", printer.Options{})
		}
	}

	key, _, err := m.CreateAPIKey(org, apiKey, description, owner, &apiKeyName, rulesModel)
	return cyout.PrintWithOptions(cmd, key, err, "unable to create API key", printer.Options{})
}
