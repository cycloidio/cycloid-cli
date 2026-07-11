package environments

import (
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/apiclient"
	"github.com/cycloidio/cycloid-cli/cmd/common"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
	"github.com/cycloidio/cycloid-cli/printer"
)

func NewUpdateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "update",
		Args:    cobra.NoArgs,
		Short:   "update an environment",
		Example: `cy --org my-org environment update --env my-environment --name "New Name"`,
		RunE:    update,
	}

	cyargs.AddEnvFlag(cmd)
	_ = cmd.MarkFlagRequired("env")
	cyargs.AddNameFlag(cmd)
	cyargs.AddEnvironmentTypeFlag(cmd)
	cyargs.AddDescriptionFlag(cmd)
	cyargs.AddEnvironmentOwnerFlag(cmd)
	cyargs.AddCloudAccountCanonicalsFlag(cmd)
	cyargs.AddEnvironmentVariablesFlag(cmd)
	cyargs.AddEnvironmentVariablesFileFlag(cmd)
	_ = cmd.Flags().MarkDeprecated(cyargs.AddColorFlag(cmd), "color now lives on environment-type and will be ignored")
	return cmd
}

func update(cmd *cobra.Command, args []string) error {
	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}

	env, err := cyargs.GetEnv(cmd)
	if err != nil {
		return err
	}

	api := common.NewAPI()
	m := apiclient.NewAPIClient(api)

	result, err := createOrUpdateEnvironment(cmd, m, org, env, true)
	return cyout.PrintWithOptions(cmd, result, err, "failed to update environment", printer.Options{})
}
