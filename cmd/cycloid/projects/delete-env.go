package projects

import (
	"fmt"

	"github.com/cycloidio/youdeploy-cli/client/client/organization_projects"
	root "github.com/cycloidio/youdeploy-cli/cmd/cycloid"
	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/common"

	"github.com/spf13/cobra"
)

func NewDeleteEnvCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "delete-env",
		Short: "...",
		Long:  `........ . . .... .. .. ....`,
		RunE:  deleteEnv,
	}
	common.RequiredPersistentFlag(common.WithFlagEnv, cmd)
	common.RequiredPersistentFlag(common.WithFlagProject, cmd)

	return cmd
}

func deleteEnv(cmd *cobra.Command, args []string) error {
	api := root.NewAPI()

	org, err := cmd.Flags().GetString("org")
	if err != nil {
		return err
	}
	project, err := cmd.Flags().GetString("project")
	if err != nil {
		return err
	}
	env, err := cmd.Flags().GetString("env")
	if err != nil {
		return err
	}

	params := organization_projects.NewDeleteProjectEnvironmentParams()
	params.SetOrganizationCanonical(org)
	params.SetProjectCanonical(project)
	params.SetEnvironmentCanonical(env)

	resp, err := api.OrganizationProjects.DeleteProjectEnvironment(params, root.ClientCredentials())
	if err != nil {
		return err
	}

	fmt.Println(resp)
	fmt.Printf("%+v\n", err)
	return nil
}

// /organizations/{organization_canonical}/projects/{project_canonical}/environments/{environment_canonical}
// delete: deleteProjectEnvironment
// Delete a project environment of the organization, and the project itself if it's the last environment.
