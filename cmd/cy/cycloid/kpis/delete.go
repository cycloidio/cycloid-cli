package kpis

import (
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/internal"
)

func NewDeleteCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "delete",
		Args:  cobra.NoArgs,
		Short: "delete a kpi",
		Example: `
	# delete a kpi
	cy --org my-org kpi delete \
		--canonical kpi
`,
		// RunE:    delete,
		RunE:    func(cmd *cobra.Command, args []string) error { panic("Not implemented") }, //create,
		PreRunE: internal.CheckAPIAndCLIVersion,
	}

	common.RequiredFlag(common.WithFlagCan, cmd)

	return cmd
}

// func delete(cmd *cobra.Command, args []string) error {
// 	api := common.NewAPI()
// 	m := middleware.NewMiddleware(api)
//
// 	org, err := cy_args.GetOrg(cmd)
// 	if err != nil {
// 		return err
// 	}
// 	can, err := cmd.Flags().GetString("canonical")
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
// 	err = m.DeleteKpi(org, can)
// 	return printer.SmartPrint(p, nil, err, "unable to delete kpi", printer.Options{}, cmd.OutOrStdout())
// }
