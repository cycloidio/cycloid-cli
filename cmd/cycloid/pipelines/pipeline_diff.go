package pipelines

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/internal"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/printer/factory"
)

func NewPipelineDiffCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:     "diff",
		Short:   "diff the remote pipeline with a local config",
		RunE:    diff,
		PreRunE: internal.CheckAPIAndCLIVersion,
		Args:    cobra.NoArgs,
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

	output, err := cyargs.GetOutput(cmd)
	if err != nil {
		return errors.Wrap(err, "unable to get output flag")
	}

	// fetch the printer from the factory
	p, err := factory.GetPrinter(output)
	if err != nil {
		return errors.Wrap(err, "unable to get printer")
	}

	rawPipeline, err := os.ReadFile(pipelinePath)
	if err != nil {
		return fmt.Errorf("failed to open pipeline config at path '%s': %s", pipelinePath, err.Error())
	}

	rawVars, err := os.ReadFile(varsPath)
	if err != nil {
		return fmt.Errorf("failed to read pipeline variables files at path '%s': %s", varsPath, err.Error())
	}

	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	pipelineDiff, err := m.DiffPipeline(org, project, env, component, pipeline, string(rawPipeline), string(rawVars), false)
	if err != nil {
		return printer.SmartPrint(p, nil, err, "failed to get pipeline diff", printer.Options{}, cmd.OutOrStderr())
	}

	return printer.SmartPrint(p, pipelineDiff, nil, "", printer.Options{}, cmd.OutOrStdout())
}
