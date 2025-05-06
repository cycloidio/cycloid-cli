package components

import (
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
		Short: "Create a component",
		RunE:  createComponent,
	}
	cy_args.AddCyContext(cmd)
	cy_args.AddComponentNameFlag(cmd)
	cy_args.AddDescriptionFlag(cmd)
	cy_args.AddUseCaseFlag(cmd)
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

	name, err := cy_args.GetComponentName(cmd)
	if err != nil {
		return err
	}
	if name == nil && *name == "" {
		// if name is empty, use the canonical
		name = &component
	}

	description, err := cy_args.GetDescription(cmd)
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

	output, err := cmd.Flags().GetString("output")
	if err != nil {
		return err
	}

	inputs, err := cy_args.GetStackformsVars(cmd)
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
		componentOutput, err = m.UpdateComponent(org, project, env, component, *description, name, useCase, inputs)
		if err != nil {
			return printer.SmartPrint(p, nil, err, "failed to update component '"+component+"'", printer.Options{}, cmd.OutOrStderr())
		}
	} else {
		componentOutput, err = m.CreateComponent(org, project, env, component, *description, name, stackRef, useCase, cloudProvider, inputs)
		if err != nil {
			return printer.SmartPrint(p, nil, err, "failed to create component '"+component+"'", printer.Options{}, cmd.OutOrStderr())
		}
	}

	return printer.SmartPrint(p, componentOutput, nil, "", printer.Options{}, cmd.OutOrStderr())
}

func NewUpdateComponentCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update -p project -e env -c component",
		Short: "update an existing component",
		RunE:  updateComponent,
	}
	cy_args.AddCyContext(cmd)
	cy_args.AddComponentNameFlag(cmd)
	cy_args.AddDescriptionFlag(cmd)
	cy_args.AddUseCaseFlag(cmd)
	cy_args.AddStackFormsInputFlags(cmd)
	return cmd
}

func updateComponent(cmd *cobra.Command, args []string) error {
	org, project, env, component, err := cy_args.GetCyContext(cmd)
	if err != nil {
		return err
	}

	name, err := cy_args.GetComponentName(cmd)
	if err != nil {
		return err
	}

	description, err := cy_args.GetDescription(cmd)
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

	inputs, err := cy_args.GetStackformsVars(cmd)
	if err != nil {
		return err
	}

	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	p, err := factory.GetPrinter(output)
	if err != nil {
		return errors.Wrap(err, "unable to get printer")
	}

	updatedComponent, err := m.UpdateComponent(org, project, env, component, *description, name, useCase, inputs)
	if err != nil {
		return printer.SmartPrint(p, nil, err, "failed to update component '"+component+"'", printer.Options{}, cmd.OutOrStderr())
	}

	return printer.SmartPrint(p, updatedComponent, nil, "", printer.Options{}, cmd.OutOrStderr())
}
