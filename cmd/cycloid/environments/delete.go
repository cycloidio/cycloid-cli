package environments

import (
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
)

func NewDeleteCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:     "delete [canonical...]",
		Args:    cyargs.RequireArgsOrFlag("env"),
		Aliases: []string{"del", "rm"},
		Short:   "delete an environment",
		Example: `
	# delete an environment using the --env flag
	cy --org my-org environment delete --project my-proj --env my-env

	# delete multiple environments using positional args
	cy --org my-org environment delete --project my-proj env-a env-b
`,
		RunE: deleteEnvironment,
	}

	cyargs.AddProjectFlag(cmd)
	cyargs.AddEnvFlag(cmd)
	return cmd
}

func deleteEnvironment(cmd *cobra.Command, args []string) error {
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

	var envs []string
	if len(args) > 0 {
		envs = args
	} else {
		env, err := cyargs.GetEnv(cmd)
		if err != nil {
			return err
		}
		envs = []string{env}
	}

	for _, env := range envs {
		_, err = m.DeleteEnv(org, project, env)
		if err != nil {
			return cyout.Print(cmd, nil, err, "unable to delete environment "+env)
		}
	}
	return nil
}
