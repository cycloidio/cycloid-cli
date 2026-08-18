package apikey

import (
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/apiclient"
	"github.com/cycloidio/cycloid-cli/cmd/common"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/gen/models"
)

var apiKeyTableOptions = printer.Options{
	Columns:    []string{"Canonical", "Name", "LastSeven", "LastUsed"},
	Identifier: "Canonical",
}

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
	cyout.RegisterModel(cmd, models.APIKey{})
	return cmd
}

// list the generated tokens
func list(cmd *cobra.Command, args []string) error {
	api := common.NewAPI()
	m := apiclient.NewAPIClient(api)

	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}

	keys, _, err := m.ListAPIKeys(org)
	return cyout.PrintWithOptions(cmd, keys, err, "unable to list API keys", apiKeyTableOptions)
}
