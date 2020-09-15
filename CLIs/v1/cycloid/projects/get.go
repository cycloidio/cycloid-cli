package projects

import (
	"fmt"

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

	common.RequiredPersistentFlag(common.WithFlagProject, cmd)

	return cmd
}

func get(cmd *cobra.Command, args []string) error {
	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	org, err := cmd.Flags().GetString("org")
	if err != nil {
		return err
	}
	project, err := cmd.Flags().GetString("project")
	if err != nil {
		return err
	}

	d, err := m.GetProject(org, project)
	if err != nil {
		return err
	}

	fmt.Printf("cannonical: %s    svcat: %s    name: %s  \n", *d.Canonical, *d.ServiceCatalogRef, *d.Name)
	fmt.Printf("    envs: %s\n", d.Environments)

	fmt.Printf("%+v\n", err)
	return nil
}
