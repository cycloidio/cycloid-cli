package pipelines

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/common"
	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/middleware"
)

func NewPauseCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "pause",
		Short: "pause a pipeline",
		Example: `
	# pause pipeline my-project-env
	cy --org my-org pipeline pause --project my-project --env env
`,
		RunE: pause,
	}

	common.RequiredPersistentFlag(common.WithFlagProject, cmd)
	common.RequiredPersistentFlag(common.WithFlagEnv, cmd)

	return cmd
}

func pause(cmd *cobra.Command, args []string) error {
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

	if err := m.PausePipeline(org, project, env); err != nil {
		return errors.Wrap(err, "unable to pause pipeline")
	}

	return nil
}
