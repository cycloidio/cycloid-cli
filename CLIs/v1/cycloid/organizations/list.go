package organizations

import (
	"fmt"

	"github.com/spf13/cobra"

	root "github.com/cycloidio/youdeploy-cli/cmd/cycloid"
	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/middleware"
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

	orgs, err := m.ListOrganizations()
	if err != nil {
		return err
	}

	for _, d := range orgs {
		fmt.Printf("canonical: %s    name: %s     Blocked: %s    team: %s  \n", *d.Canonical, *d.Name, d.Blocked, *d.CiTeamName)
	}
	fmt.Println(orgs)
	fmt.Printf("%+v\n", err)
	return nil
}
