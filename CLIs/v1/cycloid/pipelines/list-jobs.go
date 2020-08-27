package pipelines

import (
	"fmt"

	root "github.com/cycloidio/youdeploy-cli/cmd/cycloid"
	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/common"
	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/middleware"
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

	js, err := m.ListPipelineJobs(org, project, env)
	if err != nil {
		return err
	}

	for _, d := range js {
		fmt.Printf("Name: %s    Paused: %s  \n", *d.Name, d.Paused)
		fmt.Printf("    FinishedBuild: %s\n", d.FinishedBuild)
	}

	fmt.Printf("%+v\n", err)
	fmt.Printf("%+v\n", js)
	return nil
}

// /organizations/{organization_canonical}/projects/{project_canonical}/pipelines/{inpath_pipeline_name}/jobs
// get: getJobs
// Get the jobs of the pipeline that the authenticated user has access to.
