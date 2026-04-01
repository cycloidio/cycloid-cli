// Package buildwatch implements pipeline build watch: SSE event streaming and polling until the build finishes.
// Used by cmd/cycloid/pipelines (build trigger --watch). See docs/pipeline-build-watch-output.md.
package buildwatch

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"sync/atomic"
	"time"

	"github.com/cycloidio/cycloid-cli/client/models"
	"golang.org/x/term"
)

const (
	exitCodeAborted     = 2
	exitCodeFailed      = 3
	exitCodeTimeout     = 4
	exitCodeInterrupted = 130 // second Ctrl+C: stop watching (128 + SIGINT)

	defaultPollInterval = 3 * time.Second
	defaultReconnect    = time.Second
)

// OutputMode selects how SSE data lines are written when watching a build.
type OutputMode int

const (
	// OutputHuman prints parsed, filtered lines ([build], [task], logs, etc.).
	OutputHuman OutputMode = iota
	// OutputRaw writes each SSE data payload as NDJSON (for --output json).
	OutputRaw
)

// Client is the minimal middleware surface required for watch (implemented by middleware.Middleware).
type Client interface {
	GetBuild(org, project, env, component, pipeline, job, buildID string) (*models.Build, *http.Response, error)
	AbortBuild(org, project, env, component, pipeline, job, buildID string) (*http.Response, error)
	OpenBuildEventsStream(ctx context.Context, org, project, env, component, pipeline, buildID, lastEventID string) (io.ReadCloser, *http.Response, error)
}

// ExitError is returned when the build ends in aborted, failed, or timed-out states (non-zero CLI exit).
type ExitError struct {
	code    int
	message string
}

func (e *ExitError) Error() string {
	return e.message
}

// ExitCode implements CLI exit code handling in main.
func (e *ExitError) ExitCode() int {
	return e.code
}

// Options configures Watch.
type Options struct {
	Timeout         time.Duration
	CancelOnTimeout bool
	LogWriter       io.Writer
	PollInterval    time.Duration
	ReconnectDelay  time.Duration

	// OutputHuman vs OutputRaw; see docs/pipeline-build-watch-output.md.
	Output OutputMode
	// Verbose human output (all event types), e.g. when CLI verbosity is debug.
	Verbose bool
	// Theme for human output; empty for no ANSI.
	Theme StreamTheme
	// Strip ANSI inside log payloads when not a TTY.
	StripLogANSI bool
	// StatusWriter receives short interrupt hints (first/second Ctrl+C). Often cmd.ErrOrStderr().
	StatusWriter io.Writer
	// DisableInterruptHandler skips OS signal handling (tests).
	DisableInterruptHandler bool
}

// IsTerminalWriter reports whether w is an *os.File open on a terminal.
func IsTerminalWriter(w io.Writer) bool {
	f, ok := w.(*os.File)
	if !ok {
		return false
	}
	return term.IsTerminal(int(f.Fd()))
}

// Watch polls build status until completion while streaming build events to LogWriter.
func Watch(
	ctx context.Context,
	m Client,
	org, project, env, component, pipeline, job, buildID string,
	options Options,
) error {
	if options.PollInterval <= 0 {
		options.PollInterval = defaultPollInterval
	}
	if options.ReconnectDelay <= 0 {
		options.ReconnectDelay = defaultReconnect
	}
	if options.LogWriter == nil {
		options.LogWriter = io.Discard
	}

	watchCtx, cancel := context.WithCancel(ctx)
	defer cancel()

	var userStop atomic.Bool

	if !options.DisableInterruptHandler {
		sigCh := make(chan os.Signal, 8)
		signal.Notify(sigCh, os.Interrupt)
		go func() {
			defer signal.Stop(sigCh)
			interrupts := 0
			for {
				select {
				case <-watchCtx.Done():
					return
				case <-sigCh:
					interrupts++
					if interrupts == 1 {
						_, err := m.AbortBuild(org, project, env, component, pipeline, job, buildID)
						if options.StatusWriter != nil {
							if err != nil {
								fmt.Fprintf(options.StatusWriter, "warning: build abort request failed: %v\n", err)
							}
							fmt.Fprintln(options.StatusWriter, "Interrupt: press Ctrl+C again to stop watching immediately")
						}
					} else {
						userStop.Store(true)
						cancel()
						return
					}
				}
			}
		}()
	}

	lineFmt := lineFormatHuman
	if options.Output == OutputRaw {
		lineFmt = lineFormatRaw
	}
	evtOpts := consumeOpts{
		format:       lineFmt,
		verbose:      options.Verbose,
		theme:        options.Theme,
		stripLogANSI: options.StripLogANSI,
	}

	streamErrCh := make(chan error, 1)
	go func() {
		streamErrCh <- streamEventsWithReconnect(watchCtx, m, options, evtOpts, org, project, env, component, pipeline, buildID)
	}()

	waitErr := waitForCompletion(
		watchCtx,
		m,
		org, project, env, component, pipeline, job, buildID,
		options.Timeout,
		options.CancelOnTimeout,
		options.PollInterval,
	)

	cancel()

	streamErr := <-streamErrCh
	if waitErr != nil {
		if errors.Is(waitErr, context.Canceled) && userStop.Load() {
			return &ExitError{
				code:    exitCodeInterrupted,
				message: "watch interrupted",
			}
		}
		return waitErr
	}
	if streamErr != nil && !errors.Is(streamErr, context.Canceled) {
		return streamErr
	}

	return nil
}

func waitForCompletion(
	ctx context.Context,
	m Client,
	org, project, env, component, pipeline, job, buildID string,
	timeout time.Duration,
	cancelOnTimeout bool,
	pollInterval time.Duration,
) error {
	start := time.Now()

	for {
		if timeout > 0 && time.Since(start) >= timeout {
			if cancelOnTimeout {
				if _, err := m.AbortBuild(org, project, env, component, pipeline, job, buildID); err != nil {
					return fmt.Errorf("timed out waiting for build completion and failed to abort build: %w", err)
				}
			}

			return &ExitError{
				code:    exitCodeTimeout,
				message: fmt.Sprintf("build timed out after %s", timeout),
			}
		}

		build, _, err := m.GetBuild(org, project, env, component, pipeline, job, buildID)
		if err != nil {
			return fmt.Errorf("failed to get build status: %w", err)
		}
		if build == nil || build.Status == nil {
			return errors.New("build status is empty")
		}

		status := strings.ToLower(strings.TrimSpace(*build.Status))
		switch status {
		case "succeeded":
			return nil
		case "aborted":
			return &ExitError{
				code:    exitCodeAborted,
				message: "build was aborted",
			}
		case "failed", "errored":
			return &ExitError{
				code:    exitCodeFailed,
				message: "build failed",
			}
		case "pending", "started":
		default:
			return fmt.Errorf("unknown build status %q", status)
		}

		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(pollInterval):
		}
	}
}

func streamEventsWithReconnect(
	ctx context.Context,
	m Client,
	opts Options,
	evtOpts consumeOpts,
	org, project, env, component, pipeline, buildID string,
) error {
	lastEventID := ""
	reconnectDelay := opts.ReconnectDelay
	if reconnectDelay <= 0 {
		reconnectDelay = defaultReconnect
	}
	writer := opts.LogWriter
	if writer == nil {
		writer = io.Discard
	}

	for {
		select {
		case <-ctx.Done():
			return nil
		default:
		}

		stream, _, err := m.OpenBuildEventsStream(ctx, org, project, env, component, pipeline, buildID, lastEventID)
		if err != nil {
			if errors.Is(ctx.Err(), context.Canceled) {
				return nil
			}
			if err := sleepWithContext(ctx, reconnectDelay); err != nil {
				return nil
			}
			continue
		}

		nextEventID, err := consumeBuildEvents(ctx, stream, writer, lastEventID, evtOpts)
		closeErr := stream.Close()
		if nextEventID != "" {
			lastEventID = nextEventID
		}

		if err != nil && !errors.Is(err, io.EOF) && !errors.Is(err, context.Canceled) {
			if sleepErr := sleepWithContext(ctx, reconnectDelay); sleepErr != nil {
				return nil
			}
			continue
		}
		if closeErr != nil && !errors.Is(closeErr, context.Canceled) {
			if sleepErr := sleepWithContext(ctx, reconnectDelay); sleepErr != nil {
				return nil
			}
			continue
		}

		if err := sleepWithContext(ctx, reconnectDelay); err != nil {
			return nil
		}
	}
}

func consumeBuildEvents(ctx context.Context, stream io.Reader, writer io.Writer, lastEventID string, evtOpts consumeOpts) (string, error) {
	scanner := bufio.NewScanner(stream)
	scanner.Buffer(make([]byte, 0, 64*1024), 1024*1024)

	eventID := ""
	dataLines := make([]string, 0, 4)
	currentLastEventID := lastEventID

	flushEvent := func() error {
		if len(dataLines) == 0 {
			return nil
		}

		payload := strings.Join(dataLines, "\n")
		if err := formatAndWriteBuildEvent(writer, payload, evtOpts); err != nil {
			return err
		}

		if eventID != "" {
			currentLastEventID = eventID
		}

		eventID = ""
		dataLines = dataLines[:0]
		return nil
	}

	for scanner.Scan() {
		select {
		case <-ctx.Done():
			return currentLastEventID, ctx.Err()
		default:
		}

		line := scanner.Text()
		if line == "" {
			if err := flushEvent(); err != nil {
				return currentLastEventID, err
			}
			continue
		}

		if strings.HasPrefix(line, ":") {
			continue
		}
		if strings.HasPrefix(line, "id:") {
			eventID = strings.TrimSpace(strings.TrimPrefix(line, "id:"))
			continue
		}
		if strings.HasPrefix(line, "data:") {
			data := strings.TrimPrefix(strings.TrimPrefix(line, "data:"), " ")
			dataLines = append(dataLines, data)
		}
	}

	if err := scanner.Err(); err != nil {
		return currentLastEventID, err
	}
	if err := flushEvent(); err != nil {
		return currentLastEventID, err
	}

	return currentLastEventID, nil
}

func sleepWithContext(ctx context.Context, delay time.Duration) error {
	timer := time.NewTimer(delay)
	defer timer.Stop()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-timer.C:
		return nil
	}
}
