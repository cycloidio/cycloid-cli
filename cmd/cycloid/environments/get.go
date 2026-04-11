package environments

import (
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/client/models"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
)

func NewGetCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "get",
		Args:  cobra.NoArgs,
		Short: "get a environment",
		Example: `
	# get a environment in YAML format
	cy --org my-org environment get --environment my-environment -o yaml
`,
		RunE: get,
	}

	cyargs.AddProjectFlag(cmd)
	cyargs.AddEnvFlag(cmd)
	cyout.RegisterModel(cmd, models.Environment{})
	return cmd
}

func get(cmd *cobra.Command, args []string) error {
	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

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

	proj, _, err := m.GetEnv(org, project, env)
	return cyout.PrintWithOptions(cmd, proj, err, "unable to get environment", environmentTableOptions)
}
