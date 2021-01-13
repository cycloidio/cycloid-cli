package externalBackends

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	models "github.com/cycloidio/cycloid-cli/client/models"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/printer/factory"
)

func createInfraView(cmd *cobra.Command, args []string) error {
	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	var (
		purpose           = "remote_tfstate"
		err               error
		org, project, env string
		ebC               models.ExternalBackendConfiguration
		engine            = cmd.CalledAs()
	)

	output, err := cmd.Flags().GetString("output")
	if err != nil {
		return errors.Wrap(err, "unable to get output flag")
	}
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
	cred, err := cmd.Flags().GetString("cred")
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

	switch engine {
	case "AWSRemoteTFState":
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

	case "SwiftRemoteTFState":

		ebC = &models.SwiftRemoteTFState{
			Container:     &bucketName,
			Object:        &bucketpath,
			SkipVerifySsl: skipSSL,
			Region:        &region,
		}

	case "GCPRemoteTFState":

		ebC = &models.GCPRemoteTFState{
			Bucket: &bucketName,
			Object: &bucketpath,
		}
	default:
		return fmt.Errorf("Unexpected backend name")
	}

	// fetch the printer from the factory
	p, err := factory.GetPrinter(output)
	if err != nil {
		return errors.Wrap(err, "unable to get printer")
	}

	resp, err := m.CreateExternalBackends(org, project, env, purpose, cred, ebC)
	if err != nil {
		// print the result on the standard output
		if err := p.Print(err, printer.Options{}, os.Stdout); err != nil {
			return errors.Wrap(err, "unable to print result")
		}
		return errors.Wrap(err, "unable to create external backend")
	}

	// print the result on the standard output
	if err := p.Print(resp, printer.Options{}, os.Stdout); err != nil {
		return errors.Wrap(err, "unable to print result")
	}

	return nil
}
