package config

import (
	"fmt"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/printer/factory"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func NewInterpolateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   `interpolate -s service:catalog -u use_case -p project -e env -c component -v '{"section": {"group": {"var": "my_var"}}}'`,
		Args:  cobra.NoArgs,
		Short: "Generate a set of configs based on the forms input but without creating anything",
		RunE:  interpolate,
	}

	cyargs.AddUseCaseFlag(cmd)
	cyargs.AddStackRefFlag(cmd)
	cyargs.AddCyContext(cmd)
	cyargs.AddStackFormsInputFlags(cmd)
	return cmd
}

func interpolate(cmd *cobra.Command, args []string) error {
	org, project, env, component, err := cyargs.GetCyContext(cmd)
	if err != nil {
		return err
	}

	useCase, err := cyargs.GetUseCase(cmd)
	if err != nil {
		return err
	}

	stackRef, err := cyargs.GetStackRef(cmd)
	if err != nil {
		return err
	}

	output, err := cyargs.GetOutput(cmd)
	if err != nil {
		return err
	}

	// This endpoint doesn't make sense in table mode
	if output == "table" {
		output = "json"
	}

	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	// Get default to stacks
	stackConfig, err := m.GetStackConfig(org, stackRef)
	if err != nil {
		return err
	}

	useCaseConfig, err := common.FormUseCaseToFormVars(stackConfig, *useCase)
	if err != nil {
		return fmt.Errorf("failed to parse default value for stack '%s' with use-case '%s': %s", stackRef, *useCase, err)
	}

	inputs, err := cyargs.GetStackformsVars(cmd, useCaseConfig)
	if err != nil {
		return err
	}

	p, err := factory.GetPrinter(output)
	if err != nil {
		return errors.Wrap(err, "unable to get printer")
	}

	config, err := m.InterpolateFormsConfig(org, project, env, component, stackRef, *useCase, inputs)
	if err != nil {
		return printer.SmartPrint(p, nil, err, "failed to interpolate config", printer.Options{}, cmd.OutOrStderr())
	}

	return printer.SmartPrint(p, config, nil, "failed to interpolate config", printer.Options{}, cmd.OutOrStdout())
}
