package pipelines

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/apiclient"
	"github.com/cycloidio/cycloid-cli/cmd/common"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
	"github.com/cycloidio/cycloid-cli/gen/models"
)

func NewBuildGetCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "Get a build information.",
		RunE:  getBuild,
		Args:  cobra.NoArgs,
	}

	cyargs.AddCyContext(cmd)
	cyargs.AddPipeline(cmd)
	cyargs.AddPipelineJob(cmd)
	cyargs.AddPipelineBuildID(cmd)
	cyout.RegisterModel(cmd, models.Build{})
	return cmd
}

func getBuild(cmd *cobra.Command, args []string) error {
	org, project, env, component, err := cyargs.GetCyContext(cmd)
	if err != nil {
		return err
	}

	pipeline, err := cyargs.GetPipeline(cmd)
	if err != nil {
		return err
	}

	job, err := cyargs.GetPipelineJob(cmd)
	if err != nil {
		return err
	}

	buildID, err := cyargs.GetPipelineBuildID(cmd)
	if err != nil {
		return err
	}

	api := common.NewAPI()
	m := apiclient.NewMiddleware(api)

	build, _, err := m.GetBuild(org, project, env, component, pipeline, job, buildID)
	errMsg := fmt.Sprintf("failed to fetch build with ID '%s', in pipeline '%s' of project '%s', env '%s', component '%s'",
		buildID, pipeline, project, env, component)
	return cyout.PrintWithOptions(cmd, build, err, errMsg, buildTableOptions)
}
