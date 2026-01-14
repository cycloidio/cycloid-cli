package components

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/client/models"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/printer/factory"
)

func NewUpdateComponentCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "update",
		Args:    cobra.NoArgs,
		Short:   "update an existing component",
		Example: `cy component update -p project -e env -c component -V section.group.var="value-str" -u new-use-case`,
		RunE:    updateComponent,
	}

	cyargs.AddCyContext(cmd)
	cyargs.AddNameFlag(cmd)
	cyargs.AddComponentDescriptionFlag(cmd)
	cyargs.AddUseCaseFlag(cmd)
	cyargs.AddStackVersionFlags(cmd)
	cyargs.AddStackFormsInputFlags(cmd)
	cyargs.AddStackRefFlag(cmd)
	cyargs.AddCloudProviderFlag(cmd)

	return cmd
}

func updateComponent(cmd *cobra.Command, args []string) error {
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

	stackRef, err := cyargs.GetStackRef(cmd)
	if err != nil {
		return err
	}

	cloudProvider, err := cyargs.GetCloudProvider(cmd)
	if err != nil {
		return err
	}

	useCase, err := cyargs.GetUseCase(cmd)
	if err != nil {
		return err
	}

	output, err := cyargs.GetOutput(cmd)
	if err != nil {
		return err
	}

	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	p, err := factory.GetPrinter(output)
	if err != nil {
		return errors.Wrap(err, "unable to get printer")
	}

	currentComponent, err := m.GetComponent(org, project, env, component)
	if err != nil {
		return printer.SmartPrint(p, nil, err, "failed to update component '"+component+"', cannot get current component", printer.Options{}, cmd.OutOrStderr())
	}

	if stackRef == "" {
		stackRef = *currentComponent.ServiceCatalog.Ref
	}

	if useCase == "" && currentComponent.UseCase == "" {
		return errors.New("cannot determine the use case, please fill it with -(-u)se-case flag")
	} else if useCase == "" {
		useCase = currentComponent.UseCase
	}

	// Get the stack version flags
	tag, branch, hash, err := cyargs.GetStackVersionFlags(cmd)
	if err != nil {
		return errors.Wrap(err, "failed to read stack version flags")
	}

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

	// CreateComponent will reconfigure the component if it already exists
	updatedComponent, err := m.CreateAndConfigureComponent(org, project, env, component, *description, name, stackRef, tag, branch, hash, useCase, *cloudProvider, inputs)
	if err != nil {
		return printer.SmartPrint(p, nil, err, "failed to configure component '"+component+"'", printer.Options{}, cmd.OutOrStderr())
	}

	return printer.SmartPrint(p, updatedComponent, nil, "", printer.Options{}, cmd.OutOrStdout())
}
