package pipelines

import (
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/internal"
)

func NewPauseJobCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "pause-job",
		Short: "pause a pipeline job",
		Example: `
	# pause job 'my-job'
	cy --org my-org pp pause-job --project my-project --env env --job my-job
`,
		// RunE:    pauseJob,
		RunE:    func(cmd *cobra.Command, args []string) error { panic("TODO: not implemented") },
		PreRunE: internal.CheckAPIAndCLIVersion,
	}

	common.RequiredPersistentFlag(common.WithFlagProject, cmd)
	common.RequiredPersistentFlag(common.WithFlagEnv, cmd)
	common.RequiredFlag(WithFlagJob, cmd)

	return cmd
}

// func pauseJob(cmd *cobra.Command, args []string) error {
// 	api := common.NewAPI()
// 	m := middleware.NewMiddleware(api)
//
// 	org, err := cy_args.GetOrg(cmd)
// 	if err != nil {
// 		return err
// 	}
// 	project, err := cmd.Flags().GetString("project")
// 	if err != nil {
// 		return err
// 	}
// 	env, err := cmd.Flags().GetString("env")
// 	if err != nil {
// 		return err
// 	}
// 	job, err := cmd.Flags().GetString("job")
// 	if err != nil {
// 		return err
// 	}
// 	output, err := cmd.Flags().GetString("output")
// 	if err != nil {
// 		return errors.Wrap(err, "unable to get output flag")
// 	}
//
// 	// fetch the printer from the factory
// 	p, err := factory.GetPrinter(output)
// 	if err != nil {
// 		return errors.Wrap(err, "unable to get printer")
// 	}
//
// 	err = m.PausePipelineJob(org, project, env, job)
// 	return printer.SmartPrint(p, nil, err, "unable to pause the job", printer.Options{}, cmd.OutOrStdout())
// }
