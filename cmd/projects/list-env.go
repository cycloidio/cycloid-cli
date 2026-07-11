package projects

import (
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/apiclient"
	"github.com/cycloidio/cycloid-cli/cmd/common"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
)

func NewListEnvCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "list-env",
		Args:    cobra.NoArgs,
		Short:   "List environments in the current project",
		Example: `cy --org my-org projects list-env -p project -o json`,
		RunE:    listEnv,
	}

	cyargs.AddProjectFlag(cmd)
	return cmd
}

func listEnv(cmd *cobra.Command, args []string) error {
	api := common.NewAPI()
	m := apiclient.NewAPIClient(api)

	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}

	project, err := cyargs.GetProject(cmd)
	if err != nil {
		return err
	}

	projectEnvs, _, err := m.ListProjectEnvs(org, project)
	return cyout.Print(cmd, projectEnvs, err, "unable to listenv project")
}
