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
	cyargs.AddDeleteFlags(cmd)
	return cmd
}

func deleteEnvironment(cmd *cobra.Command, args []string) error {
	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}

	project, err := cyargs.GetProject(cmd)
	if err != nil {
		return err
	}

	force, skipHooks, ignoreConfigFilesErr, err := cyargs.GetDeleteFlags(cmd)
	if err != nil {
		return err
	}

	if envFlag, err := cyargs.GetEnv(cmd); err != nil {
		return err
	} else if envFlag != "" {
		found := false
		for _, a := range args {
			if a == envFlag {
				found = true
				break
			}
		}
		if !found {
			args = append(args, envFlag)
		}
	}

	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	opts := middleware.DeleteOptions{Force: force, SkipHooks: skipHooks, IgnoreConfigFilesErr: ignoreConfigFilesErr}
	for _, env := range args {
		_, err = m.DeleteEnv(org, project, env, opts)
		if err != nil {
			return cyout.Print(cmd, nil, err, "unable to delete environment "+env)
		}
	}
	return nil
}
