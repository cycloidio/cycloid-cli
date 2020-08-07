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

	var infraViewCmd = &cobra.Command{
		Use: "infraview",
		// RunE: create,
	}
	common.RequiredPersistentFlag(common.WithFlagOrg, infraViewCmd)
	// purpose = remote_tfstate

	// WithFlagPurpose(cmd)
	// log.Flags().String("aws-region", "default-p", "Purpose")

	logsCmd := newLogsCommand()
	eventsCmd := newEventsCommand()

	cmd.AddCommand(logsCmd, eventsCmd, infraViewCmd)

	return cmd
}

func newEventsCommand() *cobra.Command {
	var eventCmd = &cobra.Command{
		Use: "events [backend]",
	}
	common.RequiredPersistentFlag(common.WithFlagOrg, eventCmd)

	// Aws CW logs
	var eventsAWSCloudWatchLogsCmd = &cobra.Command{
		Use:  "AWSCloudWatchLogs",
		RunE: createEvents,
	}
	common.RequiredFlag(WithFlagAwsRegion, eventsAWSCloudWatchLogsCmd)
	common.RequiredFlag(common.WithFlagCred, eventsAWSCloudWatchLogsCmd)

	eventCmd.AddCommand(eventsAWSCloudWatchLogsCmd)
	return eventCmd
}

func newLogsCommand() *cobra.Command {
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
	common.RequiredFlag(WithFlagUrl, logsElasticsearchLogsCmd)
	common.RequiredFlag(WithFlagPrefilter, logsElasticsearchLogsCmd)
	common.WithFlagCred(logsElasticsearchLogsCmd)
	logsElasticsearchLogsCmd.Flags().StringVar(&esIndex, "index", "default", "")
	// Mapping flags
	logsElasticsearchLogsCmd.Flags().StringVar(&hostMapping, "host-mapping", "hostname", "")
	logsElasticsearchLogsCmd.Flags().StringVar(&messageMapping, "message-mapping", "message", "")
	logsElasticsearchLogsCmd.Flags().StringVar(&timestampMapping, "timestamp-mapping", "timestamp", "")
	var dd uint32
	_ = dd
	logsElasticsearchLogsCmd.Flags().Uint32("foo", 0, "")
	// logsElasticsearchLogsCmd.Flags().Uint32("foo", dd, "")

	// Note if one exist it should trigger an update (to append the new source to sources if the source does not exist yet)
	// Will require list sources maybe ?

	logsCmd.AddCommand(logsAWSCloudWatchLogsCmd, logsElasticsearchLogsCmd)
	return logsCmd
}

// /organizations/{organization_canonical}/external_backends
// post: createExternalBackend
// Save information about the external backend
