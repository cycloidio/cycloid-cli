package pipelines

import (
	"fmt"

	"github.com/cycloidio/youdeploy-cli/client/client/organization_pipelines_jobs"
	root "github.com/cycloidio/youdeploy-cli/cmd/cycloid"
	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/common"
	"github.com/spf13/cobra"
)

func NewListJobsCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "list-jobs",
		Short: "...",
		Long:  `........ . . .... .. .. ....`,
		RunE:  listJobs,
	}

	common.RequiredPersistentFlag(common.WithFlagProject, cmd)
	common.RequiredPersistentFlag(common.WithFlagEnv, cmd)

	return cmd
}

func listJobs(cmd *cobra.Command, args []string) error {
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

	pipelineName := fmt.Sprintf("%s-%s", project, env)

	params := organization_pipelines_jobs.NewGetJobsParams()
	params.SetOrganizationCanonical(org)
	params.SetProjectCanonical(project)
	params.SetInpathPipelineName(pipelineName)

	resp, err := api.OrganizationPipelinesJobs.GetJobs(params, root.ClientCredentials())
	if err != nil {
		return err
	}

	p := resp.GetPayload()
	// TODO this validate have been removed https://github.com/cycloidio/youdeploy-http-api/issues/2262
	// err = p.Validate(strfmt.Default)
	// if err != nil {
	// 	return err
	// }

	for _, d := range p.Data {
		fmt.Printf("Name: %s    Paused: %s  \n", *d.Name, d.Paused)
		fmt.Printf("    FinishedBuild: %s\n", d.FinishedBuild)
	}

	fmt.Printf("%+v\n", err)
	fmt.Printf("%+v\n", resp)
	return nil
}

// /organizations/{organization_canonical}/projects/{project_canonical}/pipelines/{inpath_pipeline_name}/jobs
// get: getJobs
// Get the jobs of the pipeline that the authenticated user has access to.
