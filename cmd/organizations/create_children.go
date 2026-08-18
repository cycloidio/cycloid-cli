package organizations

import (
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/apiclient"
	"github.com/cycloidio/cycloid-cli/cmd/common"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
	"github.com/cycloidio/cycloid-cli/printer"
)

func NewCreateChildCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:        "create-child",
		Args:       cobra.NoArgs,
		Short:      "create a child organization",
		RunE:       createChild,
		Deprecated: "this function is replaced by `cy org create --org <org> --child-of <parent-org>`",
		Hidden:     true,
	}

	cmd.PersistentFlags().String("parent-org", "", "parent organization canonical")
	cmd.MarkPersistentFlagRequired("parent-org")

	return cmd
}

func createChild(cmd *cobra.Command, args []string) error {
	api := common.NewAPI()
	m := apiclient.NewAPIClient(api)

	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}

	porg, err := cmd.Flags().GetString("parent-org")
	if err != nil {
		return err
	}

	oc, _, err := m.CreateOrganizationChild(org, porg, nil)
	return cyout.PrintWithOptions(cmd, oc, err, "unable to create a child organization", printer.Options{})
}
