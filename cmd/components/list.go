package components

import (
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/apiclient"
	"github.com/cycloidio/cycloid-cli/cmd/common"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
	"github.com/cycloidio/cycloid-cli/gen/models"
)

func NewGetComponentsCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list -p project -e env",
		Args:  cobra.NoArgs,
		Short: "List components in a project",
		RunE:  getComponents,
	}
	cyargs.AddProjectFlag(cmd)
	cyargs.AddEnvFlag(cmd)
	cyout.RegisterModel(cmd, models.Component{})
	return cmd
}

func getComponents(cmd *cobra.Command, args []string) error {
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

	api := common.NewAPI()
	m := apiclient.NewAPIClient(api)

	components, _, err := m.ListComponents(org, project, env)
	errMsg := "failed to fetch list of components in '" + project + "', '" + env + "'"
	return cyout.PrintWithOptions(cmd, components, err, errMsg, componentTableOptions)
}
