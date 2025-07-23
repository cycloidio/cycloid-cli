package external_backends

import (
	"fmt"

	strfmt "github.com/go-openapi/strfmt"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	models "github.com/cycloidio/cycloid-cli/client/models"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/printer/factory"
)

func createLogs(cmd *cobra.Command, args []string) error {
	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	var (
		ebC     models.ExternalBackendConfiguration
		engine  = cmd.CalledAs()
		cred    string
		purpose = "logs"
	)

	project, err := cmd.Flags().GetString("project")
	if err != nil {
		return err
	}
	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}
	output, err := cmd.Flags().GetString("output")
	if err != nil {
		return errors.Wrap(err, "unable to get output flag")
	}
	cred, err = cmd.Flags().GetString("cred")
	if err != nil {
		return err
	}

	switch engine {
	case "AWSCloudWatchLogs":
		region, err := cmd.Flags().GetString("region")
		if err != nil {
			return err
		}

		ebC = &models.AWSCloudWatchLogs{
			Region: &region,
		}

	case "ElasticsearchLogs":
		prefilters, err := cmd.Flags().GetStringToString("prefilter")
		if err != nil {
			return err
		}
		urls, err := cmd.Flags().GetStringSlice("url")
		if err != nil {
			return err
		}
		env, err := cmd.Flags().GetString("env")
		if err != nil {
			return err
		}
		// TODO Note: are those useful as I also can directly access to the local defined var for the flag
		hostMapping, err := cmd.Flags().GetString("host-mapping")
		if err != nil {
			return err
		}
		messageMapping, err := cmd.Flags().GetString("message-mapping")
		if err != nil {
			return err
		}
		timestampMapping, err := cmd.Flags().GetString("timestamp-mapping")
		if err != nil {
			return err
		}
		esIndex, err := cmd.Flags().GetString("index")
		if err != nil {
			return err
		}

		// In this case we know there is exactly one arg
		var sourceName = args[0]

		esM := models.ElasticsearchLogsSourcesAnonMapping{
			Host:      &hostMapping,
			Message:   &messageMapping,
			Timestamp: &timestampMapping,
		}
		err = esM.Validate(strfmt.Default)
		if err != nil {
			return err
		}

		esS := models.ElasticsearchLogsSourcesAnon{
			Index:      esIndex,
			Mapping:    &esM,
			Prefilters: prefilters,
		}
		err = esS.Validate(strfmt.Default)
		if err != nil {
			return err
		}

		sources := map[string]map[string]models.ElasticsearchLogsSourcesAnon{
			env: map[string]models.ElasticsearchLogsSourcesAnon{
				sourceName: esS,
			},
		}

		version := "7"
		ebC = &models.ElasticsearchLogs{
			Version: &version,
			Urls:    urls,
			Sources: sources,
		}

	default:
		return fmt.Errorf("Unexpected backend name")
	}

	// fetch the printer from the factory
	p, err := factory.GetPrinter(output)
	if err != nil {
		return errors.Wrap(err, "unable to get printer")
	}

	// Set env to empty cause is not used to create log eb
	envP := ""
	resp, err := m.CreateExternalBackends(org, project, envP, purpose, cred, noDefault, ebC)
	return printer.SmartPrint(p, resp, err, "unable to create external backend", printer.Options{}, cmd.OutOrStdout())
}
