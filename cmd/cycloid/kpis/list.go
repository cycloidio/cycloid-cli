package kpis

import (
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/internal"
)

func NewListCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "list",
		Short: "list kpis",
		Example: `
	# list kpis
	cy --org my-org kpi list
`,
		// RunE:    list,
		RunE:    func(cmd *cobra.Command, args []string) error { panic("Not implemented") }, //create,
		PreRunE: internal.CheckAPIAndCLIVersion,
	}

	common.WithFlagProject(cmd)
	common.WithFlagEnv(cmd)

	return cmd
}

// func list(cmd *cobra.Command, args []string) error {
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
// 	kpis, err := m.ListKpi(org, project, env)
// 	return printer.SmartPrint(p, kpis, err, "unable to list kpis", printer.Options{}, cmd.OutOrStdout())
// }
