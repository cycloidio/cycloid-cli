package organizations

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/common"
	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/middleware"
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
	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	org, err := cmd.Flags().GetString("org")
	if err != nil {
		return err
	}

	ws, err := m.ListOrganizationWorkers(org)
	if err != nil {
		return err
	}

	for _, d := range ws {
		fmt.Printf("Name: %s    StartTime: %s     State: %s    Ephemeral: %s  \n", *d.Name, *d.StartTime, *d.State, *d.Ephemeral)
		fmt.Printf("  ActiveContainers: %s    ActiveVolumes: %s     Tags: %s    Version: %s  \n", *d.ActiveContainers, *d.ActiveVolumes, d.Tags, *d.Version)
	}
	fmt.Println(ws)
	fmt.Printf("%+v\n", err)
	return nil
}
