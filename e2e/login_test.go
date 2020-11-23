//+build e2e

package e2e

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/login"
)

func TestLogin(t *testing.T) {
	t.Run("SuccessUserLogin", func(t *testing.T) {

		buf := new(bytes.Buffer)

		cmd := login.NewCommands()
		cmd.SetOut(buf)
		cmd.SetArgs([]string{
			"--email", CY_TEST_EMAIL,
			"--password", CY_TEST_PASSWORD,
		})

		cmd.Execute()

		out, err := ioutil.ReadAll(buf)
		require.Nil(t, err)

		assert.Equal(t, out, []byte(""))
	})
	t.Run("SuccessOrgLogin", func(t *testing.T) {

		buf := new(bytes.Buffer)

		cmd := login.NewCommands()
		cmd.SetOut(buf)
		cmd.SetArgs([]string{
			"--org", CY_TEST_ORG,
			"--email", CY_TEST_EMAIL,
			"--password", CY_TEST_PASSWORD,
		})

		cmd.Execute()

		out, err := ioutil.ReadAll(buf)
		require.Nil(t, err)

		assert.Equal(t, out, []byte(""))
	})
	t.Run("ErrorMissingRequiredFlag", func(t *testing.T) {

		buf := new(bytes.Buffer)

		cmd := login.NewCommands()
		cmd.SetOut(buf)
		cmd.SetArgs([]string{
			"--org", CY_TEST_ORG,
			"--password", CY_TEST_PASSWORD,
		})

		cmd.Execute()

		out, err := ioutil.ReadAll(buf)
		require.Nil(t, err)

		assert.Contains(t, string(out), string([]byte(`required flag(s) "email" not set`)))
	})
}
