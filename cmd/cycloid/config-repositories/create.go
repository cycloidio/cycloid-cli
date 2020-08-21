package configRepositories

import (
	"fmt"
	"time"

	"github.com/cycloidio/youdeploy-cli/client/client/organization_config_repositories"
	"github.com/cycloidio/youdeploy-cli/client/models"
	root "github.com/cycloidio/youdeploy-cli/cmd/cycloid"
	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/common"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/spf13/cobra"
)

func NewCreateCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "create",
		Short: "...",
		Long:  `........ . . .... .. .. ....`,
		RunE:  createConfigRepository,
	}

	common.RequiredFlag(common.WithFlagCred, cmd)
	common.RequiredFlag(WithFlagName, cmd)
	common.RequiredFlag(WithFlagBranch, cmd)
	common.RequiredFlag(WithFlagURL, cmd)
	WithFlagDefault(cmd)

	return cmd
}

// /organizations/{organization_canonical}/config_repositories
// post: createConfigRepository
// Creates a config repository

func createConfigRepository(cmd *cobra.Command, args []string) error {
	api := root.NewAPI()

	org, err := cmd.Flags().GetString("org")
	if err != nil {
		return err
	}

	name, err := cmd.Flags().GetString("name")
	if err != nil {
		return err
	}

	url, err := cmd.Flags().GetString("url")
	if err != nil {
		return err
	}

	branch, err := cmd.Flags().GetString("branch")
	if err != nil {
		return err
	}

	setDefault, err := cmd.Flags().GetBool("default")
	if err != nil {
		return err
	}

	cred, err := cmd.Flags().GetUint32("cred")
	if err != nil {
		return err
	}

	params := organization_config_repositories.NewCreateConfigRepositoryParams()
	params.SetOrganizationCanonical(org)

	body := &models.CreateConfigRepository{
		Branch:       &branch,
		CredentialID: &cred,
		Default:      &setDefault,
		Name:         &name,
		URL:          &url,
	}

	params.SetBody(body)
	err = body.Validate(strfmt.Default)
	if err != nil {
		return err
	}

	resp, err := api.OrganizationConfigRepositories.CreateConfigRepository(params, root.ClientCredentials())
	if err != nil {
		return err
	}

	cr := resp.GetPayload()

	fmt.Printf("id: %d    name: %s    url: %s    branch: %s    default: %t    credential_id: %d\n", *cr.Data.ID, *cr.Data.Name, *cr.Data.URL, cr.Data.Branch, *cr.Data.Default, cr.Data.CredentialID)
	fmt.Printf("created_at: %v    updated_at: %v\n", time.Unix(*cr.Data.CreatedAt, 0), time.Unix(*cr.Data.UpdatedAt, 0))

	fmt.Println(resp)
	fmt.Printf("%+v\n", err)

	return nil
}
