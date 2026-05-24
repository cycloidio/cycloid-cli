package version

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
)

// resolveRegistryAndPlugin resolves --registry and --plugin flags to their numeric IDs.
func resolveRegistryAndPlugin(org string, cmd *cobra.Command, m middleware.Middleware) (registryID, pluginID uint32, err error) {
	registryStr, err := cyargs.GetRegistry(cmd)
	if err != nil {
		return 0, 0, err
	}
	registryID, err = cyargs.ResolvePluginRegistryID(org, registryStr, m)
	if err != nil {
		return 0, 0, err
	}
	pluginStr, err := cyargs.GetPlugin(cmd)
	if err != nil {
		return 0, 0, err
	}
	pluginID, err = cyargs.ResolveRegistryPluginID(org, registryID, pluginStr, m)
	if err != nil {
		return 0, 0, err
	}
	return registryID, pluginID, nil
}

// parseVersionID parses a version ID string to uint32.
func parseVersionID(s string) (uint32, error) {
	n, err := strconv.ParseUint(s, 10, 32)
	if err != nil {
		return 0, fmt.Errorf("invalid version ID %q: must be a positive integer", s)
	}
	return uint32(n), nil
}
