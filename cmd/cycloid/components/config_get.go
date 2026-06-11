package components

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
	"github.com/cycloidio/cycloid-cli/printer"
)

func NewComponentConfigGetCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "get",
		Args:    cobra.NoArgs,
		Short:   "Fetch the current Stackforms variables of a component in JSON format.",
		RunE:    getComponentConfig,
		Example: "cy config get -p project -e env -c component",
	}
	cyargs.AddCyContext(cmd)
	cyargs.AddStackVersionFlags(cmd)
	return cmd
}

func getComponentConfig(cmd *cobra.Command, args []string) error {
	org, project, env, component, err := cyargs.GetCyContext(cmd)
	if err != nil {
		return err
	}

	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	// version:<id> passes the catalog version ID directly to the API, bypassing
	// stack ref resolution. Useful when debugging a known version ID.
	var tag, branch, hash string
	var versionID uint32
	rawVersion, _ := cmd.Flags().GetString("stack-version")
	if idStr, ok := strings.CutPrefix(rawVersion, "version:"); ok {
		id64, err := strconv.ParseUint(idStr, 10, 32)
		if err != nil {
			return fmt.Errorf("--stack-version=version:<id>: expected a numeric ID, got %q", idStr)
		}
		versionID = uint32(id64)
	} else {
		tag, branch, hash, err = cyargs.ResolveStackVersionArg(cmd, m, org, "")
		if err != nil {
			return err
		}
	}

	config, _, err := m.GetComponentConfig(org, project, env, component, tag, branch, hash, versionID)
	return cyout.PrintWithOptions(cmd, config, err, "failed to fetch config of component '"+component+"'", printer.Options{})
}
