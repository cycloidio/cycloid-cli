package projects

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"slices"
	"strings"

	"dario.cat/mergo"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/client/models"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/internal"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/printer/factory"
)

func NewCreateEnvCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:     "create-env",
		Aliases: []string{"create-stackforms-env", "create-raw-env"},
		Short:   "create an environment within a project using StackForms.",
		Long: `Create or update (with --update) an environment within a project using StackForms.

By default, the command will fetch the stack's default value for you to override.
You can cancel this with the --no-fetch-defaults flag

You can use the following ways to fill in the stackforms configuration (in the order of precedence):
1. --var-file (-f) flag       -> accept any valid JSON file, if the filename is "-", read from stdin (can be set multiple times)
2. CY_STACKFORMS_VAR env var  -> accept any valid JSON string (can be multiple json objects)
3. --json-vars (-j) flag      -> accept any valid JSON string (can be set multiple times)
4. --var (-V) flag            -> update a variable using a field=value syntax (e.g. -V section.group.key=value).
																for string types, add quotes like: 'section.group.string="mystring"'

The output will be the generated configuration of the project.`,
		Example: `
# create 'prod' environment in 'my-project'
cy project create-env \
  --org my-org \
  --project my-project \
  --env prod \
  --use-case usecase-1 \
  --var-file vars.yml \
  --json-vars '{"myRaw": "vars"}' \
  --var section.group.string="value"

# Update a project with some values from another environement
# using -V to override some variables.
cy project get-env-config --org my-org --project my-project --env prod \
    | cy project create-env --update \
    --project my-project --env staging --use-case aws \
    --var-file "-" \
    -V "pipeline.database.dump_version=staging"`,
		PreRunE: internal.CheckAPIAndCLIVersion,

		RunE: createEnvParseArgs,
	}

	cmd.PersistentFlags().StringP("project", "p", "", "project name")
	cmd.MarkFlagRequired("project")
	cmd.PersistentFlags().StringP("env", "e", "", "environment name")
	cmd.MarkFlagRequired("env")
	cmd.PersistentFlags().StringP("use-case", "u", "", "the selected use case of the stack")
	cmd.MarkFlagRequired("use-case")
	cmd.PersistentFlags().StringArrayP("var-file", "f", nil, "path to a JSON file containing variables, can be '-' for stdin, can be set multiple times.")
	cmd.PersistentFlags().StringArrayP("json-vars", "j", nil, "JSON string containing variables, can be set multiple times.")
	cmd.PersistentFlags().StringToStringP("var", "V", nil, `update a variable using a section.group.var=value syntax - Add quotes on the value for strings.`)
	cmd.PersistentFlags().Bool("update", false, "allow to override existing environment")
	cmd.PersistentFlags().Bool("no-fetch-defaults", false, "disable the fetching of the stacks default values")

	// TODO
	// Handle legacy createEnv, we create the flags to detect
	// env creation without stackforms and redirect user to the old command
	cmd.Flags().String("pipeline", "", "[deprecated] path to a pipeline file.")
	cmd.Flags().MarkHidden("pipeline")
	cmd.Flags().String("vars", "", "[deprecated] path to a pipeline config file.")
	cmd.Flags().MarkHidden("vars")
	cmd.Flags().StringToString("config", nil, "[deprecated] config key=val for legacy stacks")
	cmd.Flags().MarkHidden("config")

	return cmd
}

func createEnvParseArgs(cmd *cobra.Command, args []string) error {
	org, err := common.GetOrg(cmd)
	if err != nil {
		return err
	}

	project, err := cmd.Flags().GetString("project")
	if err != nil {
		return err
	}

	if len(project) < 2 {
		return errors.New("project must be at least 2 characters long")
	}

	env, err := cmd.Flags().GetString("env")
	if err != nil {
		return err
	}

	if len(env) < 2 {
		return errors.New("env must be at least 2 characters long")
	}

	useCase, err := cmd.Flags().GetString("use-case")
	if err != nil {
		return err
	}

	if useCase == "" {
		return errors.New("use-case is empty, please specify an use-case with --use-case")
	}

	update, err := cmd.Flags().GetBool("update")
	if err != nil {
		return err
	}

	varsFiles, err := cmd.Flags().GetStringArray("var-file")
	if err != nil {
		return err
	}

	extraVar, err := cmd.Flags().GetStringToString("var")
	if err != nil {
		return err
	}

	noFetchDefault, err := cmd.Flags().GetBool("no-fetch-defaults")
	if err != nil {
		return err
	}

	output, err := cmd.Flags().GetString("output")
	if err != nil {
		return errors.Wrap(err, "unable to get output flag")
	}

	// Handle legacy flags
	// We will detect the stacks V2 flags and use the legacy function here
	legacyPipeline, err := cmd.Flags().GetString("pipeline")
	if err != nil {
		return err
	}

	legacyVars, err := cmd.Flags().GetString("vars")
	if err != nil {
		return err
	}

	legacyConfig, err := cmd.Flags().GetStringToString("config")
	if err != nil {
		return err
	}

	if (legacyPipeline + legacyVars) != "" {
		internal.Warning(cmd.ErrOrStderr(), "You are using a legacy V2 stack and should migrate to use stackforms.")
		internal.Warning(cmd.ErrOrStderr(), "This way of creating env will be deprecated in the future")
		return createLegacyEnv(cmd, org, project, env, useCase, legacyVars, legacyPipeline, output, legacyConfig)
	}

	return createEnv(cmd, org, project, env, useCase, output, update, noFetchDefault, varsFiles, extraVar)
}

type FormVars = map[string]map[string]map[string]interface{}

// Merge variable in correct order of precedence for createEnv and updateEnv
func mergeVars(defaultValues FormVars, varsFiles []string, jsonVars []string, keyValVars map[string]string) (FormVars, error) {
	var vars = make(FormVars)

	// We merge default values first
	mergo.Merge(&vars, defaultValues, mergo.WithOverride)

	// Fetch vars from files and stdin
	for _, varFile := range varsFiles {
		var decoder *json.Decoder

		if varFile == "-" {
			decoder = json.NewDecoder(os.Stdin)
		} else if varFile != "" {
			reader, err := os.Open(varFile)
			if err != nil {
				return nil, errors.Errorf("failed to read input vars from stdin: %v", err)
			}

			defer reader.Close()
			decoder = json.NewDecoder(reader)
		}

		// Files can contain one or more object, so we scan for all with a decoder
		for {
			var extractedVars = make(FormVars)
			err := decoder.Decode(&extractedVars)
			if err == io.EOF {
				internal.Debug("finished reading input vars from", varFile)
				break
			}

			if err != nil {
				if varFile == "-" {
					varFile = "stdin"
				}

				return nil, fmt.Errorf("failed to read input vars from "+varFile+": %v", err)
			}

			if err := mergo.Merge(&vars, extractedVars, mergo.WithOverride); err != nil {
				return nil, errors.Errorf("failed to merge input vars from "+varFile+": %v", err)
			}
		}
	}

	for _, varInput := range jsonVars {
		if varInput == "" {
			continue
		}

		var extractedVars = make(FormVars)

		err := json.Unmarshal([]byte(varInput), &extractedVars)
		if err != nil {
			return nil, errors.Errorf("failed to parse json-var input '"+varInput+"' as JSON: %s", err)
		}

		if err := mergo.Merge(&vars, extractedVars, mergo.WithOverride); err != nil {
			return nil, errors.Errorf("failed to merge input vars from json-var input: %v\nerr: %v", extractedVars, err)
		}
	}

	// Merge key/val from --var
	for k, v := range keyValVars {
		err := common.UpdateMapField(k, v, vars)
		if err != nil {
			return nil, err
		}
	}

	return vars, nil
}

func createEnv(cmd *cobra.Command, org, project, env, useCase, output string, update, noFetchDefault bool, varsFiles []string, extraVar map[string]string) error {
	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	// We need the project data first to get the stack ref
	projectData, err := m.GetProject(org, project)
	if err != nil {
		return err
	}

	var defaultValues FormVars

	if !noFetchDefault {
		// First we fetch the stack's default
		stack, err := m.GetStackConfig(org, *projectData.ServiceCatalog.Ref)
		if err != nil {
			return errors.Wrap(err, "failed to retrieve stack's defaults values")
		}

		var stackConfig map[string]struct {
			Forms common.UseCase `json:"forms"`
		}

		errMsg := `failed to serialize API response for stack default value fetched with getServiceCatalogConfig.`
		stackJson, err := json.MarshalIndent(stack, "", "  ")
		if err != nil {
			return errors.Wrap(err, errMsg)
		}

		err = json.Unmarshal(stackJson, &stackConfig)
		if err != nil {
			return errors.Wrap(err, errMsg)
		}

		defaultValues = common.UseCaseToFormInput(stackConfig[useCase].Forms, true)

	}

	// Get variables via CLI arguments --json-vars
	cliVars, err := cmd.Flags().GetStringArray("json-vars")
	if err != nil {
		return err
	}

	var jsonVars []string

	envConfig, exists := os.LookupEnv("CY_STACKFORMS_VARS")
	if exists {
		jsonVars = append(jsonVars, envConfig)
	}

	jsonVars = append(jsonVars, cliVars...)
	vars, err := mergeVars(defaultValues, varsFiles, jsonVars, extraVar)
	if err != nil {
		return errors.Wrapf(err, "Failed to merge variables: %v", vars)
	}

	// fetch the printer from the factory
	if output == "table" {
		output = "json"
	}

	p, err := factory.GetPrinter(output)
	if err != nil {
		return errors.Wrap(err, "unable to get printer")
	}

	// Infer icon and color based on usecase
	var cloudProviderCanonical, icon, color string
	switch strings.ToLower(useCase) {
	case "aws":
		cloudProviderCanonical = "aws"
		icon = "mdi-aws"
		color = "staging"
	case "azure":
		cloudProviderCanonical = "azurerm"
		icon = "mdi-azure"
		color = "prod"
	case "gcp":
		cloudProviderCanonical = "google"
		icon = "mdi-google-cloud"
		color = "dev"
	default:
		cloudProviderCanonical = ""
		icon = "extension"
		color = "default"
	}

	// TODO: add the same color/icon as frontend for prod/prd staging/stg/preprod
	inputs := models.FormInput{
		EnvironmentCanonical: &env,
		UseCase:              &useCase,
		Vars:                 vars,
	}

	// TODO: Add support for resource pool canonical in case of resource quotas
	err = m.CreateEnv(
		org,
		project,
		env,
		useCase,
		cloudProviderCanonical,
		color,
		icon,
		&inputs,
	)

	// we extract the potential apiErr to check to http error code 409
	var apiErr *middleware.ApiError
	if update && errors.As(err, &apiErr) && apiErr.HttpCode == "409" {
		_, err = m.UpdateEnv(
			org,
			project,
			env,
			useCase,
			cloudProviderCanonical,
			color,
			icon,
			&inputs,
		)

		if err != nil {
			return errors.Wrapf(err, "failed to update env '%s': ", env)
		}
	}

	if err != nil {
		return printer.SmartPrint(p, nil, err, "failed to create env", printer.Options{}, cmd.OutOrStdout())
	}

	// return the config understood by the backend
	resp, err := m.GetProjectConfig(org, project, env)
	if err != nil {
		// we didn't got correct response from backend but we can return our inputs
		return printer.SmartPrint(p, inputs, err, "cannot read current config for current env in backend.", printer.Options{}, cmd.OutOrStdout())
	}

	useCaseIndex := slices.IndexFunc(resp.Forms.UseCases, func(useCase *models.FormUseCase) bool {
		if useCase.Name == nil || resp.UseCase == nil {
			return false
		}
		return *useCase.Name == *resp.UseCase
	})
	if useCaseIndex == -1 {
		return printer.SmartPrint(p, resp, errors.Errorf("Failed to find usecase '%s' for env '%s'.", *resp.UseCase, env), "", printer.Options{}, cmd.ErrOrStderr())
	}

	data, err := json.Marshal(resp.Forms.UseCases[useCaseIndex])
	if err != nil {
		// we didn't got correct response from backend but we can return our inputs
		return printer.SmartPrint(p, inputs, err, "", printer.Options{}, cmd.OutOrStdout())
	}

	var formsUseCase common.UseCase
	err = json.Unmarshal(data, &formsUseCase)
	if err != nil {
		return printer.SmartPrint(p, nil, err, "fail to get env config", printer.Options{}, cmd.OutOrStdout())
	}

	return printer.SmartPrint(p, common.UseCaseToFormInput(formsUseCase, true), nil, "", printer.Options{}, cmd.OutOrStdout())
}
