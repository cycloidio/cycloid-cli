package pipelines

import (
	"fmt"

	"github.com/cycloidio/youdeploy-cli/client/client"
	"github.com/cycloidio/youdeploy-cli/client/client/organization_pipelines"
	root "github.com/cycloidio/youdeploy-cli/cmd/cycloid"
	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/common"

	"github.com/spf13/cobra"
)

func NewPauseCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "pause",
		Short: "...",
		Long:  `........ . . .... .. .. ....`,
		RunE:  pause,
	}

	common.RequiredPersistentFlag(common.WithFlagProject, cmd)
	common.RequiredPersistentFlag(common.WithFlagEnv, cmd)

	return cmd
}

func pause(cmd *cobra.Command, args []string) error {
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

	err = Pause(api, org, project, env)
	if err != nil {
		return err
	}

	// fmt.Println(resp)
	fmt.Printf("%+v\n", err)
	return nil
}

func Pause(api *client.APIClient, org string, project string, env string) error {

	pipelineName := fmt.Sprintf("%s-%s", project, env)

	params := organization_pipelines.NewPausePipelineParams()
	params.SetOrganizationCanonical(org)
	params.SetProjectCanonical(project)
	params.SetInpathPipelineName(pipelineName)

	_, err := api.OrganizationPipelines.PausePipeline(params, root.ClientCredentials())

	return err
}

// /organizations/{organization_canonical}/projects/{project_canonical}/pipelines/{inpath_pipeline_name}/pause
// put: pausePipeline
// pause a pipeline
