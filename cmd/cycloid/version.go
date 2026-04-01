package cycloid

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/printer/factory"
)

func NewVersionCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "version",
		Args:  cobra.NoArgs,
		Short: "Get the version of the consumed API",
		Example: `
	# get the version in JSON format
	cy version -o json
`,
		RunE: getVersion,
	}
	return cmd
}

func getVersion(cmd *cobra.Command, args []string) error {
	output, err := cyargs.GetOutput(cmd)
	if err != nil {
		return err
	}

	p, err := factory.GetPrinter(output)
	if err != nil {
		return errors.Wrap(err, "unable to get printer")
	}

	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	d, _, err := m.GetAppVersion()
	if err != nil {
		return printer.SmartPrint(p, nil, err, "unable to get API version", printer.Options{}, cmd.ErrOrStderr())
	}

	return printer.SmartPrint(p, d, nil, "", printer.Options{}, cmd.OutOrStdout())
}
