package config

import (
	"fmt"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cy_args"
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

	cmd.Flags().StringP("use-case", "u", "", "specify the use case canonical")
	cmd.Flags().StringP("stack-ref", "s", "", "stack ref (sometimes called service_catalog_ref")
	cy_args.AddCyContext(cmd)
	cy_args.AddStackFormsInputFlags(cmd)
	return cmd
}

func interpolate(cmd *cobra.Command, args []string) error {
	org, project, env, component, err := cy_args.GetCyContext(cmd)
	if err != nil {
		return err
	}

	useCase, err := cy_args.GetUseCase(cmd)
	if err != nil {
		return err
	}

	stackRef, err := cy_args.GetStackRef(cmd)
	if err != nil {
		return err
	}

	output, err := cy_args.GetOutput(cmd)
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
	stackConfig, err := m.GetStackConfig(org, *stackRef)
	if err != nil {
		return err
	}

	useCaseConfig, err := common.FormUseCaseToFormVars(stackConfig, *useCase)
	if err != nil {
		return fmt.Errorf("failed to parse default value for stack '%s' with use-case '%s': %s", *stackRef, *useCase, err)
	}

	inputs, err := cy_args.GetStackformsVars(cmd, useCaseConfig)
	if err != nil {
		return err
	}

	p, err := factory.GetPrinter(output)
	if err != nil {
		return errors.Wrap(err, "unable to get printer")
	}

	config, err := m.InterpolateFormsConfig(org, project, env, component, *stackRef, *useCase, inputs)
	if err != nil {
		return printer.SmartPrint(p, nil, err, "failed to interpolate config", printer.Options{}, cmd.OutOrStderr())
	}

	return printer.SmartPrint(p, config, nil, "failed to interpolate config", printer.Options{}, cmd.OutOrStdout())
}
