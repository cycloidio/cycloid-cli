package projects

import (
	"encoding/json"
	"fmt"
	"slices"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/client/models"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/internal"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/printer/factory"
)

func NewGetEnvCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use: "get-env <project?> <env?>",
		Aliases: []string{
			"get-env-config",
		},
		Short: "Get the default stackforms config of a project's env",
		Long: `This command will fetch the configuration of an environment in a project.

Output is in JSON by default.

The output object will be the same format required as input for 'cy project create-stackform-env' like the following: 

{
  "mySection": {
    "myGroup1": {
      "myVar1": "myValue"
    }
  }
}

The values are generated as following:

- If no current value is present, we get the default
`,
		Example: `# Get the configuration as json (default)
cy --org my-org project get-env-config -p my-project-canonical -e my-env-canonical

# Get the configuration as yaml
cy --org my-org project get-env-config my-project my-project use_case -o yaml`,
		PreRunE: internal.CheckAPIAndCLIVersion,
		RunE:    getEnvConfig,
		Args:    cobra.RangeArgs(0, 2),
	}

	cmd.Flags().StringP("project", "p", "", "specify the project")
	cmd.Flags().StringP("env", "e", "", "specify the env")
	cmd.Flags().BoolP("default", "d", false, "if set, will fetch the default value from the stack on top of the current ones.")

	// This will display flag in the order declared above
	cmd.Flags().SortFlags = false

	return cmd
}

func getEnvConfig(cmd *cobra.Command, args []string) error {
	// Flags have precedence over args
	project, _ := cmd.Flags().GetString("project")
	if len(args) >= 1 && project == "" {
		project = args[0]
	} else if project == "" {
		return fmt.Errorf("missing project argument")
	}

	env, _ := cmd.Flags().GetString("env")
	if len(args) == 2 && env == "" {
		env = args[1]
	} else if env == "" {
		return fmt.Errorf("missing use case argument")
	}

	useDefaults, err := cmd.Flags().GetBool("default")
	if err != nil {
		return err
	}

	internal.Debug("project:", project, "| env:", env)

	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	org, err := common.GetOrg(cmd)
	if err != nil {
		return err
	}

	output, err := cmd.Flags().GetString("output")
	if err != nil {
		return errors.Wrap(err, "unable to get output flag")
	}

	// output as json by default
	if output == "table" {
		output = "json"
	}

	// fetch the printer from the factory
	p, err := factory.GetPrinter(output)
	if err != nil {
		return errors.Wrap(err, "unable to get printer")
	}

	resp, err := m.GetProjectConfig(org, project, env)
	if err != nil {
		return printer.SmartPrint(p, nil, err, fmt.Sprint("failed to fetch project '", project, "' config for env '", env, "' in org '", org, "'"), printer.Options{}, cmd.OutOrStderr())
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
		return errors.New("failed to marshall API response.")
	}

	var useCase common.UseCase
	err = json.Unmarshal(data, &useCase)

	vars := common.UseCaseToFormInput(useCase, useDefaults)
	return printer.SmartPrint(p, vars, err, "failed to get stack config", printer.Options{}, cmd.OutOrStdout())
}
