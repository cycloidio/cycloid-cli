package registry

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
	"github.com/cycloidio/cycloid-cli/printer"
)

func NewCreateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "create",
		Aliases: []string{"add"},
		Args:    cobra.NoArgs,
		Short:   "Create a plugin registry",
		Example: `
  cy plugin registry create --name my-registry --url https://registry.example.com
  cy plugin registry create --name my-registry --url https://registry.example.com --update
`,
		RunE: createPluginRegistry,
	}

	_ = cmd.MarkFlagRequired(cyargs.AddNameFlag(cmd))
	cyargs.AddURLFlag(cmd, "URL of the plugin registry (required)")
	_ = cmd.MarkFlagRequired("url")
	cmd.Flags().Bool("update", false, "update the plugin registry if it already exists")
	return cmd
}

func createPluginRegistry(cmd *cobra.Command, args []string) error {
	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}

	name, err := cyargs.GetName(cmd)
	if err != nil {
		return err
	}

	url, err := cyargs.GetURL(cmd)
	if err != nil {
		return errors.Wrap(err, "unable to get --url flag")
	}

	update, err := cmd.Flags().GetBool("update")
	if err != nil {
		return err
	}

	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	// Check if a registry with this name already exists.
	existingID, resolveErr := cyargs.ResolvePluginRegistryID(org, name, m)
	exists := resolveErr == nil
	if resolveErr != nil && !cyargs.IsNoMatchError(resolveErr) {
		// List failure or ambiguous match — surface it.
		return cyout.PrintWithOptions(cmd, nil, resolveErr, "failed to check if plugin registry exists", printer.Options{})
	}

	if exists && !update {
		return cyout.PrintWithOptions(cmd, nil,
			fmt.Errorf("plugin registry %q already exists; use --update or `cy plugin registry update`", name),
			"unable to create plugin registry", printer.Options{})
	}

	if exists {
		if cmd.Flags().Changed("url") {
			fmt.Fprintln(cmd.ErrOrStderr(), "note: --url cannot be changed via the API; existing URL is kept")
		}
		result, _, err := m.UpdatePluginRegistry(org, existingID, name)
		return cyout.PrintWithOptions(cmd, result, err, "unable to update plugin registry", printer.Options{})
	}

	result, _, err := m.CreatePluginRegistry(org, name, url)
	return cyout.PrintWithOptions(cmd, result, err, "unable to create plugin registry", printer.Options{})
}
