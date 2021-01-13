package pipelines

import (
	"io/ioutil"
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
	cy --org my-org pp update --project my-project --env my-env --vars /path/to/vars.yml --pipeline /path/to/pipeline.yml --config /path/config.tf
`,
		RunE:    update,
		PreRunE: internal.CheckAPIAndCLIVersion,
	}
	common.RequiredPersistentFlag(common.WithFlagProject, cmd)
	common.RequiredPersistentFlag(common.WithFlagEnv, cmd)
	common.RequiredFlag(WithFlagPipeline, cmd)
	common.RequiredFlag(WithFlagVars, cmd)
	WithFlagConfig(cmd)

	return cmd
}

func update(cmd *cobra.Command, args []string) error {
	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	var err error

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
	varsPath, err := cmd.Flags().GetString("vars")
	if err != nil {
		return err
	}
	pipelinePath, err := cmd.Flags().GetString("pipeline")
	if err != nil {
		return err
	}
	configs, err := cmd.Flags().GetStringToString("config")
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

	rawPipeline, err := ioutil.ReadFile(pipelinePath)
	if err != nil {
		return errors.Wrap(err, "unable to read pipeline file")
	}
	pipeline := string(rawPipeline)

	rawVars, err := ioutil.ReadFile(varsPath)
	if err != nil {
		return errors.Wrap(err, "unable to read variables file")
	}
	variables := string(rawVars)

	resp, err := m.UpdatePipeline(org, project, env, pipeline, variables)
	if err != nil {
		// print the result on the standard output
		if err := p.Print(err, printer.Options{}, os.Stdout); err != nil {
			return errors.Wrap(err, "unable to print result")
		}
		return errors.Wrap(err, "unable to update pipeline")
	}

	// print the result on the standard output
	if err := p.Print(resp, printer.Options{}, os.Stdout); err != nil {
		return errors.Wrap(err, "unable to print result")
	}

	//
	// PUSH CONFIG If pipeline update succeeded
	//

	if len(configs) > 0 {

		cfs := make(map[string]strfmt.Base64)

		for fp, dest := range configs {
			var c strfmt.Base64
			c, err = ioutil.ReadFile(fp)
			if err != nil {
				return errors.Wrap(err, "unable to read config file")
			}
			cfs[dest] = c
		}

		if err := m.PushConfig(org, project, env, cfs); err != nil {
			// print the result on the standard output
			if err := p.Print(err, printer.Options{}, os.Stdout); err != nil {
				return errors.Wrap(err, "unable to print result")
			}
			return errors.Wrap(err, "unable to push config")
		}
	}

	return nil
}
