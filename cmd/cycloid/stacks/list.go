package stacks

import (
	"fmt"

	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/common"
	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/middleware"

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
	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	org, err := cmd.Flags().GetString("org")
	if err != nil {
		return err
	}

	d, err := m.ListStacks(org)

	for _, s := range d {
		fmt.Printf("ref: %s    name: %s    status: %s  \n", *s.Ref, *s.Name, s.Status)
		fmt.Printf("  author: %s    describ: %s  \n", *s.Author, *s.Description)
	}
	fmt.Printf("%+v\n", err)
	return nil
}

// /organizations/{organization_canonical}/service_catalogs
// get: getServiceCatalogs
// Return all the service catalogs
