package cyargs_test

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/cycloidio/cycloid-cli/client/models"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func TestGetStackformsVars(t *testing.T) {
	// vars
	anyArray := []any{"string", 1, 1.1, true, []string{"hello"}, map[string]string{"hello": "world"}}
	anyMap := map[string]any{
		"string": "string",
		"int":    1,
		"float":  1.1,
		"bool":   true,
		"array":  anyArray,
		"map": map[string]any{
			"hello": "world",
		},
	}
	expected := models.FormVariables{
		"section": {
			"group": {
				"string": "string",
				"int":    1,
				"float":  1.1,
				"bool":   true,
				"array":  anyArray,
				"map":    anyMap,
			},
		},
	}
	expectedJSON, err := json.Marshal(expected)
	if err != nil {
		t.Fatalf("failed setup for test, failed to encode expected value '%v' to json: %s", expected, err)
	}

	t.Run("VarsStdin", func(t *testing.T) {
		cmd := &cobra.Command{}
		cyargs.AddStackFormsInputFlags(cmd)

		tempFile, err := os.CreateTemp("", "json-test")
		if err != nil {
			t.Fatalf("failed to setup test, cannot create temp file: %s", err)
		}
		defer os.Remove(tempFile.Name())

		_, err = tempFile.Write(expectedJSON)
		if err != nil {
			t.Fatalf("failed to setup test, cannot write to test file %s: %s", tempFile.Name(), err)
		}

		_, err = tempFile.Seek(0, 0)
		if err != nil {
			t.Fatalf("failed to setup test, cannot change offset in file reader: %s", err)
		}

		oldStdin := os.Stdin
		defer func() { os.Stdin = oldStdin }()

		os.Stdin = tempFile

		cmd.ParseFlags([]string{
			"--json-file", "-",
		})

		var defaults = make(models.FormVariables)
		output, err := cyargs.GetStackformsVars(cmd, &defaults)
		if err != nil {
			t.Fatalf("stackform var parsing failed: %s", err)
		}

		gotJSON, err := json.Marshal(output)
		if err != nil {
			t.Fatalf("failed to serialize json output '%v': %s", &output, err)
		}
		assert.Equal(t, string(expectedJSON), string(gotJSON), "should be equal")
	})

	t.Run("NilDefaultsShouldFail", func(t *testing.T) {
		cmd := &cobra.Command{}
		cyargs.AddStackFormsInputFlags(cmd)

		_, err := cyargs.GetStackformsVars(cmd, nil)
		if err == nil {
			t.Fatal("Using cy_args.GetStackformsVars should fail with nil defaults")
		}
	})

	t.Run("EmptyDefaultsEmptyVars", func(t *testing.T) {
		cmd := &cobra.Command{}
		cyargs.AddStackFormsInputFlags(cmd)

		var defaults = make(models.FormVariables)
		output, err := cyargs.GetStackformsVars(cmd, &defaults)
		if err != nil {
			t.Fatalf("Empty default should work: %s", err)
		}

		assert.Equal(t, defaults, *output, "should be empty")
	})

	t.Run("EnvVar", func(t *testing.T) {
		cmd := &cobra.Command{}
		cyargs.AddStackFormsInputFlags(cmd)

		os.Setenv(cyargs.StackformsEnvVarName, string(expectedJSON))
		defer os.Unsetenv(cyargs.StackformsEnvVarName)
		var defaults = make(models.FormVariables)
		output, err := cyargs.GetStackformsVars(cmd, &defaults)
		if err != nil {
			t.Fatalf("Empty default should work: %s", err)
		}

		gotJSON, err := json.Marshal(output)
		if err != nil {
			t.Fatalf("failed to serialize json output '%v': %s", &output, err)
		}
		assert.Equal(t, expectedJSON, gotJSON, "should be equal")
	})

	t.Run("VarsJSON", func(t *testing.T) {
		cmd := &cobra.Command{}
		cyargs.AddStackFormsInputFlags(cmd)

		cmd.ParseFlags([]string{
			"--json-vars", string(expectedJSON),
		})

		var defaults = make(models.FormVariables)
		output, err := cyargs.GetStackformsVars(cmd, &defaults)
		if err != nil {
			t.Fatalf("Empty default should work: %s", err)
		}

		gotJSON, err := json.Marshal(output)
		if err != nil {
			t.Fatalf("failed to serialize json output '%v': %s", &output, err)
		}
		assert.Equal(t, string(expectedJSON), string(gotJSON), "should be equal")
	})

	t.Run("VarsFile", func(t *testing.T) {
		cmd := &cobra.Command{}
		cyargs.AddStackFormsInputFlags(cmd)

		tempFile, err := os.CreateTemp("", "json-test")
		if err != nil {
			t.Fatalf("failed to setup test, cannot create temp file: %s", err)
		}
		defer os.Remove(tempFile.Name())

		_, err = tempFile.Write(expectedJSON)
		if err != nil {
			t.Fatalf("failed to setup test, cannot write to test file %s: %s", tempFile.Name(), err)
		}

		cmd.ParseFlags([]string{
			"--json-file", tempFile.Name(),
		})

		var defaults = make(models.FormVariables)
		output, err := cyargs.GetStackformsVars(cmd, &defaults)
		if err != nil {
			t.Fatalf("Empty default should work: %s", err)
		}

		gotJSON, err := json.Marshal(output)
		if err != nil {
			t.Fatalf("failed to serialize json output '%v': %s", &output, err)
		}
		assert.Equal(t, string(expectedJSON), string(gotJSON), "should be equal")
	})

}
