package plugin

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
	"github.com/cycloidio/cycloid-cli/printer"
)

func NewCreateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "create <registry-id-or-name>",
		Args:              cobra.ExactArgs(1),
		ValidArgsFunction: cyargs.CompletePluginRegistryID,
		Short:             "Create a plugin in a registry",
		Example: `
  cy plugin registry plugin create my-registry --name my-plugin
  cy plugin registry plugin create my-registry --name my-plugin --update
`,
		RunE: createRegistryPlugin,
	}

	cmd.MarkFlagRequired(cyargs.AddNameFlag(cmd))
	cmd.Flags().Bool("update", false, "update the plugin if it already exists in the registry")
	return cmd
}

func createRegistryPlugin(cmd *cobra.Command, args []string) error {
	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}

	name, err := cyargs.GetName(cmd)
	if err != nil {
		return err
	}

	update, err := cmd.Flags().GetBool("update")
	if err != nil {
		return err
	}

	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	registryID, err := cyargs.ResolvePluginRegistryID(org, args[0], m)
	if err != nil {
		return err
	}

	// Check if a plugin with this name already exists in the registry.
	existingID, resolveErr := cyargs.ResolveRegistryPluginID(org, registryID, name, m)
	exists := resolveErr == nil
	if resolveErr != nil && !isNoMatchError(resolveErr) {
		return cyout.PrintWithOptions(cmd, nil, resolveErr, "failed to check if registry plugin exists", printer.Options{})
	}

	if exists && !update {
		return cyout.PrintWithOptions(cmd, nil,
			fmt.Errorf("plugin %q already exists in this registry; use --update or `cy plugin registry plugin update`", name),
			"unable to create registry plugin", printer.Options{})
	}

	if exists {
		result, _, err := m.UpdateRegistryPlugin(org, registryID, existingID, name)
		return cyout.PrintWithOptions(cmd, result, err, "unable to update registry plugin", printer.Options{})
	}

	result, _, err := m.CreateRegistryPlugin(org, registryID, name)
	return cyout.PrintWithOptions(cmd, result, err, "unable to create registry plugin", printer.Options{})
}

// isNoMatchError returns true when err is the "no X found matching" sentinel
// from cyargs.resolveUnique — indicating the resource simply does not exist yet.
func isNoMatchError(err error) bool {
	return err != nil && strings.Contains(err.Error(), "found matching")
}
