package components

import (
	"fmt"

	"github.com/cycloidio/cycloid-cli/client/models"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cy_args"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/printer/factory"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func NewCreateComponentCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create -p project -e env -c component",
		Args:  cobra.NoArgs,
		Short: "Create a component",
		RunE:  createComponent,
	}
	cy_args.AddCyContext(cmd)
	cy_args.AddNameFlag(cmd)
	cy_args.AddComponentDescriptionFlag(cmd)
	useCaseFlag := cy_args.AddUseCaseFlag(cmd)
	cmd.MarkFlagRequired(useCaseFlag)
	cy_args.AddCloudProviderFlag(cmd)
	cy_args.AddStackRefFlag(cmd)
	cy_args.AddStackFormsInputFlags(cmd)
	cmd.Flags().Bool("update", false, "If the component exists, update it.")
	return cmd
}

func createComponent(cmd *cobra.Command, args []string) error {
	org, project, env, component, err := cy_args.GetCyContext(cmd)
	if err != nil {
		return err
	}

	name, err := cy_args.GetName(cmd)
	if err != nil {
		return err
	}
	if name == "" {
		// if name is empty, use the canonical
		name = component
	}

	description, err := cy_args.GetComponentDescription(cmd)
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

	cloudProvider, err := cy_args.GetCloudProvider(cmd)
	if err != nil {
		return err
	}

	output, err := cy_args.GetOutput(cmd)
	if err != nil {
		return err
	}

	update, err := cmd.Flags().GetBool("update")
	if err != nil {
		return err
	}

	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	p, err := factory.GetPrinter(output)
	if err != nil {
		return errors.Wrap(err, "unable to get printer")
	}

	var componentOutput *models.Component
	if update {
		_, err := m.GetComponent(org, project, env, component)
		if err == nil {
			// Fetch base forms value from current component
			config, err := m.GetComponentConfig(org, project, env, component)
			if err != nil {
				return printer.SmartPrint(p, nil, err, "failed to update component '"+component+"', cannot get current config.", printer.Options{}, cmd.OutOrStderr())
			}

			inputs, err := cy_args.GetStackformsVars(cmd, config)
			if err != nil {
				return err
			}

			componentOutput, err = m.UpdateComponent(org, project, env, component, *description, &name, useCase, inputs)
			if err != nil {
				return printer.SmartPrint(p, nil, err, "failed to update component '"+component+"'", printer.Options{}, cmd.OutOrStderr())
			}
			return printer.SmartPrint(p, componentOutput, nil, "", printer.Options{}, cmd.OutOrStdout())
		}
	}

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

	componentOutput, err = m.CreateComponent(org, project, env, component, *description, &name, stackRef, useCase, cloudProvider, inputs)
	if err != nil {
		return printer.SmartPrint(p, nil, err, "failed to create component '"+component+"'", printer.Options{}, cmd.OutOrStderr())
	}

	return printer.SmartPrint(p, componentOutput, nil, "", printer.Options{}, cmd.OutOrStdout())
}
