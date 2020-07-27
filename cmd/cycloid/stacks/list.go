package stacks

import (
	"fmt"

	"github.com/cycloidio/youdeploy-cli/client/client/service_catalogs"

	root "github.com/cycloidio/youdeploy-cli/cmd/cycloid"

	"github.com/spf13/cobra"
)

func NewListCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "list",
		Short: "...",
		Long:  `........ . . .... .. .. ....`,
		RunE:  list,
	}
	return cmd

}

func list(cmd *cobra.Command, args []string) error {
	api := root.NewAPI()

	org, err := cmd.Flags().GetString("org")
	if err != nil {
		return err
	}

	params := service_catalogs.NewGetServiceCatalogsParams()
	params.SetOrganizationCanonical(org)

	resp, err := api.ServiceCatalogs.GetServiceCatalogs(params, root.ClientCredentials())
	if err != nil {
		return err
	}

	fmt.Println("...")
	p := resp.GetPayload()

	// TODO this validate have been removed https://github.com/cycloidio/youdeploy-http-api/issues/2262
	// err = p.Validate(strfmt.Default)
	// if err != nil {
	// 	return err
	// }

	for _, d := range p.Data {
		fmt.Printf("ref: %s    name: %s    status: %s  \n", *d.Ref, *d.Name, d.Status)
		fmt.Printf("  author: %s    describ: %s  \n", *d.Author, *d.Description)
	}
	fmt.Println(resp)
	fmt.Printf("%+v\n", err)
	return nil
}

// /organizations/{organization_canonical}/service_catalogs
// get: getServiceCatalogs
// Return all the service catalogs
