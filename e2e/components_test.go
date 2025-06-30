package e2e_test

import (
	"encoding/json"
	"os"
	"strings"
	"testing"

	"github.com/cycloidio/cycloid-cli/client/models"
	"github.com/stretchr/testify/assert"
)

func TestComponentCmd(t *testing.T) {
	m := config.Middleware

	var (
		projectName      = "Test E2E component"
		project          = randomCanonical("test-e2e-component")
		description      = "Testing components"
		configRepository = *config.ConfigRepo.Canonical
		owner            = ""
		team             = ""
		color            = "blue"
		icon             = "planet"
	)

	defer func() {
		err := m.DeleteProject(config.Org, project)
		if err != nil {
			t.Fatalf("Failed to cleanup project '%s' for test '%s': %v", project, t.Name(), err)
		}
	}()

	createdProject, err := m.CreateProject(config.Org, projectName, project, description, configRepository, owner, team, color, icon)
	if err != nil {
		t.Fatalf("Failed to create pre-requisite project '%s' for test '%s': %v", project, t.Name(), err)
	}

	var (
		env      = "test"
		envName  = "Test"
		envColor = "red"
	)

	defer func() {
		err := m.DeleteEnv(config.Org, project, env)
		if err != nil {
			t.Fatalf("Failed to delete env '%s' for test '%s': %v", env, t.Name(), err)
		}
	}()

	_, err = m.CreateEnv(config.Org, *createdProject.Canonical, env, envName, envColor)
	if err != nil {
		t.Fatalf("Failed to create env '%s': %v", env, err)
	}
	// end setup

	var (
		componentName        = "Test Component"
		component            = randomCanonical("e2e-component")
		componentDescription = "My cool component"
		stackRef             = "cycloid:stack-e2e-stackforms"
	)

	t.Run("CreateReadListDelete", func(t *testing.T) {
		// create
		testJSON := `{"types": {"tests": {"map": {"hello": "world", "int": 1, "bool": true}}}}`
		testJSONFileContent := `{"types": {"tests": {"string": "myString"}}}`
		testJSONStdin := `{"types": {"tests": {"array": ["hello", false, 1, 1.1]}}}`
		filename, err := WriteTempFile(testJSONFileContent)
		if err != nil {
			t.Fatalf("comp create setup failed: %v", err)
		}

		args := []string{
			"--output", "json",
			"--org", config.Org,
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
				"--org", config.Org,
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
			"--org", config.Org,
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

		assert.NotNil(t, comp)
		assert.NotNil(t, comp.Canonical)
		assert.Equal(t, component, *comp.Canonical)
		assert.Equal(t, componentName, *comp.Name)

		// list
		args = []string{
			"--output", "json",
			"--org", config.Org,
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

	t.Run("CreateWithUpdateNew", func(t *testing.T) {
		var newComp = randomCanonical("e2e-new")
		args := []string{
			"--output", "json",
			"--org", config.Org,
			"components", "create",
			"--name", componentName,
			"-p", project,
			"-e", env,
			"-c", newComp,
			"-d", componentDescription,
			"-V", `section with spaces.group with spaces.no_spaces="new"`,
			"-s", stackRef,
			"-u", "default",
			"--update",
		}
		stdout, stderr, err := executeCommandStdin("", args)
		if err != nil {
			t.Fatalf("component creation failed: %v\nstdout:\n%s\nstderr\n%s", err, stdout, stderr)
		}
		defer t.Run("DeleteCreateWithUpdateComp", func(t *testing.T) {
			out, err := executeCommand([]string{
				"--org", config.Org,
				"components", "delete",
				"-p", project,
				"-e", env,
				"-c", newComp,
			})
			if err != nil {
				t.Fatalf("failed to delete and cleanup component '%s' from '%s' test: %v\nstdout: %s", component, t.Name(), err, out)
			}
		})
	})

	t.Run("CreateWithUpdateAndConfig", func(t *testing.T) {
		// create
		testJSON := `{"types": {"tests": {"map": {"hello": "world", "int": 1, "bool": true}}}}`
		testJSONFileContent := `{"types": {"tests": {"string": "myString"}}}`
		testJSONStdin := `{"types": {"tests": {"array": ["hello", false, 1, 1.1]}}}`
		filename, err := WriteTempFile(testJSONFileContent)
		if err != nil {
			t.Fatalf("comp create setup failed: %v", err)
		}
		defer os.Remove(filename)

		args := []string{
			"--output", "json",
			"--org", config.Org,
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
			"--update",
		}
		stdout, stderr, err := executeCommandStdin(testJSONStdin, args)
		if err != nil {
			t.Fatalf("component creation failed: %v\nstdout:\n%s\nstderr\n%s", err, stdout, stderr)
		}
		defer t.Run("DeleteCreateWithUpdateComp", func(t *testing.T) {
			out, err := executeCommand([]string{
				"--org", config.Org,
				"components", "delete",
				"-p", project,
				"-e", env,
				"-c", component,
			})
			if err != nil {
				t.Fatalf("failed to delete and cleanup component '%s' from '%s' test: %v\nstdout: %s", component, t.Name(), err, out)
			}
		})

		// create with update
		testJSON = `{"types": {"tests": {"map": {"update": true}}}}`
		testJSONFileContent = `{"types": {"tests": {"string": "updated"}}}`
		testJSONStdin = `{"types": {"tests": {"array": []}}}`
		filename, err = WriteTempFile(testJSONFileContent)
		if err != nil {
			t.Fatalf("comp create setup failed: %v", err)
		}
		defer os.Remove(filename)

		args = []string{
			"--output", "json",
			"--org", config.Org,
			"components", "create",
			"--name", componentName,
			"-p", project,
			"-e", env,
			"-c", component,
			"-d", description,
			// test raw var flag
			"-j", testJSON,
			// test stdin
			"-f", "-",
			// test file flag
			"-f", filename,
			// Test var=value flag
			"-V", `section with spaces.group with spaces.no_spaces=osef`,
			"-s", stackRef,
			"-u", "default",
			"--update",
		}

		stdout, stderr, err = executeCommandStdin(testJSONStdin, args)
		if err != nil {
			t.Fatalf("component put failed: %v\nstdout:\n%s\nstderr\n%s", err, stdout, stderr)
		}

		// get config
		args = []string{
			"--output", "json",
			"components", "config", "get",
			"-p", project,
			"-e", env,
			"-c", component,
		}
		stdout, stderr, err = executeCommandStdin(testJSONStdin, args)
		if err != nil {
			t.Fatalf("component put failed: %v\nstdout:\n%s\nstderr\n%s", err, stdout, stderr)
		}

		var outVars models.FormVariables
		err = json.Unmarshal([]byte(stdout), &outVars)
		if err != nil {
			t.Fatalf("failed to parse output of CLI as JSON:\n%s\n%s", stdout, err)
		}

		value, _ := outVars["types"]["tests"]["map"].(map[string]any)["update"].(bool)
		assert.Equal(t, true, value)
		assert.Equal(t, "updated", outVars["types"]["tests"]["string"])
		assert.Equal(t, "osef", outVars["section with spaces"]["group with spaces"]["no_spaces"])

		// update
		testJSON = `{"types": {"tests": {"string": "update2"}}}`
		testJSONFileContent = `{"types": {"tests": {"integer": 14}}}`
		testJSONStdin = `{"types": {"tests": {"float": 2.2}}}`
		filename, err = WriteTempFile(testJSONFileContent)
		if err != nil {
			t.Fatalf("comp update setup failed: %v", err)
		}
		defer os.Remove(filename)

		args = []string{
			"--output", "json",
			"--org", config.Org,
			"components", "update",
			"--name", componentName,
			"-p", project,
			"-e", env,
			"-c", component,
			"-d", description,
			// test raw var flag
			"-j", testJSON,
			// test stdin
			"-f", "-",
			// test file flag
			"-f", filename,
			// Test var=value flag
			"-V", `section with spaces.group with spaces.no_spaces=update2`,
			"-u", "default",
		}

		stdout, stderr, err = executeCommandStdin(testJSONStdin, args)
		if err != nil {
			t.Fatalf("component update failed: %v\nstdout:\n%s\nstderr\n%s", err, stdout, stderr)
		}

		// check config after update
		args = []string{
			"--output", "json",
			"components", "config", "get",
			"-p", project,
			"-e", env,
			"-c", component,
		}
		stdout, stderr, err = executeCommandStdin(testJSONStdin, args)
		if err != nil {
			t.Fatalf("component put failed: %v\nstdout:\n%s\nstderr\n%s", err, stdout, stderr)
		}

		outVars = make(models.FormVariables)
		err = json.Unmarshal([]byte(stdout), &outVars)
		if err != nil {
			t.Fatalf("failed to parse output of CLI as JSON:\n%s\n%s", stdout, err)
		}

		assert.Equal(t, "update2", outVars["types"]["tests"]["string"])
		assert.Equal(t, float64(14), outVars["types"]["tests"]["integer"].(float64))
		assert.Equal(t, 2.2, outVars["types"]["tests"]["float"])
		assert.Equal(t, "update2", outVars["section with spaces"]["group with spaces"]["no_spaces"])
	})
}
