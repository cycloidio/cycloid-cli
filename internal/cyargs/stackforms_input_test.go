package cyargs_test

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"testing"

	"github.com/cycloidio/cycloid-cli/client/models"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/sanity-io/litter"
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
		output, err := cyargs.GetStackformsVars(cmd, defaults)
		if err != nil {
			t.Fatalf("stackform var parsing failed: %s", err)
		}

		gotJSON, err := json.Marshal(output)
		if err != nil {
			t.Fatalf("failed to serialize json output '%v': %s", &output, err)
		}
		assert.Equal(t, string(expectedJSON), string(gotJSON), "should be equal")
	})

	t.Run("EmptyDefaultsEmptyVars", func(t *testing.T) {
		cmd := &cobra.Command{}
		cyargs.AddStackFormsInputFlags(cmd)

		var defaults = make(models.FormVariables)
		output, err := cyargs.GetStackformsVars(cmd, defaults)
		if err != nil {
			t.Fatalf("Empty default should work: %s", err)
		}

		assert.Equal(t, defaults, output, "should be empty")
	})

	t.Run("EnvVar", func(t *testing.T) {
		cmd := &cobra.Command{}
		cyargs.AddStackFormsInputFlags(cmd)

		os.Setenv(cyargs.StackformsEnvVarName, string(expectedJSON))
		defer os.Unsetenv(cyargs.StackformsEnvVarName)
		var defaults = make(models.FormVariables)
		output, err := cyargs.GetStackformsVars(cmd, defaults)
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
		output, err := cyargs.GetStackformsVars(cmd, defaults)
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
		output, err := cyargs.GetStackformsVars(cmd, defaults)
		if err != nil {
			t.Fatalf("Empty default should work: %s", err)
		}

		gotJSON, err := json.Marshal(output)
		if err != nil {
			t.Fatalf("failed to serialize json output '%v': %s", &output, err)
		}
		assert.Equal(t, string(expectedJSON), string(gotJSON), "should be equal")
	})

	t.Run("VarFlag", func(t *testing.T) {
		cmd := &cobra.Command{}
		cyargs.AddStackFormsInputFlags(cmd)

		cmd.ParseFlags([]string{
			"--var", `types.string.double_quote="1"`,
			"--var", `types.string.single_quote='1'`,
			"--var", `types.float.no_quote=1`,
			"--var", `types.boolean.no_quote=true`,
			"--var", `types.array.no_quote=["string", "hello_there"]`,
			"--var", `types.map.no_quote={"string": "hello_there"}`,
		})

		var defaults = make(models.FormVariables)
		parsedVars, err := cyargs.GetStackformsVars(cmd, defaults)
		if err != nil {
			t.Fatalf("Empty default should work: %s", err)
		}

		litter.Dump(parsedVars)
		value, ok := parsedVars["types"]["string"]["double_quote"].(string)
		assert.True(t, ok, "type cast to string should be okay")
		assert.Equal(t, "1", value, "the output should be a string 1 with no quotes")

		value, ok = parsedVars["types"]["string"]["single_quote"].(string)
		assert.True(t, ok, "type cast to string should be okay")
		assert.Equal(t, "1", value, "the output should be a string 1 with no quotes")

		valueFloat, ok := parsedVars["types"]["float"]["no_quote"].(float64)
		assert.True(t, ok, "type cast to float should be okay")
		assert.Equal(t, 1.0, valueFloat, "the output should be a float")

		valueBool, ok := parsedVars["types"]["boolean"]["no_quote"].(bool)
		assert.True(t, ok, "type cast to boolean should be okay")
		assert.Equal(t, true, valueBool, "the output should be a bool")

		fmt.Println(reflect.TypeOf(parsedVars["types"]["array"]["no_quote"]))
		valueArray, ok := parsedVars["types"]["array"]["no_quote"]
		assert.True(t, ok, "type cast to array should be okay")
		assert.Equal(t, []any{"string", "hello_there"}, valueArray, "the output should be a array")

		valueMap, ok := parsedVars["types"]["map"]["no_quote"]
		assert.True(t, ok, "type cast to map should be okay")
		assert.Equal(t, map[string]any{"string": "hello_there"}, valueMap, "the output should be a map")
	})
}
