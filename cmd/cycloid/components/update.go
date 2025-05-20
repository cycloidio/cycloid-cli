package components

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

func NewUpdateComponentCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update -p project -e env -c component",
		Args:  cobra.NoArgs,
		Short: "update an existing component",
		RunE:  updateComponent,
	}
	cy_args.AddCyContext(cmd)
	cy_args.AddNameFlag(cmd)
	cy_args.AddComponentDescriptionFlag(cmd)
	cy_args.AddUseCaseFlag(cmd)
	cy_args.AddStackFormsInputFlags(cmd)
	return cmd
}

func updateComponent(cmd *cobra.Command, args []string) error {
	org, project, env, component, err := cy_args.GetCyContext(cmd)
	if err != nil {
		return err
	}

	name, err := cy_args.GetName(cmd)
	if err != nil {
		return err
	}

	description, err := cy_args.GetComponentDescription(cmd)
	if err != nil {
		return err
	}

	useCase, err := cy_args.GetUseCase(cmd)
	if err != nil {
		return err
	}

	output, err := cmd.Flags().GetString("output")
	if err != nil {
		return err
	}

	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	p, err := factory.GetPrinter(output)
	if err != nil {
		return errors.Wrap(err, "unable to get printer")
	}

	// Fetch base forms value from current component
	config, err := m.GetComponentConfig(org, project, env, component)
	if err != nil {
		return printer.SmartPrint(p, nil, err, "failed to update component '"+component+"', cannot get current config.", printer.Options{}, cmd.OutOrStderr())
	}

	inputs, err := cy_args.GetStackformsVars(cmd, config)
	if err != nil {
		return err
	}

	// fetch the current component to fill unspecified values by the user
	current, err := m.GetComponent(org, project, env, component)
	if err != nil {
		return fmt.Errorf("failed to update component, target '%s' doesn't seems to exists: %s", component, err)
	}

	if name == "" {
		name = *current.Name
	}

	if *description == "" && current.Description == "" {
		description = &current.Description
	}

	if *useCase == "" {
		useCase = current.UseCase
	}

	updatedComponent, err := m.UpdateComponent(org, project, env, component, *description, &name, useCase, inputs)
	if err != nil {
		return printer.SmartPrint(p, nil, err, "failed to update component '"+component+"'", printer.Options{}, cmd.OutOrStderr())
	}

	return printer.SmartPrint(p, updatedComponent, nil, "", printer.Options{}, cmd.OutOrStdout())
}
