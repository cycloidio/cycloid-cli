package pipelines

import (
	"context"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/buildwatch"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/printer/factory"
)

func NewBuildCreateCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use: "trigger",
		Aliases: []string{
			"create",
			"run",
		},
		Short:   "trigger a pipeline build",
		Example: `cy pipeline build trigger --project my-project --env my-env --component my-component --job my-job`,
		RunE:    createBuild,
		Args:    cobra.NoArgs,
	}

	cyargs.AddCyContext(cmd)
	cyargs.AddPipeline(cmd)
	cmd.MarkFlagRequired(cyargs.AddPipelineJob(cmd))
	cmd.Flags().Bool("watch", false, "watch the build until completion")
	cmd.Flags().Int("timeout", 0, "timeout in seconds when used with --watch (0 means no timeout)")
	cmd.Flags().Bool("watch-cancel-on-timeout", false, "with --watch, abort the build when --timeout is reached")
	return cmd
}

func createBuild(cmd *cobra.Command, args []string) error {
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

	output, err := cyargs.GetOutput(cmd)
	if err != nil {
		return errors.Wrap(err, "unable to get output flag")
	}

	watch, err := cmd.Flags().GetBool("watch")
	if err != nil {
		return errors.Wrap(err, "unable to get watch flag")
	}
	timeoutSeconds, err := cmd.Flags().GetInt("timeout")
	if err != nil {
		return errors.Wrap(err, "unable to get timeout flag")
	}
	cancelOnTimeout, err := cmd.Flags().GetBool("watch-cancel-on-timeout")
	if err != nil {
		return errors.Wrap(err, "unable to get watch-cancel-on-timeout flag")
	}

	if timeoutSeconds < 0 {
		return errors.New("--timeout must be greater than or equal to 0")
	}
	if !watch && timeoutSeconds > 0 {
		return errors.New("--timeout can only be used with --watch")
	}
	if !watch && cancelOnTimeout {
		return errors.New("--watch-cancel-on-timeout can only be used with --watch")
	}
	if watch && cancelOnTimeout && timeoutSeconds == 0 {
		return errors.New("--watch-cancel-on-timeout requires --timeout to be set")
	}

	// fetch the printer from the factory
	p, err := factory.GetPrinter(output)
	if err != nil {
		return errors.Wrap(err, "unable to get printer")
	}

	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	build, _, err := m.CreateBuild(org, project, env, component, pipeline, job)
	if err != nil {
		return fmt.Errorf("failed to trigger build in context project %q, env %q, component %q with pipeline %q in job %q: %w", project, env, component, pipeline, job, err)
	}

	if watch {
		if build == nil || build.ID == nil {
			return errors.New("triggered build payload is missing id")
		}

		// Human vs raw watch stream: docs/pipeline-build-watch-output.md (quick rollback: force OutputRaw).
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

		watchVerbose := strings.EqualFold(viper.GetString("verbosity"), "debug")

		buildID := fmt.Sprintf("%d", *build.ID)
		writeBuildWatchIntro(cmd.ErrOrStderr(), resolveConsoleURL(), org, project, env, component, pipeline, job, buildID)
		fmt.Fprintln(cmd.ErrOrStderr(), "Note: first Ctrl+C requests a remote build abort; second Ctrl+C exits watch immediately.")

		return buildwatch.Watch(
			context.Background(),
			m,
			org,
			project,
			env,
			component,
			pipeline,
			job,
			buildID,
			buildwatch.Options{
				Timeout:         time.Duration(timeoutSeconds) * time.Second,
				CancelOnTimeout: cancelOnTimeout,
				LogWriter:       cmd.OutOrStdout(),
				Output:          outMode,
				Verbose:         watchVerbose,
				Theme:           theme,
				StripLogANSI:    stripLogANSI,
				StatusWriter:    cmd.ErrOrStderr(),
			},
		)
	}

	return printer.SmartPrint(p, build, nil, "", printer.Options{}, cmd.OutOrStdout())
}

// resolveConsoleURL returns CY_CONSOLE_URL when set, otherwise the default (https://console.cycloid.io).
func resolveConsoleURL() string {
	return strings.TrimSpace(viper.GetString("console_url"))
}

func writeBuildWatchIntro(w io.Writer, consoleBase, org, project, env, component, pipeline, job, buildID string) {
	if link, ok := common.PipelineBuildConsoleURL(consoleBase, org, project, env, component, pipeline, job, buildID); ok {
		fmt.Fprintf(w, "Watching build %s: %s\n", buildID, link)
		return
	}
	fmt.Fprintln(w, "No console URL is available (CY_CONSOLE_URL is empty); cannot print a link to this build.")
}
