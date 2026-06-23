package config

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/apiclient"
	"github.com/cycloidio/cycloid-cli/cmd/common"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
	"github.com/cycloidio/cycloid-cli/printer"
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
	cyargs.AddStackVersionFlags(cmd)
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

	api := common.NewAPI()
	m := apiclient.NewMiddleware(api)

	// Resolve stack version: --stack-version (new) or legacy flags.
	tag, branch, hash, err := cyargs.ResolveStackVersionArg(cmd, m, org, stackRef)
	if err != nil {
		return fmt.Errorf("failed to read stack version flags: %w", err)
	}

	// Get default to stacks
	stackConfig, _, err := m.GetComponentStackConfig(org, project, env, component, useCase, tag, branch, hash)
	if err != nil {
		return err
	}

	useCaseConfig, err := common.FormUseCaseToFormVars(stackConfig, useCase)
	if err != nil {
		return fmt.Errorf("failed to parse default value for stack %q with use-case %q: %w", stackRef, useCase, err)
	}

	inputs, err := cyargs.GetStackformsVars(cmd, useCaseConfig)
	if err != nil {
		return err
	}

	config, _, err := m.InterpolateFormsConfig(org, project, env, component, stackRef, useCase, inputs)
	return cyout.PrintWithOptions(cmd, config, err, "failed to interpolate config", printer.Options{})
}
