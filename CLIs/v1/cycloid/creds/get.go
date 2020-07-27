package creds

import (
	"fmt"
	"reflect"

	"github.com/cycloidio/youdeploy-cli/client/client/organization_credentials"
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

	common.RequiredFlag(common.WithFlagID, cmd)
	return cmd
}

func get(cmd *cobra.Command, args []string) error {
	api := root.NewAPI()

	org, err := cmd.Flags().GetString("org")
	if err != nil {
		return err
	}
	id, err := cmd.Flags().GetUint32("id")
	if err != nil {
		return err
	}

	credP := organization_credentials.NewGetCredentialParams()
	credP.SetOrganizationCanonical(org)
	credP.SetCredentialID(id)

	resp, err := api.OrganizationCredentials.GetCredential(credP, root.ClientCredentials())
	if err != nil {
		return err
	}

	p := resp.GetPayload()

	// TODO this validate have been removed https://github.com/cycloidio/youdeploy-http-api/issues/2262
	// err = p.Validate(strfmt.Default)
	// if err != nil {
	// 	return err
	// }

	c := p.Data
	fmt.Printf("id: %d  Name: %s  type: %s    path: %s  \n", *c.ID, *c.Name, *c.Type, *c.Path)

	fields := reflect.TypeOf(*c.Raw)
	values := reflect.ValueOf(*c.Raw)
	for i := 0; i < fields.NumField(); i++ {
		field := fields.Field(i)
		value := values.Field(i)
		if value.Kind() != reflect.String {
			continue
		}
		if value.String() != "" {
			fmt.Print("    ", field.Name, "=", value, "\n")
		}
	}

	fmt.Println(resp)
	fmt.Printf("%+v\n", err)
	return nil
}

// /organizations/{organization_canonical}/credentials/{credential_id}
// get: getCredential
// Get the information of the Credential.
