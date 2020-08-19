package creds

import (
	"fmt"

	"github.com/cycloidio/youdeploy-cli/client/client/organization_credentials"
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

	common.RequiredFlag(common.WithFlagID, cmd)
	return cmd
}

func delete(cmd *cobra.Command, args []string) error {
	api := root.NewAPI()

	org, err := cmd.Flags().GetString("org")
	if err != nil {
		return err
	}
	id, err := cmd.Flags().GetUint32("id")
	if err != nil {
		return err
	}

	credP := organization_credentials.NewDeleteCredentialParams()
	credP.SetOrganizationCanonical(org)
	credP.SetCredentialID(id)

	resp, err := api.OrganizationCredentials.DeleteCredential(credP, root.ClientCredentials())
	if err != nil {
		return err
	}

	fmt.Println(resp)
	fmt.Printf("%+v\n", err)
	return nil
}

// /organizations/{organization_canonical}/credentials/{credential_id}
// delete: deleteCredential
// Delete the Credential.
