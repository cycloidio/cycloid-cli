package widget

import (
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
	"github.com/cycloidio/cycloid-cli/printer"
)

func NewListCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Args:  cobra.NoArgs,
		Short: "List plugin widgets for a component",
		Example: `
  cy plugin component widget list --project my-project --env production --component my-component
`,
		RunE: listComponentPluginWidgets,
	}

	cyargs.AddProjectFlag(cmd)
	cyargs.AddEnvFlag(cmd)
	cyargs.AddComponentFlag(cmd)
	return cmd
}

func listComponentPluginWidgets(cmd *cobra.Command, args []string) error {
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

	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	result, _, err := m.ListComponentPluginWidgets(org, project, env, component)
	return cyout.PrintWithOptions(cmd, result, err, "unable to list component plugin widgets", printer.Options{})
}
