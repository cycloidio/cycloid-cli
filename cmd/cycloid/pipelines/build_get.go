package pipelines

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/printer/factory"
)

func NewBuildGetCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "get",
		Short: "Get a build information.",
		RunE:  getBuild,
		Args:  cobra.NoArgs,
	}

	cyargs.AddCyContext(cmd)
	cyargs.AddPipeline(cmd)
	cyargs.AddPipelineJob(cmd)
	cyargs.AddPipelineBuildID(cmd)
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

	output, err := cyargs.GetOutput(cmd)
	if err != nil {
		return err
	}

	p, err := factory.GetPrinter(output)
	if err != nil {
		return err
	}

	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	build, err := m.GetBuild(org, project, env, component, pipeline, job, buildID)
	if err != nil {
		return printer.SmartPrint(p, nil, err,
			fmt.Sprintf("failed to fetch build with ID '%s', in pipeline '%s' of project '%s', env '%s', component '%s'",
				buildID, pipeline, project, env, component),
			printer.Options{}, cmd.OutOrStderr())
	}

	return printer.SmartPrint(p, build, nil, "", printer.Options{}, cmd.OutOrStdout())
}
