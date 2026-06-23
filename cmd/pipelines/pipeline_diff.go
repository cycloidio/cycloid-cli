package pipelines

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/apiclient"
	"github.com/cycloidio/cycloid-cli/cmd/common"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
	"github.com/cycloidio/cycloid-cli/printer"
)

func NewPipelineDiffCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "diff",
		Short: "diff the remote pipeline with a local config",
		RunE:  diff,
		Args:  cobra.NoArgs,
	}

	cyargs.AddCyContext(cmd)
	cyargs.AddPipeline(cmd)
	cyargs.AddPipelineConfig(cmd)
	cyargs.AddPipelineVars(cmd)
	return cmd
}

func diff(cmd *cobra.Command, args []string) error {
	org, project, env, component, err := cyargs.GetCyContext(cmd)
	if err != nil {
		return err
	}

	pipeline, err := cyargs.GetPipeline(cmd)
	if err != nil {
		return err
	}

	varsPath, err := cyargs.GetPipelineVars(cmd)
	if err != nil {
		return err
	}

	pipelinePath, err := cyargs.GetPipelineConfig(cmd)
	if err != nil {
		return err
	}

	rawPipeline, err := os.ReadFile(pipelinePath)
	if err != nil {
		return fmt.Errorf("failed to open pipeline config at path %q: %w", pipelinePath, err)
	}

	rawVars, err := os.ReadFile(varsPath)
	if err != nil {
		return fmt.Errorf("failed to read pipeline variables files at path %q: %w", varsPath, err)
	}

	api := common.NewAPI()
	m := apiclient.NewMiddleware(api)

	pipelineDiff, _, err := m.DiffPipeline(org, project, env, component, pipeline, string(rawPipeline), string(rawVars), false)
	return cyout.PrintWithOptions(cmd, pipelineDiff, err, "failed to get pipeline diff", printer.Options{})
}
