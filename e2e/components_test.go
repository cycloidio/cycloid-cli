package e2e_test

import (
	"encoding/json"
	"errors"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/cycloidio/cycloid-cli/client/models"
)

func TestComponentCmd(t *testing.T) {
	var (
		componentName        = "Test Component"
		component            = randomCanonical("e2e-component")
		componentDescription = "My cool component"
		stackRef             = config.Org + ":stack-e2e-stackforms"
		description          = "Testing components"
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

		var stdout, stderr string
		args := []string{
			"--output", "json",
			"--org", config.Org,
			"components", "create",
			"--name", componentName,
			"-p", *config.Project.Canonical,
			"-e", *config.Environment.Canonical,
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
		stdout, stderr, err = executeCommandStdin(testJSONStdin, args)
		if err != nil {
			t.Logf("component creation failed: %v\nstdout:\n%s\nstderr\n%s", err, stdout, stderr)
			t.FailNow()
		}

		// delete
		defer t.Run("DeleteCreateComp", func(t *testing.T) {
			out, err := executeCommand([]string{
				"--org", config.Org,
				"components", "delete",
				"-p", *config.Project.Canonical,
				"-e", *config.Environment.Canonical,
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
			"-p", *config.Project.Canonical,
			"-e", *config.Environment.Canonical,
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
			"-p", *config.Project.Canonical,
			"-e", *config.Environment.Canonical,
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
	})

	t.Run("CreateWithUpdateNew", func(t *testing.T) {
		var newComp = randomCanonical("e2e-new")
		args := []string{
			"--output", "json",
			"--org", config.Org,
			"components", "create",
			"--name", componentName,
			"-p", *config.Project.Canonical,
			"-e", *config.Environment.Canonical,
			"-c", newComp,
			"-d", componentDescription,
			"-V", `section with spaces.group with spaces.no_spaces="new"`,
			"-s", stackRef,
			"-u", "default",
			"--update",
		}
		var err, errList error
		var stdout, stderr string
		for range 3 {
			stdout, stderr, err = executeCommandStdin("", args)
			if err != nil {
				errList = errors.Join(errList, err)
				continue
			}
			errList = nil
			break
		}

		if errList != nil {
			t.Fatalf("component creation failed: %v\nstdout:\n%s\nstderr\n%s", err, stdout, stderr)
		}

		defer t.Run("DeleteCreateWithUpdateComp", func(t *testing.T) {
			out, err := executeCommand([]string{
				"--org", config.Org,
				"components", "delete",
				"-p", *config.Project.Canonical,
				"-e", *config.Environment.Canonical,
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
			"-p", *config.Project.Canonical,
			"-e", *config.Environment.Canonical,
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
				"-p", *config.Project.Canonical,
				"-e", *config.Environment.Canonical,
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
			"-p", *config.Project.Canonical,
			"-e", *config.Environment.Canonical,
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
			"-p", *config.Project.Canonical,
			"-e", *config.Environment.Canonical,
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
			"-p", *config.Project.Canonical,
			"-e", *config.Environment.Canonical,
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
			"-p", *config.Project.Canonical,
			"-e", *config.Environment.Canonical,
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

	t.Run("TestVarsInvalidSectionsAndGroup", func(t *testing.T) {
		args := []string{
			"--output", "json",
			"--org", config.Org,
			"components", "create",
			"--update",
			"--name", componentName,
			"-p", *config.Project.Canonical,
			"-e", *config.Environment.Canonical,
			"-c", component,
			"-d", description,
			"-s", stackRef,
			"-u", "default",
			"-V", `section with spaces.thisgroupdoesnotexists.no_spaces=update2`,
			"-V", `sectiondoesnotexists.thisgroupdoesnotexists.no_spaces=true`,
		}

		cmdOut, cmdErr := executeCommand(args)
		if cmdErr != nil {
			// We just check that it doesn't panic for now
			t.Fatalf("component update failed, stdout:\n%s\nstderr\n%s", cmdOut, cmdErr)
		}
	})

	t.Run("TestCreateWithUpdateOnNonConfiguredComponentOk", func(t *testing.T) {
		component := randomCanonical("created-not-configured")

		m := config.Middleware
		created, err := m.CreateComponent(config.Org, *config.Project.Canonical,
			*config.Environment.Canonical, component, description, &component,
			&stackRef, "")
		if err != nil {
			t.Logf("test setup failed: component creation %q reported err: %v", component, err)
			t.FailNow()
		}
		defer m.DeleteComponent(config.Org, *created.Project.Canonical,
			*created.Environment.Canonical, *created.Canonical)

		args := []string{
			"--output", "json",
			"--org", config.Org,
			"components", "create",
			"--update",
			"--name", component,
			"-p", *config.Project.Canonical,
			"-e", *config.Environment.Canonical,
			"-c", component,
			"-d", description,
			"-s", stackRef,
			"-u", "default",
		}

		cmdOut, cmdErr := executeCommand(args)
		if cmdErr != nil {
			// We just check that it doesn't panic for now
			t.Fatalf("component update failed, stdout:\n%s\nstderr\n%s", cmdOut, cmdErr)
		}
	})

	t.Run("TestUpdateOnNonConfiguredComponentOk", func(t *testing.T) {
		component := randomCanonical("created-not-configured")

		m := config.Middleware
		created, err := m.CreateComponent(config.Org, *config.Project.Canonical,
			*config.Environment.Canonical, component, description, &component,
			&stackRef, "")
		if err != nil {
			t.Logf("test setup failed: component creation %q reported err: %v", component, err)
			t.FailNow()
		}
		defer m.DeleteComponent(config.Org, *created.Project.Canonical,
			*created.Environment.Canonical, *created.Canonical)

		args := []string{
			"--output", "json",
			"--org", config.Org,
			"components", "update",
			"--name", component,
			"-p", *config.Project.Canonical,
			"-e", *config.Environment.Canonical,
			"-c", component,
			"-d", description,
			"-u", "default",
		}

		cmdOut, cmdErr := executeCommand(args)
		if cmdErr != nil {
			// We just check that it doesn't panic for now
			t.Fatalf("component update failed, stdout:\n%s\nstderr\n%s", cmdOut, cmdErr)
		}
	})
}
