package e2e

import (
	"encoding/json"
	"os"
	"strings"
	"testing"

	"github.com/cycloidio/cycloid-cli/client/models"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/stretchr/testify/assert"
)

func TestComponentCmd(t *testing.T) {
	t.Parallel()
	api := common.NewAPI(
		common.WithInsecure(true),
		common.WithURL(TestAPIURL),
		common.WithToken(TestAPIKey),
	)
	m := middleware.NewMiddleware(api)

	var (
		projectName      = "Test E2E component"
		project          = randomCanonical("test-e2e-component")
		description      = "Testing components"
		configRepository = CyTestConfigRepo
		owner            = ""
		team             = ""
		color            = "blue"
		icon             = "planet"
	)

	defer func() {
		err := m.DeleteProject(TestRootOrg, project)
		if err != nil {
			t.Fatalf("Failed to cleanup project '%s' for test '%s': %v", project, t.Name(), err)
		}
	}()

	createdProject, err := m.CreateProject(TestRootOrg, projectName, project, description, configRepository, owner, team, color, icon)
	if err != nil {
		t.Fatalf("Failed to create pre-requisite project '%s' for test '%s': %v", project, t.Name(), err)
	}

	var (
		env      = "test"
		envName  = "Test"
		envColor = "red"
	)

	defer func() {
		err := m.DeleteEnv(TestRootOrg, project, env)
		if err != nil {
			t.Fatalf("Failed to delete env '%s' for test '%s': %v", env, t.Name(), err)
		}
	}()

	_, err = m.CreateEnv(TestRootOrg, *createdProject.Canonical, env, envName, envColor)
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

	t.Run("CreateReadListDelete", func(t *testing.T) {
		// create
		testJSON := `{"types": "tests": "map": {"hello": "world", "int": 1, "bool": true}}`
		testJSONFileContent := `{"types": "tests": "string": "myString"}`
		testJSONStdin := `{"types": "tests": "array": ["hello", false, 1, 1.1]}`
		filename, err := WriteTempFile(testJSONFileContent)
		if err != nil {
			t.Fatalf("comp create setup failed: %v", err)
		}

		args := []string{
			"--output", "json",
			"--org", TestRootOrg,
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

		// delete
		defer t.Run("DeleteCreateComp", func(t *testing.T) {
			out, err := executeCommand([]string{
				"--org", TestRootOrg,
				"components", "delete",
				"-p", project,
				"-e", env,
				"-c", component,
			})
			if err != nil {
				t.Fatalf("failed to delete and cleanup component '%s' from Create test: %v\nstdout: %s", component, err, out)
			}
		})

		// get
		args = []string{
			"--output", "json",
			"--org", TestRootOrg,
			"components", "get",
			"-p", project,
			"-e", env,
			"-c", component,
		}
		out, err := executeCommand(args)
		if err != nil {
			t.Fatalf("failed to execute cmd '%s': %s", strings.Join(args, " "), err)
		}

		var comp models.Component
		err = json.Unmarshal([]byte(out), &comp)
		if err != nil {
			t.Fatalf("failed to parse output of cy comp get command: %v\noutput: %s", err, out)
		}

		assert.Equal(t, component, *comp.Canonical)
		assert.Equal(t, componentName, *comp.Name)

		// list
		args = []string{
			"--output", "json",
			"--org", TestRootOrg,
			"components", "list",
			"-p", project,
			"-e", env,
		}
		out, err = executeCommand(args)
		if err != nil {
			t.Fatalf("failed to execute cmd '%s': %s", strings.Join(args, " "), err)
		}

		var comps []models.Component
		err = json.Unmarshal([]byte(out), &comps)
		if err != nil {
			t.Fatalf("failed to parse output of cy comp get command: %v\noutput: %s", err, out)
		}

		assert.Equal(t, []models.Component{comp}, comps)
	})

	t.Run("CreateWithUpdateAndConfig", func(t *testing.T) {
		// create
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
			"--org", TestRootOrg,
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
				"--org", TestRootOrg,
				"components", "delete",
				"-p", project,
				"-e", env,
				"-c", component + "-update",
			})
			if err != nil {
				t.Fatalf("failed to delete and cleanup component '%s' from '%s' test: %v\nstdout: %s", component, t.Name(), err, out)
			}
		})

		// create with update
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
			"--org", TestRootOrg,
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
			t.Fatalf("component put failed: %v\nstdout:\n%s\nstderr\n%s", err, stdout, stderr)
		}

		// update
		testJSON = `{"types": "tests": "map": {"update": true}}`
		testJSONFileContent = `{"types": "tests": "string": "update"}`
		testJSONStdin = `{"types": "tests": "array": ["update"]}`
		filename, err = WriteTempFile(testJSONFileContent)
		if err != nil {
			t.Fatalf("comp update setup failed: %v", err)
		}
		defer os.Remove(filename)

		args = []string{
			"--output", "json",
			"--org", TestRootOrg,
			"components", "update",
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
			"-V", `section with spaces.group with spaces.no_spaces="update"`,
			"-u", "default",
			"--update",
		}

		stdout, stderr, err = executeCommandStdin(testJSONStdin, args)
		if err != nil {
			t.Fatalf("component update failed: %v\nstdout:\n%s\nstderr\n%s", err, stdout, stderr)
		}

		// config get
		args = []string{
			// By default, output should be in json, we test that
			// "--output", "json",
			"--org", TestRootOrg,
			"components", "update",
			"--name", componentName,
			"-p", project,
			"-e", env,
			"-c", component + "-update",
		}
		out, err := executeCommand(args)
		if err != nil {
			t.Fatalf("failed to get config from component '%s': %v", component+"-update", err)
		}

		var outVars models.FormVariables
		err = json.Unmarshal([]byte(out), &outVars)
		if err != nil {
			t.Fatalf("failed to parse '%s' CLI output as JSON vars:\noutput:\n%s\nerr:\n%s", strings.Join(args, " "), out, err)
		}

		assert.Equal(t, true, outVars["types"]["tests"]["map"].(map[string]bool)["update"])
		assert.Equal(t, "update", outVars["types"]["tests"]["string"].(string))
		assert.Equal(t, []string{"update"}, outVars["types"]["tests"]["string"].([]string))
	})
}
