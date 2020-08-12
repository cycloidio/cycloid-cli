package externalBackends

import (
	"errors"
	"fmt"

	"github.com/cycloidio/youdeploy-cli/client/client/organization_external_backends"
	models "github.com/cycloidio/youdeploy-cli/client/models"
	root "github.com/cycloidio/youdeploy-cli/cmd/cycloid"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/spf13/cobra"
)

func createInfraView(cmd *cobra.Command, args []string) error {
	api := root.NewAPI()
	var purpose = "remote_tfstate"
	var err error
	var org, project, env string
	var body *models.NewExternalBackend
	var ebC models.ExternalBackendConfiguration
	var engine = cmd.CalledAs()

	org, err = cmd.Flags().GetString("org")
	if err != nil {
		return err
	}
	project, err = cmd.Flags().GetString("project")
	if err != nil {
		return err
	}
	env, err = cmd.Flags().GetString("env")
	if err != nil {
		return err
	}

	ebParams := organization_external_backends.NewCreateExternalBackendParams()
	ebParams.SetOrganizationCanonical(org)
	//Common
	cred, err := cmd.Flags().GetUint32("cred")
	if err != nil {
		return err
	}

	bucketName, err := cmd.Flags().GetString("bucket-name")
	if err != nil {
		return err
	}
	if bucketName == "" {
		bucketName = fmt.Sprintf("%s-terraform-remote-state", org)
	}

	bucketpath, err := cmd.Flags().GetString("bucket-path")
	if err != nil {
		return err
	}
	if bucketpath == "" {
		bucketpath = fmt.Sprintf("%s/%s/%s-%s.tfstate", project, env, project, env)
	}
	skipSSL, err := cmd.Flags().GetBool("skip-verify-ssl")
	if err != nil {
		return err
	}
	region, err := cmd.Flags().GetString("region")
	if err != nil {
		return err
	}

	// AWSRemoteTFState
	if engine == "AWSRemoteTFState" {
		endpoint, err := cmd.Flags().GetString("endpoint")
		if err != nil {
			return err
		}

		forcePathStyle, err := cmd.Flags().GetBool("s3-force-path-style")
		if err != nil {
			return err
		}

		if endpoint != "" {
			ebC = &models.AWSRemoteTFState{
				Region:           &region,
				Bucket:           &bucketName,
				Key:              &bucketpath,
				Endpoint:         endpoint,
				S3ForcePathStyle: forcePathStyle,
				SkipVerifySsl:    skipSSL,
			}
		} else {
			ebC = &models.AWSRemoteTFState{
				Region: &region,
				Bucket: &bucketName,
				Key:    &bucketpath,
			}
		}

		// SwiftRemoteTFState
	} else if engine == "SwiftRemoteTFState" {

		ebC = &models.SwiftRemoteTFState{
			Container:     &bucketName,
			Object:        &bucketpath,
			SkipVerifySsl: skipSSL,
			Region:        &region,
		}

		// GCPRemoteTFState
	} else if engine == "GCPRemoteTFState" {

		ebC = &models.GCPRemoteTFState{
			Bucket: &bucketName,
			Object: &bucketpath,
		}

	} else {
		return errors.New("Unexpected backend name")
	}

	err = ebC.Validate(strfmt.Default)
	if err != nil {
		return err
	}

	body = &models.NewExternalBackend{
		Purpose:              &purpose,
		ProjectCanonical:     project,
		EnvironmentCanonical: env,
		CredentialID:         cred,
	}

	body.SetConfiguration(ebC)
	ebParams.SetBody(body)
	err = body.Validate(strfmt.Default)
	if err != nil {
		return err
	}

	resp, err := api.OrganizationExternalBackends.CreateExternalBackend(ebParams, root.ClientCredentials())
	if err != nil {
		return err
	}
	fmt.Println(resp)

	return nil
}
