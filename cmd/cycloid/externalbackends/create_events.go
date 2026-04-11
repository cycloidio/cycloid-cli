package externalbackends

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/client/models"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
	"github.com/cycloidio/cycloid-cli/printer"
)

func createEvents(cmd *cobra.Command, args []string) error {
	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	var (
		err     error
		cred    string
		purpose = "events"
		ebC     models.ExternalBackendConfiguration
		engine  = cmd.CalledAs()
	)

	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}

	// AWS CW logs
	if engine == "AWSCloudWatchLogs" {
		region, err := cmd.Flags().GetString("region")
		if err != nil {
			return err
		}
		cred, err = cmd.Flags().GetString("cred")
		if err != nil {
			return err
		}

		ebC = &models.AWSCloudWatchLogs{
			Region: &region,
		}
	} else {
		return fmt.Errorf("unexpected backend name")
	}

	// Set project and env to empty cause events are not linked to a project
	project := ""
	env := ""
	resp, _, err := m.CreateExternalBackends(org, project, env, purpose, cred, noDefault, ebC)
	return cyout.PrintWithOptions(cmd, resp, err, "unable to create external backend", printer.Options{})
}
