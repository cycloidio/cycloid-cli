package e2e_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/cycloidio/cycloid-cli/client/models"
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

	t.Run("SuccessEventsList", func(t *testing.T) {
		cmdOut, cmdErr := executeCommand([]string{
			"--output", "json",
			"--org", config.Org,
			"events",
			"list",
		})

		require.Nil(t, cmdErr)
		var events []*models.Event
		err := json.Unmarshal([]byte(cmdOut), &events)
		require.Nil(t, err, "CLI output should deserialize to []*models.Event, out: %s", cmdOut)
	})
}
