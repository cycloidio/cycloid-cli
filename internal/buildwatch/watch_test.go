package buildwatch

import (
	"context"
	"errors"
	"io"
	"net/http"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/cycloidio/cycloid-cli/client/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type fakeClient struct {
	mu sync.Mutex

	statuses      []string
	getBuildCalls int
	getBuildErr   error

	abortCalls int
	abortErr   error

	streamBodies []string
	openErr      error
	openCalls    int
}

func (f *fakeClient) GetBuild(_, _, _, _, _, _, _ string) (*models.Build, *http.Response, error) {
	f.mu.Lock()
	defer f.mu.Unlock()

	if f.getBuildErr != nil {
		return nil, nil, f.getBuildErr
	}
	if len(f.statuses) == 0 {
		return nil, nil, errors.New("no status configured")
	}

	idx := f.getBuildCalls
	if idx >= len(f.statuses) {
		idx = len(f.statuses) - 1
	}
	f.getBuildCalls++

	status := f.statuses[idx]
	return &models.Build{Status: &status}, nil, nil
}

func (f *fakeClient) AbortBuild(_, _, _, _, _, _, _ string) (*http.Response, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.abortCalls++
	return nil, f.abortErr
}

func (f *fakeClient) OpenBuildEventsStream(_ context.Context, _, _, _, _, _, _, _ string) (io.ReadCloser, *http.Response, error) {
	f.mu.Lock()
	defer f.mu.Unlock()

	if f.openErr != nil {
		return nil, nil, f.openErr
	}

	idx := f.openCalls
	if idx >= len(f.streamBodies) {
		idx = len(f.streamBodies) - 1
	}
	f.openCalls++

	if len(f.streamBodies) == 0 {
		return io.NopCloser(strings.NewReader("")), nil, nil
	}

	return io.NopCloser(strings.NewReader(f.streamBodies[idx])), nil, nil
}

func TestWatchSucceeded(t *testing.T) {
	t.Parallel()

	client := &fakeClient{
		statuses:     []string{"pending", "started", "succeeded"},
		streamBodies: []string{"data: first line\n\ndata: second line\n\n"},
	}

	var out strings.Builder
	err := Watch(
		context.Background(),
		client,
		"org",
		"project",
		"env",
		"component",
		"pipeline",
		"job",
		"42",
		Options{
			LogWriter:               &out,
			PollInterval:            time.Millisecond,
			ReconnectDelay:          50 * time.Millisecond,
			Output:                  OutputRaw,
			DisableInterruptHandler: true,
		},
	)
	require.NoError(t, err)
	assert.Contains(t, out.String(), "first line")
	assert.Contains(t, out.String(), "second line")
}

func TestWatchAborted(t *testing.T) {
	t.Parallel()

	client := &fakeClient{
		statuses: []string{"aborted"},
	}

	err := Watch(
		context.Background(),
		client,
		"org",
		"project",
		"env",
		"component",
		"pipeline",
		"job",
		"42",
		Options{
			PollInterval:            time.Millisecond,
			ReconnectDelay:          50 * time.Millisecond,
			DisableInterruptHandler: true,
		},
	)
	require.Error(t, err)

	var codedErr *ExitError
	require.ErrorAs(t, err, &codedErr)
	assert.Equal(t, exitCodeAborted, codedErr.ExitCode())
}

func TestWatchFailed(t *testing.T) {
	t.Parallel()

	client := &fakeClient{
		statuses: []string{"failed"},
	}

	err := Watch(
		context.Background(),
		client,
		"org",
		"project",
		"env",
		"component",
		"pipeline",
		"job",
		"42",
		Options{
			PollInterval:            time.Millisecond,
			ReconnectDelay:          50 * time.Millisecond,
			DisableInterruptHandler: true,
		},
	)
	require.Error(t, err)

	var codedErr *ExitError
	require.ErrorAs(t, err, &codedErr)
	assert.Equal(t, exitCodeFailed, codedErr.ExitCode())
}

func TestWatchTimeoutWithCancel(t *testing.T) {
	t.Parallel()

	client := &fakeClient{
		statuses: []string{"pending"},
	}

	err := Watch(
		context.Background(),
		client,
		"org",
		"project",
		"env",
		"component",
		"pipeline",
		"job",
		"42",
		Options{
			Timeout:                 5 * time.Millisecond,
			CancelOnTimeout:         true,
			PollInterval:            time.Millisecond,
			ReconnectDelay:          50 * time.Millisecond,
			DisableInterruptHandler: true,
		},
	)
	require.Error(t, err)

	var codedErr *ExitError
	require.ErrorAs(t, err, &codedErr)
	assert.Equal(t, exitCodeTimeout, codedErr.ExitCode())
	assert.Equal(t, 1, client.abortCalls)
}

func TestConsumeBuildEventsTracksLastEventID(t *testing.T) {
	t.Parallel()

	stream := strings.NewReader("id: 10\ndata: one\n\ndata: two\n\nid: 11\ndata: three\n\n")
	var out strings.Builder

	rawOpts := consumeOpts{format: lineFormatRaw}
	lastID, err := consumeBuildEvents(context.Background(), stream, &out, "", rawOpts)
	require.NoError(t, err)
	assert.Equal(t, "11", lastID)
	assert.Equal(t, "one\ntwo\nthree\n", out.String())
}

func TestFormatHumanBuildEventSample(t *testing.T) {
	t.Parallel()

	opts := consumeOpts{
		format:       lineFormatHuman,
		verbose:      false,
		theme:        StreamTheme{},
		stripLogANSI: true,
	}

	statusStarted := `{"data":{"status":"started","time":1774994819},"event":"status","version":"1.0","event_id":"0"}`
	out, ok := formatHumanBuildEvent(statusStarted, opts)
	require.True(t, ok)
	assert.Equal(t, "[build] status: started\n", out)

	logOut := `{"data":{"time":1774994823,"origin":{"id":"69b7cd6b","source":"stdout"},"payload":"hello world\r\n"},"event":"log","version":"5.1","event_id":"16"}`
	out, ok = formatHumanBuildEvent(logOut, opts)
	require.True(t, ok)
	assert.Contains(t, out, "[stdout]")
	assert.Contains(t, out, "hello world")

	startTask := `{"data":{"time":1774994823,"origin":{"id":"69b7cd6b"}},"event":"start-task","version":"5.0","event_id":"15"}`
	out, ok = formatHumanBuildEvent(startTask, opts)
	require.True(t, ok)
	assert.Equal(t, "[task] 69b7cd6b: started\n", out)

	finishTask := `{"data":{"time":1774994823,"exit_status":0,"origin":{"id":"69b7cd6b"}},"event":"finish-task","version":"4.0","event_id":"17"}`
	out, ok = formatHumanBuildEvent(finishTask, opts)
	require.True(t, ok)
	assert.Equal(t, "[task] 69b7cd6b: finished (exit 0)\n", out)

	imageCheck := `{"data":{"time":1},"event":"image-check","version":"1.1","event_id":"2"}`
	_, ok = formatHumanBuildEvent(imageCheck, opts)
	assert.False(t, ok)

	opts.verbose = true
	out, ok = formatHumanBuildEvent(imageCheck, opts)
	require.True(t, ok)
	assert.Contains(t, out, "[res]")
	assert.Contains(t, out, "image-check")
}

func TestFormatHumanBuildEventRawPassthroughInvalidJSON(t *testing.T) {
	t.Parallel()

	opts := consumeOpts{format: lineFormatHuman, theme: StreamTheme{}}
	out, ok := formatHumanBuildEvent("not-json", opts)
	require.True(t, ok)
	assert.Equal(t, "not-json\n", out)
}

func TestFormatHumanBuildEventLogSkipsEmptyAndWhitespaceLines(t *testing.T) {
	t.Parallel()

	opts := consumeOpts{
		format:       lineFormatHuman,
		theme:        StreamTheme{},
		stripLogANSI: true,
	}

	multi := `{"data":{"origin":{"source":"stdout"},"payload":"hello\n\n\nworld\n"},"event":"log","version":"1.0"}`
	out, ok := formatHumanBuildEvent(multi, opts)
	require.True(t, ok)
	assert.Contains(t, out, "[stdout] hello")
	assert.Contains(t, out, "[stdout] world")
	assert.NotContains(t, out, "[stdout] \n")
	lineCount := strings.Count(out, "[stdout]")
	assert.Equal(t, 2, lineCount)

	emptyPayload := `{"data":{"origin":{"source":"stdout"},"payload":"\n"},"event":"log","version":"1.0"}`
	_, ok = formatHumanBuildEvent(emptyPayload, opts)
	assert.False(t, ok)

	onlySpaces := `{"data":{"origin":{"source":"stderr"},"payload":"   \n\t\n"},"event":"log","version":"1.0"}`
	_, ok = formatHumanBuildEvent(onlySpaces, opts)
	assert.False(t, ok)
}
