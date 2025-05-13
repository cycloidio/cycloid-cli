package cy_args_test

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/cycloidio/cycloid-cli/client/models"
	"github.com/cycloidio/cycloid-cli/internal/cy_args"
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

	cmd := &cobra.Command{}
	cy_args.AddStackFormsInputFlags(cmd)

	t.Run("NilDefaultsShouldFail", func(t *testing.T) {
		_, err := cy_args.GetStackformsVars(cmd, nil)
		if err == nil {
			t.Fatal("Using cy_args.GetStackformsVars should fail with nil defaults")
		}
	})

	t.Run("EmptyDefaultsEmptyVars", func(t *testing.T) {
		var defaults = make(models.FormVariables)
		output, err := cy_args.GetStackformsVars(cmd, &defaults)
		if err != nil {
			t.Fatalf("Empty default should work: %s", err)
		}

		assert.Equal(t, defaults, *output, "should be empty")
	})

	t.Run("EnvVar", func(t *testing.T) {
		os.Setenv(cy_args.StackformsEnvVarName, string(expectedJSON))
		defer os.Unsetenv(cy_args.StackformsEnvVarName)
		var defaults = make(models.FormVariables)
		output, err := cy_args.GetStackformsVars(cmd, &defaults)
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
		cmd.ParseFlags([]string{
			"--json-vars", string(expectedJSON),
		})

		var defaults = make(models.FormVariables)
		output, err := cy_args.GetStackformsVars(cmd, &defaults)
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
		output, err := cy_args.GetStackformsVars(cmd, &defaults)
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
