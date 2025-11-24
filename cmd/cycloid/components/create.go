package components

import (
	"fmt"
	"slices"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/client/models"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/printer/factory"
)

func NewCreateComponentCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "create",
		Args:    cobra.NoArgs,
		Short:   "Create a component",
		RunE:    createComponent,
		Example: `cy component create -p project -e env -c component -V section.group.var="value-str" -s "stack:ref" -u new-use-case`,
	}
	cyargs.AddCyContext(cmd)
	cyargs.AddNameFlag(cmd)
	cyargs.AddComponentDescriptionFlag(cmd)
	cmd.MarkFlagRequired(cyargs.AddUseCaseFlag(cmd))
	cyargs.AddStackVersionFlags(cmd)
	cyargs.AddCloudProviderFlag(cmd)
	cmd.MarkFlagRequired(cyargs.AddComponentStackRefFlag(cmd))
	cyargs.AddStackFormsInputFlags(cmd)
	cmd.Flags().Bool("update", false, "If the component exists, update it.")
	return cmd
}

func createComponent(cmd *cobra.Command, args []string) error {
	org, project, env, component, err := cyargs.GetCyContext(cmd)
	if err != nil {
		return err
	}

	name, err := cyargs.GetName(cmd)
	if err != nil {
		return err
	}

	if name == "" {
		// if name is empty, use the canonical
		name = component
	}

	description, err := cyargs.GetComponentDescription(cmd)
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

	cloudProvider, err := cyargs.GetCloudProvider(cmd)
	if err != nil {
		return err
	}

	output, err := cyargs.GetOutput(cmd)
	if err != nil {
		return err
	}

	update, err := cmd.Flags().GetBool("update")
	if err != nil {
		return err
	}

	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	// Get the stack version flags
	tag, branch, hash, err := cyargs.GetStackVersionFlags(cmd)
	if err != nil {
		return errors.Wrap(err, "failed to read stack version flags")
	}

	p, err := factory.GetPrinter(output)
	if err != nil {
		return errors.Wrap(err, "unable to get printer")
	}

	if update {
		components, err := m.ListComponents(org, project, env)
		if err != nil {
			return fmt.Errorf("failed to create --update component, cannot check existing component %q: %w", component, err)
		}

		componentIndex := slices.IndexFunc(components, func(c *models.Component) bool {
			return *c.Canonical == component
		})
		if componentIndex != -1 {
			currentComponent := components[componentIndex]

			// Fetch base forms value from current component
			var currentConfig = make(models.FormVariables)
			if currentComponent.IsConfigured {
				currentConfig, err = m.GetComponentConfig(org, project, env, component)
				if err != nil {
					return printer.SmartPrint(p, nil, err, "failed to update component '"+component+"', cannot get current config.", printer.Options{}, cmd.OutOrStderr())
				}
			}

			inputs, err := cyargs.GetStackformsVars(cmd, currentConfig)
			if err != nil {
				return err
			}

			// ConfigureComponent will reconfigure the component
			componentOutput, err := m.CreateAndConfigureComponent(org, project, env, component, *description, name, stackRef, tag, branch, hash, useCase, *cloudProvider, inputs)
			if err != nil {
				return printer.SmartPrint(p, nil, err, "failed to configure component '"+component+"'", printer.Options{}, cmd.OutOrStderr())
			}

			return printer.SmartPrint(p, componentOutput, nil, "", printer.Options{}, cmd.OutOrStdout())
		}
	}

	componentEntity, err := m.GetComponent(org, project, env, component)
	if err == nil {
		return printer.SmartPrint(p, componentEntity, fmt.Errorf("component %q already exists, to update it, use the --update flag", component), "failed to create component", printer.Options{}, cmd.OutOrStderr())
	}

	inputs, err := cyargs.GetStackformsVars(cmd, nil)
	if err != nil {
		return err
	}

	componentOutput, err := m.CreateAndConfigureComponent(org, project, env, component, *description, name, stackRef, tag, branch, hash, useCase, *cloudProvider, inputs)
	if err != nil {
		return printer.SmartPrint(p, nil, err, "failed to create and configure component '"+component+"'", printer.Options{}, cmd.OutOrStderr())
	}

	return printer.SmartPrint(p, componentOutput, nil, "", printer.Options{}, cmd.OutOrStdout())
}
