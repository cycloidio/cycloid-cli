package external_backends

import (
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
)

const noDefault = false

var hostMapping, messageMapping, timestampMapping, esIndex string

func NewCreateCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "create [log|infraview]",
		Args:  cobra.NoArgs,
		Short: "create an external-backend",
		Example: `
       # create AWS external backend for logs
       cy --org my-org eb create logs AWSCloudWatchLogs --project foo --org bar --region eu-west-1

       # create Elasticsearch external backend for logs
       cy --org my-org eb create logs ElasticsearchLogs source-name --project foo --org bar --env foo --cred my-credential --url http://elasticsearch.local.tld --prefilter foo=bar
`,
	}

	infraViewCmd := newInfraViewCommand()
	logsCmd := newLogsCommand()
	eventsCmd := newEventsCommand()

	cmd.AddCommand(logsCmd, eventsCmd, infraViewCmd)
	return cmd
}

func newInfraViewCommand() *cobra.Command {

	var infraViewCmd = &cobra.Command{
		Use: "infraview [backend]",
		// RunE: create,
	}
	common.RequiredPersistentFlag(common.WithFlagProject, infraViewCmd)
	common.RequiredPersistentFlag(common.WithFlagEnv, infraViewCmd)

	// AWSRemoteTFState
	var aWSRemoteTFState = &cobra.Command{
		Use:  "AWSRemoteTFState",
		Args: cobra.NoArgs,
		RunE: createInfraView,
	}
	WithFlagAwsRegion(aWSRemoteTFState)
	common.RequiredFlag(common.WithFlagCred, aWSRemoteTFState)
	WithFlagBucketName(aWSRemoteTFState)
	WithFlagBucketPath(aWSRemoteTFState)
	WithFlagS3BucketEndpoint(aWSRemoteTFState)
	WithFlagS3ForcePathStyle(aWSRemoteTFState)
	WithFlagSkipVerifySSL(aWSRemoteTFState)
	WithFlagDefault(aWSRemoteTFState)

	var gCPRemoteTFState = &cobra.Command{
		Use:  "GCPRemoteTFState",
		Args: cobra.NoArgs,
		RunE: createInfraView,
	}
	common.RequiredFlag(common.WithFlagCred, gCPRemoteTFState)
	WithFlagBucketName(gCPRemoteTFState)
	WithFlagBucketPath(gCPRemoteTFState)
	WithFlagDefault(gCPRemoteTFState)

	var swiftRemoteTFState = &cobra.Command{
		Use:  "SwiftRemoteTFState",
		Args: cobra.NoArgs,
		RunE: createInfraView,
	}
	common.RequiredFlag(common.WithFlagCred, swiftRemoteTFState)
	WithFlagBucketName(swiftRemoteTFState)
	WithFlagBucketPath(swiftRemoteTFState)
	common.RequiredFlag(WithFlagRegion, swiftRemoteTFState)
	WithFlagSkipVerifySSL(swiftRemoteTFState)
	WithFlagDefault(swiftRemoteTFState)

	infraViewCmd.AddCommand(
		aWSRemoteTFState,
		gCPRemoteTFState,
		swiftRemoteTFState,
	)
	return infraViewCmd
}

func newEventsCommand() *cobra.Command {
	var eventCmd = &cobra.Command{
		Use:    "events [backend]",
		Hidden: true,
	}

	// Aws CW logs
	var eventsAWSCloudWatchLogsCmd = &cobra.Command{
		Use:  "AWSCloudWatchLogs",
		Args: cobra.NoArgs,
		RunE: createEvents,
	}
	WithFlagAwsRegion(eventsAWSCloudWatchLogsCmd)
	common.RequiredFlag(common.WithFlagCred, eventsAWSCloudWatchLogsCmd)

	eventCmd.AddCommand(eventsAWSCloudWatchLogsCmd)
	return eventCmd
}

func newLogsCommand() *cobra.Command {
	var logsCmd = &cobra.Command{
		Use: "logs [backend]",
	}

	// Aws CW logs
	var logsAWSCloudWatchLogsCmd = &cobra.Command{
		Use:  "AWSCloudWatchLogs",
		RunE: createLogs,
	}
	WithFlagAwsRegion(logsAWSCloudWatchLogsCmd)
	common.RequiredPersistentFlag(common.WithFlagProject, logsAWSCloudWatchLogsCmd)
	common.RequiredFlag(common.WithFlagCred, logsAWSCloudWatchLogsCmd)

	// Elasticsearch
	var logsElasticsearchLogsCmd = &cobra.Command{
		Use:  "ElasticsearchLogs [SourceName]",
		Args: cobra.ExactArgs(1),
		RunE: createLogs,
	}
	common.RequiredPersistentFlag(common.WithFlagProject, logsElasticsearchLogsCmd)
	common.RequiredPersistentFlag(common.WithFlagEnv, logsElasticsearchLogsCmd)
	common.RequiredFlag(WithFlagUrl, logsElasticsearchLogsCmd)
	common.RequiredFlag(WithFlagPrefilter, logsElasticsearchLogsCmd)
	common.WithFlagCred(logsElasticsearchLogsCmd)
	logsElasticsearchLogsCmd.Flags().StringVar(&esIndex, "index", "default", "")
	logsElasticsearchLogsCmd.Flags().StringVar(&hostMapping, "host-mapping", "hostname", "")
	logsElasticsearchLogsCmd.Flags().StringVar(&messageMapping, "message-mapping", "message", "")
	logsElasticsearchLogsCmd.Flags().StringVar(&timestampMapping, "timestamp-mapping", "timestamp", "")

	// Note if one exist it should trigger an update (to append the new source to sources if the source does not exist yet)
	// Will require list sources maybe ?

	logsCmd.AddCommand(logsAWSCloudWatchLogsCmd, logsElasticsearchLogsCmd)
	return logsCmd
}
