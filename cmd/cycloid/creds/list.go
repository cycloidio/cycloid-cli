package creds

import (
	"fmt"

	"github.com/cycloidio/youdeploy-cli/client/client/organization_credentials"
	root "github.com/cycloidio/youdeploy-cli/cmd/cycloid"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/spf13/cobra"
)

func NewListCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "list",
		Short: "...",
		Long:  `........ . . .... .. .. ....`,
		RunE:  list,
	}

	WithFlagType(cmd)
	return cmd
}

func list(cmd *cobra.Command, args []string) error {
	api := root.NewAPI()

	org, err := cmd.Flags().GetString("org")
	if err != nil {
		return err
	}
	credT, err := cmd.Flags().GetString("type")
	if err != nil {
		return err
	}

	credP := organization_credentials.NewGetCredentialsParams()
	credP.SetOrganizationCanonical(org)

	if credT != "" {
		credP.SetCredentialType(&credT)
	}

	// ebP.SetEnvironment(&env)
	// ebP.SetProject(&project)
	resp, err := api.OrganizationCredentials.GetCredentials(credP, root.ClientCredentials())
	if err != nil {
		return err
	}

	fmt.Println("...")
	p := resp.GetPayload()
	err = p.Validate(strfmt.Default)
	if err != nil {
		return err
	}

	for _, c := range p.Data {
		fmt.Printf("id: %d    type: %s    path: %s  \n", *c.ID, *c.Type, *c.Path)

	}
	fmt.Println(resp)
	fmt.Printf("%+v\n", err)
	return nil
}
