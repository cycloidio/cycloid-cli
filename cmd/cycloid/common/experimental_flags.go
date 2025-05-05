package common

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"dario.cat/mergo"
	"github.com/cycloidio/cycloid-cli/client/models"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// Experimental helpers to add support for env var to manage cycloid context
func GetCyContextFlagSet() *pflag.FlagSet {
	flagSet := pflag.NewFlagSet("cycloid-context", pflag.ContinueOnError)
	flagSet.StringP("project", "p", "", "the project canonical, can also be set with the CY_PROJECT env var")
	viper.BindPFlag("project", flagSet.Lookup("project"))
	flagSet.StringP("env", "e", "", "the environment canonical, can also be set with the CY_ENV env var")
	viper.BindPFlag("env", flagSet.Lookup("env"))
	flagSet.StringP("component", "c", "", "the component canonical, can also be set with the CY_COMPONENT env var")
	viper.BindPFlag("component", flagSet.Lookup("component"))
	return flagSet
}

func GetCyContext(cmd *cobra.Command) (org, project, env, component string, err error) {
	org, err = GetOrg(cmd)
	if err != nil {
		return "", "", "", "", err
	}

	project, err = GetProject(cmd)
	if err != nil {
		return "", "", "", "", err
	}

	env, err = GetEnv(cmd)
	if err != nil {
		return "", "", "", "", err
	}

	component, err = GetComponent(cmd)
	if err != nil {
		return "", "", "", "", err
	}

	return org, project, env, component, nil
}

func GetProject(cmd *cobra.Command) (project string, err error) {
	project = viper.GetString("project")
	if project == "" {
		return "", errors.New("project is not set, use --project flag or CY_PROJECT env var")
	}

	return project, nil
}

func GetEnv(cmd *cobra.Command) (env string, err error) {
	env = viper.GetString("env")
	if env == "" {
		return "", errors.New("env is not set, use --env flag or CY_ENV env var")
	}

	return env, nil
}

func GetComponent(cmd *cobra.Command) (component string, err error) {
	component = viper.GetString("component")
	if component == "" {
		return "", errors.New("component is not set, use --component flag or CY_COMPONENT env var")
	}

	return component, nil
}

// Let's also make some flags to handle forms variables
func WithStackFormsFlagSet(cmd *cobra.Command) {
	cmd.Flags().StringArrayP("json-file", "f", []string{}, "path to a JSON file containing Stackform input. Can be '-' to read from stdin. This flag can be set multiple times.")
	cmd.MarkFlagFilename("json-file")
	cmd.Flags().StringArrayP("json-vars", "j", []string{}, "JSON string containing Stackform input. This flag can be set multiple times.")
	cmd.Flags().StringToStringP("var", "V", nil, `specify a StackForms variable using a section.group.key=value notation. The value will be parsed to try to validate the type. To force a string, add double quotes " to the value`)
}

// GetStackformsVarsFromFlags wrap the flag parsing and the merge of the variables
// set by the user.
func GetStackformsVarsFromFlags(cmd *cobra.Command) (*models.FormVariables, error) {
	varFiles, err := cmd.Flags().GetStringArray("json-file")
	if err != nil {
		return nil, err
	}

	varJson, err := cmd.Flags().GetStringArray("json-vars")
	if err != nil {
		return nil, err
	}

	varField, err := cmd.Flags().GetStringToString("var")
	if err != nil {
		return nil, err
	}

	output, err := MergeStackformsVars(varFiles, varJson, varField)
	if err != nil {
		return nil, err
	}

	return &output, nil
}

// MergeStackformsVars will parse and merge all variables inputs in the following order of
// precedence:
// file < jsonString < keyValueField
func MergeStackformsVars(jsonFiles, jsonStrings []string, keyValueField map[string]string) (models.FormVariables, error) {
	var output = make(models.FormVariables)
	jsonFileVars, err := MergeJsonFileVars(jsonFiles)
	if err != nil {
		return nil, err
	}

	err = mergo.Merge(&output, jsonFileVars)
	if err != nil {
		return nil, err
	}

	jsonVars, err := MergeJsonVars(jsonStrings)
	if err != nil {
		return nil, err
	}

	err = mergo.Merge(&output, jsonVars)
	if err != nil {
		return nil, err
	}

	return output, nil
}

// MergeJsonFileVars will read and merge the Stackforms vars from the `json-file` arg
func MergeJsonFileVars(jsonFiles []string) (*models.FormVariables, error) {
	var output = make(models.FormVariables)

	for _, file := range jsonFiles {
		var decoder *json.Decoder
		// Check if the user specified stdin
		if file == "-" {
			file = "stdin" // for error msg
			decoder = json.NewDecoder(os.Stdin)
		} else {
			reader, err := os.Open(file)
			if err != nil {
				return nil, fmt.Errorf("Failed to open var file at path '%s': %v", file, err)
			}
			defer reader.Close()

			decoder = json.NewDecoder(reader)
		}

		for {
			var extractedVars = make(models.FormVariables)
			err := decoder.Decode(&extractedVars)
			if err == io.EOF {
				break
			}
			if err != nil {
				return nil, fmt.Errorf("failed to read StackForms variables from '%s': %v", file, err)
			}

			if err := mergo.Merge(&output, extractedVars, mergo.WithOverride); err != nil {
				return nil, fmt.Errorf("failed to merge variables from '%s': %v", file, err)
			}
		}
	}

	return &output, nil
}

// MergeJsonVars expect an array of valid JSON string as stackforms input and return a models.FormVariables
func MergeJsonVars(jsonVars []string) (*models.FormVariables, error) {
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

	return &output, nil
}

// UpdateFormVar will take a Stackform variable ref in the format section.group.var
// and update its value. The value is passed as string but can be any valid 'JSON' type.
func UpdateFormVar(field string, value string, vars models.FormVariables) error {
	keys := strings.Split(field, ".")
	if len(keys) != 3 {
		return errors.New("key=val update failed, you can only update a value using `section.group.var=value` syntax")
	}

	if vars == nil {
		vars = make(models.FormVariables)
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

		vars[keys[0]][keys[1]][keys[2]] = data
		return nil
	}

	// We will prioritize the use of quotes to explicitly define strings values
	// This allow users to circumvent issues in case of strings that could be parsed
	// as other types
	if strings.HasPrefix(trimmedValue, "\"") && strings.HasSuffix(trimmedValue, "\"") ||
		strings.HasPrefix(trimmedValue, "'") && strings.HasSuffix(trimmedValue, "'") {
		vars[keys[0]][keys[1]][keys[2]] = trimmedValue[1 : len(trimmedValue)-1]
		return nil
	}

	// Detect standard types
	// numbers, we do all as float since JSON doesn't care
	// Important! We parse number firsts, since 1 and 0 are considered bools by strconv.ParseBool
	float, err := strconv.ParseFloat(value, 64)
	if err == nil {
		vars[keys[0]][keys[1]][keys[2]] = float
		return nil
	}

	// bools
	boolean, err := strconv.ParseBool(value)
	if err == nil {
		vars[keys[0]][keys[1]][keys[2]] = boolean
		return nil
	}

	// null
	if strings.ToLower(value) == "null" {
		vars[keys[0]][keys[1]][keys[2]] = nil
		return nil
	}

	// if all type conversion failed, consider the value as string
	vars[keys[0]][keys[1]][keys[2]] = value
	return nil
}

// End
