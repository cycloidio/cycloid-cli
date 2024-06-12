package projects

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
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
		Use:   "create-stackforms-env",
		Short: "create an environment within a project using StackForms.",
		Long: `
You can provide stackforms variables via files, env var and the --vars flag
The precedence order for variable provisioning is as follows:
- --var-file (-f) flag
- env vars CY_STACKFORMS_VAR
- --vars (-V) flag
- --extra-var (-x) flag

--vars accept json encoded values.

You can provide values fron stdin using the '--var-file -' flag.

The output will be the generated configuration of the project.
`,
		Example: `
# create 'prod' environment in 'my-project'
cy --org my-org project create-stackforms-env \
  --project my-project \
  --env prod \
  --use-case usecase-1 \
  --var-file vars.yml \
  --vars '{"myRaw": "vars"}'`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			internal.Warning(cmd.ErrOrStderr(),
				"This command will replace `cy project create-env` soon.\n"+
					"Please see https://github.com/cycloidio/cycloid-cli/issues/268 for more information.\n")
			return internal.CheckAPIAndCLIVersion(cmd, args)
		},

		RunE: createEnv,
	}

	common.WithFlagOrg(cmd)
	cmd.PersistentFlags().StringP("project", "p", "", "the selected project")
	cmd.MarkFlagRequired("project")
	cmd.PersistentFlags().StringP("use-case", "u", "", "the selected use case of the stack")
	cmd.MarkFlagRequired("use-case")
	cmd.PersistentFlags().StringP("env", "e", "", "the environment name of the stack")
	cmd.MarkFlagRequired("env")
	cmd.PersistentFlags().StringArrayP("var-file", "f", nil, "path to a JSON file containing variables, can be '-' for stdin, can be set multiple times.")
	cmd.PersistentFlags().StringArrayP("vars", "V", nil, "JSON string containing variables, can be set multiple times.")
	cmd.PersistentFlags().StringToStringP("extra-var", "x", nil, "extra variable to be added to the environment in the -e key=value -e key=value format")
	cmd.PersistentFlags().Bool("update", false, "Allow to override existing environment")
	cmd.Flags().SortFlags = false

	return cmd
}

func createEnv(cmd *cobra.Command, args []string) error {
	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	var err error

	org, err := cmd.Flags().GetString("org")
	if err != nil {
		return err
	}

	// Right now, in the way orgs are handled, there is a possibility
	// of an empty org.
	// TODO: Fix in another PR about orgs.
	if org == "" {
		return errors.New("org is empty, please specify an org with --org")
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

	usecase, err := cmd.Flags().GetString("use-case")
	if err != nil {
		return err
	}

	if usecase == "" {
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

	extraVar, err := cmd.Flags().GetStringToString("extra-var")
	if err != nil {
		return err
	}

	//
	// Variable merge
	//

	var vars = make(map[string]interface{})

	// Fetch vars from files and stdin
	for _, varFile := range varsFiles {
		internal.Debug("found var file", varFile)
		var decoder *json.Decoder
		if varFile == "-" {
			decoder = json.NewDecoder(os.Stdin)
		} else {
			reader, err := os.Open(varFile)
			if err != nil {
				return fmt.Errorf("failed to read input vars from stdin: %v", err)
			}
			defer reader.Close()
			decoder = json.NewDecoder(reader)
		}

		// Files can contain one or more object, so we scan for all with a decoder
		for {
			var extractedVars = make(map[string]interface{})
			err := decoder.Decode(&extractedVars)
			if err == io.EOF {
				internal.Debug("finished reading input vars from", varFile)
				break
			}

			if err != nil {
				if varFile == "-" {
					varFile = "stdin"
				}
				return fmt.Errorf("failed to read input vars from "+varFile+": %v", err)
			}

			if err := mergo.Merge(&vars, extractedVars, mergo.WithOverride); err != nil {
				log.Fatalf("failed to merge input vars from "+varFile+": %v", err)
			}
		}
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

		if err := mergo.Merge(&vars, envVars, mergo.WithOverride); err != nil {
			log.Fatalf("failed to merge input vars from environment: %v", err)
		}
	}

	// Get variables via CLI arguments --vars
	cliVars, err := cmd.Flags().GetStringArray("vars")
	if err != nil {
		return err
	}

	for _, varInput := range cliVars {
		internal.Debug("found var input", varInput)
		var extractedVars = make(map[string]interface{})
		err = json.Unmarshal([]byte(varInput), &extractedVars)
		if err != nil {
			return fmt.Errorf("failed to parse var input '"+varInput+"' as JSON: %s", err)
		}

		if err := mergo.Merge(&vars, extractedVars, mergo.WithOverride); err != nil {
			log.Fatalf("failed to merge input vars from environment: %v", err)
		}
	}

	// Merge key/val from extraVar
	for k, v := range extraVar {
		common.UpdateMapField(k, v, vars)
	}

	projectData, err := m.GetProject(org, project)
	if err != nil {
		return err
	}

	output, err := cmd.Flags().GetString("output")
	if err != nil {
		return errors.Wrap(err, "unable to get output flag")
	}

	// fetch the printer from the factory
	p, err := factory.GetPrinter(output)
	if err != nil {
		return errors.Wrap(err, "unable to get printer")
	}

	envs := make([]*models.NewEnvironment, len(projectData.Environments))

	for i, e := range projectData.Environments {
		if *e.Canonical == env && !update {
			return fmt.Errorf("environment %s exists already in %s\nIf you want to update it, add the --update flag.", env, project)
		}

		if e.Canonical == nil {
			return fmt.Errorf("missing canonical for environment %v", e)
		}

		cloudProviderCanonical := ""
		if e.CloudProvider != nil {
			cloudProviderCanonical = *e.CloudProvider.Canonical
		}

		color := "default"
		if e.Color != nil {
			color = *e.Color
		}

		icon := "extension"
		if e.Icon != nil {
			icon = *e.Icon
		}

		envs[i] = &models.NewEnvironment{
			Canonical:              e.Canonical,
			CloudProviderCanonical: cloudProviderCanonical,
			Color:                  color,
			Icon:                   icon,
		}
	}

	// Infer icon and color based on usecase
	var cloudProviderCanonical, icon, color string
	switch strings.ToLower(usecase) {
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

	// finally add the new environment
	envs = append(envs, &models.NewEnvironment{
		// TODO: https://github.com/cycloidio/cycloid-cli/issues/67
		Canonical:              &env,
		CloudProviderCanonical: cloudProviderCanonical,
		Color:                  color,
		Icon:                   icon,
	})

	inputs := []*models.FormInput{
		{
			EnvironmentCanonical: &env,
			UseCase:              &usecase,
			Vars:                 vars,
		},
	}

	// Send the updateProject call
	// TODO: Add support for resource pool canonical in case of resource quotas
	resp, err := m.UpdateProject(org,
		*projectData.Name,
		project,
		envs,
		projectData.Description,
		*projectData.ServiceCatalog.Ref,
		*projectData.Owner.Username,
		projectData.ConfigRepositoryCanonical,
		inputs,
		*projectData.UpdatedAt,
	)

	return printer.SmartPrint(p, resp, err, "", printer.Options{}, cmd.OutOrStdout())
}
