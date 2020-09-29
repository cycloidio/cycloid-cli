package pipelines

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/common"
	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/middleware"
)

func NewTriggerBuildCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use: "trigger-build",
		Example: `
	# trigger a pipeline build for 'my-job'
	cy --org my-org pp trigger-build --project my-project --env my-env --job my-job
`,
		Short: "trigger a pipeline build",
		RunE:  createBuild,
	}

	common.RequiredPersistentFlag(common.WithFlagProject, cmd)
	common.RequiredPersistentFlag(common.WithFlagEnv, cmd)
	common.RequiredFlag(WithFlagJob, cmd)

	return cmd
}

func createBuild(cmd *cobra.Command, args []string) error {
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
	job, err := cmd.Flags().GetString("job")
	if err != nil {
		return err
	}

	if err := m.TriggerPipelineBuild(org, project, env, job); err != nil {
		return errors.Wrap(err, "unable to trigger pipeline build")
	}

	return nil
}
