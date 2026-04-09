package plugins

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/printer/factory"
)

func NewUpgradeCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "upgrade <id-or-name>",
		Aliases:           []string{"update"},
		Args:              cobra.ExactArgs(1),
		ValidArgsFunction: cyargs.CompletePluginInstallID,
		Short:             "[beta] Upgrade a plugin to a new version",
		Example: `
  cy beta plugin upgrade 42 --version-id 7 --config host=localhost
  cy beta plugin upgrade my-plugin --version-id 7 --config-file ./config.json
`,
		RunE: upgradePlugin,
	}

	cmd.Flags().Uint32("version-id", 0, "ID of the target plugin version (required)")
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

	output, err := cyargs.GetOutput(cmd)
	if err != nil {
		return errors.Wrap(err, "unable to get output flag")
	}

	p, err := factory.GetPrinter(output)
	if err != nil {
		return errors.Wrap(err, "unable to get printer")
	}

	result, _, err := m.UpdatePlugin(org, id, versionID, config)
	return printer.SmartPrint(p, result, err, "unable to upgrade plugin", printer.Options{}, cmd.OutOrStdout())
}
