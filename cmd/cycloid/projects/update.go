package projects

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/printer/factory"
)

func NewUpdateCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:     "update",
		Args:    cobra.NoArgs,
		Short:   "update a project",
		Example: `cy --org my-org project update --project "my-project" --name "NewName"`,
		RunE:    update,
	}

	cyargs.AddNameFlag(cmd)
	cyargs.AddProjectFlag(cmd)
	cyargs.AddDescriptionFlag(cmd)
	cyargs.AddIconFlag(cmd)
	cyargs.AddColorFlag(cmd)
	cyargs.AddOwnerFlag(cmd)
	cyargs.AddConfigRepositoryFlag(cmd)
	return cmd
}

func update(cmd *cobra.Command, args []string) error {
	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}

	project, err := cyargs.GetProject(cmd)
	if err != nil {
		return err
	}

	name, err := cyargs.GetName(cmd)
	if err != nil {
		return err
	}

	description, err := cyargs.GetDescription(cmd)
	if err != nil {
		return err
	}

	owner, err := cyargs.GetOwner(cmd)
	if err != nil {
		return err
	}

	color, err := cyargs.GetColor(cmd)
	if err != nil {
		return err
	}

	icon, err := cyargs.GetIcon(cmd)
	if err != nil {
		return err
	}

	configRepository, err := cyargs.GetDefaultConfigRepository(cmd)
	if err != nil {
		return err
	}

	output, err := cyargs.GetOutput(cmd)
	if err != nil {
		return errors.Wrap(err, "unable to get output flag")
	}

	// fetch the printer from the factory
	p, err := factory.GetPrinter(output)
	if err != nil {
		return errors.Wrap(err, "unable to get printer")
	}

	currentProject, err := m.GetProject(org, project)
	if err != nil {
		return printer.SmartPrint(p, currentProject, err, "project not found", printer.Options{}, cmd.OutOrStderr())
	}

	if name == "" {
		name = *currentProject.Name
	}

	if description == "" {
		description = currentProject.Description
	}

	if color == "" && currentProject.Color != nil {
		color = *currentProject.Color
	}

	if icon == "" && currentProject.Icon != nil {
		icon = *currentProject.Icon
	}

	if owner == "" && currentProject.Owner != nil {
		owner = *currentProject.Owner.Username
	}

	if configRepository == "" {
		configRepository = currentProject.ConfigRepositoryCanonical
	}

	resp, err := m.UpdateProject(org, name, project, description, configRepository, owner, "", color, icon, "")
	if err != nil {
		return printer.SmartPrint(p, resp, err, "", printer.Options{}, cmd.OutOrStderr())
	}

	return printer.SmartPrint(p, resp, nil, "", printer.Options{}, cmd.OutOrStdout())
}
