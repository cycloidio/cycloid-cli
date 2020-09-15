package externalBackends

import (
	"errors"
	"fmt"

	models "github.com/cycloidio/youdeploy-cli/client/models"
	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/common"
	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/middleware"
	"github.com/spf13/cobra"
)

func createInfraView(cmd *cobra.Command, args []string) error {
	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	var purpose = "remote_tfstate"
	var err error
	var org, project, env string
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

	resp, err := m.CreateExternalBackends(org, project, env, purpose, cred, ebC)
	if err != nil {
		return err
	}

	fmt.Println(resp)

	return nil
}
