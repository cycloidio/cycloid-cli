package inventory

import (
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
)

func newOutputsListCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Args:  cobra.NoArgs,
		Short: "List inventory outputs",
		Example: `
  # List all outputs for an org
  cy --org my-org inventory outputs list

  # Filter by project and environment
  cy --org my-org inventory outputs list --project my-project --env prod

  # Show only pinned outputs as JSON
  cy --org my-org inventory outputs list --pinned -o json

  # Use a raw LHS filter
  cy --org my-org inventory outputs list --filter 'output_type[eq]=string'
`,
		RunE: outputsList,
	}

	addScopingFlags(cmd)
	cyargs.AddLHSFilterFlag(cmd)
	cmd.Flags().String("key", "", "Filter by output key name")
	cmd.Flags().String("type", "", "Filter by output type (e.g. string, number, bool, list, map, object)")
	cmd.Flags().Bool("pinned", false, "Show only pinned outputs")
	cyout.RegisterModel(cmd, middleware.InventoryOutput{})

	return cmd
}

func outputsList(cmd *cobra.Command, args []string) error {
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

	if key, _ := cmd.Flags().GetString("key"); key != "" {
		filters = append(filters, middleware.LHSFilter{Attribute: "output_key", Condition: "eq", Value: key})
	}
	if outputType, _ := cmd.Flags().GetString("type"); outputType != "" {
		filters = append(filters, middleware.LHSFilter{Attribute: "output_type", Condition: "eq", Value: outputType})
	}
	if cmd.Flags().Changed("pinned") {
		pinned, _ := cmd.Flags().GetBool("pinned")
		val := "false"
		if pinned {
			val = "true"
		}
		filters = append(filters, middleware.LHSFilter{Attribute: "output_is_pinned", Condition: "eq", Value: val})
	}

	// Step 2: API + middleware.
	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	// Step 3: call + print.
	outputs, _, err := m.ListInventoryOutputs(org, filters...)
	return cyout.PrintWithOptions(cmd, outputs, err, "unable to list inventory outputs", outputsTableOptions)
}
