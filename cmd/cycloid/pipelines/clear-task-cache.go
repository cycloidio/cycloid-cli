package pipelines

import (
	root "github.com/cycloidio/youdeploy-cli/cmd/cycloid"
	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/common"
	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/middleware"
	"github.com/spf13/cobra"
)

func NewClearTaskCacheCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "clear-task-cache",
		Short: "...",
		Long:  `........ . . .... .. .. ....`,
		RunE:  cleartaskCache,
	}

	common.RequiredPersistentFlag(common.WithFlagProject, cmd)
	common.RequiredPersistentFlag(common.WithFlagEnv, cmd)
	common.RequiredFlag(WithFlagJob, cmd)
	common.RequiredFlag(WithFlagTask, cmd)

	return cmd
}

func cleartaskCache(cmd *cobra.Command, args []string) error {
	api := root.NewAPI()
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

	err = m.ClearTaskCachePipeline(org, project, env, job, task)

	return err
}

// /organizations/{organization_canonical}/projects/{project_canonical}/pipelines/{inpath_pipeline_name}/jobs/{job_name}/tasks/{step_name}/cache
// delete: clearTaskCache
// Clear task cache
