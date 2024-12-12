package projects

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/client/models"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/internal"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/printer/factory"
)

var (
	varFiles  []string
	jsonVars  []string
	extraVars []string
)

func NewUpdateEnvCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "update-env",
		Short: "update an environment within a project using StackForms.",
		Long: `Create or update (with --update) an environment within a project using StackForms.

By default, the command will fetch the stack's default value for you to override.
You can cancel this with the --no-fetch-defaults flag

You can use the following ways to fill in the stackforms configuration (in the order of precedence):
1. --var-file (-f) flag       -> accept any valid JSON file, if the filename is "-", read from stdin (can be set multiple times)
2. CY_STACKFORMS_VARS env var  -> accept any valid JSON string (can be multiple json objects)
3. --json-vars (-j) flag      -> accept any valid JSON string (can be set multiple times)
4. --var (-V) flag            -> update a variable using a field=value syntax (e.g. -V section.group.key=value)

The output will be the generated configuration of the project.`,
		Example: `
# create 'prod' environment in 'my-project'
cy project update \
  --org my-org \
  --project my-project \
  --env prod \
  --use-case usecase-1 \
  --var-file vars.yml \
  --json-vars '{"myRaw": "vars"}' \
  --var section.group.key=value
`,
		PreRunE: internal.CheckAPIAndCLIVersion,

		RunE: updateEnv,
	}

	cmd.PersistentFlags().StringP("project", "p", "", "project name")
	cmd.MarkFlagRequired("project")
	cmd.PersistentFlags().StringP("env", "e", "", "environment name")
	cmd.MarkFlagRequired("env")
	cmd.PersistentFlags().StringP("use-case", "u", "", "the selected use case of the stack")
	cmd.MarkFlagRequired("use-case")

	cmd.Flags().StringArrayVarP(&varFiles, "var-file", "f", []string{}, "path to a JSON file containing variables, can be '-' for stdin, can be set multiple times.")
	cmd.Flags().StringArrayVarP(&jsonVars, "json-vars", "j", []string{}, "JSON string containing variables, can be set multiple times.")
	cmd.Flags().StringArrayVarP(&extraVars, "var", "V", []string{}, `update a variable using a section.group.var=value syntax - JSON values aren't supported for this flag.`)

	cmd.PersistentFlags().Bool("no-fetch-defaults", false, "disable the fetching of the stacks default values")
	return cmd
}

func updateEnv(cmd *cobra.Command, args []string) error {
	fmt.Println(args)
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

	var keyValVars = make(map[string]string)
	for _, keyVal := range extraVars {
		key, val, found := strings.Cut(keyVal, "=")
		if !found {
			return errors.New("invalid update with --var (-V) flag, format should be section.group.var=value")
		}
		keyValVars[key] = val
	}

	noFetchDefault, err := cmd.Flags().GetBool("no-fetch-defaults")
	if err != nil {
		return err
	}

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

	// Get vars via the CY_STACKFORMS_VARS env var
	envConfig, exists := os.LookupEnv("CY_STACKFORMS_VARS")
	if exists {
		internal.Debug("found config via env var", envConfig)
		var envVars = make(map[string]interface{})
		err := json.Unmarshal([]byte(envConfig), &envVars)

		// TODO: does this should error if parsing fail, of should we just put a warning ?
		if err != nil {
			return fmt.Errorf("failed to parse env var config '"+envConfig+"' as JSON: %s", err)
		}
	}

	vars, err := mergeVars(defaultValues, varFiles, append([]string{envConfig}, cliVars...), keyValVars)
	if err != nil {
		return err
	}

	// Handle output options
	output, err := cmd.Flags().GetString("output")
	if err != nil {
		return errors.Wrap(err, "unable to get output flag")
	}

	if output == "table" {
		output = "json"
	}

	p, err := factory.GetPrinter(output)
	if err != nil {
		return errors.Wrap(err, "unable to get printer")
	}

	inputs := models.FormInput{
		EnvironmentCanonical: &env,
		UseCase:              &useCase,
		Vars:                 vars,
	}

	_, err = m.UpdateEnv(
		org,
		project,
		env,
		useCase,
		"",
		"",
		"",
		&inputs,
	)

	if err != nil {
		return err
	}

	// return the config understood by the backend
	resp, err := m.GetProjectConfig(org, project, env)
	if err != nil {
		return errors.Wrap(err, "fail to get current configuration ")
	}

	data, err := json.Marshal(resp.Forms.UseCases[0])
	if err != nil {
		return errors.New("failed to marshall API response.")
	}

	var envData common.UseCase
	err = json.Unmarshal(data, &envData)
	if err != nil {
		// we didn't got correct response from backend but we can return our inputs
		return printer.SmartPrint(p, nil, err, "failed to repond env", printer.Options{}, cmd.OutOrStdout())
	}

	return printer.SmartPrint(p, common.UseCaseToFormInput(envData, false), err, "", printer.Options{}, cmd.OutOrStdout())
}
