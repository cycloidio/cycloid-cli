package e2e

import (
	"os"
	"testing"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
)

func TestComponentCmd(t *testing.T) {
	t.Parallel()
	api := common.NewAPI(
		common.WithInsecure(true),
		common.WithURL(CY_API_URL),
		common.WithToken(CY_TEST_API_KEY),
	)
	m := middleware.NewMiddleware(api)

	var (
		projectName      = "Test E2E component"
		project          = "test-e2e-componetn"
		description      = "Testing components"
		configRepository = CY_TEST_GIT_CR_URL
		owner            = ""
		team             = ""
		color            = "blue"
		icon             = "planet"
	)

	defer func() {
		err := m.DeleteProject(CY_TEST_ROOT_ORG, project)
		if err != nil {
			t.Fatalf("Failed to cleanup project '%s' for test '%s': %v", project, t.Name(), err)
		}
	}()

	createdProject, err := m.CreateProject(CY_TEST_ROOT_ORG, projectName, project, description, configRepository, owner, team, color, icon)
	if err != nil {
		t.Fatalf("Failed to create pre-requisite project '%s' for test '%s': %v", project, t.Name(), err)
	}

	var (
		env      = "test"
		envName  = "Test"
		envColor = "red"
	)

	defer func() {
		err := m.DeleteEnv(CY_TEST_ROOT_ORG, project, env)
		if err != nil {
			t.Fatalf("Failed to delete env '%s' for test '%s': %v", env, t.Name(), err)
		}
	}()

	_, err = m.CreateEnv(CY_TEST_ROOT_ORG, *createdProject.Canonical, env, envName, envColor)
	if err != nil {
		t.Fatalf("Failed to create env '%s': %v", env, err)
	}
	// end setup

	var (
		componentName        = "Test Component"
		component            = "test-component"
		componentDescription = "My cool component"
		stackRef             = "cycloid:stack-e2e-stackforms"
	)

	t.Run("Create", func(t *testing.T) {
		testJSON := `{"types": "tests": "map": {"hello": "world", "int": 1, "bool": true}}`
		testJSONFileContent := `{"types": "tests": "string": "myString"}`
		testJSONStdin := `{"types": "tests": "array": ["hello", false, 1, 1.1]}`
		filename, err := WriteTempFile(testJSONFileContent)
		if err != nil {
			t.Fatalf("comp create setup failed: %v", err)
		}

		args := []string{
			"--output", "json",
			"--org", CY_TEST_ROOT_ORG,
			"components", "create",
			"--name", componentName,
			"-p", project,
			"-e", env,
			"-c", component,
			"-d", componentDescription,
			// test raw var flag
			"-j", testJSON,
			// test stdin
			"-f", "-",
			// test file flag
			"-f", filename,
			// Test var=value flag
			"-V", `section with spaces.group with spaces.no_spaces="osef"`,
			"-s", stackRef,
			"-u", "default",
		}
		stdout, stderr, err := executeCommandStdin(testJSONStdin, args)
		if err != nil {
			t.Fatalf("component creation failed: %v\nstdout:\n%s\nstderr\n%s", err, stdout, stderr)
		}

		defer t.Run("DeleteCreateComp", func(t *testing.T) {
			out, err := executeCommand([]string{
				"--org", CY_TEST_ROOT_ORG,
				"components", "delete",
				"-p", project,
				"-e", env,
				"-c", component,
			})
			if err != nil {
				t.Fatalf("failed to delete and cleanup component '%s' from Create test: %v\nstdout: %s", component, err, out)
			}
		})
	})

	t.Run("CreateWithUpdate", func(t *testing.T) {
		testJSON := `{"types": "tests": "map": {"hello": "world", "int": 1, "bool": true}}`
		testJSONFileContent := `{"types": "tests": "string": "myString"}`
		testJSONStdin := `{"types": "tests": "array": ["hello", false, 1, 1.1]}`
		filename, err := WriteTempFile(testJSONFileContent)
		if err != nil {
			t.Fatalf("comp create setup failed: %v", err)
		}
		defer os.Remove(filename)

		args := []string{
			"--output", "json",
			"--org", CY_TEST_ROOT_ORG,
			"components", "create",
			"--name", componentName,
			"-p", project,
			"-e", env,
			"-c", component + "-update",
			"-d", componentDescription,
			// test raw var flag
			"-j", testJSON,
			// test stdin
			"-f", "-",
			// test file flag
			"-f", filename,
			// Test var=value flag
			"-V", `section with spaces.group with spaces.no_spaces="osef"`,
			"-s", stackRef,
			"-u", "default",
			"--update",
		}
		stdout, stderr, err := executeCommandStdin(testJSONStdin, args)
		if err != nil {
			t.Fatalf("component creation failed: %v\nstdout:\n%s\nstderr\n%s", err, stdout, stderr)
		}
		defer t.Run("DeleteCreateWithUpdateComp", func(t *testing.T) {
			out, err := executeCommand([]string{
				"--org", CY_TEST_ROOT_ORG,
				"components", "delete",
				"-p", project,
				"-e", env,
				"-c", component + "-update",
			})
			if err != nil {
				t.Fatalf("failed to delete and cleanup component '%s' from '%s' test: %v\nstdout: %s", component, t.Name(), err, out)
			}
		})

		testJSON = `{"types": "tests": "map": {"update": true}}`
		testJSONFileContent = `{"types": "tests": "string": "updated"}`
		testJSONStdin = `{"types": "tests": "array": []}`
		filename, err = WriteTempFile(testJSONFileContent)
		if err != nil {
			t.Fatalf("comp create setup failed: %v", err)
		}
		defer os.Remove(filename)

		args = []string{
			"--output", "json",
			"--org", CY_TEST_ROOT_ORG,
			"components", "create",
			"--name", componentName,
			"-p", project,
			"-e", env,
			"-c", component + "-update",
			"-d", description,
			// test raw var flag
			"-j", testJSON,
			// test stdin
			"-f", "-",
			// test file flag
			"-f", filename,
			// Test var=value flag
			"-V", `section with spaces.group with spaces.no_spaces="osef"`,
			"-s", stackRef,
			"-u", "default",
			"--update",
		}

		stdout, stderr, err = executeCommandStdin(testJSONStdin, args)
		if err != nil {
			t.Fatalf("component creation failed: %v\nstdout:\n%s\nstderr\n%s", err, stdout, stderr)
		}
	})
}
