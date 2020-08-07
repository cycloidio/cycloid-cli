package externalBackends

import (
	"fmt"

	"github.com/cycloidio/youdeploy-cli/client/client/organization_external_backends"
	root "github.com/cycloidio/youdeploy-cli/cmd/cycloid"
	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/common"

	"github.com/spf13/cobra"
)

var ebProject string

func NewListCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "list",
		Short: "...",
		Long:  `........ . . .... .. .. ....`,
		// Run: func(cmd *cobra.Command, args []string) {
		Run: list,
	}

	common.WithFlagProject(cmd)
	common.WithFlagEnv(cmd)
	common.WithFlagOrg(cmd)

	// cmd.Flags().String("pproject", "pp", "Project name")
	// viper.BindPFlag("pproject", cmd.Flags().Lookup("pproject"))

	// cmd.Flags().String("project", "default-p", "Project name")

	// viper.BindPFlag("project", cmd.Flags().Lookup("project"))

	// viper.BindPFlag("pproject", cmd.Flags().Lookup("pproject"))
	// viper.RegisterAlias("pproject", "project")

	return cmd
}

func list(cmd *cobra.Command, args []string) {
	api := root.NewAPI()

	var project, org, env string
	project, _ = cmd.Flags().GetString("project")
	org, _ = cmd.Flags().GetString("organization")
	env, _ = cmd.Flags().GetString("environment")
	// project = viper.GetString("project")

	ebP := organization_external_backends.NewGetExternalBackendsParams()
	ebP.SetEnvironment(&env)
	ebP.SetOrganizationCanonical(org)
	ebP.SetProject(&project)
	resp, err := api.OrganizationExternalBackends.GetExternalBackends(ebP, nil)
	// api.OrganizationExternalBackends.GetExternalBackends(params *GetExternalBackendsParams, authInfo runtime.ClientAuthInfoWriter)
	fmt.Println("...")
	fmt.Println(resp)
	fmt.Printf("%+v\n", err)
}

// /organizations/{organization_canonical}/external_backends
// get: getExternalBackends
// Get the list of organization external backends
