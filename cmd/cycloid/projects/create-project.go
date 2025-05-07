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

func NewCreateCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "create",
		Short: "create a project",
		Example: `
	# create a project
	cy --org my-org project create \
		--name my-project \
		--description "an awesome project" \
		--owner "username" \
		--stack-ref my-stack-ref \
		--config-repo config-repo-canonical
`,
		RunE:    create,
		PreRunE: internal.CheckAPIAndCLIVersion,
	}
	common.RequiredFlag(WithFlagName, cmd)
	common.RequiredFlag(WithFlagStackRef, cmd)
	common.RequiredFlag(WithFlagConfigRepository, cmd)
	cmd.Flags().StringP("env", "e", "", "[deprecated] add an environment with project creation")
	cmd.Flags().MarkHidden("env")
	cmd.Flags().String("vars", "", "[deprecated] path to a variable file for the env creation")
	cmd.Flags().MarkHidden("vars")
	cmd.Flags().String("pipeline", "", "[deprecated] path to a pipeline file for the env creation")
	cmd.Flags().MarkHidden("pipeline")
	cmd.Flags().StringToString("config", nil, "[deprecated] path to a config file to inject in the config repo")
	cmd.Flags().MarkHidden("config")
	cmd.Flags().String("usecase", "", "[deprecated] the usecase for the env creation")
	cmd.Flags().MarkHidden("usecase")
	cmd.Flags().String("owner", "", "the owner username")
	cmd.Flags().String("team", "", "the team")
	cmd.Flags().String("icon", "", "icon of the project")
	cmd.Flags().String("color", "", "the name of the color")

	WithFlagDescription(cmd)
	WithFlagCanonical(cmd)
	return cmd
}

func create(cmd *cobra.Command, args []string) error {
	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	org, err := cy_args.GetOrg(cmd)
	if err != nil {
		return err
	}

	name, err := cmd.Flags().GetString("name")
	if err != nil {
		return err
	}

	canonical, err := cmd.Flags().GetString("canonical")
	if err != nil {
		return err
	}

	description, err := cmd.Flags().GetString("description")
	if err != nil {
		return err
	}

	configRepo, err := cmd.Flags().GetString("config-repo")
	if err != nil {
		return err
	}

	owner, err := cmd.Flags().GetString("owner")
	if err != nil {
		return err
	}

	team, err := cmd.Flags().GetString("team")
	if err != nil {
		return err
	}

	icon, err := cmd.Flags().GetString("icon")
	if err != nil {
		return err
	}

	color, err := cmd.Flags().GetString("color")
	if err != nil {
		return err
	}

	output, err := cmd.Flags().GetString("output")
	if err != nil {
		return errors.Wrap(err, "unable to get output flag")
	}

	// fetch the printer from the factory
	p, err := factory.GetPrinter(output)
	if err != nil {
		return errors.Wrap(err, "unable to get printer")
	}

	project, err := m.CreateProject(org, name, canonical, description, configRepo, owner, team, color, icon)
	return printer.SmartPrint(p, project, err, "", printer.Options{}, cmd.OutOrStdout())
}
