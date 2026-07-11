package manager

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/apiclient"
	"github.com/cycloidio/cycloid-cli/cmd/common"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
	"github.com/cycloidio/cycloid-cli/printer"
)

func NewCreateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create",
		Args:  cobra.NoArgs,
		Short: "Register a plugin manager with the organization",
		Example: `
  cy plugin manager create --name my-manager --url https://pm.example.com
  cy plugin manager create --name my-manager --url https://pm.example.com --update
`,
		RunE: createPluginManager,
	}

	_ = cmd.MarkFlagRequired(cyargs.AddNameFlag(cmd))
	cyargs.AddURLFlag(cmd, "URL of the plugin manager instance (required)")
	_ = cmd.MarkFlagRequired("url")
	cmd.Flags().Bool("update", false, "if the plugin manager already exists, return it without failing (idempotent create)")
	return cmd
}

func createPluginManager(cmd *cobra.Command, args []string) error {
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
	m := apiclient.NewAPIClient(api)

	// Check if a manager with this name already exists.
	existingID, resolveErr := cyargs.ResolvePluginManagerID(org, name, m)
	exists := resolveErr == nil
	if resolveErr != nil && !cyargs.IsNoMatchError(resolveErr) {
		return cyout.PrintWithOptions(cmd, nil, resolveErr, "failed to check if plugin manager exists", printer.Options{})
	}

	if exists && !update {
		return cyout.PrintWithOptions(cmd, nil,
			fmt.Errorf("plugin manager %q already exists; use --update for an idempotent create", name),
			"unable to create plugin manager", printer.Options{})
	}

	if exists {
		if cmd.Flags().Changed("url") {
			fmt.Fprintln(cmd.ErrOrStderr(), "note: --url cannot be changed via the API; existing URL is kept")
		}
		result, _, err := m.GetPluginManager(org, existingID)
		return cyout.PrintWithOptions(cmd, result, err, "unable to get existing plugin manager", printer.Options{})
	}

	result, _, err := m.CreatePluginManager(org, name, url, true)
	return cyout.PrintWithOptions(cmd, result, err, "unable to create plugin manager", printer.Options{})
}
