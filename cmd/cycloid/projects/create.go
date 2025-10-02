package projects

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

func NewCreateCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:     "create",
		Args:    cobra.NoArgs,
		Short:   "create a project",
		Example: `cy --org my-org project create --project "my-project"`,
		RunE:    create,
	}

	cyargs.AddNameFlag(cmd)
	cyargs.AddProjectFlag(cmd)
	cyargs.AddDescriptionFlag(cmd)
	cyargs.AddIconFlag(cmd)
	cyargs.AddColorFlag(cmd)
	cyargs.AddOwnerFlag(cmd)
	cyargs.AddConfigRepositoryFlag(cmd)
	cmd.Flags().Bool("update", false, "if set, will update the project if it exists.")
	return cmd
}

func create(cmd *cobra.Command, args []string) error {
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

	icon, err := cyargs.GetIcon(cmd)
	if err != nil {
		return err
	}

	color, err := cyargs.GetColor(cmd)
	if err != nil {
		return err
	}

	owner, err := cyargs.GetOwner(cmd)
	if err != nil {
		return err
	}

	configRepository, err := cyargs.GetDefaultConfigRepository(cmd)
	if err != nil {
		return err
	}

	update, err := cmd.Flags().GetBool("update")
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

	if update {
		projects, err := m.ListProjects(org)
		if err != nil {
			return fmt.Errorf("failed to create --update project, cannot check is project %q exists: %w", project, err)
		}

		currentIndex := slices.IndexFunc(projects, func(p *models.Project) bool { return *p.Canonical == project })
		if currentIndex != -1 {
			// Make the update use the current color if not explicitly set by the user
			current := projects[currentIndex]
			if color == cyargs.DefaultColor {
				if current.Color != nil {
					color = *current.Color
				} else {
					// Use a random one if none is set
					color = cyargs.PickRandomColor(nil)
				}
			}

			if icon == cyargs.DefaultIcon {
				if current.Icon != nil {
					icon = *current.Icon
				} else {
					// Use a random one if none is set
					icon = cyargs.PickRandomIcon(nil)
				}
			}

			resp, err := m.UpdateProject(org, name, project, description, configRepository, owner, "", color, icon, "")
			if err != nil {
				return printer.SmartPrint(p, nil, err, "", printer.Options{}, cmd.OutOrStderr())
			}

			return printer.SmartPrint(p, resp, nil, "", printer.Options{}, cmd.OutOrStdout())
		}
	}

	if color == cyargs.DefaultColor {
		color = cyargs.PickRandomColor(nil)
	}

	if icon == cyargs.DefaultIcon {
		icon = cyargs.PickRandomIcon(nil)
	}

	resp, err := m.CreateProject(org, name, project, description, configRepository, owner, "", color, icon)
	if err != nil {
		return printer.SmartPrint(p, nil, err, "", printer.Options{}, cmd.OutOrStderr())
	}

	return printer.SmartPrint(p, resp, nil, "", printer.Options{}, cmd.OutOrStdout())
}
