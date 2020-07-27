package projects

import (
	"fmt"

	"github.com/cycloidio/youdeploy-cli/client/client/organization_projects"
	root "github.com/cycloidio/youdeploy-cli/cmd/cycloid"
	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/common"
	"github.com/spf13/cobra"
)

func NewDeleteCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "delete",
		Short: "...",
		Long:  `........ . . .... .. .. ....`,
		RunE:  delete,
	}

	common.RequiredPersistentFlag(common.WithFlagProject, cmd)
	return cmd
}

func delete(cmd *cobra.Command, args []string) error {
	api := root.NewAPI()

	org, err := cmd.Flags().GetString("org")
	if err != nil {
		return err
	}
	project, err := cmd.Flags().GetString("project")
	if err != nil {
		return err
	}

	params := organization_projects.NewDeleteProjectParams()
	params.SetOrganizationCanonical(org)
	params.SetProjectCanonical(project)

	resp, err := api.OrganizationProjects.DeleteProject(params, root.ClientCredentials())
	if err != nil {
		return err
	}

	fmt.Println(resp)
	fmt.Printf("%+v\n", err)
	return nil
}

// /organizations/{organization_canonical}/projects/{project_canonical}
// delete: deleteProject
// Delete a project of the organization.
