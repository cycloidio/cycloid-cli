package organizations

import (
	"fmt"

	"github.com/cycloidio/youdeploy-cli/client/client/organizations"
	"github.com/spf13/cobra"

	root "github.com/cycloidio/youdeploy-cli/cmd/cycloid"
	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/common"
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

	org, err := cmd.Flags().GetString("org")
	if err != nil {
		return err
	}

	params := organizations.NewGetOrgParams()
	params.SetOrganizationCanonical(org)

	resp, err := api.Organizations.GetOrg(params, root.ClientCredentials())
	if err != nil {
		return err
	}

	fmt.Println("...")
	p := resp.GetPayload()

	// TODO this validate have been removed https://github.com/cycloidio/youdeploy-http-api/issues/2262
	// err = p.Validate(strfmt.Default)
	// if err != nil {
	// 	return err
	// }

	d := p.Data
	fmt.Printf("canonical: %s    name: %s     Blocked: %s    team: %s  \n", *d.Canonical, *d.Name, d.Blocked, *d.CiTeamName)

	fmt.Println(resp)
	fmt.Printf("%+v\n", err)
	return nil
}

// /organizations/{organization_canonical}
// get: getOrg
// Get the information of the organization.
