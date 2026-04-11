package components

import (
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/client/models"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
)

func NewGetComponentCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get -p project -e env -c component",
		Args:  cobra.NoArgs,
		Short: "Get the state of a current component.",
		RunE:  getComponent,
	}
	cyargs.AddCyContext(cmd)
	cyout.RegisterModel(cmd, models.Component{})
	return cmd
}

func getComponent(cmd *cobra.Command, args []string) error {
	org, project, env, component, err := cyargs.GetCyContext(cmd)
	if err != nil {
		return err
	}

	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	componentState, _, err := m.GetComponent(org, project, env, component)
	return cyout.PrintWithOptions(cmd, componentState, err,
		"failed to fetch state of component '"+component+"'", componentTableOptions)
}
