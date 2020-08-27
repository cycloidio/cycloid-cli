package creds

import (
	"fmt"
	"reflect"

	"github.com/cycloidio/youdeploy-cli/client/client/organization_credentials"
	root "github.com/cycloidio/youdeploy-cli/cmd/cycloid"
	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/common"
	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/middleware"
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
	m := middleware.NewMiddleware(api)

	org, err := cmd.Flags().GetString("org")
	if err != nil {
		return err
	}
	id, err := cmd.Flags().GetUint32("id")
	if err != nil {
		return err
	}

	params := organization_credentials.NewGetCredentialParams()
	params.SetOrganizationCanonical(org)
	params.SetCredentialID(id)

	c, err := m.GetCredential(org, id)
	if err != nil {
		return err
	}

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

	fmt.Println(c)
	fmt.Printf("%+v\n", err)
	return nil
}
