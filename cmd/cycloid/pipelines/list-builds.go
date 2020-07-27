package pipelines

import (
	"fmt"

	"github.com/cycloidio/youdeploy-cli/client/client/organization_pipelines_jobs_build"
	root "github.com/cycloidio/youdeploy-cli/cmd/cycloid"
	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/common"
	"github.com/spf13/cobra"
)

func NewGetListBuildsCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "list-builds",
		Short: "...",
		Long:  `........ . . .... .. .. ....`,
		RunE:  listBuilds,
	}

	common.RequiredPersistentFlag(common.WithFlagProject, cmd)
	common.RequiredPersistentFlag(common.WithFlagEnv, cmd)
	common.RequiredFlag(WithFlagJob, cmd)

	return cmd
}

func listBuilds(cmd *cobra.Command, args []string) error {
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

	params := organization_pipelines_jobs_build.NewGetBuildsParams()
	params.SetOrganizationCanonical(org)
	params.SetProjectCanonical(project)
	params.SetInpathPipelineName(pipelineName)
	params.SetJobName(job)

	resp, err := api.OrganizationPipelinesJobsBuild.GetBuilds(params, root.ClientCredentials())
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
		fmt.Printf("Name: %s    Status: %s  \n", *d.Name, *d.Status)
		fmt.Printf("    StartTime: %s    EndTime: %s  \n", d.StartTime, d.EndTime)
	}

	fmt.Printf("%+v\n", err)
	fmt.Printf("%+v\n", resp)
	return nil
}

// /organizations/{organization_canonical}/projects/{project_canonical}/pipelines/{inpath_pipeline_name}/jobs/{job_name}/builds
// get: getBuilds
// Get the pipeline job's builds that the authenticated user has access to.
