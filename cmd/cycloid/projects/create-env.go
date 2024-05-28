package projects

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"maps"
	"os"
	"time"

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
- --var-file flag
- env vars  
- --vars flag

--vars accept json encoded values.

You can provide values fron stdin using the '--var-file -' flag.
`,
		Example: `
# create 'prod' environment in 'my-project'
 cy --org my-org project create-raw-env \
  --project my-project \
  --env prod \
  --usecase usecase-1 \
  --var-file vars.yml \
  --vars '{"myRaw": "vars"}'
`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			internal.Warning(cmd.ErrOrStderr(),
				"This command will replace `cy project create-env` soon.\n"+
					"Please see https://github.com/cycloidio/cycloid-cli/issues/268 for more information.\n")
			return internal.CheckAPIAndCLIVersion(cmd, args)
		},

		RunE: createEnv,
	}

	common.RequiredPersistentFlag(common.WithFlagProject, cmd)
	common.RequiredPersistentFlag(common.WithFlagEnv, cmd)
	WithFlagConfig(cmd)
	WithFlagUsecase(cmd)
	cmd.PersistentFlags().StringArrayP("var-file", "f", nil, "path to a JSON file containing variables, can be '-' for stdin")
	cmd.PersistentFlags().StringArray("vars", nil, "JSON string containing variables")

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

	project, err := cmd.Flags().GetString("project")
	if err != nil {
		return err
	}

	env, err := cmd.Flags().GetString("env")
	if err != nil {
		return err
	}

	usecase, err := cmd.Flags().GetString("usecase")
	if err != nil {
		return err
	}

	varsFiles, err := cmd.Flags().GetStringArray("var-file")
	if err != nil {
		return err
	}

	// init the final variable map
	// TODO: We need to implement a recursive merge function
	// To be able to merge submaps
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
				log.Fatalf("failed to read input vars from "+varFile+": %v", err)
				break
			}
			maps.Copy(vars, extractedVars)
		}
	}

	internal.Debug("found theses vars via files:", vars)

	// Get vars via the CY_CREATE_ENV_VARS env var
	envConfig, exists := os.LookupEnv("CY_CREATE_ENV_VARS")
	if exists {
		internal.Debug("found config via env var", envConfig)
		var envVars = make(map[string]interface{})
		err := json.Unmarshal([]byte(envConfig), &envVars)

		// TODO: does this should error if parsing fail, of should we just put a warning ?
		if err != nil {
			return fmt.Errorf("failed to parse env var config '"+envConfig+"' as JSON: %s", err)
		}

		maps.Copy(vars, envVars)
	}

	// Get variables via CLI arguments
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

		maps.Copy(vars, extractedVars)
	}

	internal.Debug("vars after CLI merge:", vars)

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

	// need to conver the environment to "new environment" as required
	// by the API
	envs := make([]*models.NewEnvironment, len(projectData.Environments))

	for i, e := range projectData.Environments {
		if *e.Canonical == env {
			return fmt.Errorf("environment %s exists already in %s", env, project)
		}
		envs[i] = &models.NewEnvironment{
			Canonical: e.Canonical,
		}
	}

	// finally add the new environment
	envs = append(envs, &models.NewEnvironment{
		// TODO: https://github.com/cycloidio/cycloid-cli/issues/67
		Canonical: &env,
	})

	inputs := models.FormInputs{
		Inputs: []models.FormInput{
			{
				EnvironmentCanonical: &env,
				UseCase:              &usecase,
				Vars:                 vars,
			},
		},
		ServiceCatalogRef: &projectData.ServiceCatalog.Ref,
	}
	//
	// UPDATE PROJECT
	//
	// TODO: Add support for resource pool canonical in case of resource quotas
	timestamp := time.Now()
	_, err = m.UpdateProject(org,
		*projectData.Name,
		project,
		envs,
		projectData.Description,
		*projectData.ServiceCatalog.Ref,
		*projectData.Owner.Username,
		projectData.ConfigRepositoryCanonical,
		inputs,
		timestamp,
	)

	err = printer.SmartPrint(p, nil, err, "unable to update project", printer.Options{}, cmd.OutOrStdout())
	if err != nil {
		return err
	}

	return nil

	// //
	// // PIPELINE UNPAUSE
	// //
	// err = m.UnpausePipeline(org, project, env)
	// err = printer.SmartPrint(p, nil, err, "unable to unpause pipeline", printer.Options{}, cmd.OutOrStdout())
	// if err != nil {
	// 	return err
	// }
	//
	// return printer.SmartPrint(p, resp, err, "", printer.Options{}, cmd.OutOrStdout())
}
