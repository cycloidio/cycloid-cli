//+build e2e

package e2e

import (
	"testing"
	"fmt"
	"regexp"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var createdKpi string

func TestKpis(t *testing.T) {
	LoginToRootOrg()

	// Prepare a running project
	t.Run("CleanupAndPrepare", func(t *testing.T) {

		// Clean kpi if exist
		cmdOut, cmdErr := executeCommand([]string{
			"--output", "json",
			"--org", CY_TEST_ROOT_ORG,
			"kpis",
			"list",
		})
		require.Nil(t, cmdErr)
		cs, err := JsonListExtractFields(cmdOut, "canonical", "canonical", "^test-.*")
		require.Nil(t, err)

		for _, c := range cs {
			executeCommand([]string{
				"--output", "json",
				"--org", CY_TEST_ROOT_ORG,
				"kpis",
				"delete",
				"--canonical", c,
			})
		}

		// Create ssh cred
		WriteFile("/tmp/test_cli-ssh", TestGitSshKey)
		executeCommand([]string{
			"--output", "json",
			"--org", CY_TEST_ROOT_ORG,
			"creds",
			"create",
			"ssh",
			"--name", "git-project-creds",
			"--ssh-key", "/tmp/test_cli-ssh",
		})

		// Create config repo
		executeCommand([]string{
			"--output", "json",
			"--org", CY_TEST_ROOT_ORG,
			"config-repo",
			"create",
			"--name", "project-config",
			"--branch", CY_TEST_GIT_CR_BRANCH,
			"--cred", "git-project-creds",
			"--url", CY_TEST_GIT_CR_URL,
		})

		// Provide service catalog public
		executeCommand([]string{
			"--output", "json",
			"--org", CY_TEST_ROOT_ORG,
			"catalog-repository",
			"create",
			"--branch", "master",
			"--url", "https://github.com/cycloid-community-catalog/stack-dummy.git",
			"--name", "dummy",
		})

		// Create project
		WriteFile("/tmp/test_cli-pp-vars", TestPipelineVariables)
		WriteFile("/tmp/test_cli-pp", TestPipelineSample)
		executeCommand([]string{
			"--output", "json",
			"--org", CY_TEST_ROOT_ORG,
			"project",
			"create",
			"--name", "kpi-test",
			"--description", "this is a test project",
			"--stack-ref", fmt.Sprintf("%s:stack-dummy",  CY_TEST_ROOT_ORG),
			"--config-repo", "project-config",
			"--env", "test",
			"--usecase", "default",
			"--vars", "/tmp/test_cli-pp-vars",
			"--pipeline", "/tmp/test_cli-pp",
			"--config", "/tmp/test_cli-pp=/snowy/test/test_cli-pp",
		})

		// Ensure the catalog is present
		cmdOut, cmdErr = executeCommand([]string{
			"--output", "json",
			"--org", CY_TEST_ROOT_ORG,
			"project",
			"get",
			"--project", "kpi-test",
		})

		assert.Nil(t, cmdErr)
		require.Contains(t, cmdOut, "canonical\":\"kpi-test")
	})

	t.Run("SuccessKpisCreate", func(t *testing.T) {
		cmdOut, cmdErr := executeCommand([]string{
			"--output", "json",
			"--org", CY_TEST_ROOT_ORG,
			"kpis",
			"create",
			"--name", "test",
			"--type", "build_history",
			"--widget", "history",
			"--project", "kpi-test",
			"--env", "test",
			"--job", "job-hello-world",
		})

		require.Nil(t, cmdErr)
		require.Contains(t, cmdOut, "canonical\":\"test-")
	})

	t.Run("SuccessKpisList", func(t *testing.T) {
		cmdOut, cmdErr := executeCommand([]string{
			"--output", "json",
			"--org", CY_TEST_ROOT_ORG,
			"kpis",
			"list",
		})

		require.Nil(t, cmdErr)
		require.Contains(t, cmdOut, "canonical\":\"test-")

		re := regexp.MustCompile(`canonical":"(test-[^"]+)"`)
		createdKpi = re.FindAllStringSubmatch(cmdOut, 1)[0][1]
	})

	t.Run("SuccessKpisDelete", func(t *testing.T) {
		cmdOut, cmdErr := executeCommand([]string{
			"--output", "json",
			"--org", CY_TEST_ROOT_ORG,
			"kpis",
			"delete",
			"--canonical", createdKpi,
		})

		require.Nil(t, cmdErr)
		assert.Equal(t, "", cmdOut)
	})
}
