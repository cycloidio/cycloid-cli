package projects

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/internal"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cy_args"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/printer/factory"
)

func NewUpdateCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:     "update",
		Short:   "update a project",
		Example: `cy --org my-org project update --project "my-project" --name "NewName"`,
		RunE:    update,
		PreRunE: internal.CheckAPIAndCLIVersion,
	}

	cy_args.AddNameFlag(cmd)
	cy_args.AddProjectFlag(cmd)
	cy_args.AddDescriptionFlag(cmd)
	cy_args.AddIconFlag(cmd)
	cy_args.AddColorFlag(cmd)
	cy_args.AddOwnerFlag(cmd)
	cy_args.AddConfigRepositoryFlag(cmd)
	return cmd
}

func update(cmd *cobra.Command, args []string) error {
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

	owner, err := cy_args.GetOwner(cmd)
	if err != nil {
		return err
	}

	color, err := cy_args.GetColor(cmd)
	if err != nil {
		return err
	}

	icon, err := cy_args.GetIcon(cmd)
	if err != nil {
		return err
	}

	configRepository, err := cy_args.GetConfigRepository(cmd, org, m)
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

	projectResp, err := m.GetProject(org, project)
	if err != nil {
		return printer.SmartPrint(p, projectResp, err, "project not found", printer.Options{}, cmd.OutOrStdout())
	}

	resp, err := m.UpdateProject(org, name, project, description, configRepository, owner, "", color, icon, "", projectResp.UpdatedAt)
	return printer.SmartPrint(p, resp, err, "", printer.Options{}, cmd.OutOrStdout())
}
