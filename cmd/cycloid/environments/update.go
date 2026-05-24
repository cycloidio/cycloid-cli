package environments

import (
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
	"github.com/cycloidio/cycloid-cli/printer"
)

func NewUpdateCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:     "update",
		Args:    cobra.NoArgs,
		Short:   "update a environment",
		Example: `cy --org my-org environment update --env"my-environment" --name "NewName"`,
		RunE:    update,
	}

	cyargs.AddNameFlag(cmd)
	cyargs.AddProjectFlag(cmd)
	cyargs.AddEnvFlag(cmd)
	cyargs.AddColorFlag(cmd)
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

	env, err := cyargs.GetEnv(cmd)
	if err != nil {
		return err
	}

	name, err := cyargs.GetName(cmd)
	if err != nil {
		return err
	}

	color, err := cyargs.GetColor(cmd)
	if err != nil {
		return err
	}

	currentEnv, _, err := m.GetEnv(org, project, env)
	if err != nil {
		return cyout.PrintWithOptions(cmd, currentEnv, err, "environment not found", printer.Options{})
	}

	// Make the update use the current color if not explicitly set by the user
	if color == cyargs.DefaultColor {
		if currentEnv.Color != nil {
			color = *currentEnv.Color
		} else {
			// Use a random one if none is set
			color = cyargs.PickRandomColor(&env)
		}
	}

	if name == "" {
		name = currentEnv.Name
	}

	resp, _, err := m.UpdateEnv(org, project, env, name, color)
	return cyout.PrintWithOptions(cmd, resp, err, "", printer.Options{})
}
