package inventory

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
)

func newOutputsGetCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get <key>",
		Args:  cobra.ExactArgs(1),
		Short: "Get the value of a single inventory output",
		Long: `Get the value of a named Terraform state output.

By default this prints only the raw value (via the "value" field output),
which is convenient for scripting:

  IP=$(cy --org my-org inventory outputs get instance_ip --project p --env e)

Output keys are scoped to a project/environment/component; if a key matches
more than one output, narrow the result with --project / --env / --component.

Override the default with -o json (full object), -o yaml, or -o jq=<expr>.`,
		Example: `
  # Print the raw value (default — scriptable)
  cy --org my-org inventory outputs get instance_ip --project my-project --env prod

  # Get the full object as JSON
  cy --org my-org inventory outputs get instance_ip -o json
`,
		RunE: outputsGet,
	}

	addScopingFlags(cmd)
	cyargs.AddLHSFilterFlag(cmd)
	cyout.RegisterModel(cmd, middleware.InventoryOutput{})

	return cmd
}

func outputsGet(cmd *cobra.Command, args []string) error {
	key := args[0]

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
	filters = append(filters, middleware.LHSFilter{Attribute: "output_key", Condition: "eq", Value: key})

	// Step 2: API + middleware.
	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	outputs, _, err := m.ListInventoryOutputs(org, filters...)
	if err != nil {
		return cyout.PrintWithOptions(cmd, nil, err, "unable to fetch inventory output", outputsTableOptions)
	}

	switch len(outputs) {
	case 0:
		return fmt.Errorf("output %q not found in org %q (try --project / --env / --component to scope the search)", key, org)

	case 1:
		// Default to the raw value (field printer) unless the user asked for a
		// specific format via --output / --jq / CY_OUTPUT / config.
		effective, err := cyargs.GetOutput(cmd)
		if err != nil {
			return err
		}
		if effective == "table" && !cmd.Flags().Changed("output") && !cmd.Flags().Changed("jq") {
			if err := cmd.Flags().Set("output", "value"); err != nil {
				return err
			}
		}
		return cyout.PrintWithOptions(cmd, outputs[0], nil, "", outputsTableOptions)

	default:
		return fmt.Errorf("output %q matched %d results across multiple scopes; narrow with --project / --env / --component", key, len(outputs))
	}
}
