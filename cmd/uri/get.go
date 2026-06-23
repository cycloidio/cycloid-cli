package uri

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/interpolator"
	"github.com/cycloidio/cycloid-cli/interpolator/resolvers/httpresolver"
)

func NewGetCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get [uri]",
		Short: "Get Cycloid resources by using its URI",
		Long: strings.Join([]string{
			"Get a cycloid resource by its URI Path.",
			interpolator.Docs,
		}, "\n"),
		Args: cobra.ExactArgs(1),
		Example: `Fetch a credential by its canonical as JSON
  cy get cy://organizations/my-org/credentials/my-cred?output=json

Fetch an ssh_key
  cy get cy://organizations/my-org/credentials/my-ssh?key=.raw.ssh_key
`,
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
