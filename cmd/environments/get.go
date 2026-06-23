package environments

import (
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/apiclient"
	"github.com/cycloidio/cycloid-cli/cmd/common"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
	"github.com/cycloidio/cycloid-cli/gen/models"
)

func NewGetCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get [canonical...]",
		Args:  cyargs.RequireArgsOrFlag("env"),
		Short: "get an environment",
		Example: `
	# get an environment using the --env flag
	cy --org my-org environment get --env my-env -o yaml

	# get multiple environments using positional args
	cy --org my-org environment get env-a env-b
`,
		RunE: get,
	}

	cyargs.AddEnvFlag(cmd)
	cyout.RegisterModel(cmd, models.Environment{})
	return cmd
}

func get(cmd *cobra.Command, args []string) error {
	api := common.NewAPI()
	m := apiclient.NewMiddleware(api)

	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}

	if envFlag, err := cyargs.GetEnvOrEmpty(cmd); err != nil {
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
		e, _, err := m.GetOrgEnv(org, args[0])
		return cyout.PrintWithOptions(cmd, e, err, "unable to get environment", environmentTableOptions)
	}

	results := make([]*models.Environment, 0, len(args))
	for _, env := range args {
		e, _, err := m.GetOrgEnv(org, env)
		if err != nil {
			return cyout.PrintWithOptions(cmd, nil, err, "unable to get environment "+env, environmentTableOptions)
		}
		results = append(results, e)
	}
	return cyout.PrintWithOptions(cmd, results, nil, "", environmentTableOptions)
}

func projectEnvToEnvironment(env *models.ProjectEnvironment) *models.Environment {
	if env == nil {
		return nil
	}
	return &models.Environment{
		Canonical:       env.Canonical,
		CloudAccounts:   env.CloudAccounts,
		Components:      env.Components,
		CreatedAt:       env.CreatedAt,
		Description:     env.Description,
		EnvironmentType: env.EnvironmentType,
		ID:              env.ID,
		Name:            env.Name,
		Owner:           env.Owner,
		UpdatedAt:       env.UpdatedAt,
		VersionStatus:   env.VersionStatus,
	}
}
