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

func NewCreateChildCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "create-child",
		Args:  cobra.NoArgs,
		Short: "create a child organization",
		RunE:  createChild,
	}
	common.RequiredPersistentFlag(WithFlagParentOrganization, cmd)

	return cmd
}

func createChild(cmd *cobra.Command, args []string) error {
	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}

	porg, err := cmd.Flags().GetString("parent-org")
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

	oc, err := m.CreateOrganizationChild(org, porg, nil)
	return printer.SmartPrint(p, oc, err, "unable to create a child organization", printer.Options{}, cmd.OutOrStdout())
}
