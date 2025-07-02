package pipelines

import (
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

func NewPipelineUpdateCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "update",
		Short: "update a running pipeline",
		Example: `
	# update a running pipeline
	cy --org my-org pp update --project my-project --env my-env --vars /path/to/vars.yml --pipeline /path/to/pipeline.yml
`,
		RunE:    update,
		PreRunE: internal.CheckAPIAndCLIVersion,
		Args:    cobra.NoArgs,
	}
	cyargs.AddCyContext(cmd)
	cyargs.AddPipeline(cmd)
	cyargs.AddPipelineConfig(cmd)
	cyargs.AddPipelineVars(cmd)
	return cmd
}

func update(cmd *cobra.Command, args []string) error {
	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	org, project, env, component, err := cyargs.GetCyContext(cmd)
	if err != nil {
		return err
	}

	pipeline, err := cyargs.GetPipeline(cmd)
	if err != nil {
		return err
	}

	pipelinePath, err := cyargs.GetPipelineConfig(cmd)
	if err != nil {
		return err
	}

	pipelineVarsPath, err := cyargs.GetPipelineVars(cmd)
	if err != nil {
		return err
	}

	output, err := cmd.Flags().GetString("output")
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
		return errors.Wrap(err, "unable to read pipeline file")
	}

	rawVars, err := os.ReadFile(pipelineVarsPath)
	if err != nil {
		return errors.Wrap(err, "unable to read variables file")
	}

	resp, err := m.UpdatePipeline(org, project, env, component, pipeline, string(rawPipeline), string(rawVars), false)
	if err != nil {
		return err
	}

	return printer.SmartPrint(p, resp, err, "", printer.Options{}, cmd.OutOrStdout())
}
