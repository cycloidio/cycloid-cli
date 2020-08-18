package externalBackends

import (
	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/common"
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
		RunE: createInfraView,
	}
	WithFlagAwsRegion(aWSRemoteTFState)
	common.RequiredFlag(common.WithFlagCred, aWSRemoteTFState)
	WithFlagBucketName(aWSRemoteTFState)
	WithFlagBucketPath(aWSRemoteTFState)
	WithFlagS3BucketEndpoint(aWSRemoteTFState)
	WithFlagS3ForcePathStyle(aWSRemoteTFState)
	WithFlagSkipVerifySSL(aWSRemoteTFState)

	var gCPRemoteTFState = &cobra.Command{
		Use:  "GCPRemoteTFState",
		RunE: createInfraView,
	}
	common.RequiredFlag(common.WithFlagCred, gCPRemoteTFState)
	WithFlagBucketName(gCPRemoteTFState)
	WithFlagBucketPath(gCPRemoteTFState)

	var swiftRemoteTFState = &cobra.Command{
		Use:  "SwiftRemoteTFState",
		RunE: createInfraView,
	}
	common.RequiredFlag(common.WithFlagCred, swiftRemoteTFState)
	WithFlagBucketName(swiftRemoteTFState)
	WithFlagBucketPath(swiftRemoteTFState)
	common.RequiredFlag(WithFlagRegion, swiftRemoteTFState)
	WithFlagSkipVerifySSL(swiftRemoteTFState)

	infraViewCmd.AddCommand(aWSRemoteTFState,
		gCPRemoteTFState,
		swiftRemoteTFState)
	return infraViewCmd
}

func newEventsCommand() *cobra.Command {
	var eventCmd = &cobra.Command{
		Use: "events [backend]",
	}

	// Aws CW logs
	var eventsAWSCloudWatchLogsCmd = &cobra.Command{
		Use:  "AWSCloudWatchLogs",
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

// /organizations/{organization_canonical}/external_backends
// post: createExternalBackend
// Save information about the external backend
