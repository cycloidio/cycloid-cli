package externalbackends

import (
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/apiclient"
	"github.com/cycloidio/cycloid-cli/cmd/common"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
	"github.com/cycloidio/cycloid-cli/printer"
)

var (
	example = `
	# List all the external backends within my-org organization in JSON output format
	cy --org my-org external-backends list --output=json

	# List all the external backends within my-org organization in YAML output format
	cy --org my-org external-backends list --output=yaml
`
	short = "Get the list of organization external backends"
	long  = short
)

func NewListCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "list",
		Args:    cobra.NoArgs,
		Example: example,
		Short:   short,
		Long:    long,
		RunE:    list,
	}

	return cmd
}

func list(cmd *cobra.Command, args []string) error {
	api := common.NewAPI()
	m := apiclient.NewMiddleware(api)

	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}

	ebs, _, err := m.ListExternalBackends(org)
	return cyout.PrintWithOptions(cmd, ebs, err, "unable to list external backends", printer.Options{})
}
