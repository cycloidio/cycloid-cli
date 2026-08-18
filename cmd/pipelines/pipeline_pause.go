package pipelines

import (
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/apiclient"
	"github.com/cycloidio/cycloid-cli/cmd/common"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
	"github.com/cycloidio/cycloid-cli/printer"
)

func NewPipelinePauseCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "pause",
		Short:   "pause a pipeline",
		Example: `cy --org my-org pipeline pause --project my-project --env env --component component --pipeline pipeline-name`,
		RunE:    pause,
		Args:    cobra.NoArgs,
	}

	cyargs.AddCyContext(cmd)
	cyargs.AddPipeline(cmd)
	return cmd
}

func pause(cmd *cobra.Command, args []string) error {
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

	_, err = m.PausePipeline(org, project, env, component, pipeline)
	return cyout.PrintWithOptions(cmd, nil, err, "failed to pause pipeline", printer.Options{})
}
