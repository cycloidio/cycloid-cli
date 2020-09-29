package externalBackends

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	models "github.com/cycloidio/youdeploy-cli/client/models"
	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/common"
	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/youdeploy-cli/printer"
	"github.com/cycloidio/youdeploy-cli/printer/factory"
)

func createEvents(cmd *cobra.Command, args []string) error {
	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	var (
		cred    uint32 = 0
		err     error
		org     string
		purpose = "events"
		ebC     models.ExternalBackendConfiguration
		engine  = cmd.CalledAs()
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
		cred, err = cmd.Flags().GetUint32("cred")
		if err != nil {
			return err
		}

		ebC = &models.AWSCloudWatchLogs{
			Region: &region,
		}
	} else {
		return fmt.Errorf("Unexpected backend name")
	}

	// Set project and env to empty cause events are not linked to a project
	project := ""
	env := ""
	resp, err := m.CreateExternalBackends(org, project, env, purpose, cred, ebC)
	if err != nil {
		return errors.Wrap(err, "unable to create external backend")
	}

	// fetch the printer from the factory
	p, err := factory.GetPrinter(output)
	if err != nil {
		return errors.Wrap(err, "unable to get printer")
	}

	// print the result on the standard output
	if err := p.Print(resp, printer.Options{}, os.Stdout); err != nil {
		return errors.Wrap(err, "unable to print result")
	}

	return nil
}
