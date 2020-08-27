package externalBackends

import (
	"errors"
	"fmt"

	models "github.com/cycloidio/youdeploy-cli/client/models"
	root "github.com/cycloidio/youdeploy-cli/cmd/cycloid"
	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/middleware"
	"github.com/spf13/cobra"
)

func createEvents(cmd *cobra.Command, args []string) error {
	api := root.NewAPI()
	m := middleware.NewMiddleware(api)

	var purpose = "events"
	var cred uint32
	cred = 0

	var err error
	var org string
	var ebC models.ExternalBackendConfiguration
	var engine = cmd.CalledAs()

	org, err = cmd.Flags().GetString("org")
	if err != nil {
		return err
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
		return errors.New("Unexpected backend name")
	}

	// Set project and env to empty cause events is not linked to a project
	project := ""
	env := ""
	resp, err := m.CreateExternalBackends(org, project, env, purpose, cred, ebC)
	if err != nil {
		return err
	}

	fmt.Println(resp)

	return nil
}
