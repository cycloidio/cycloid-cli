package uri

import (
	"fmt"
	"strings"

	"github.com/cycloidio/cycloid-cli/interpolator"
	"github.com/cycloidio/cycloid-cli/interpolator/resolvers/httpresolver"
	"github.com/spf13/cobra"
)

func NewGetCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get [uris...]",
		Short: "Get Cycloid resources by using its URI",
		Long: strings.Join([]string{
			"Get a cycloid resource by its URI Path.",
			interpolator.Docs,
		}, "\n"),
		Args: cobra.MinimumNArgs(1),
		Example: strings.Join([]string{
			"Fetch a credential by its canonical as JSON\n",
			"  cy get cy://organizations/<org>/credentials/<canonical>?format=json",
		}, "\n"),
		RunE: get,
	}

	return cmd
}

func get(cmd *cobra.Command, args []string) error {
	resolver, err := httpresolver.NewHTTPResolver()
	if err != nil {
		return err
	}

	out, err := resolver.Interpolate(args[0])
	if err != nil {
		return err
	}

	fmt.Fprint(cmd.OutOrStdout(), out)
	return nil
}
