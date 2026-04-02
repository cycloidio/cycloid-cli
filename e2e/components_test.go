package e2e_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/cycloidio/cycloid-cli/client/models"
)

func TestComponentCmd(t *testing.T) {
	var (
		componentName        = "Test Component"
		componentDescription = "My cool component"
		stackRef             = config.Org + ":stack-e2e-stackforms"
		description          = "Testing components"
	)

	t.Run("CreateReadListDelete", func(t *testing.T) {
		component := randomCanonical("e2e-component")

		// create
		testJSON := `{"types": {"tests": {"map": {"hello": "world", "int": 1, "bool": true}}}}`
		testJSONFileContent := `{"types": {"tests": {"string": "myString"}}}`
		testJSONStdin := `{"types": {"tests": {"array": ["hello", false, 1, 1.1]}}}`
		filename, err := WriteTempFile(testJSONFileContent)
		if err != nil {
			t.Errorf("comp create setup failed: %v", err)
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
			"--stack-commit-hash", *config.CatalogRepoVersionStacks.CommitHash,
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
				t.Errorf("failed to delete and cleanup component '%s' from Create test: %v\nstdout: %s", component, err, out)
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
			t.Errorf("failed to execute cmd '%s': %s", strings.Join(args, " "), err)
		}

		var comp models.Component
		err = json.Unmarshal([]byte(out), &comp)
		if err != nil {
			t.Errorf("failed to parse output of cy comp get command: %v\noutput: %s", err, out)
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
			t.Errorf("failed to execute cmd '%s': %s", strings.Join(args, " "), err)
		}

		var comps []models.Component
		err = json.Unmarshal([]byte(out), &comps)
		if err != nil {
			t.Errorf("failed to parse output of cy comp get command: %v\noutput: %s", err, out)
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
			"--stack-commit-hash", *config.CatalogRepoVersionStacks.CommitHash,
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
			t.Errorf("component creation failed: %v\nstdout:\n%s\nstderr\n%s", err, stdout, stderr)
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
				t.Errorf("failed to delete and cleanup component '%s' from '%s' test: %v\nstdout: %s", newComp, t.Name(), err, out)
			}
		})
	})

	t.Run("CreateWithUpdateAndConfig", func(t *testing.T) {
		component := randomCanonical("e2e-component")

		// create
		testJSON := `{"types": {"tests": {"map": {"hello": "world", "int": 1, "bool": true}}}}`
		testJSONFileContent := `{"types": {"tests": {"string": "myString"}}}`
		testJSONStdin := `{"types": {"tests": {"array": ["hello", false, 1, 1.1]}}}`
		filename, err := WriteTempFile(testJSONFileContent)
		if err != nil {
			t.Errorf("comp create setup failed: %v", err)
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
			"--stack-branch", config.CatalogRepo.Branch,
		}
		stdout, stderr, err := executeCommandStdin(testJSONStdin, args)
		if err != nil {
			t.Errorf("component creation failed: %v\nstdout:\n%s\nstderr\n%s", err, stdout, stderr)
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
				t.Errorf("failed to delete and cleanup component '%s' from '%s' test: %v\nstdout: %s", component, t.Name(), err, out)
			}
		})

		// create with update
		testJSON = `{"types": {"tests": {"map": {"update": true}}}}`
		testJSONFileContent = `{"types": {"tests": {"string": "updated"}}}`
		testJSONStdin = `{"types": {"tests": {"array": []}}}`
		filename, err = WriteTempFile(testJSONFileContent)
		if err != nil {
			t.Errorf("comp create setup failed: %v", err)
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
			"--stack-commit-hash", *config.CatalogRepoVersionStacks.CommitHash,
		}

		stdout, stderr, err = executeCommandStdin(testJSONStdin, args)
		if err != nil {
			t.Errorf("component put failed: %v\nstdout:\n%s\nstderr\n%s", err, stdout, stderr)
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
			t.Errorf("component put failed: %v\nstdout:\n%s\nstderr\n%s", err, stdout, stderr)
		}

		var outVars models.FormVariables
		err = json.Unmarshal([]byte(stdout), &outVars)
		if err != nil {
			t.Errorf("failed to parse output of CLI as JSON:\n%s\n%s", stdout, err)
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
			t.Errorf("comp update setup failed: %v", err)
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
			"--stack-commit-hash", *config.CatalogRepoVersionStacks.CommitHash,
		}

		stdout, stderr, err = executeCommandStdin(testJSONStdin, args)
		if err != nil {
			t.Errorf("component update failed: %v\nstdout:\n%s\nstderr\n%s", err, stdout, stderr)
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
			t.Errorf("component put failed: %v\nstdout:\n%s\nstderr\n%s", err, stdout, stderr)
		}

		outVars = make(models.FormVariables)
		err = json.Unmarshal([]byte(stdout), &outVars)
		if err != nil {
			t.Errorf("failed to parse output of CLI as JSON:\n%s\n%s", stdout, err)
		}

		assert.Equal(t, "update2", outVars["types"]["tests"]["string"])
		assert.Equal(t, float64(14), outVars["types"]["tests"]["integer"].(float64))
		assert.Equal(t, 2.2, outVars["types"]["tests"]["float"])
		assert.Equal(t, "update2", outVars["section with spaces"]["group with spaces"]["no_spaces"])
	})

	t.Run("TestVarsInvalidSectionsAndGroup", func(t *testing.T) {
		component := randomCanonical("e2e-component")

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
			"--stack-commit-hash", *config.CatalogRepoVersionStacks.CommitHash,
			"-V", `section with spaces.thisgroupdoesnotexists.no_spaces=update2`,
			"-V", `sectiondoesnotexists.thisgroupdoesnotexists.no_spaces=true`,
		}

		cmdOut, cmdErr := executeCommand(args)
		if cmdErr != nil {
			// We just check that it doesn't panic for now
			t.Errorf("component update failed, stdout:\n%s\nstderr\n%s", cmdOut, cmdErr)
		}
	})

	t.Run("TestCreateWithUpdateOnNonConfiguredComponentOk", func(t *testing.T) {
		component := randomCanonical("created-not-configured")

		m := config.Middleware
		created, _, err := m.CreateOrUpdateComponent(config.Org, *config.Project.Canonical,
			*config.Environment.Canonical, component, description, component,
			stackRef, "", "", *config.CatalogRepoVersionStacks.CommitHash, "default", "", nil)
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
			"--stack-commit-hash", *config.CatalogRepoVersionStacks.CommitHash,
		}

		cmdOut, cmdErr := executeCommand(args)
		if cmdErr != nil {
			// We just check that it doesn't panic for now
			t.Errorf("component update failed, stdout:\n%s\nstderr\n%s", cmdOut, cmdErr)
		}
	})

	t.Run("TestUpdateOnNonConfiguredComponentOk", func(t *testing.T) {
		component := randomCanonical("created-not-configured")

		m := config.Middleware
		created, _, err := m.CreateOrUpdateComponent(config.Org, *config.Project.Canonical,
			*config.Environment.Canonical, component, description, component,
			stackRef, "", "", *config.CatalogRepoVersionStacks.CommitHash, "default", "", nil)
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
			"--stack-commit-hash", *config.CatalogRepoVersionStacks.CommitHash,
		}

		cmdOut, cmdErr := executeCommand(args)
		if cmdErr != nil {
			// We just check that it doesn't panic for now
			t.Errorf("component update failed, stdout:\n%s\nstderr\n%s", cmdOut, cmdErr)
		}
	})
}

// TestBackendComponentConcurrency is a stress test targeting the backend API.
// It ramps up concurrent component creations from 2 to 30, one level at a time.
// Each level creates N components simultaneously then deletes them before moving on.
// If no component is created within a 30-second window at a given level, the test
// fails and reports the concurrency level at which the backend stopped responding.
func TestBackendComponentConcurrency(t *testing.T) {
	const (
		maxConcurrency = 30
		timeout        = 30 * time.Second
	)

	stackRef := config.Org + ":stack-e2e-stackforms"
	m := config.Middleware

	type result struct {
		canonical string
		comp      *models.Component
		err       error
		elapsed   time.Duration
	}

	for n := 2; n <= maxConcurrency; n++ {
		t.Logf("=== level %d: starting %d concurrent component creations ===", n, n)
		levelStart := time.Now()

		results := make([]result, n)
		// firstDone is closed as soon as the first goroutine completes (success or failure),
		// allowing us to detect a full 30s stall.
		firstDone := make(chan struct{})
		var firstDoneOnce sync.Once
		var wg sync.WaitGroup
		wg.Add(n)

		for i := range n {
			go func(idx int) {
				defer wg.Done()
				canonical := randomCanonical(fmt.Sprintf("e2e-c%d-%d", n, idx))
				t.Logf("[n=%02d goroutine %02d] creating %s ...", n, idx, canonical)
				createStart := time.Now()
				comp, resp, err := m.CreateOrUpdateComponent(
					config.Org,
					*config.Project.Canonical,
					*config.Environment.Canonical,
					canonical,
					fmt.Sprintf("Concurrent component n=%d idx=%d", n, idx),
					canonical,
					stackRef,
					"", "",
					*config.CatalogRepoVersionStacks.CommitHash,
					"default",
					"",
					nil,
				)
				elapsed := time.Since(createStart)
				firstDoneOnce.Do(func() { close(firstDone) })
				if err != nil {
					status := 0
					if resp != nil {
						status = resp.StatusCode
					}
					t.Logf("[n=%02d goroutine %02d] FAILED %s in %s (HTTP %d): %v", n, idx, canonical, elapsed, status, err)
				} else {
					t.Logf("[n=%02d goroutine %02d] OK     %s in %s", n, idx, canonical, elapsed)
				}
				results[idx] = result{canonical: canonical, comp: comp, err: err, elapsed: elapsed}
			}(i)
		}

		// Wait for the first goroutine to finish, or timeout.
		select {
		case <-firstDone:
			// At least one request came back — wait for the rest normally.
			wg.Wait()
		case <-time.After(timeout):
			t.Errorf("STALL at concurrency level %d: no component created or failed within %s — backend appears deadlocked", n, timeout)
			// Let goroutines finish in background to avoid leaks, but stop the test.
			go wg.Wait()
			return
		}

		levelElapsed := time.Since(levelStart)
		var succeeded, failed int
		for _, r := range results {
			if r.err != nil {
				failed++
			} else {
				succeeded++
			}
		}
		t.Logf("=== level %d: %d/%d OK, %d/%d FAILED in %s ===", n, succeeded, n, failed, n, levelElapsed)

		// Cleanup
		for i, r := range results {
			if r.comp == nil {
				continue
			}
			t.Logf("[n=%02d cleanup] deleting %s ...", n, r.canonical)
			_, err := m.DeleteComponent(config.Org, *r.comp.Project.Canonical,
				*r.comp.Environment.Canonical, *r.comp.Canonical)
			if err != nil {
				t.Logf("[n=%02d cleanup] WARNING: failed to delete component %d (%s): %v", n, i, r.canonical, err)
			}
		}
	}

	t.Logf("Ramp-up complete: backend handled up to %d concurrent component creations", maxConcurrency)
}
