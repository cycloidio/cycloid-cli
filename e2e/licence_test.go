package e2e_test

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/matryer/is"
	"github.com/stretchr/testify/require"

	"github.com/cycloidio/cycloid-cli/client/models"
)

func TestLicence(t *testing.T) {
	licenceKey, ok := os.LookupEnv("API_LICENCE_KEY")
	require.True(t, ok, "API_LICENCE_KEY must be set for licence e2e tests")

	t.Run("SuccessLicenceActivateAndGet", func(t *testing.T) {
		is := is.New(t)

		_, cmdErr := executeCommand([]string{
			"--org", config.Org,
			"organization", "licence", "activate",
			"--key", licenceKey,
		})
		is.NoErr(cmdErr)

		cmdOut, cmdErr := executeCommand([]string{
			"--output", "json",
			"--org", config.Org,
			"organization", "licence", "get",
		})
		is.NoErr(cmdErr)

		var out models.Licence
		err := json.Unmarshal([]byte(cmdOut), &out)
		is.NoErr(err)
		require.NotNil(t, out.Key)
		is.Equal(licenceKey, *out.Key)
	})

	t.Run("SuccessLicenceActivateFromFile", func(t *testing.T) {
		is := is.New(t)

		dir := t.TempDir()
		licenceFile := filepath.Join(dir, "licence.jwt")
		err := os.WriteFile(licenceFile, []byte(licenceKey), 0o600)
		require.NoError(t, err)

		_, cmdErr := executeCommand([]string{
			"--org", config.Org,
			"organization", "licence", "activate",
			"--key-file", licenceFile,
		})
		is.NoErr(cmdErr)
	})

	t.Run("SuccessLicenceActivateFromStdin", func(t *testing.T) {
		is := is.New(t)

		_, _, err := executeCommandStdin(licenceKey, []string{
			"--org", config.Org,
			"organization", "licence", "activate",
		})
		is.NoErr(err)
	})

	t.Run("SuccessLicenceActivateOverwrite", func(t *testing.T) {
		is := is.New(t)

		_, cmdErr := executeCommand([]string{
			"--org", config.Org,
			"organization", "licence", "activate",
			"--key", licenceKey,
		})
		is.NoErr(cmdErr)

		_, cmdErr = executeCommand([]string{
			"--org", config.Org,
			"organization", "licence", "activate",
			"--key", licenceKey,
		})
		is.NoErr(cmdErr)

		cmdOut, cmdErr := executeCommand([]string{
			"--output", "json",
			"--org", config.Org,
			"organization", "licence", "get",
		})
		is.NoErr(cmdErr)
		is.True(strings.Contains(cmdOut, licenceKey))
	})
}
