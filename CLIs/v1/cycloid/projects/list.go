package projects

import (
	"fmt"

	"github.com/cycloidio/youdeploy-cli/client/client/organization_projects"
	root "github.com/cycloidio/youdeploy-cli/cmd/cycloid"
	"github.com/spf13/cobra"
)

func NewListCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "list",
		Short: "...",
		Long:  `........ . . .... .. .. ....`,
		RunE:  list,
	}
	return cmd

}

// /organizations/{organization_canonical}/projects
// get: getProjects
// Get list of projects of the organization.

func list(cmd *cobra.Command, args []string) error {
	api := root.NewAPI()

	org, err := cmd.Flags().GetString("org")
	if err != nil {
		return err
	}

	params := organization_projects.NewGetProjectsParams()
	params.SetOrganizationCanonical(org)

	resp, err := api.OrganizationProjects.GetProjects(params, root.ClientCredentials())
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

	for _, pr := range p.Data {
		fmt.Printf("cannonical: %s    svcat: %s    name: %s  \n", *pr.Canonical, pr.ServiceCatalogName, *pr.Name)
	}
	fmt.Println(resp)
	fmt.Printf("%+v\n", err)
	return nil
}
