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
		Use:   "get [canonical...]",
		Args:  cyargs.RequireArgsOrFlag("env"),
		Short: "get an environment",
		Example: `
	# get an environment using the --env flag
	cy --org my-org environment get --project my-proj --env my-env -o yaml

	# get multiple environments using positional args
	cy --org my-org environment get --project my-proj env-a env-b
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

	if len(args) == 1 {
		e, _, err := m.GetEnv(org, project, args[0])
		return cyout.PrintWithOptions(cmd, e, err, "unable to get environment", environmentTableOptions)
	}

	results := make([]*models.Environment, 0, len(args))
	for _, env := range args {
		e, _, err := m.GetEnv(org, project, env)
		if err != nil {
			return cyout.PrintWithOptions(cmd, nil, err, "unable to get environment "+env, environmentTableOptions)
		}
		results = append(results, e)
	}
	return cyout.PrintWithOptions(cmd, results, nil, "", environmentTableOptions)
}
