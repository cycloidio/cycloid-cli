package cyargs

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"

	"dario.cat/mergo"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/client/models"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
)

var (
	StackformsEnvVarName = "CY_STACKFORMS_VARS"
)

func AddStackFormsInputFlags(cmd *cobra.Command) {
	cmd.Flags().StringArrayP("json-file", "f", []string{}, "path to a JSON file containing Stackform input. Can be '-' to read from stdin. This flag can be set multiple times.")
	cmd.MarkFlagFilename("json-file", "json")
	cmd.Flags().StringArrayP("json-vars", "j", []string{}, "JSON string containing Stackform input. This flag can be set multiple times. Can also be set using "+StackformsEnvVarName+" env var")
	cmd.Flags().StringToStringP("var", "V", nil, `specify a StackForms variable using a section.group.key=value notation. The value will be parsed to try to validate the type. To force a string, add double quotes " to the value`)
}

// GetStackformsVars wrap the flag parsing and the merge of the variables
// set by the user. The caller must provide the defaults values as he only knows
// if it must be fetched from a stack or a current component.
func GetStackformsVars(cmd *cobra.Command, defaults models.FormVariables) (models.FormVariables, error) {
	if defaults == nil {
		defaults = make(models.FormVariables)
	}

	varFiles, err := cmd.Flags().GetStringArray("json-file")
	if err != nil {
		return nil, err
	}

	// We need to check for the '-' filename at this stage to be able to
	// dump stdin from cmd.InOrStdin() in a temp file.
	index := slices.Index(varFiles, "-")
	if index != -1 {
		if !common.DetectStdinInput() {
			return nil, fmt.Errorf("stdin looks empty, please fill stdin when using '-' argument")
		}

		tempFile, err := os.CreateTemp("", "cy-stdin-*")
		if err != nil {
			return nil, fmt.Errorf("failed to write temp file for stdin: %v", err)
		}
		defer func() {
			closeErr := tempFile.Close()
			rmErr := os.Remove(tempFile.Name())
			if closeErr != nil || rmErr != nil {
				log.Fatalf("failed to purge temp file with stdin content: %s: %s", closeErr, rmErr)
			}
		}()

		stdin, err := io.ReadAll(cmd.InOrStdin())
		if err != nil {
			return nil, fmt.Errorf("failed to read from stdin: %s", err)
		}

		_, err = tempFile.Write(stdin)
		if err != nil {
			return nil, fmt.Errorf("failed to write stdin to temp file: %s", err)
		}

		varFiles[index] = tempFile.Name()
	}

	varJSON, err := cmd.Flags().GetStringArray("json-vars")
	if err != nil {
		return nil, err
	}

	jsonFromEnv, ok := os.LookupEnv(StackformsEnvVarName)
	var varsFromEnv = make(models.FormVariables)
	if ok {
		decoder := json.NewDecoder(strings.NewReader(jsonFromEnv))
		err := decoder.Decode(&varsFromEnv)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON variables in '%s' env var: %s", StackformsEnvVarName, err)
		}
	}

	varField, err := cmd.Flags().GetStringToString("var")
	if err != nil {
		return nil, err
	}

	output, err := MergeStackformsVars(defaults, varsFromEnv, varFiles, varJSON, varField)
	if err != nil {
		return nil, err
	}

	if output == nil {
		return nil, errors.New("invalid user input: stackforms vars must not be empty.")
	}

	return output, nil
}

// MergeStackformsVars will parse and merge all variables inputs in the following order of
// precedence:
// file < jsonString < keyValueField
func MergeStackformsVars(defaults models.FormVariables, envVars models.FormVariables, jsonFiles, jsonStrings []string, keyValueField map[string]string) (models.FormVariables, error) {
	if defaults == nil {
		defaults = make(models.FormVariables)
	}

	err := mergo.Merge(&defaults, envVars, mergo.WithOverride)
	if err != nil {
		return nil, fmt.Errorf("failed to merge JSON from env var: %w", err)
	}

	jsonFileVars, err := MergeJSONFileVars(jsonFiles)
	if err != nil {
		return nil, err
	}

	err = mergo.Merge(&defaults, jsonFileVars, mergo.WithOverride)
	if err != nil {
		return nil, fmt.Errorf("failed to merge JSON vars files: %w", err)
	}

	jsonVars, err := MergeJSONVars(jsonStrings)
	if err != nil {
		return nil, err
	}

	err = mergo.Merge(&defaults, jsonVars, mergo.WithOverride)
	if err != nil {
		return nil, fmt.Errorf("failed to merge JSON Vars: %w", err)
	}

	for k, v := range keyValueField {
		err = UpdateFormVar(k, v, defaults)
		if err != nil {
			return nil, fmt.Errorf("failed to update vars '%s' with field '%s': %s", k, v, err)
		}
	}

	return defaults, nil
}

// MergeJSONFileVars will read and merge the Stackforms vars from the `json-file` arg
func MergeJSONFileVars(jsonFiles []string) (models.FormVariables, error) {
	var output = make(models.FormVariables)

	for _, filename := range jsonFiles {
		var decoder *json.Decoder
		reader, err := os.Open(filename)
		if err != nil {
			return nil, fmt.Errorf("failed to open var file at path '%s': %v", filename, err)
		}
		defer reader.Close()

		decoder = json.NewDecoder(reader)

		for {
			var extractedVars = make(models.FormVariables)
			err := decoder.Decode(&extractedVars)
			if err == io.EOF {
				break
			}
			if err != nil {
				return nil, fmt.Errorf("failed to read StackForms variables from '%s': %v", filename, err)
			}

			if err := mergo.Merge(&output, extractedVars, mergo.WithOverride); err != nil {
				return nil, fmt.Errorf("failed to merge variables from '%s': %v", filename, err)
			}
		}
	}

	return output, nil
}

// MergeJSONVars expect an array of valid JSON string as stackforms input and return a models.FormVariables
func MergeJSONVars(jsonVars []string) (models.FormVariables, error) {
	var output = make(models.FormVariables)
	for _, jsonString := range jsonVars {
		if jsonString == "" {
			continue
		}

		var extractedVars = make(models.FormVariables)

		err := json.Unmarshal([]byte(jsonString), &extractedVars)
		if err != nil {
			return nil, errors.Errorf("failed to parse json-var input as valid Stackform JSON:\n%s\n%v", jsonString, err)
		}

		if err := mergo.Merge(&output, extractedVars, mergo.WithOverride); err != nil {
			return nil, errors.Errorf("failed to merge input vars from json-var input:\n%v\n%v", extractedVars, err)
		}
	}

	return output, nil
}

// UpdateFormVar will take a Stackform variable ref in the format section.group.var
// and update its value. The value is passed as string but can be any valid 'JSON' type.
func UpdateFormVar(field string, value string, vars models.FormVariables) error {
	keys := strings.Split(field, ".")
	if len(keys) != 3 {
		return errors.New("key=val update failed, you can only update a value using `section.group.var=value` syntax")
	}

	var section, group, key = keys[0], keys[1], keys[2]

	if vars == nil {
		vars = models.FormVariables{section: {group: {key: nil}}}
	}

	if vars[section] == nil {
		vars[section] = map[string]map[string]any{
			group: {key: nil},
		}
	}

	if vars[section][group] == nil {
		vars[section][group] = map[string]any{
			key: nil,
		}
	}

	// Try to detect JSON first
	// we strip value for space and newline in begin/end of the string
	trimmedValue := strings.TrimSpace(value)
	// We check if we have a valid JSON array or object by looking up first and last char.
	if strings.HasPrefix(trimmedValue, "[") && strings.HasSuffix(trimmedValue, "]") || strings.HasPrefix(trimmedValue, "{") && strings.HasSuffix(trimmedValue, "}") {
		var data any
		err := json.Unmarshal([]byte(trimmedValue), &data)
		if err != nil {
			return errors.Wrapf(err, "invalid JSON value in key=val update with value '%s'", trimmedValue)
		}

		vars[section][group][key] = data
		return nil
	}

	// We will prioritize the use of quotes to explicitly define strings values
	// This allow users to circumvent issues in case of strings that could be parsed
	// as other types
	// Note: cobra seems to already trim trailling double quotes.
	if strings.HasPrefix(trimmedValue, `"`) {
		vars[section][group][key] = trimmedValue[1:]
		return nil
	} else if strings.HasPrefix(trimmedValue, "'") && strings.HasSuffix(trimmedValue, "'") {
		vars[section][group][key] = trimmedValue[1 : len(trimmedValue)-1] // single quote won't be trimmed by cobra
		return nil
	}

	// Detect standard types
	// numbers, we do all as float since JSON doesn't care
	// Important! We parse number firsts, since 1 and 0 are considered bools by strconv.ParseBool
	float, err := strconv.ParseFloat(value, 64)
	if err == nil {
		vars[section][group][key] = float
		return nil
	}

	// bools
	boolean, err := strconv.ParseBool(value)
	if err == nil {
		vars[section][group][key] = boolean
		return nil
	}

	// null
	if strings.ToLower(value) == "null" {
		vars[section][group][key] = nil
		return nil
	}

	// if all type conversion failed, consider the value as string
	vars[section][group][key] = value
	return nil
}
