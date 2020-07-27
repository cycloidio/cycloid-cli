package externalBackends

import (
	"fmt"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/cycloidio/youdeploy-cli/client/client/organization_external_backends"
	root "github.com/cycloidio/youdeploy-cli/cmd/cycloid"
	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/common"

	"github.com/spf13/cobra"
)

func NewListCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "list",
		Short: "...",
		Long:  `........ . . .... .. .. ....`,
		// Run: func(cmd *cobra.Command, args []string) {
		RunE: list,
	}

	common.RequiredPersistentFlag(common.WithFlagOrg, cmd)

	return cmd
}

func list(cmd *cobra.Command, args []string) error {
	api := root.NewAPI()

	// org := viper.GetString("org")
	org, err := cmd.Flags().GetString("org")
	if err != nil {
		return err
	}

	ebP := organization_external_backends.NewGetExternalBackendsParams()
	ebP.SetOrganizationCanonical(org)

	// ebP.SetEnvironment(&env)
	// ebP.SetProject(&project)
	resp, err := api.OrganizationExternalBackends.GetExternalBackends(ebP, root.ClientCredentials())
	if err != nil {
		return err
	}

	// api.OrganizationExternalBackends.GetExternalBackends(params *GetExternalBackendsParams, authInfo runtime.ClientAuthInfoWriter)
	fmt.Println("...")
	p := resp.GetPayload()
	err = p.Validate(strfmt.Default)
	if err != nil {
		return err
	}

	for _, eb := range p.Data {
		fmt.Printf("id: %d    project: %s - env: %s\nPurpose: %s      Config: %s\n\n", eb.ID, eb.ProjectCanonical, eb.EnvironmentCanonical, *eb.Purpose, eb.Configuration().Engine())

	}
	fmt.Println(resp)
	fmt.Printf("%+v\n", err)
	return nil
}

// /organizations/{organization_canonical}/external_backends
// get: getExternalBackends
// Get the list of organization external backends
