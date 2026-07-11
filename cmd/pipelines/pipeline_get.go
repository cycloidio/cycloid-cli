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

func NewPipelineGetCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "get",
		Short:   "get a pipeline",
		Example: `cy pipeline get --project my-project --env env --component component --pipeline pipeline_name`,
		RunE:    getPipeline,
		Args:    cobra.NoArgs,
	}

	cyargs.AddCyContext(cmd)
	cyargs.AddPipeline(cmd)
	cyout.RegisterModel(cmd, models.Pipeline{})
	return cmd
}

func getPipeline(cmd *cobra.Command, args []string) error {
	org, project, env, component, err := cyargs.GetCyContext(cmd)
	if err != nil {
		return err
	}

	pipeline, err := cyargs.GetPipeline(cmd)
	if err != nil {
		return err
	}

	api := common.NewAPI()
	m := apiclient.NewAPIClient(api)

	outPipeline, _, err := m.GetPipeline(org, project, env, component, pipeline)
	errMsg := fmt.Sprintf("failed to get pipeline %q in context project %q, env %q, component %q", pipeline, project, env, component)
	return cyout.PrintWithOptions(cmd, outPipeline, err, errMsg, pipelineTableOptions)
}
