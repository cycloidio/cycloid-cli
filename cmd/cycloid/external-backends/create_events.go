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

func createEvents(cmd *cobra.Command, args []string) error {
	api := root.NewAPI()
	var purpose = "events"
	var err error
	var org string
	var body *models.NewExternalBackend
	var ebC models.ExternalBackendConfiguration
	var engine = cmd.CalledAs()

	org, err = cmd.Flags().GetString("org")
	if err != nil {
		return err
	}

	ebParams := organization_external_backends.NewCreateExternalBackendParams()
	ebParams.SetOrganizationCanonical(org)

	// AWS CW logs
	if engine == "AWSCloudWatchLogs" {
		region, err := cmd.Flags().GetString("region")
		if err != nil {
			return err
		}
		cred, err := cmd.Flags().GetUint32("cred")
		if err != nil {
			return err
		}

		ebC = &models.AWSCloudWatchLogs{
			Region: &region,
		}
		body = &models.NewExternalBackend{
			Purpose:      &purpose,
			CredentialID: cred,
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

	resp, err := api.OrganizationExternalBackends.CreateExternalBackend(ebParams, root.ClientCredentials())
	// TODO create a error handeling function to format our error with a better display
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
