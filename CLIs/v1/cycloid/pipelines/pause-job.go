package pipelines

import (
	"fmt"

	"github.com/cycloidio/youdeploy-cli/client/client/organization_pipelines_jobs"
	root "github.com/cycloidio/youdeploy-cli/cmd/cycloid"
	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/common"
	"github.com/spf13/cobra"
)

func NewPauseJobCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "pause-job",
		Short: "...",
		Long:  `........ . . .... .. .. ....`,
		RunE:  pauseJob,
	}

	common.RequiredPersistentFlag(common.WithFlagProject, cmd)
	common.RequiredPersistentFlag(common.WithFlagEnv, cmd)
	common.RequiredFlag(WithFlagJob, cmd)

	return cmd
}

func pauseJob(cmd *cobra.Command, args []string) error {
	api := root.NewAPI()

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

	pipelineName := fmt.Sprintf("%s-%s", project, env)

	params := organization_pipelines_jobs.NewPauseJobParams()
	params.SetOrganizationCanonical(org)
	params.SetProjectCanonical(project)
	params.SetInpathPipelineName(pipelineName)
	params.SetJobName(job)

	resp, err := api.OrganizationPipelinesJobs.PauseJob(params, root.ClientCredentials())
	if err != nil {
		return err
	}

	fmt.Println(resp)
	fmt.Printf("%+v\n", err)
	return nil
}

// /organizations/{organization_canonical}/projects/{project_canonical}/pipelines/{inpath_pipeline_name}/jobs/{job_name}/pause
// put: pauseJob
// pause a job
