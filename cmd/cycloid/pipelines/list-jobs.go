package pipelines

import (
	"os"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/common"
	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/youdeploy-cli/printer"
	"github.com/cycloidio/youdeploy-cli/printer/factory"
)

func NewListJobsCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "list-jobs",
		Short: "list pipeline jobs",
		Example: `
	# list jobs in pipeline my-project-env and display it in JSON
	cy --org my-org pp list-jobs --project my-project --env env -o json
`,
		RunE: listJobs,
	}

	common.RequiredPersistentFlag(common.WithFlagProject, cmd)
	common.RequiredPersistentFlag(common.WithFlagEnv, cmd)

	return cmd
}

func listJobs(cmd *cobra.Command, args []string) error {
	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

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
	output, err := cmd.Flags().GetString("output")
	if err != nil {
		return errors.Wrap(err, "unable to get output flag")
	}

	jobs, err := m.ListPipelineJobs(org, project, env)
	if err != nil {
		return errors.Wrap(err, "unable to list pipeline jobs")
	}

	// fetch the printer from the factory
	p, err := factory.GetPrinter(output)
	if err != nil {
		return errors.Wrap(err, "unable to get printer")
	}

	// print the result on the standard output
	if err := p.Print(jobs, printer.Options{}, os.Stdout); err != nil {
		return errors.Wrap(err, "unable to print result")
	}
	return nil
}
