package buildwatch

import (
	"context"
	"io"
	"strings"
	"time"
)

const defaultIdleTimeout = 3 * time.Second

// StreamLogs fetches the event log for a specific build.
//
// Dump mode (options.Watch == false): opens the SSE stream once, formats events,
// and returns when the stream closes (completed build) or after a period of
// inactivity (in-progress build — prints whatever has been logged so far).
//
// Watch mode (options.Watch == true): delegates to Watch with ReadOnly: true —
// reconnects on disconnect, polls build status, and exits cleanly when the build
// finishes or on Ctrl+C (exit 130, no abort).
func StreamLogs(
	ctx context.Context,
	m Client,
	org, project, env, component, pipeline, job, buildID string,
	options Options,
) error {
	if options.Watch {
		return Watch(ctx, m, org, project, env, component, pipeline, job, buildID, Options{
			LogWriter:               options.LogWriter,
			PollInterval:            options.PollInterval,
			ReconnectDelay:          options.ReconnectDelay,
			Output:                  options.Output,
			Verbose:                 options.Verbose,
			Theme:                   options.Theme,
			StripLogANSI:            options.StripLogANSI,
			StatusWriter:            options.StatusWriter,
			DisableInterruptHandler: options.DisableInterruptHandler,
			ReadOnly:                true,
		})
	}

	// Dump mode: open SSE stream once and consume events until the stream closes
	// or idles out (for in-progress builds).
	stream, _, err := m.OpenBuildEventsStream(ctx, org, project, env, component, pipeline, buildID, "")
	if err != nil {
		return err
	}
	wrapped := newIdleCloseRC(stream, defaultIdleTimeout)
	defer wrapped.Close()

	writer := options.LogWriter
	if writer == nil {
		writer = io.Discard
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

	_, err = consumeBuildEvents(ctx, wrapped, writer, "", evtOpts)
	// Ignore errors caused by the idle-timeout close (closed reader / net error).
	// Those are expected and mean we've caught up with the current build state.
	if err != nil && !isClosedErr(err) {
		return err
	}
	return nil
}

// idleCloseRC wraps an io.ReadCloser and closes it automatically after a period
// of inactivity. Each successful Read call resets the idle timer.
// This is used in dump mode to stop reading from in-progress SSE streams.
type idleCloseRC struct {
	rc      io.ReadCloser
	timeout time.Duration
	timer   *time.Timer
}

func newIdleCloseRC(rc io.ReadCloser, timeout time.Duration) *idleCloseRC {
	icr := &idleCloseRC{rc: rc, timeout: timeout}
	icr.timer = time.AfterFunc(timeout, func() { _ = rc.Close() })
	return icr
}

func (r *idleCloseRC) Read(p []byte) (int, error) {
	n, err := r.rc.Read(p)
	if n > 0 {
		r.timer.Reset(r.timeout)
	}
	return n, err
}

func (r *idleCloseRC) Close() error {
	r.timer.Stop()
	return r.rc.Close()
}

// isClosedErr reports whether the error is a result of a closed connection,
// which is the expected outcome when idleCloseRC fires its timer.
func isClosedErr(err error) bool {
	if err == nil {
		return false
	}
	s := err.Error()
	// Match common "use of closed network connection" and pipe errors.
	return strings.Contains(s, "closed") || strings.Contains(s, "EOF") || strings.Contains(s, "broken pipe")
}
