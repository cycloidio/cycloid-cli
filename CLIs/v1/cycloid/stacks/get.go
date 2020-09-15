package stacks

import (
	"fmt"

	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/common"
	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/middleware"

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
	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	org, err := cmd.Flags().GetString("org")
	if err != nil {
		return err
	}

	ref, err := cmd.Flags().GetString("ref")
	if err != nil {
		return err
	}

	d, err := m.GetStack(org, ref)

	fmt.Printf("ref: %s    name: %s    status: %s  \n", *d.Ref, *d.Name, d.Status)
	fmt.Printf("  author: %s    describ: %s  \n", *d.Author, *d.Description)

	fmt.Printf("%+v\n", err)
	return nil
}

// /organizations/{organization_canonical}/service_catalogs/{service_catalog_ref}
// get: getServiceCatalog
// Get the information of the service catalog
