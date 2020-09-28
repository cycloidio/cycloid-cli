package projects

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/common"
	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/middleware"
)

func NewDeleteEnvCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "delete-env",
		Short: "delete an environment within a project",
		Example: `
	# delete env 'my-env' in 'my-project'
	cy --org my-org project --project my-project --env my-env
`,
		RunE: deleteEnv,
	}
	common.RequiredPersistentFlag(common.WithFlagEnv, cmd)
	common.RequiredPersistentFlag(common.WithFlagProject, cmd)

	return cmd
}

func deleteEnv(cmd *cobra.Command, args []string) error {
	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	org, err := cmd.Flags().GetString("org")
	if err != nil {
		return err
	}
	project, err := cmd.Flags().GetString("project")
	if err != nil {
		return err
	}
	env, err := cmd.Flags().GetString("env")
	if err != nil {
		return err
	}

	if err := m.DeleteProjectEnv(org, project, env); err != nil {
		return errors.Wrap(err, "unable to delete environment")
	}

	return nil
}
