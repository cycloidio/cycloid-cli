package pipelines

import (
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/internal"
)

func NewUnpauseCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "unpause",
		Short: "unpause a pipeline",
		Example: `
# unpause pipeline my-project-env
cy --org my-org pipeline unpause --project my-project --env env
`,
		// RunE:    unpause,
		RunE:    func(cmd *cobra.Command, args []string) error { panic("TODO: not implemented") },
		PreRunE: internal.CheckAPIAndCLIVersion,
	}

	common.RequiredPersistentFlag(common.WithFlagProject, cmd)
	common.RequiredPersistentFlag(common.WithFlagEnv, cmd)

	return cmd
}

// func unpause(cmd *cobra.Command, args []string) error {
// 	api := common.NewAPI()
// 	m := middleware.NewMiddleware(api)
//
// 	org, err := cy_args.GetOrg(cmd)
// 	if err != nil {
// 		return err
// 	}
//
// 	project, err := cmd.Flags().GetString("project")
// 	if err != nil {
// 		return err
// 	}
//
// 	env, err := cmd.Flags().GetString("env")
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
// 	err = m.UnpausePipeline(org, project, env)
// 	return printer.SmartPrint(p, nil, err, "unable to unpause pipeline", printer.Options{}, cmd.OutOrStdout())
// }
