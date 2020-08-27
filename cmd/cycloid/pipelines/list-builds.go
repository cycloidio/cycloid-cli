package pipelines

import (
	"fmt"

	root "github.com/cycloidio/youdeploy-cli/cmd/cycloid"
	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/common"
	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/middleware"
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

	bs, err := m.ListPipelineJobsBuilds(org, project, env, job)
	if err != nil {
		return err
	}

	for _, d := range bs {
		fmt.Printf("Name: %s    Status: %s  \n", *d.Name, *d.Status)
		fmt.Printf("    StartTime: %s    EndTime: %s  \n", d.StartTime, d.EndTime)
	}

	fmt.Printf("%+v\n", err)
	fmt.Printf("%+v\n", bs)
	return nil
}

// /organizations/{organization_canonical}/projects/{project_canonical}/pipelines/{inpath_pipeline_name}/jobs/{job_name}/builds
// get: getBuilds
// Get the pipeline job's builds that the authenticated user has access to.
