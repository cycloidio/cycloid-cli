package organizations

import (
	stderrors "errors"
	"fmt"
	"net/http"

	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/apiclient"
	"github.com/cycloidio/cycloid-cli/cmd/common"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/gen/models"
)

// This command have been Hidden because it is not compatible with API key login.
// Advanced user still can use it passing a user token in CY_API_KEY env var during a login.
func NewCreateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create",
		Args:  cobra.NoArgs,
		Short: "create an organization",
		Long: `Create an organization in the Cycloid console.
Created organization are at root level by default.

If you want to create a child org, you need to specify the parent organization canonical
using the --parent-canonical (-p) flag.

Check the documentation at: https://docs.cycloid.io/reference/organizations/

See examples below.`,
		Example: `# create an organization foo
cy organization create --name foo

# create a child organization bar with parent foo
cy organization create --name bar --parent-canonical foo
`,
		RunE: create,
	}

	cmd.MarkFlagRequired(cyargs.AddOrgNameFlag(cmd))
	cyargs.AddOrgChildOfFlag(cmd)
	cmd.Flags().Bool("update", false, "update the organization display name if it already exists (same canonical as derived from --name)")
	return cmd
}

func create(cmd *cobra.Command, args []string) error {
	name, err := cyargs.GetOrgName(cmd)
	if err != nil {
		return err
	}

	parentOrg, err := cyargs.GetOrgParentCanonical(cmd)
	if err != nil {
		return err
	}

	update, err := cmd.Flags().GetBool("update")
	if err != nil {
		return err
	}

	canonical := common.GenerateCanonical(name)

	api := common.NewAPI()
	m := apiclient.NewMiddleware(api)

	_, _, getErr := m.GetOrganization(canonical)
	exists := getErr == nil
	if getErr != nil {
		var apiErr *apiclient.APIResponseError
		if !stderrors.As(getErr, &apiErr) || apiErr.StatusCode != http.StatusNotFound {
			return cyout.PrintWithOptions(cmd, nil, getErr, "failed to check if organization exists", printer.Options{})
		}
	}

	if exists && !update {
		return cyout.PrintWithOptions(cmd, nil,
			fmt.Errorf("organization %q already exists; use --update to change its display name", canonical),
			"failed to create organization", printer.Options{})
	}

	var outOrg *models.Organization
	if exists {
		outOrg, _, err = m.UpdateOrganization(canonical, name)
		return cyout.PrintWithOptions(cmd, outOrg, err, "failed to update organization", printer.Options{})
	}

	if parentOrg != "" {
		outOrg, _, err = m.CreateOrganizationChild(parentOrg, canonical, &name)
	} else {
		outOrg, _, err = m.CreateOrganization(name)
	}
	return cyout.PrintWithOptions(cmd, outOrg, err, "failed to create org named '"+name+"'", printer.Options{})
}
