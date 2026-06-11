package components

import (
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
	"github.com/cycloidio/cycloid-cli/printer"
)

func NewComponentConfigGetCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "get",
		Args:    cobra.NoArgs,
		Short:   "Fetch the current Stackforms variables of a component in JSON format.",
		RunE:    getComponentConfig,
		Example: "cy config get -p project -e env -c component",
	}
	cyargs.AddCyContext(cmd)
	cmd.Flags().Uint32("service-catalog-source-version-id", 0, "service catalog source version ID (default: latest)")
	return cmd
}

func getComponentConfig(cmd *cobra.Command, args []string) error {
	org, project, env, component, err := cyargs.GetCyContext(cmd)
	if err != nil {
		return err
	}

	versionID, err := cmd.Flags().GetUint32("service-catalog-source-version-id")
	if err != nil {
		return err
	}

	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	config, _, err := m.GetComponentConfig(org, project, env, component, versionID)
	return cyout.PrintWithOptions(cmd, config, err, "failed to fetch config of component '"+component+"'", printer.Options{})
}
