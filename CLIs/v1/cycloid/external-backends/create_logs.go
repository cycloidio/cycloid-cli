package externalBackends

import (
	"errors"
	"fmt"

	models "github.com/cycloidio/youdeploy-cli/client/models"
	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/common"
	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/middleware"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/spf13/cobra"
)

func createLogs(cmd *cobra.Command, args []string) error {
	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	var purpose = "logs"
	var cred uint32
	cred = 0

	project, err := cmd.Flags().GetString("project")
	if err != nil {
		return err
	}
	org, err := cmd.Flags().GetString("org")
	if err != nil {
		return err
	}

	var ebC models.ExternalBackendConfiguration
	var engine = cmd.CalledAs()

	// AWS CW logs
	if engine == "AWSCloudWatchLogs" {
		region, err := cmd.Flags().GetString("region")
		if err != nil {
			return err
		}

		ebC = &models.AWSCloudWatchLogs{
			Region: &region,
		}

		// Elasticsearch
	} else if engine == "ElasticsearchLogs" {
		prefilters, err := cmd.Flags().GetStringToString("prefilter")
		if err != nil {
			return err
		}
		urls, err := cmd.Flags().GetStringSlice("url")
		if err != nil {
			return err
		}
		cred, err = cmd.Flags().GetUint32("cred")
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

	} else {
		return errors.New("Unexpected backend name")
	}

	// Set env to empty cause is not used to create log eb
	envP := ""
	resp, err := m.CreateExternalBackends(org, project, envP, purpose, cred, ebC)
	if err != nil {
		return err
	}

	// if err != nil {
	// 	// *errors.Validation, not *runtime.APIError
	// 	apiErr, ok := err.(*runtime.APIError)
	// 	if ok {
	// 		spew.Dump(apiErr.Error())
	// 		r := apiErr.Response.(runtime.ClientResponse)
	// 		spew.Dump(r.Message())
	// 	}
	// 	// fmt.Printf("%+v\n", err.Error())
	// 	return err
	// }
	fmt.Println(resp)

	return nil
}
