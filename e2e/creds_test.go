package e2e_test

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/cycloidio/cycloid-cli/client/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreds(t *testing.T) {
	t.Run("SuccessCredsList", func(t *testing.T) {
		cmdOut, cmdErr := executeCommand([]string{
			"--output", "json",
			"--org", config.Org,
			"creds",
			"list",
		})

		assert.Nil(t, cmdErr)
		require.Contains(t, cmdOut, "canonical\": \"vault")
	})

	t.Run("SuccessCredsGet", func(t *testing.T) {
		cmdOut, cmdErr := executeCommand([]string{
			"--output", "json",
			"--org", config.Org,
			"creds",
			"get",
			"--canonical", "vault",
		})

		assert.Nil(t, cmdErr)
		require.Contains(t, cmdOut, "canonical\": \"vault")
	})

	t.Run("SuccessCredsCreateCustom", func(t *testing.T) {
		cmdOut, cmdErr := executeCommand([]string{
			"--output", "json",
			"--org", config.Org,
			"creds",
			"create",
			"custom",
			"--name", "cli-custom",
			"--canonical", "cli-custom-canonical",
			"--path", "cli-custom-path",
			"--field", "foo=bar",
			"--field", "int=1",
		})

		defer t.Run("SuccessDelete", func(t *testing.T) {
			_, err := executeCommand([]string{
				"--output", "json",
				"--org", config.Org,
				"creds",
				"delete",
				"--canonical", "cli-custom-canonical",
			})
			if err != nil {
				t.Fatalf("failed to delete cred cli-custom: %s", err.Error())
			}
		})

		assert.Nil(t, cmdErr)
		var outCred *models.Credential
		err := json.Unmarshal([]byte(cmdOut), &outCred)
		if err != nil {
			t.Fatalf("should be able to marshal cli output to a credential, cmdOut: %s\ncmdErr: %s\nerr: %s", cmdOut, cmdErr, err.Error())
		}

		// Test that the flag works
		assert.NotNil(t, outCred, "cli output should contain a credential:", cmdOut, cmdErr)
		assert.NotNil(t, outCred.Name, "cli output should contain a name:", cmdOut, cmdErr, outCred)
		assert.NotNil(t, outCred.Canonical, "cli output should contain a canonical:", cmdOut, cmdErr, outCred)
		assert.NotNil(t, outCred.Path, "cli output should contain a path:", cmdOut, cmdErr, outCred)
		assert.Equal(t, "cli-custom", *outCred.Name)
		assert.Equal(t, "cli-custom-canonical", *outCred.Canonical)
		assert.Equal(t, "cli-custom-path", *outCred.Path)

		t.Run("Update", func(t *testing.T) {
			cmdOut, cmdErr := executeCommand([]string{
				"--output", "json",
				"--org", config.Org,
				"creds",
				"update",
				"custom",
				"--canonical", "cli-custom-canonical",
				"--name", "cli-custom",
				"--field", "foo=bar",
				"--field", "int=1",
				"--field", "new=field",
			})

			assert.Nil(t, cmdErr)
			var outCred *models.Credential
			err := json.Unmarshal([]byte(cmdOut), &outCred)
			if err != nil {
				t.Fatalf("should be able to marshal cli output to a credential, cmdOut: %s\ncmdErr: %s\nerr: %s", cmdOut, cmdErr, err.Error())
			}

			cmdOut, cmdErr = executeCommand([]string{
				"--output", "json",
				"--org", config.Org,
				"creds",
				"get",
				"--canonical", "cli-custom-canonical",
			})
			assert.Nil(t, cmdErr)
			require.Contains(t, cmdOut, "new\": \"field")
		})

		t.Run("SuccessCredsCreateCustomWithFile", func(t *testing.T) {
			fileContent := "hello world"
			fileName, err := WriteTempFile(fileContent)
			if err != nil {
				t.Fatalf("failed to setup test, temp file write failed: %s", err.Error())
			}
			defer os.Remove(fileName)

			cmdOut, cmdErr := executeCommand([]string{
				"--output", "json",
				"--org", config.Org,
				"creds",
				"create",
				"custom",
				"--name", "cli-custom-file",
				"--field", "foo=bar",
				"--field-file", "key=" + fileName,
			})
			assert.Nil(t, cmdErr)

			defer t.Run("SuccessDelete", func(t *testing.T) {
				_, err := executeCommand([]string{
					"--output", "json",
					"--org", config.Org,
					"creds",
					"delete",
					"--canonical", "cli-custom-file",
				})
				if err != nil {
					t.Fatalf("failed to delete cred 'cli-custom-file'")
				}
			})

			var outCred *models.Credential
			err = json.Unmarshal([]byte(cmdOut), &outCred)
			if err != nil {
				t.Fatalf("should be able to marshal cli output to a credential, cmdOut: %s\ncmdErr: %s\nerr: %s", cmdOut, cmdErr, err.Error())
			}
		})

		t.Run("SuccessCredsCreateSSH", func(t *testing.T) {
			defer t.Run("SuccessDelete", func(t *testing.T) {
				_, err := executeCommand([]string{
					"--output", "json",
					"--org", config.Org,
					"creds",
					"delete",
					"--canonical", "cli-ssh",
				})
				if err != nil {
					t.Fatalf("failed to delete cred 'cli-ssh'")
				}
			})

			fileName, err := WriteTempFile(string(TestGitSSHKey))
			if err != nil {
				t.Fatalf("failed to setup test, temp file write failed: %s", err.Error())
			}
			defer os.Remove(fileName)
			cmdOut, cmdErr := executeCommand([]string{
				"--output", "json",
				"--org", config.Org,
				"creds",
				"create",
				"ssh",
				"--name", "cli-ssh",
				"--ssh-key", fileName,
			})

			assert.Nil(t, cmdErr)
			var outCred *models.Credential
			err = json.Unmarshal([]byte(cmdOut), &outCred)
			if err != nil {
				t.Fatalf("should be able to marshal cli output to a credential, cmdOut: %s\ncmdErr: %s\nerr: %s", cmdOut, cmdErr, err.Error())
			}
		})
	})
	t.Run("SuccessCredsCreateCustomWithFile", func(t *testing.T) {
		// Cleanup just in case
		executeCommand([]string{
			"--output", "json",
			"--org", config.Org,
			"creds",
			"delete",
			"--canonical", "cli-custom-file",
		})

		fileContent := "hello world"
		fileName, err := WriteTempFile(fileContent)
		if err != nil {
			t.Fatalf("failed to setup test, temp file write failed: %s", err.Error())
		}
		defer os.Remove(fileName)

		cmdOut, cmdErr := executeCommand([]string{
			"--output", "json",
			"--org", config.Org,
			"creds",
			"create",
			"custom",
			"--name", "cli-custom-file",
			"--field", "foo=bar",
			"--field-file", "key=" + fileName,
		})

		assert.Nil(t, cmdErr)
		require.Equal(t, "", cmdOut)
	})

	t.Run("SuccessCredsCreateSSH", func(t *testing.T) {
		// Cleanup just in case
		executeCommand([]string{
			"--output", "json",
			"--org", config.Org,
			"creds",
			"delete",
			"--canonical", "cli-ssh",
		})

		fileName, err := WriteTempFile(string(TestGitSSHKey))
		if err != nil {
			t.Fatalf("failed to setup test, temp file write failed: %s", err.Error())
		}
		defer os.Remove(fileName)
		cmdOut, cmdErr := executeCommand([]string{
			"--output", "json",
			"--org", config.Org,
			"creds",
			"create",
			"ssh",
			"--name", "cli-ssh",
			"--ssh-key", fileName,
		})

		assert.Nil(t, cmdErr)
		require.Equal(t, "", cmdOut)
	})

}
