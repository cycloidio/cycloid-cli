package pipelines

import (
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/internal"
)

func NewGetCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:    "get",
		Hidden: true,
		Short:  "get a pipeline",
		Example: `
	# get a pipeline in 'my-org'
	cy --org my-org pipeline get --project my-project --env env
`, PreRunE: internal.CheckAPIAndCLIVersion,
		// RunE: get,
		RunE: func(cmd *cobra.Command, args []string) error { panic("TODO: not implemented") },
	}

	common.RequiredPersistentFlag(common.WithFlagProject, cmd)
	common.RequiredPersistentFlag(common.WithFlagEnv, cmd)

	return cmd
}

// func get(cmd *cobra.Command, args []string) error {
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
// 	pp, err := m.GetPipeline(org, project, env)
// 	return printer.SmartPrint(p, pp, err, "unable to get pipeline", printer.Options{}, cmd.OutOrStdout())
// }
