package pipelines

import (
	"os"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/internal"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/printer/factory"
)

func NewUnpauseJobCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:     "unpause-job",
		Short:   "unpause a pipeline job",
		RunE:    unpauseJob,
		PreRunE: internal.CheckAPIAndCLIVersion,
		Example: `
	# unpause job 'my-job'
	cy --org my-org pp unpause-job --project my-project --env env --job my-job
`,
	}

	common.RequiredPersistentFlag(common.WithFlagProject, cmd)
	common.RequiredPersistentFlag(common.WithFlagEnv, cmd)
	common.RequiredFlag(WithFlagJob, cmd)

	return cmd
}

func unpauseJob(cmd *cobra.Command, args []string) error {
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
	job, err := cmd.Flags().GetString("job")
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

	if err := m.UnpausePipelineJob(org, project, env, job); err != nil {
		// print the result on the standard output
		if err := p.Print(err, printer.Options{}, os.Stdout); err != nil {
			return errors.Wrap(err, "unable to print result")
		}
		return errors.Wrap(err, "unable to unpause pipeline's job")
	}

	return nil
}
