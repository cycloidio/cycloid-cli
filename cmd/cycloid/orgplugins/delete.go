package orgplugins

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/printer/factory"
)

func NewDeleteCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "delete",
		Args:  cobra.NoArgs,
		Short: "delete an installed plugin",
		Example: `
	# Delete installed plugin with install ID 5
	cy --org my-org plugin delete --install-id 5
`,
		RunE: deletePlugin,
	}

	cmd.MarkFlagRequired(cyargs.AddPluginInstallIDFlag(cmd))

	return cmd
}

func deletePlugin(cmd *cobra.Command, args []string) error {
	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}
	installID, err := cyargs.GetPluginInstallID(cmd)
	if err != nil {
		return err
	}
	output, err := cyargs.GetOutput(cmd)
	if err != nil {
		return errors.Wrap(err, "unable to get output flag")
	}

	p, err := factory.GetPrinter(output)
	if err != nil {
		return errors.Wrap(err, "unable to get printer")
	}

	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	_, err = m.DeletePlugin(org, installID)
	return printer.SmartPrint(p, nil, err, "unable to delete plugin", printer.Options{}, cmd.OutOrStdout())
}
