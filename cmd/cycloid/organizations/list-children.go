package organizations

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/printer/factory"
)

func NewListChildrensCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:  "list-children",
		Args: cobra.NoArgs,
		Aliases: []string{
			"list-childrens",
		},
		Short: "list the organization children",
		RunE:  listChildrens,
	}

	return cmd
}

func listChildrens(cmd *cobra.Command, args []string) error {
	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	org, err := cyargs.GetOrg(cmd)
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

	oc, err := m.ListOrganizationChildrens(org)
	return printer.SmartPrint(p, oc, err, "unable to list organization childrens", printer.Options{}, cmd.OutOrStdout())
}
