package environments

import (
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/apiclient"
	"github.com/cycloidio/cycloid-cli/cmd/common"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
	"github.com/cycloidio/cycloid-cli/printer"
)

func NewCreateCommand() *cobra.Command {
	cmd := &cobra.Command{
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
	_ = cmd.Flags().MarkDeprecated(cyargs.AddColorFlag(cmd), "color now lives on environment-type and will be ignored")
	cyargs.AddUpdateFlag(cmd, "if set, will update the environment if it exists.")
	cyargs.AddNameFlag(cmd)
	cyargs.AddEnvFlag(cmd)
	cmd.MarkFlagsOneRequired("name", "env")
	return cmd
}

func create(cmd *cobra.Command, args []string) error {
	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}

	env, err := cyargs.GetEnvOrEmpty(cmd)
	if err != nil {
		return err
	}

	api := common.NewAPI()
	m := apiclient.NewMiddleware(api)

	result, err := createOrUpdateEnvironment(cmd, m, org, env, cyargs.GetUpdate(cmd))
	return cyout.PrintWithOptions(cmd, result, err, "failed to create environment", printer.Options{})
}
