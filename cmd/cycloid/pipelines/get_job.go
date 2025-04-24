package pipelines

import (
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/internal"
)

func NewGetJobCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "get-job",
		Short: "get a pipeline's job",
		Example: `
	# get the job 'my-job' in my-project-env pipeline in JSON format
	cy --org my-org pp get-job --project my-project --env env --job my-job -o json
`,
		//RunE:    getJob,
		RunE:    func(cmd *cobra.Command, args []string) error { panic("TODO: not implemented") },
		PreRunE: internal.CheckAPIAndCLIVersion,
	}

	common.RequiredPersistentFlag(common.WithFlagProject, cmd)
	common.RequiredPersistentFlag(common.WithFlagEnv, cmd)
	common.RequiredFlag(WithFlagJob, cmd)

	return cmd
}

// func getJob(cmd *cobra.Command, args []string) error {
// 	api := common.NewAPI()
// 	m := middleware.NewMiddleware(api)
//
// 	org, err := common.GetOrg(cmd)
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
//
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
// 	j, err := m.GetPipelineJob(org, project, env, job)
// 	return printer.SmartPrint(p, j, err, "unable to get job", printer.Options{}, cmd.OutOrStdout())
// }
