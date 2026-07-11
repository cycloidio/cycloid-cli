package environments

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/apiclient"
	"github.com/cycloidio/cycloid-cli/cmd/common"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
)

func NewDeleteCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "delete [canonical...]",
		Args:    cyargs.RequireArgsOrFlag("env"),
		Aliases: []string{"del", "rm"},
		Short:   "delete an environment",
		Example: `
	# delete an environment using the --env flag
	cy --org my-org environment delete --env my-env

	# delete multiple environments using positional args
	cy --org my-org environment delete env-a env-b
`,
		RunE: deleteEnvironment,
	}

	cyargs.AddEnvFlag(cmd)
	cyargs.AddProjectFlag(cmd)
	return cmd
}

func deleteEnvironment(cmd *cobra.Command, args []string) error {
	if cyargs.IsSet(cmd, "project") {
		return fmt.Errorf(`--project is no longer accepted on environment delete.
Did you mean cy environment unlink --project %s --env <env>?
Use cy environment delete --env <env> to destroy the org-level environment`, cmd.Flag("project").Value.String())
	}

	org, err := cyargs.GetOrg(cmd)
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
	m := apiclient.NewAPIClient(api)

	for _, env := range args {
		_, err = m.DeleteOrgEnv(org, env)
		if err != nil {
			return cyout.Print(cmd, nil, err, "unable to delete environment "+env)
		}
	}
	return nil
}
