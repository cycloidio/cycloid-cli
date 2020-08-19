package organizations

import (
	"fmt"

	"github.com/cycloidio/youdeploy-cli/client/client/organization_workers"
	"github.com/spf13/cobra"

	root "github.com/cycloidio/youdeploy-cli/cmd/cycloid"
	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/common"
)

func NewListWorkersCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "list-workers",
		Short: "...",
		Long:  `........ . . .... .. .. ....`,
		RunE:  listWorkers,
	}
	common.RequiredPersistentFlag(common.WithFlagOrg, cmd)

	return cmd
}

func listWorkers(cmd *cobra.Command, args []string) error {
	api := root.NewAPI()

	org, err := cmd.Flags().GetString("org")
	if err != nil {
		return err
	}

	params := organization_workers.NewGetWorkersParams()
	params.SetOrganizationCanonical(org)

	resp, err := api.OrganizationWorkers.GetWorkers(params, root.ClientCredentials())
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
		fmt.Printf("Name: %s    StartTime: %s     State: %s    Ephemeral: %s  \n", *d.Name, *d.StartTime, *d.State, *d.Ephemeral)
		fmt.Printf("  ActiveContainers: %s    ActiveVolumes: %s     Tags: %s    Version: %s  \n", *d.ActiveContainers, *d.ActiveVolumes, d.Tags, *d.Version)
	}
	fmt.Println(resp)
	fmt.Printf("%+v\n", err)
	return nil
}

// /organizations/{organization_canonical}/workers
// get: getWorkers
// Get the workers that the authenticated user has access to.
