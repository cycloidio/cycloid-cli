package pipelines

import (
	"os"

	strfmt "github.com/go-openapi/strfmt"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/internal"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/printer/factory"
)

func NewUpdateCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "update",
		Short: "update a running pipeline",
		Example: `
	# update a running pipeline
	cy --org my-org pp update --project my-project --env my-env --vars /path/to/vars.yml --pipeline /path/to/pipeline.yml
`,
		RunE:    update,
		PreRunE: internal.CheckAPIAndCLIVersion,
	}
	common.RequiredPersistentFlag(common.WithFlagProject, cmd)
	common.RequiredPersistentFlag(common.WithFlagEnv, cmd)
	common.RequiredFlag(WithFlagPipeline, cmd)
	common.RequiredFlag(WithFlagVars, cmd)

	return cmd
}

func update(cmd *cobra.Command, args []string) error {
	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	var err error

	org, err := common.GetOrg(cmd)
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
	varsPath, err := cmd.Flags().GetString("vars")
	if err != nil {
		return err
	}
	pipelinePath, err := cmd.Flags().GetString("pipeline")
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
	pipeline := string(rawPipeline)

	rawVars, err := os.ReadFile(varsPath)
	if err != nil {
		return errors.Wrap(err, "unable to read variables file")
	}
	variables := string(rawVars)

	resp, err := m.UpdatePipeline(org, project, env, pipeline, variables)
	err = printer.SmartPrint(p, resp, err, "unable to update pipeline", printer.Options{}, cmd.OutOrStdout())
	if err != nil {
		return err
	}

	//
	// PUSH pipeline vars into CONFIG
	//
	crVarsPath, err := GetPipelineVarsPath(m, org, project, *resp.UseCase)
	if err != nil {
		printer.SmartPrint(p, nil, err, "unable to get pipeline variables destination path", printer.Options{}, cmd.OutOrStdout())
	}

	cfs := make(map[string]strfmt.Base64)
	cfs[crVarsPath] = rawVars
	err = m.PushConfig(org, project, env, cfs)

	return printer.SmartPrint(p, nil, err, "unable to push config", printer.Options{}, cmd.OutOrStdout())
}
