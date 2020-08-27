package pipelines

import (
	"fmt"

	root "github.com/cycloidio/youdeploy-cli/cmd/cycloid"
	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/common"
	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/middleware"
	"github.com/spf13/cobra"
)

func NewGetJobCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "get-job",
		Short: "...",
		Long:  `........ . . .... .. .. ....`,
		RunE:  getJob,
	}

	common.RequiredPersistentFlag(common.WithFlagProject, cmd)
	common.RequiredPersistentFlag(common.WithFlagEnv, cmd)
	common.RequiredFlag(WithFlagJob, cmd)

	return cmd
}

func getJob(cmd *cobra.Command, args []string) error {
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

	d, err := m.GetPipelineJob(org, project, env, job)
	if err != nil {
		return err
	}

	fmt.Printf("Name: %s    Paused: %s  \n", *d.Name, d.Paused)
	fmt.Printf("    FinishedBuild: %s\n", d.FinishedBuild)

	fmt.Printf("%+v\n", err)
	fmt.Printf("%+v\n", d)
	return nil
}

// /organizations/{organization_canonical}/projects/{project_canonical}/pipelines/{inpath_pipeline_name}/jobs/{job_name}
// get: getJob
// Get the information of the job.
