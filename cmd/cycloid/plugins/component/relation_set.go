package component

import (
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
	"github.com/cycloidio/cycloid-cli/printer"
)

func NewRelationSetCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "relation-set <plugin-id-or-name>",
		Args:              cobra.ExactArgs(1),
		ValidArgsFunction: cyargs.CompletePluginInstallID,
		Short:             "Enable or disable a plugin for a component",
		Example: `
  cy plugin component relation-set my-plugin --project my-project --env production --component my-component --enabled
  cy plugin component relation-set 42 --project my-project --env production --component my-component --no-enabled
`,
		RunE: setComponentPluginRelation,
	}

	cyargs.AddProjectFlag(cmd)
	cyargs.AddEnvFlag(cmd)
	cyargs.AddComponentFlag(cmd)
	cmd.Flags().Bool("enabled", true, "Enable or disable the plugin for the component")
	return cmd
}

func setComponentPluginRelation(cmd *cobra.Command, args []string) error {
	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}

	project, err := cyargs.GetProject(cmd)
	if err != nil {
		return err
	}

	env, err := cyargs.GetEnv(cmd)
	if err != nil {
		return err
	}

	component, err := cyargs.GetComponent(cmd)
	if err != nil {
		return err
	}

	enabled, err := cmd.Flags().GetBool("enabled")
	if err != nil {
		return err
	}

	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	pluginInstallID, err := cyargs.ResolvePluginInstallID(org, args[0], m)
	if err != nil {
		return err
	}

	result, _, err := m.SetComponentPluginRelation(org, project, env, component, pluginInstallID, enabled)
	return cyout.PrintWithOptions(cmd, result, err, "unable to set component plugin relation", printer.Options{})
}
