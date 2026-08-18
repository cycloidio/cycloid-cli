package cmd

import (
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/apiclient"
	"github.com/cycloidio/cycloid-cli/cmd/common"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
	"github.com/cycloidio/cycloid-cli/printer"
)

func NewVersionCmd() *cobra.Command {
	cmd := &cobra.Command{
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
	api := common.NewAPI()
	m := apiclient.NewAPIClient(api)

	d, _, err := m.GetAppVersion()
	return cyout.PrintWithOptions(cmd, d, err, "unable to get API version", printer.Options{})
}
