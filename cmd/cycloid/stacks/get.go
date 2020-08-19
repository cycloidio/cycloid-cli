package stacks

import (
	"fmt"

	"github.com/cycloidio/youdeploy-cli/client/client/service_catalogs"

	root "github.com/cycloidio/youdeploy-cli/cmd/cycloid"

	"github.com/spf13/cobra"
)

var refFlag string

func NewGetCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "get",
		Short: "...",
		Long:  `........ . . .... .. .. ....`,
		RunE:  get,
	}

	cmd.Flags().StringVar(&refFlag, "ref", "", "...")
	cmd.MarkFlagRequired("ref")

	return cmd
}

func get(cmd *cobra.Command, args []string) error {
	api := root.NewAPI()

	org, err := cmd.Flags().GetString("org")
	if err != nil {
		return err
	}

	ref, err := cmd.Flags().GetString("ref")
	if err != nil {
		return err
	}

	params := service_catalogs.NewGetServiceCatalogParams()
	params.SetOrganizationCanonical(org)
	params.SetServiceCatalogRef(ref)

	resp, err := api.ServiceCatalogs.GetServiceCatalog(params, root.ClientCredentials())
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

	d := p.Data
	fmt.Printf("ref: %s    name: %s    status: %s  \n", *d.Ref, *d.Name, d.Status)
	fmt.Printf("  author: %s    describ: %s  \n", *d.Author, *d.Description)

	fmt.Println(resp)
	fmt.Printf("%+v\n", err)
	return nil
}

// /organizations/{organization_canonical}/service_catalogs/{service_catalog_ref}
// get: getServiceCatalog
// Get the information of the service catalog
