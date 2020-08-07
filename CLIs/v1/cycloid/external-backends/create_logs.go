package externalBackends

import (
	"errors"
	"fmt"

	"github.com/cycloidio/youdeploy-cli/client/client/organization_external_backends"
	models "github.com/cycloidio/youdeploy-cli/client/models"
	root "github.com/cycloidio/youdeploy-cli/cmd/cycloid"
	"github.com/davecgh/go-spew/spew"
	"github.com/go-openapi/runtime"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/spf13/cobra"
)

func createLogs(cmd *cobra.Command, args []string) error {
	api := root.NewAPI()
	var purpose = "logs"
	var err error

	var project, org string

	project, err = cmd.Flags().GetString("project")
	if err != nil {
		return err
	}
	org, err = cmd.Flags().GetString("org")
	if err != nil {
		return err
	}

	ebParams := organization_external_backends.NewCreateExternalBackendParams()
	ebParams.SetOrganizationCanonical(org)
	var body *models.NewExternalBackend
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
		body = &models.NewExternalBackend{
			ProjectCanonical: project,
			Purpose:          &purpose,
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
		cred, err := cmd.Flags().GetUint32("cred")
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

		// This env param exist but does not seems to be used by api
		// I feel it's a swagger issue
		// ebParams.SetEnvironment(&env)

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

		ebParams.SetProject(&project)
		version := "7"
		ebC = &models.ElasticsearchLogs{
			Version: &version,
			Urls:    urls,
			Sources: sources,
		}

		if cred != 0 {
			body = &models.NewExternalBackend{
				ProjectCanonical: project,
				Purpose:          &purpose,
				CredentialID:     cred,
			}
		} else {
			body = &models.NewExternalBackend{
				ProjectCanonical: project,
				Purpose:          &purpose,
			}
		}
	} else {
		return errors.New("Unexpected backend name")
	}

	err = ebC.Validate(strfmt.Default)
	if err != nil {
		return err
	}

	body.SetConfiguration(ebC)
	ebParams.SetBody(body)
	err = body.Validate(strfmt.Default)
	if err != nil {
		return err
	}

	resp, err := api.OrganizationExternalBackends.CreateExternalBackend(ebParams, nil)
	if err != nil {
		// *errors.Validation, not *runtime.APIError
		apiErr, ok := err.(*runtime.APIError)
		if ok {
			spew.Dump(apiErr.Error())
			r := apiErr.Response.(runtime.ClientResponse)
			spew.Dump(r.Message())
		}
		// fmt.Printf("%+v\n", err.Error())
		return err
	}
	fmt.Println(resp)

	return nil
}
