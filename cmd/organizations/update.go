package organizations

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/apiclient"
	"github.com/cycloidio/cycloid-cli/cmd/common"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
	"github.com/cycloidio/cycloid-cli/printer"
)

// This command have been Hidden because it is not compatible with API key login.
// Advanced user still can use it passing a user token in CY_API_KEY env var during a login.
func NewUpdateCommand() *cobra.Command {
	cmd := &cobra.Command{
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
	m := apiclient.NewMiddleware(api)

	name, err := cyargs.GetOrgName(cmd)
	if err != nil {
		return err
	}

	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return fmt.Errorf("unable get org flag: %w", err)
	}

	o, _, err := m.UpdateOrganization(org, name)
	return cyout.PrintWithOptions(cmd, o, err, "unable to update organization", printer.Options{})
}
