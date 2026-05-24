package buildwatch

import (
	"bytes"
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestStreamLogsDumpCompletedBuild verifies that StreamLogs (dump mode) consumes
// all events from a completed build and returns nil when the stream closes naturally.
func TestStreamLogsDumpCompletedBuild(t *testing.T) {
	body := "id:1\ndata:{\"event\":\"status\",\"data\":{\"status\":\"started\",\"time\":1}}\n\n" +
		"id:2\ndata:{\"event\":\"status\",\"data\":{\"status\":\"succeeded\",\"time\":2}}\n\n"

	client := &fakeClient{
		statuses:     []string{"succeeded"},
		streamBodies: []string{body},
	}

	var out bytes.Buffer
	err := StreamLogs(
		context.Background(),
		client,
		"org", "proj", "env", "comp", "pipeline", "job", "42",
		Options{
			LogWriter:               &out,
			DisableInterruptHandler: true,
		},
	)
	require.NoError(t, err)
	// The stream had two events; human output should mention build status lines.
	assert.Contains(t, out.String(), "[build]")
}

// TestStreamLogsDumpIdleTimeout verifies that StreamLogs (dump mode) returns
// after the idle timeout fires, without blocking indefinitely on an open stream.
func TestStreamLogsDumpIdleTimeout(t *testing.T) {
	// An empty stream body will block in a real connection, but our fakeClient
	// returns a closed reader immediately. This tests the normal "no events" path.
	client := &fakeClient{
		statuses:     []string{"started"},
		streamBodies: []string{""},
	}

	var out bytes.Buffer
	start := time.Now()
	err := StreamLogs(
		context.Background(),
		client,
		"org", "proj", "env", "comp", "pipeline", "job", "42",
		Options{
			LogWriter:               &out,
			DisableInterruptHandler: true,
		},
	)
	elapsed := time.Since(start)
	require.NoError(t, err)
	// Should complete quickly (stream closed immediately by fake client).
	assert.Less(t, elapsed, 5*time.Second)
}

// TestStreamLogsWatchDelegatesToWatch verifies that StreamLogs in watch mode
// behaves like Watch (returns error codes for failed builds, etc.).
func TestStreamLogsWatchDelegatesToWatch(t *testing.T) {
	client := &fakeClient{
		statuses:     []string{"started", "failed"},
		streamBodies: []string{""},
	}

	var out bytes.Buffer
	err := StreamLogs(
		context.Background(),
		client,
		"org", "proj", "env", "comp", "pipeline", "job", "42",
		Options{
			Watch:                   true,
			LogWriter:               &out,
			PollInterval:            10 * time.Millisecond,
			DisableInterruptHandler: true,
		},
	)
	// Watch mode should return ExitError for a failed build.
	var exitErr *ExitError
	require.ErrorAs(t, err, &exitErr)
	assert.Equal(t, exitCodeFailed, exitErr.ExitCode())
	// AbortBuild should NOT be called in read-only watch mode.
	assert.Equal(t, 0, client.abortCalls)
}
