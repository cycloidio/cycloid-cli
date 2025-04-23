package pipelines

import (
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/internal"
)

func NewListCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:     "list",
		Short:   "list all pipelines of the organization",
		PreRunE: internal.CheckAPIAndCLIVersion,
		// RunE:    list,
		RunE: func(cmd *cobra.Command, args []string) error { panic("TODO: not implemented") },
	}

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
// 	pps, err := m.ListPipelines(org)
// 	return printer.SmartPrint(p, pps, err, "unable to list pipelines", printer.Options{}, cmd.OutOrStdout())
// }
