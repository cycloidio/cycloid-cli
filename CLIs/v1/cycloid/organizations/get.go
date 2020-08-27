package organizations

import (
	"fmt"

	"github.com/cycloidio/youdeploy-cli/client/client/organizations"
	"github.com/spf13/cobra"

	root "github.com/cycloidio/youdeploy-cli/cmd/cycloid"
	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/common"
	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/middleware"
)

func NewGetCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "get",
		Short: "...",
		Long:  `........ . . .... .. .. ....`,
		RunE:  get,
	}
	common.RequiredPersistentFlag(common.WithFlagOrg, cmd)

	return cmd
}

func get(cmd *cobra.Command, args []string) error {
	api := root.NewAPI()
	m := middleware.NewMiddleware(api)

	org, err := cmd.Flags().GetString("org")
	if err != nil {
		return err
	}

	params := organizations.NewGetOrgParams()
	params.SetOrganizationCanonical(org)

	d, err := m.GetOrganization(org)
	if err != nil {
		return err
	}

	fmt.Printf("canonical: %s    name: %s     Blocked: %s    team: %s  \n", *d.Canonical, *d.Name, d.Blocked, *d.CiTeamName)

	fmt.Println(d)
	fmt.Printf("%+v\n", err)
	return nil
}

// /organizations/{organization_canonical}
// get: getOrg
// Get the information of the organization.
