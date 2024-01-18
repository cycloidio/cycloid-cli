//go:build e2e
// +build e2e

package e2e

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStacks(t *testing.T) {
	LoginToRootOrg()

	// Since the latest update the public catalog have been added by default
	// Here is a sample of code if we need to add a dedicated one
	// t.Run("InitPublicCatalog", func(t *testing.T) {
	// 	executeCommand([]string{
	// 		"--output", "json",
	// 		"--org", CY_TEST_ROOT_ORG,
	// 		"catalog-repository",
	// 		"create",
	// 		"--branch", "master",
	// 		"--url", "https://github.com/cycloid-community-catalog/stack-magento.git",
	// 		"--name", "magento",
	// 	})
	//
	// 	// Ensure the catalog is present
	// 	cmdOut, _ := executeCommand([]string{
	// 		"--output", "json",
	// 		"--org", CY_TEST_ROOT_ORG,
	// 		"catalog-repository",
	// 		"get",
	// 		"--canonical", "magento",
	// 	})
	//
	// 	require.Contains(t, cmdOut, "canonical\":\"magento")
	// })

	t.Run("SuccessStacksList", func(t *testing.T) {
		cmdOut, cmdErr := executeCommand([]string{
			"--output", "json",
			"--org", CY_TEST_ROOT_ORG,
			"stacks",
			"list",
		})

		require.Nil(t, cmdErr)
		assert.Contains(t, cmdOut, "canonical\":\"stack-dummy")
	})
	t.Run("SuccessStacksGet", func(t *testing.T) {
		cmdOut, cmdErr := executeCommand([]string{
			"--output", "json",
			"--org", CY_TEST_ROOT_ORG,
			"stacks",
			"get",
			"--ref", fmt.Sprintf("%s:stack-dummy", CY_TEST_ROOT_ORG),
		})

		require.Nil(t, cmdErr)
		assert.Contains(t, cmdOut, "canonical\":\"stack-dummy")
	})

	t.Run("SuccessStacksValidateForm", func(t *testing.T) {
		var TestGitSshKey = []byte(`---
default:
  pipeline:
    AWS:
      - name: "Default Region"
        key: aws_default_region
        type: string
        widget: dropdown
        description: "In which region you would like your project to run"
        default: "eu-west-1"
        values: ["eu-west-1", "eu-west-2", "eu-west3", "eu-south1", "eu-north1", "eu-central1"]
        required: true
`)
		WriteFile("/tmp/test_ci_form", TestGitSshKey)

		cmdOut, cmdErr := executeCommand([]string{
			"--output", "json",
			"--org", CY_TEST_ROOT_ORG,
			"stacks",
			"validate-form",
			"--forms", "/tmp/test_ci_form",
		})

		require.Nil(t, cmdErr)
		assert.Contains(t, cmdOut, "errors\":[]")
	})
}
