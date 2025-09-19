package stacks

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/internal"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/printer/factory"
)

func NewStacksGetComponentStackConfig() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "get-config -p project -e env -c component <use-case?> [flags]",
		Short: "Output a stack configuration in JSON",
		Example: `
cy --org my-org stacks get-config -p my-project -e my-env -c my-component my-usecase
`,
		RunE: getConfig,
		Args: cobra.RangeArgs(0, 2),
	}
	cyargs.AddUseCaseFlag(cmd)

	return cmd
}

func getConfig(cmd *cobra.Command, args []string) error {
	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}

	proj, err := cyargs.GetProject(cmd)
	if err != nil {
		return err
	}

	env, err := cyargs.GetEnv(cmd)
	if err != nil {
		return err
	}

	component, err := cyargs.GetComponent(cmd)
	if err != nil {
		return err
	}

	useCase, _ := cyargs.GetUseCase(cmd)
	if len(args) == 2 && *useCase == "" {
		useCase = &args[1]
	} else if *useCase == "" {
		return fmt.Errorf("missing use-case argument")
	}

	internal.Debug("project:", proj, "env:", env, "component:", component, "usecase:", useCase)

	output, err := cmd.Flags().GetString("output")
	if err != nil {
		return errors.Wrap(err, "unable to get output flag")
	}

	// Default output is in JSON
	if output == "table" {
		output = "json"
	}

	// fetch the printer from the factory
	p, err := factory.GetPrinter(output)
	if err != nil {
		return errors.Wrap(err, "unable to get printer")
	}

	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	stackConfigs, err := m.GetComponentStackConfig(org, proj, env, component, *useCase)
	if err != nil {
		return printer.SmartPrint(p, nil, err, "unable to get the stack configuration", printer.Options{}, cmd.OutOrStderr())
	}

	useCaseConfig, err := common.FormUseCaseToFormVars(stackConfigs, *useCase)
	if err != nil {
		return fmt.Errorf("failed to parse default form values for component '%s' with use-case '%s': %s", component, *useCase, err)
	}

	config, err := cyargs.GetStackformsVars(cmd, useCaseConfig)
	if err != nil {
		return printer.SmartPrint(p, nil, err, "failed to fetch stack config.", printer.Options{}, cmd.OutOrStderr())
	}

	return printer.SmartPrint(p, config, nil, "", printer.Options{}, cmd.OutOrStdout())
}
