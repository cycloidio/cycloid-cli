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
	cyargs.AddStackVersionFlags(cmd)
	return cmd
}

func getComponentConfig(cmd *cobra.Command, args []string) error {
	org, project, env, component, err := cyargs.GetCyContext(cmd)
	if err != nil {
		return err
	}

	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	var tag, branch, hash string
	versionID, ok, err := cyargs.GetStackVersionID(cmd)
	if err != nil {
		return err
	}
	if !ok {
		tag, branch, hash, err = cyargs.ResolveStackVersionArg(cmd, m, org, "")
		if err != nil {
			return err
		}
	}

	config, _, err := m.GetComponentConfig(org, project, env, component, tag, branch, hash, versionID)
	return cyout.PrintWithOptions(cmd, config, err, "failed to fetch config of component '"+component+"'", printer.Options{})
}
