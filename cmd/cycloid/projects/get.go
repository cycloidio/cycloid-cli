package projects

import (
	"fmt"

	"github.com/cycloidio/youdeploy-cli/client/client"
	"github.com/cycloidio/youdeploy-cli/client/client/organization_projects"
	"github.com/cycloidio/youdeploy-cli/client/models"
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

	d, err := Get(api, org, project)

	fmt.Printf("cannonical: %s    svcat: %s    name: %s  \n", *d.Canonical, *d.ServiceCatalogRef, *d.Name)
	fmt.Printf("    envs: %s\n", d.Environments)

	fmt.Printf("%+v\n", err)
	return nil
}

func Get(api *client.APIClient, org string, project string) (*models.Project, error) {

	params := organization_projects.NewGetProjectParams()
	params.SetOrganizationCanonical(org)
	params.SetProjectCanonical(project)

	resp, err := api.OrganizationProjects.GetProject(params, root.ClientCredentials())
	if err != nil {
		return nil, err
	}

	p := resp.GetPayload()
	// TODO this validate have been removed https://github.com/cycloidio/youdeploy-http-api/issues/2262
	// err = p.Validate(strfmt.Default)
	// if err != nil {
	// 	return err
	// }

	d := p.Data

	return d, err
}

// /organizations/{organization_canonical}/projects/{project_canonical}
// get: getProject
// Get a project of the organization.
