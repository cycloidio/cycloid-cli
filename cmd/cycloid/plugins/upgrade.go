package plugins

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
	"github.com/cycloidio/cycloid-cli/printer"
)

func NewUpgradeCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "upgrade <id-or-name>",
		Aliases:           []string{"update"},
		Args:              cobra.ExactArgs(1),
		ValidArgsFunction: cyargs.CompletePluginInstallID,
		Short:             "Upgrade a plugin to a new version",
		Example: `
  cy plugin upgrade 42 --version-id 7 --config host=localhost
  cy plugin upgrade my-plugin --version-id 7 --config-file ./config.json
`,
		RunE: upgradePlugin,
	}

	cyargs.AddPluginVersionIDFlag(cmd)
	cmd.MarkFlagRequired("version-id")
	cyargs.AddPluginConfigFlags(cmd)
	return cmd
}

func upgradePlugin(cmd *cobra.Command, args []string) error {
	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}

	id, err := cyargs.ResolvePluginInstallID(org, args[0], m)
	if err != nil {
		return err
	}

	versionID, err := cmd.Flags().GetUint32("version-id")
	if err != nil {
		return errors.Wrap(err, "unable to get --version-id flag")
	}

	config, err := cyargs.GetPluginConfig(cmd)
	if err != nil {
		return errors.Wrap(err, "unable to get plugin configuration")
	}

	result, _, err := m.UpdatePlugin(org, id, versionID, config)
	return cyout.PrintWithOptions(cmd, result, err, "unable to upgrade plugin", printer.Options{})
}
