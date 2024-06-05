package projects

import (
	"fmt"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/internal"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/printer/factory"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func NewGetStackformsConfigFromEnvCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "get-env-config <project?> <env?>",
		Short: "Get the default stackforms config of a project's env",
		Long: `
This command will fetch the configuration of an environment in a project.

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

- First we get the current env value if exists (unless you set --default)
- I no current value is present, we get the default
- If no default is set, we set a zeroed value in the correct type: ("", 0, [], {})
`,
		Example: `
# Get the configuration as json (default)
cy --org my-org project get-env-config -p my-project-canonical -e my-env-canonical

# Get the configuration as yaml
cy --org my-org project get-env-config my-project my-project use_case -o yaml
`,
		PreRunE: internal.CheckAPIAndCLIVersion,
		RunE:    getStackFormsConfigFromEnv,
		Args:    cobra.RangeArgs(0, 2),
	}

	common.WithFlagOrg(cmd)
	cmd.Flags().StringP("project", "p", "", "specify the project")
	cmd.Flags().StringP("env", "e", "", "specify the env")
	cmd.Flags().BoolP("default-values", "d", false, "if set, will fetch the default value from the stack instead of the current ones.")

	// This will display flag in the order declared above
	cmd.Flags().SortFlags = false

	return cmd
}

func getStackFormsConfigFromEnv(cmd *cobra.Command, args []string) error {
	// Flags have precedence over args
	project, err := cmd.Flags().GetString("project")
	if len(args) >= 1 && project == "" {
		project = args[0]
	} else if project == "" {
		return fmt.Errorf("missing project argument")
	}

	env, err := cmd.Flags().GetString("env")
	if len(args) == 2 && env == "" {
		env = args[1]
	} else if env == "" {
		return fmt.Errorf("missing use case argument")
	}

	getDefault, err := cmd.Flags().GetBool("default-values")
	if err != nil {
		return err
	}

	internal.Debug("project:", project, "| env:", env)

	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	org, err := cmd.Flags().GetString("org")
	if err != nil {
		return err
	}

	output, err := cmd.Flags().GetString("output")
	if err != nil {
		return errors.Wrap(err, "unable to get output flag")
	}

	// output as yaml by default
	if output == "table" {
		output = "json"
	}

	// fetch the printer from the factory
	p, err := factory.GetPrinter(output)
	if err != nil {
		return errors.Wrap(err, "unable to get printer")
	}

	resp, err := m.GetProjectConfig(org, project, env)
	if err != nil || resp == nil {
		return printer.SmartPrint(p, nil, err, fmt.Sprint("failed to fetch project '", project, "' config for env '", env, "' in org '", org, "'"), printer.Options{}, cmd.OutOrStderr())
	}

	formData, err := common.ParseFormsConfig(resp, *resp.UseCase, !getDefault)
	if err != nil {
		fmt.Println("failed to parse config data")
		return printer.SmartPrint(p, nil, err, "failed to get stack config", printer.Options{}, cmd.OutOrStdout())
	}

	return printer.SmartPrint(p, formData, err, "failed to get stack config", printer.Options{}, cmd.OutOrStdout())
}
