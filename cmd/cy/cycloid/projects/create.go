package projects

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cy/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cy/cycloid/internal"
	"github.com/cycloidio/cycloid-cli/cmd/cy/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cy_args"
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
		PreRunE: internal.CheckAPIAndCLIVersion,
	}

	cy_args.AddNameFlag(cmd)
	cy_args.AddProjectFlag(cmd)
	cy_args.AddDescriptionFlag(cmd)
	cy_args.AddIconFlag(cmd)
	cy_args.AddColorFlag(cmd)
	cy_args.AddOwnerFlag(cmd)
	cy_args.AddConfigRepositoryFlag(cmd)
	cmd.Flags().Bool("update", false, "if set, will update the project if it exists.")
	return cmd
}

func create(cmd *cobra.Command, args []string) error {
	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	org, err := cy_args.GetOrg(cmd)
	if err != nil {
		return err
	}

	project, err := cy_args.GetProject(cmd)
	if err != nil {
		return err
	}

	name, err := cy_args.GetName(cmd)
	if err != nil {
		return err
	}

	description, err := cy_args.GetDescription(cmd)
	if err != nil {
		return err
	}

	icon, err := cy_args.GetIcon(cmd)
	if err != nil {
		return err
	}

	color, err := cy_args.GetColor(cmd)
	if err != nil {
		return err
	}

	owner, err := cy_args.GetOwner(cmd)
	if err != nil {
		return err
	}

	configRepository, err := cy_args.GetConfigRepository(cmd, org, m)
	if err != nil {
		return err
	}

	update, err := cmd.Flags().GetBool("update")
	if err != nil {
		return err
	}

	output, err := cy_args.GetOutput(cmd)
	if err != nil {
		return errors.Wrap(err, "unable to get output flag")
	}

	// fetch the printer from the factory
	p, err := factory.GetPrinter(output)
	if err != nil {
		return errors.Wrap(err, "unable to get printer")
	}

	if update {
		current, err := m.GetProject(org, project)
		if err == nil {
			// Make the update use the current color if not explicitly set by the user
			if color == cy_args.DefaultColor {
				if current.Color != nil {
					color = *current.Color
				} else {
					// Use a random one if none is set
					color = cy_args.PickRandomColor(nil)
				}
			}

			if icon == cy_args.DefaultIcon {
				if current.Icon != nil {
					icon = *current.Icon
				} else {
					// Use a random one if none is set
					icon = cy_args.PickRandomIcon(nil)
				}
			}

			resp, err := m.UpdateProject(org, name, project, description, configRepository, owner, "", color, icon, "", current.UpdatedAt)
			if err != nil {
				return printer.SmartPrint(p, nil, err, "", printer.Options{}, cmd.OutOrStderr())
			}

			return printer.SmartPrint(p, resp, nil, "", printer.Options{}, cmd.OutOrStdout())
		}
	}

	if color == cy_args.DefaultColor {
		color = cy_args.PickRandomColor(nil)
	}

	if icon == cy_args.DefaultIcon {
		icon = cy_args.PickRandomIcon(nil)
	}

	resp, err := m.CreateProject(org, name, project, description, configRepository, owner, "", color, icon)
	if err != nil {
		return printer.SmartPrint(p, nil, err, "", printer.Options{}, cmd.OutOrStderr())
	}
	return printer.SmartPrint(p, resp, nil, "", printer.Options{}, cmd.OutOrStdout())
}
