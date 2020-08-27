package projects

import (
	"fmt"

	root "github.com/cycloidio/youdeploy-cli/cmd/cycloid"
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
	api := root.NewAPI()
	m := middleware.NewMiddleware(api)

	org, err := cmd.Flags().GetString("org")
	if err != nil {
		return err
	}

	d, err := m.ListProjects(org)

	for _, pr := range d {
		fmt.Printf("cannonical: %s    svcat: %s    name: %s  \n", *pr.Canonical, pr.ServiceCatalogName, *pr.Name)
	}
	fmt.Println(d)
	fmt.Printf("%+v\n", err)
	return nil
}
