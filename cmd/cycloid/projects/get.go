package projects

import (
	"fmt"

	"github.com/cycloidio/youdeploy-cli/client/client/organization_projects"
	root "github.com/cycloidio/youdeploy-cli/cmd/cycloid"
	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/common"

	"github.com/spf13/cobra"
)

func NewGetCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "get",
		Short: "...",
		Long:  `........ . . .... .. .. ....`,
		RunE:  get,
	}

	common.RequiredPersistentFlag(common.WithFlagProject, cmd)

	return cmd
}

// /organizations/{organization_canonical}/projects
// get: getProjects
// Get list of projects of the organization.

func get(cmd *cobra.Command, args []string) error {
	api := root.NewAPI()

	org, err := cmd.Flags().GetString("org")
	if err != nil {
		return err
	}
	project, err := cmd.Flags().GetString("project")
	if err != nil {
		return err
	}

	params := organization_projects.NewGetProjectParams()
	params.SetOrganizationCanonical(org)
	params.SetProjectCanonical(project)

	resp, err := api.OrganizationProjects.GetProject(params, root.ClientCredentials())
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

	pr := p.Data
	fmt.Printf("cannonical: %s    svcat: %s    name: %s  \n", *pr.Canonical, *pr.ServiceCatalogRef, *pr.Name)

	fmt.Println(resp)
	fmt.Printf("%+v\n", err)
	return nil
}

// /organizations/{organization_canonical}/projects/{project_canonical}
// get: getProject
// Get a project of the organization.
