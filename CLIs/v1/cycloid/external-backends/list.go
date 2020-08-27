package externalBackends

import (
	"fmt"

	root "github.com/cycloidio/youdeploy-cli/cmd/cycloid"
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

	common.RequiredPersistentFlag(common.WithFlagOrg, cmd)

	return cmd
}

func list(cmd *cobra.Command, args []string) error {
	api := root.NewAPI()
	m := middleware.NewMiddleware(api)

	// org := viper.GetString("org")
	org, err := cmd.Flags().GetString("org")
	if err != nil {
		return err
	}

	ebs, err := m.ListExternalBackends(org)
	for _, eb := range ebs {
		fmt.Printf("id: %d    project: %s - env: %s\nPurpose: %s      Config: %s\n\n", eb.ID, eb.ProjectCanonical, eb.EnvironmentCanonical, *eb.Purpose, eb.Configuration().Engine())

	}
	fmt.Println(ebs)
	fmt.Printf("%+v\n", err)
	return nil
}
