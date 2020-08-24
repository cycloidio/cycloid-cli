package catalogRepositories

import (
	"fmt"
	"time"

	"github.com/cycloidio/youdeploy-cli/client/client/organization_service_catalog_sources"
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
		RunE:  createCatalogRepository,
	}

	// create --branch test --cred 105 --url "git@github.com:foo/bla.git"  --name catalogname

	common.RequiredFlag(common.WithFlagCred, cmd)
	common.RequiredFlag(WithFlagName, cmd)
	common.RequiredFlag(WithFlagBranch, cmd)
	common.RequiredFlag(WithFlagURL, cmd)

	return cmd
}

// /organizations/{organization_canonical}/service_catalog_sources
// post: createServiceCatalogSource
// Creates a Service catalog source

func createCatalogRepository(cmd *cobra.Command, args []string) error {
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

	cred, err := cmd.Flags().GetUint32("cred")
	if err != nil {
		return err
	}

	params := organization_service_catalog_sources.NewCreateServiceCatalogSourceParams()
	params.SetOrganizationCanonical(org)

	body := &models.CreateServiceCatalogSource{
		Branch:       &branch,
		CredentialID: cred,
		Name:         &name,
		URL:          &url,
	}

	params.SetBody(body)
	err = body.Validate(strfmt.Default)
	if err != nil {
		return err
	}

	resp, err := api.OrganizationServiceCatalogSources.CreateServiceCatalogSource(params, root.ClientCredentials())
	if err != nil {
		return err
	}

	cr := resp.GetPayload()

	fmt.Printf("id: %d    name: %s    url: %s    branch: %s    credential_id: %d\n", *cr.Data.ID, *cr.Data.Name, *cr.Data.URL, cr.Data.Branch, cr.Data.CredentialID)
	fmt.Printf("created_at: %v    updated_at: %v\n", time.Unix(*cr.Data.CreatedAt, 0), time.Unix(*cr.Data.UpdatedAt, 0))

	fmt.Println(resp)
	fmt.Printf("%+v\n", err)

	return nil
}
