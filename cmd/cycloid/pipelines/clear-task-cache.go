package pipelines

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/internal"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
)

func NewClearTaskCacheCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "clear-task-cache",
		Short: "clear cache for a task",
		Example: `
	# clean cache for task 'my-task'
	cy --org my-org pp clear-task-cache --project my-project --job my-job --env my-env --task my-task
`,
		RunE:    cleartaskCache,
		PreRunE: internal.CheckAPIAndCLIVersion,
	}

	common.RequiredPersistentFlag(common.WithFlagProject, cmd)
	common.RequiredPersistentFlag(common.WithFlagEnv, cmd)
	common.RequiredFlag(WithFlagJob, cmd)
	common.RequiredFlag(WithFlagTask, cmd)

	return cmd
}

func cleartaskCache(cmd *cobra.Command, args []string) error {
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
	task, err := cmd.Flags().GetString("task")
	if err != nil {
		return err
	}

	if err := m.ClearTaskCachePipeline(org, project, env, job, task); err != nil {
		return errors.Wrap(err, "unable to clear task cache")
	}

	return nil
}
