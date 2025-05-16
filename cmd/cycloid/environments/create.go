package environments

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
		Use:     "create",
		Short:   "create a environment",
		Example: `cy --org my-org environment create --env "my-environment"`,
		RunE:    create,
		PreRunE: internal.CheckAPIAndCLIVersion,
	}

	cy_args.AddNameFlag(cmd)
	cy_args.AddProjectFlag(cmd)
	cy_args.AddEnvFlag(cmd)
	cy_args.AddColorFlag(cmd)
	cmd.Flags().Bool("update", false, "if set, will update the environment if it exists.")
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

	env, err := cy_args.GetEnv(cmd)
	if err != nil {
		return err
	}

	name, err := cy_args.GetName(cmd)
	if err != nil {
		return err
	}
	if name == "" {
		name = env
	}

	color, err := cy_args.GetColor(cmd)
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
		current, err := m.GetEnv(org, project, env)
		if err == nil {
			// Make the update use the current color if not explicitly set by the user
			if color == cy_args.DefaultColor {
				if current.Color != nil {
					color = *current.Color
				} else {
					// Use a random one if none is set
					color = cy_args.PickRandomColor(&env)
				}
			}

			resp, err := m.UpdateEnv(org, project, env, name, color)
			if err != nil {
				return printer.SmartPrint(p, nil, err, "", printer.Options{}, cmd.OutOrStderr())
			}
			return printer.SmartPrint(p, resp, err, "", printer.Options{}, cmd.OutOrStdout())
		}
	}

	if color == cy_args.DefaultColor {
		color = cy_args.PickRandomColor(&env)
	}

	resp, err := m.CreateEnv(org, project, env, name, color)
	return printer.SmartPrint(p, resp, err, "", printer.Options{}, cmd.OutOrStdout())
}
