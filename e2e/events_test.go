package e2e_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEvents(t *testing.T) {
	t.Run("SuccessEventsSend", func(t *testing.T) {
		cmdOut, cmdErr := executeCommand([]string{
			"--output", "json",
			"--org", config.Org,
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
