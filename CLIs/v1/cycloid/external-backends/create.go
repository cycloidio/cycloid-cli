package externalBackends

import (
	"fmt"

	"github.com/cycloidio/youdeploy-cli/client/client/organization_external_backends"
	models "github.com/cycloidio/youdeploy-cli/client/models"
	root "github.com/cycloidio/youdeploy-cli/cmd/cycloid"
	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/common"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/spf13/cobra"
)

var hostMapping, messageMapping, timestampMapping, esIndex string

func NewCreateCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use: "create [purpose]",
		// Args:       cobra.OnlyValidArgs,
		// ValidArgs:  []string{"log", "infraview"},
		Short: "...",
		Long:  `........ . . .... .. .. ....`,
	}

	var eventsCmd = &cobra.Command{
		Use:  "events",
		RunE: create,
	}
	common.RequiredPersistentFlag(common.WithFlagOrg, eventsCmd)
	// purpose = events

	var infraViewCmd = &cobra.Command{
		Use:  "infraview",
		RunE: create,
	}
	common.RequiredPersistentFlag(common.WithFlagOrg, infraViewCmd)
	// purpose = remote_tfstate

	// WithFlagPurpose(cmd)
	// log.Flags().String("aws-region", "default-p", "Purpose")

	logsCmd := newLogCommand()

	cmd.AddCommand(logsCmd, eventsCmd, infraViewCmd)

	return cmd
}

func newLogCommand() *cobra.Command {
	var logsCmd = &cobra.Command{
		Use: "logs [backend]",
	}
	common.RequiredPersistentFlag(common.WithFlagOrg, logsCmd)

	// Aws CW logs
	var logsAWSCloudWatchLogsCmd = &cobra.Command{
		Use:  "AWSCloudWatchLogs",
		RunE: createLogs,
	}
	common.RequiredFlag(WithFlagAwsRegion, logsAWSCloudWatchLogsCmd)
	common.RequiredPersistentFlag(common.WithFlagProject, logsAWSCloudWatchLogsCmd)

	// Elasticsearch
	var logsElasticsearchLogsCmd = &cobra.Command{
		Use:  "ElasticsearchLogs [SourceName]",
		Args: cobra.ExactArgs(1),
		RunE: createLogs,
	}
	common.RequiredPersistentFlag(common.WithFlagProject, logsElasticsearchLogsCmd)
	common.RequiredPersistentFlag(common.WithFlagEnv, logsElasticsearchLogsCmd)
	common.RequiredFlag(common.WithFlagCred, logsElasticsearchLogsCmd)
	common.RequiredFlag(WithFlagUrl, logsElasticsearchLogsCmd)
	common.RequiredFlag(WithFlagPrefilter, logsElasticsearchLogsCmd)
	logsElasticsearchLogsCmd.Flags().StringVar(&esIndex, "index", "default", "")
	// Mapping flags
	logsElasticsearchLogsCmd.Flags().StringVar(&hostMapping, "host-mapping", "hostname", "")
	logsElasticsearchLogsCmd.Flags().StringVar(&messageMapping, "message-mapping", "message", "")
	logsElasticsearchLogsCmd.Flags().StringVar(&timestampMapping, "timestamp-mapping", "timestamp", "")

	// ---
	// credential_id: 1296
	// project_canonical: testlogs
	// configuration:
	//   engine: ElasticsearchLogs
	//   version: '7'
	//   urls:
	//   - http://testurl
	//   sources:
	//     test:
	//       eb1:
	//         index: app-logs
	//         mapping:
	//           host: hostname
	//           message: message,status
	//           timestamp: time
	//         prefilters:
	//           bar: foo
	//           foo: bar
	//         urls: []
	//       eb2:
	//         index: app-logs2
	//         urls: []
	//         mapping:
	//           host: hostname
	//           timestamp: time
	//           message: message,status
	//         prefilters: {}
	//
	// ./cy  external-backends  create logs ElasticsearchLogs eb1 --index app-logs
	// --host-mapping hostname --message-mapping "message,status" --timestamp-mapping time
	// --prefilter bar=foo
	// --prefilter foo=bar
	// --url
	// --project foo --org bar --env test

	// Note if one exist it should trigger an update (to append the new source to sources if the source does not exist yet)
	// Will require list sources maybe ?

	logsCmd.AddCommand(logsAWSCloudWatchLogsCmd, logsElasticsearchLogsCmd)
	return logsCmd
}

func create(cmd *cobra.Command, args []string) error {
	api := root.NewAPI()

	var project, org, env string

	var region = "eu-west-1"
	var purpose = cmd.CalledAs()
	// 'events', 'logs', 'remote_tfstate'

	project, _ = cmd.Flags().GetString("project")
	org, _ = cmd.Flags().GetString("org")
	// project = viper.GetString("project")

	ebP := organization_external_backends.NewCreateExternalBackendParams()
	ebP.SetOrganizationCanonical(org)

	// ebP.SetEnvironment(&env)
	// ebP.SetProject(&project)

	var body *models.NewExternalBackend
	ebC := models.AWSCloudWatchLogs{
		Region: &region,
	}
	err := ebC.Validate(strfmt.Default)
	if err != nil {
		return err
	}

	body = &models.NewExternalBackend{
		CredentialID:         2,
		EnvironmentCanonical: env,
		ProjectCanonical:     project,
		Purpose:              &purpose,
	}
	body.SetConfiguration(&ebC)

	err = body.Validate(strfmt.Default)
	if err != nil {
		return err
	}

	ebP.SetBody(body)

	resp, err := api.OrganizationExternalBackends.CreateExternalBackend(ebP, nil)
	if err != nil {
		return err
	}

	// api.OrganizationExternalBackends.GetExternalBackends(params *GetExternalBackendsParams, authInfo runtime.ClientAuthInfoWriter)
	fmt.Println("...")
	fmt.Println(resp)
	// fmt.Printf("%+v\n", err)

	return nil
}

// /organizations/{organization_canonical}/external_backends
// post: createExternalBackend
// Save information about the external backend
