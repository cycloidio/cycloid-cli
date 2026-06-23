package organizations

import (
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/apiclient"
	"github.com/cycloidio/cycloid-cli/cmd/common"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
	"github.com/cycloidio/cycloid-cli/printer"
)

func NewListChildrensCommand() *cobra.Command {
	cmd := &cobra.Command{
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
	m := apiclient.NewMiddleware(api)

	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}

	oc, _, err := m.ListOrganizationChildrens(org)
	return cyout.PrintWithOptions(cmd, oc, err, "unable to list organization childrens", printer.Options{})
}
