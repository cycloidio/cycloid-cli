package version

import (
	"fmt"
	"strconv"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
)

// resolveRegistryAndPlugin resolves args[0] (registry) and args[1] (plugin) to their numeric IDs.
func resolveRegistryAndPlugin(org string, args []string, m middleware.Middleware) (registryID, pluginID uint32, err error) {
	registryID, err = cyargs.ResolvePluginRegistryID(org, args[0], m)
	if err != nil {
		return 0, 0, err
	}
	pluginID, err = cyargs.ResolveRegistryPluginID(org, registryID, args[1], m)
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
