package pipelines

import (
	"context"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/cycloidio/cycloid-cli/cmd/apiclient"
	"github.com/cycloidio/cycloid-cli/cmd/common"
	"github.com/cycloidio/cycloid-cli/internal/buildwatch"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
)

func NewBuildLogsCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "logs",
		Aliases: []string{"log"},
		Short:   "stream the event log of a pipeline build",
		Example: `
  # Dump all available logs for a past build and exit
  cy pp builds logs --project p --env e --component c --pipeline pp --job j --build-id 42

  # Tail logs for an in-progress build (keeps streaming until the build finishes)
  cy pp builds logs --project p --env e --component c --pipeline pp --job j --build-id 42 --watch
`,
		RunE: streamBuildLogs,
		Args: cobra.NoArgs,
	}

	cyargs.AddCyContext(cmd)
	cyargs.AddPipeline(cmd)
	cmd.MarkFlagRequired(cyargs.AddPipelineJob(cmd))
	cmd.MarkFlagRequired(cyargs.AddPipelineBuildID(cmd))
	cmd.Flags().Bool("watch", false, "keep tailing the event stream (useful for in-progress builds)")
	return cmd
}

func streamBuildLogs(cmd *cobra.Command, args []string) error {
	org, project, env, component, err := cyargs.GetCyContext(cmd)
	if err != nil {
		return err
	}

	pipeline, err := cyargs.GetPipeline(cmd)
	if err != nil {
		return err
	}

	job, err := cyargs.GetPipelineJob(cmd)
	if err != nil {
		return err
	}

	buildID, err := cyargs.GetPipelineBuildID(cmd)
	if err != nil {
		return err
	}

	output, err := cyargs.GetOutput(cmd)
	if err != nil {
		return fmt.Errorf("unable to get output flag: %w", err)
	}

	watch, err := cmd.Flags().GetBool("watch")
	if err != nil {
		return fmt.Errorf("unable to get --watch flag: %w", err)
	}

	api := common.NewAPI()
	m := apiclient.NewMiddleware(api)

	outMode := buildwatch.OutputHuman
	if strings.EqualFold(output, "json") {
		outMode = buildwatch.OutputRaw
	}

	theme := buildwatch.StreamTheme{}
	stripLogANSI := false
	if outMode == buildwatch.OutputHuman {
		if buildwatch.IsTerminalWriter(cmd.OutOrStdout()) {
			theme = buildwatch.DefaultStreamTheme
		} else {
			stripLogANSI = true
		}
	}

	verbose := strings.EqualFold(viper.GetString("verbosity"), "debug")

	return buildwatch.StreamLogs(
		context.Background(),
		m,
		org, project, env, component, pipeline, job, buildID,
		buildwatch.Options{
			Watch:        watch,
			LogWriter:    cmd.OutOrStdout(),
			Output:       outMode,
			Verbose:      verbose,
			Theme:        theme,
			StripLogANSI: stripLogANSI,
			StatusWriter: cmd.ErrOrStderr(),
		},
	)
}
