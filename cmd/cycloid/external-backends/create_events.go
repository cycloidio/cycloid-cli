package externalBackends

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	models "github.com/cycloidio/cycloid-cli/client/models"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/printer/factory"
)

func createEvents(cmd *cobra.Command, args []string) error {
	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	var (
		err       error
		org, cred string
		purpose   = "events"
		ebC       models.ExternalBackendConfiguration
		engine    = cmd.CalledAs()
	)

	org, err = cmd.Flags().GetString("org")
	if err != nil {
		return err
	}
	output, err := cmd.Flags().GetString("output")
	if err != nil {
		return errors.Wrap(err, "unable to get output flag")
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
		return fmt.Errorf("Unexpected backend name")
	}

	// fetch the printer from the factory
	p, err := factory.GetPrinter(output)
	if err != nil {
		return errors.Wrap(err, "unable to get printer")
	}

	// Set project and env to empty cause events are not linked to a project
	project := ""
	env := ""
	resp, err := m.CreateExternalBackends(org, project, env, purpose, cred, ebC)
	return printer.SmartPrint(p, resp, err, "unable to create external backend", printer.Options{}, cmd.OutOrStdout())
}
