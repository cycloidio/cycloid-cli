package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPipelineBuildConsoleURL(t *testing.T) {
	t.Parallel()

	u, ok := PipelineBuildConsoleURL("", "o", "p", "e", "c", "pl", "j", "42")
	assert.False(t, ok)
	assert.Empty(t, u)

	u, ok = PipelineBuildConsoleURL("https://console.example.com/", "my org", "p", "e", "c", "pl", "j", "99")
	require.True(t, ok)
	assert.Equal(t, "https://console.example.com/organizations/my%20org/projects/p/environments/e/components/c/pipelines/pl/jobs/j/builds/99", u)
}
