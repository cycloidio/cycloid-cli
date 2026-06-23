package pipelines

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/cycloidio/cycloid-cli/cmd/apiclient"
	"github.com/cycloidio/cycloid-cli/cmd/common"
	"github.com/cycloidio/cycloid-cli/internal/buildwatch"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
)

const defaultJobLogsPollInterval = 5 * time.Second

func NewJobLogsCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "logs",
		Aliases: []string{"log"},
		Short:   "stream logs for the latest build of a pipeline job",
		Example: `
  # Print available logs for the latest build of a job and exit
  cy pp job logs --project p --env e --component c --pipeline pp --job j

  # Watch: stream latest build, then tail future builds continuously (Ctrl+C to stop)
  cy pp job logs --project p --env e --component c --pipeline pp --job j --watch

  # Control poll interval when waiting for a new build in --watch mode
  cy pp job logs --project p --env e --component c --pipeline pp --job j --watch --poll-interval 10s
`,
		RunE: streamJobLogs,
		Args: cobra.NoArgs,
	}

	cyargs.AddCyContext(cmd)
	cyargs.AddPipeline(cmd)
	cmd.MarkFlagRequired(cyargs.AddPipelineJob(cmd))
	cmd.Flags().Bool("watch", false, "after the current build ends, poll for new builds and stream them")
	cmd.Flags().Duration("poll-interval", defaultJobLogsPollInterval, "how often to check for a new build when --watch is waiting")
	return cmd
}

func streamJobLogs(cmd *cobra.Command, args []string) error {
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
		return fmt.Errorf("unable to get output flag: %w", err)
	}

	watch, err := cmd.Flags().GetBool("watch")
	if err != nil {
		return fmt.Errorf("unable to get --watch flag: %w", err)
	}

	pollInterval, err := cmd.Flags().GetDuration("poll-interval")
	if err != nil {
		return fmt.Errorf("unable to get --poll-interval flag: %w", err)
	}
	if pollInterval <= 0 {
		return fmt.Errorf("--poll-interval must be a positive duration")
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

	streamOpts := buildwatch.Options{
		LogWriter:    cmd.OutOrStdout(),
		Output:       outMode,
		Verbose:      verbose,
		Theme:        theme,
		StripLogANSI: stripLogANSI,
		StatusWriter: cmd.ErrOrStderr(),
	}

	// Find the latest build for this job (sort by ID descending; highest ID = newest).
	buildID, err := latestBuildID(m, org, project, env, component, pipeline, job)
	if err != nil {
		return err
	}
	if buildID == "" {
		// No builds yet — nothing to show.
		return nil
	}

	if !watch {
		return buildwatch.StreamLogs(
			context.Background(),
			m,
			org, project, env, component, pipeline, job, buildID,
			streamOpts,
		)
	}

	// Watch mode: stream current build, then poll for new builds.
	ctx := context.Background()
	lastBuildID := buildID
	for {
		fmt.Fprintf(cmd.ErrOrStderr(), "[job logs] streaming build %s\n", lastBuildID)
		streamOpts.Watch = true
		if err := buildwatch.StreamLogs(ctx, m, org, project, env, component, pipeline, job, lastBuildID, streamOpts); err != nil {
			// ExitError with code 130 means the user pressed Ctrl+C.
			var exitErr *buildwatch.ExitError
			if errors.As(err, &exitErr) && exitErr.ExitCode() == 130 {
				return err
			}
			// Other exit errors (failed, aborted) are informational in watch mode —
			// log and continue polling for the next build.
			fmt.Fprintf(cmd.ErrOrStderr(), "[job logs] build %s ended: %v\n", lastBuildID, err)
		}

		// Poll until a new build ID appears.
		nextBuildID, err := pollForNewBuild(ctx, m, org, project, env, component, pipeline, job, lastBuildID, pollInterval, cmd.ErrOrStderr())
		if err != nil {
			return err
		}
		lastBuildID = nextBuildID
	}
}

// latestBuildID lists builds for the job and returns the ID of the most recent
// one (highest numeric build ID). Returns "" if no builds exist.
func latestBuildID(m apiclient.Middleware, org, project, env, component, pipeline, job string) (string, error) {
	builds, _, err := m.GetBuilds(org, project, env, component, pipeline, job)
	if err != nil {
		return "", fmt.Errorf("listing builds for job %q: %w", job, err)
	}
	if len(builds) == 0 {
		return "", nil
	}

	// Sort by Concourse build ID descending (highest = newest).
	sort.Slice(builds, func(i, j int) bool {
		if builds[i].ID == nil {
			return false
		}
		if builds[j].ID == nil {
			return true
		}
		return *builds[i].ID > *builds[j].ID
	})

	if builds[0].ID == nil {
		return "", nil
	}
	return fmt.Sprintf("%d", *builds[0].ID), nil
}

// pollForNewBuild polls GetBuilds until a build ID different from lastBuildID
// appears. It returns the new build ID.
func pollForNewBuild(
	ctx context.Context,
	m apiclient.Middleware,
	org, project, env, component, pipeline, job, lastBuildID string,
	interval time.Duration,
	statusWriter interface{ Write([]byte) (int, error) },
) (string, error) {
	fmt.Fprintf(statusWriter, "[job logs] waiting for next build (polling every %s)...\n", interval)
	for {
		select {
		case <-ctx.Done():
			return "", ctx.Err()
		case <-time.After(interval):
		}

		nextID, err := latestBuildID(m, org, project, env, component, pipeline, job)
		if err != nil {
			return "", err
		}
		if nextID != "" && nextID != lastBuildID {
			return nextID, nil
		}
	}
}
