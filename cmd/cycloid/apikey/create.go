package apikey

import (
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/client/models"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/printer/factory"
)

var (
	apiKeyFlag     string
	apiKeyNameFlag string
)

// NewcreateCommand returns the cobra command holding
// the create API key subcommand
func NewCreateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create",
		Args:  cobra.NoArgs,
		Short: "create an API Key",
		Example: `# Create and Admin API Key
cy --org myOrg api-key create --canonical "admin-api-key" --owner "Admin" --rules '[{"action": "organization:**", "effect": "allow", "resources": []}]'`,
		RunE: create,
	}

	apiKeyNameFlag = cyargs.AddAPIKeyNameFlag(cmd)
	apiKeyFlag = cyargs.AddAPIKeyCanonicalFlag(cmd)
	cyargs.AddOwnerFlag(cmd)
	cyargs.AddAPIKeyRulesFlag(cmd)
	cyargs.AddDescriptionFlag(cmd)

	return cmd
}

// create the generated tokens
func create(cmd *cobra.Command, args []string) error {
	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}

	var apiKey, apiKeyName string
	apiKey, _ = cyargs.GetAPIKeyCanonical(cmd)
	apiKeyName, _ = cyargs.GetAPIKeyName(cmd)
	if apiKey == "" && apiKeyName == "" {
		return fmt.Errorf("missing either '--%s' or '--%s' flags", apiKeyFlag, apiKeyNameFlag)
	}

	if apiKey == "" {
		apiKey = common.GenerateCanonical(apiKeyName)
	}

	if apiKeyName == "" {
		apiKeyName = apiKey
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

	output, err := cyargs.GetOutput(cmd)
	if err != nil {
		return fmt.Errorf("unable to get output flag: %w", err)
	}

	// fetch the printer from the factory
	p, err := factory.GetPrinter(output)
	if err != nil {
		return errors.Wrap(err, "unable to get printer")
	}

	key, err := m.CreateAPIKey(org, apiKey, description, owner, &apiKeyName, rulesModel)
	if err != nil {
		return fmt.Errorf("failed to request API Key: %w", err)
	}

	return printer.SmartPrint(p, key, nil, "", printer.Options{}, cmd.OutOrStdout())
}
