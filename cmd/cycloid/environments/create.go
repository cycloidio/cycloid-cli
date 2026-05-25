package environments

import (
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
	"github.com/cycloidio/cycloid-cli/printer"
)

func NewCreateCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:     "create",
		Args:    cobra.NoArgs,
		Short:   "create an environment",
		Example: `cy --org my-org environment create --env my-environment --name "My Environment"`,
		RunE:    create,
	}

	cyargs.AddEnvironmentTypeFlag(cmd)
	cyargs.AddDescriptionFlag(cmd)
	cyargs.AddEnvironmentOwnerFlag(cmd)
	cyargs.AddCloudAccountCanonicalsFlag(cmd)
	cyargs.AddEnvironmentVariablesFlag(cmd)
	cyargs.AddEnvironmentVariablesFileFlag(cmd)
	cyargs.AddColorFlag(cmd)
	cmd.Flags().Bool("update", false, "if set, will update the environment if it exists.")
	cyargs.AddNameFlag(cmd)
	cyargs.AddEnvFlag(cmd)
	cmd.MarkFlagsOneRequired("name", "env")
	return cmd
}

func create(cmd *cobra.Command, args []string) error {
	warnDeprecatedColor(cmd)

	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}

	env, err := cyargs.GetEnvOrEmpty(cmd)
	if err != nil {
		return err
	}

	update, err := cmd.Flags().GetBool("update")
	if err != nil {
		return err
	}

	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	result, err := createOrUpdateEnvironment(cmd, m, org, env, update)
	return cyout.PrintWithOptions(cmd, result, err, "failed to create environment", printer.Options{})
}
