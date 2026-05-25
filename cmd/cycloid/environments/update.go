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
		Short:   "update an environment",
		Example: `cy --org my-org environment update --env my-environment --name "New Name"`,
		RunE:    update,
	}

	cyargs.AddEnvFlag(cmd)
	cmd.MarkFlagRequired("env")
	cyargs.AddNameFlag(cmd)
	cyargs.AddEnvironmentTypeFlag(cmd)
	cyargs.AddDescriptionFlag(cmd)
	cyargs.AddEnvironmentOwnerFlag(cmd)
	cyargs.AddCloudAccountCanonicalsFlag(cmd)
	cyargs.AddEnvironmentVariablesFlag(cmd)
	cyargs.AddEnvironmentVariablesFileFlag(cmd)
	cyargs.AddColorFlag(cmd)
	return cmd
}

func update(cmd *cobra.Command, args []string) error {
	warnDeprecatedColor(cmd)

	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}

	env, err := cyargs.GetEnv(cmd)
	if err != nil {
		return err
	}

	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	current, _, err := m.GetOrgEnv(org, env)
	if err != nil {
		return cyout.PrintWithOptions(cmd, nil, err, "environment not found", printer.Options{})
	}

	updateBody, err := buildUpdateEnvironment(cmd, current)
	if err != nil {
		return err
	}

	result, _, err := m.UpdateOrgEnv(org, env, updateBody)
	return cyout.PrintWithOptions(cmd, result, err, "failed to update environment", printer.Options{})
}
