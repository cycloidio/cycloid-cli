//go:build e2e
// +build e2e

package e2e

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEvents(t *testing.T) {
	LoginToRootOrg()

	t.Run("SuccessEventsSend", func(t *testing.T) {
		cmdOut, cmdErr := executeCommand([]string{
			"--output", "json",
			"--org", CY_TEST_ROOT_ORG,
			"events",
			"create",
			"--tag", "foo=bar",
			"--title", "CLI e2e",
			"--message", "hello",
		})

		require.Nil(t, cmdErr)
		assert.Equal(t, "", cmdOut)
	})
}
