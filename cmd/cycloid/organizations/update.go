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

// This command have been Hidden because it is not compatible with API key login.
// Advanced user still can use it passing a user token in CY_API_KEY env var during a login.
func NewUpdateCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "update",
		Args:  cobra.NoArgs,
		Short: "update an organization",
		Example: `
	# update an organization foo
	cy organization update --org org --name foo
`,
		RunE: update,
	}

	cmd.MarkFlagRequired(cyargs.AddOrgNameFlag(cmd))

	return cmd
}

func update(cmd *cobra.Command, args []string) error {
	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	name, err := cyargs.GetOrgName(cmd)
	if err != nil {
		return err
	}

	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return errors.Wrap(err, "unable get org flag")
	}

	output, err := cyargs.GetOutput(cmd)
	if err != nil {
		return errors.Wrap(err, "unable to get output flag")
	}

	p, err := factory.GetPrinter(output)
	if err != nil {
		return errors.Wrap(err, "unable to get printer")
	}

	o, _, err := m.UpdateOrganization(org, name)
	if err != nil {
		return printer.SmartPrint(p, nil, err, "unable to update organization", printer.Options{}, cmd.OutOrStderr())
	}
	return printer.SmartPrint(p, o, nil, "", printer.Options{}, cmd.OutOrStdout())
}
