package components

import (
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/apiclient"
	"github.com/cycloidio/cycloid-cli/cmd/common"
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
	return cmd
}

func getComponentConfig(cmd *cobra.Command, args []string) error {
	org, project, env, component, err := cyargs.GetCyContext(cmd)
	if err != nil {
		return err
	}

	api := common.NewAPI()
	m := apiclient.NewMiddleware(api)

	config, _, err := m.GetComponentConfig(org, project, env, component)
	return cyout.PrintWithOptions(cmd, config, err, "failed to fetch config of component '"+component+"'", printer.Options{})
}
