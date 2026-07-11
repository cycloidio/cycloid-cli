package inventory

import (
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/apiclient"
	"github.com/cycloidio/cycloid-cli/cmd/common"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
	"github.com/cycloidio/cycloid-cli/gen/models"
)

// resourceNamedFilters maps inline list flags to their LHS attribute names.
var resourceNamedFilters = []struct {
	flag string
	attr string
	desc string
}{
	{"provider", "resources_provider", "Filter by resource provider (e.g. aws, google, azurerm)"},
	{"type", "resources_type", "Filter by resource type (e.g. aws_instance)"},
	{"name", "resources_name", "Filter by resource name (display name)"},
	{"module", "resources_module", "Filter by module name"},
	{"mode", "resources_mode", "Filter by resource mode (managed or data)"},
	{"label", "resources_label", "Filter by resource label"},
}

func newResourcesListCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Args:  cobra.NoArgs,
		Short: "List inventory resources",
		Example: `
  # List all resources for an org
  cy --org my-org inventory resources list

  # Filter by provider and type
  cy --org my-org inventory resources list --provider aws --type aws_instance

  # Filter by project and environment as JSON
  cy --org my-org inventory resources list --project my-project --env prod -o json

  # Use a raw LHS filter
  cy --org my-org inventory resources list --filter 'resources_label[eq]=compute'
`,
		RunE: resourcesList,
	}

	addScopingFlags(cmd)
	cyargs.AddLHSFilterFlag(cmd)
	for _, f := range resourceNamedFilters {
		cmd.Flags().String(f.flag, "", f.desc)
	}
	cyout.RegisterModel(cmd, models.InventoryResource{})

	return cmd
}

func resourcesList(cmd *cobra.Command, args []string) error {
	// Step 1: all flags first.
	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}

	filters, err := cyargs.GetLHSFilters(cmd)
	if err != nil {
		return err
	}
	filters, err = appendScopingFilters(cmd, filters)
	if err != nil {
		return err
	}

	for _, f := range resourceNamedFilters {
		if val, _ := cmd.Flags().GetString(f.flag); val != "" {
			filters = append(filters, apiclient.LHSFilter{Attribute: f.attr, Condition: "eq", Value: val})
		}
	}

	// Step 2: API + apiclient.
	api := common.NewAPI()
	m := apiclient.NewAPIClient(api)

	// Step 3: call + print.
	resources, _, err := m.ListInventoryResources(org, filters...)
	return cyout.PrintWithOptions(cmd, resources, err, "unable to list inventory resources", resourcesTableOptions)
}
