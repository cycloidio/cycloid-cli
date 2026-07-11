package cyargs

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/apiclient"
	"github.com/cycloidio/cycloid-cli/cmd/common"
)

// ---------------------------------------------------------------------------
// Registry context flag (--registry, used by plugin and version subcommands)
// ---------------------------------------------------------------------------

// AddRegistryFlag registers a --registry flag and returns the flag name for use
// with MarkFlagRequired.
func AddRegistryFlag(cmd *cobra.Command) string {
	cmd.Flags().String("registry", "", "plugin registry name, ID, or URL")
	return "registry"
}

// GetRegistry reads the --registry flag value.
func GetRegistry(cmd *cobra.Command) (string, error) {
	return cmd.Flags().GetString("registry")
}

// ---------------------------------------------------------------------------
// Plugin context flag (--plugin, used by version subcommands)
// ---------------------------------------------------------------------------

// AddPluginFlag registers a --plugin flag and returns the flag name for use
// with MarkFlagRequired.
func AddPluginFlag(cmd *cobra.Command) string {
	cmd.Flags().String("plugin", "", "plugin name or ID within the registry")
	return "plugin"
}

// GetPlugin reads the --plugin flag value.
func GetPlugin(cmd *cobra.Command) (string, error) {
	return cmd.Flags().GetString("plugin")
}

// ---------------------------------------------------------------------------
// URL flag (registry add, manager create, version publish)
// ---------------------------------------------------------------------------

// AddURLFlag registers a --url flag with the given usage description.
func AddURLFlag(cmd *cobra.Command, usage string) {
	cmd.Flags().String("url", "", usage)
}

// GetURL returns the value of the --url flag.
func GetURL(cmd *cobra.Command) (string, error) {
	return cmd.Flags().GetString("url")
}

// ---------------------------------------------------------------------------
// Plugin version ID flag (for install / upgrade)
// ---------------------------------------------------------------------------

// AddPluginVersionIDFlag registers an optional --version-id uint32 flag.
func AddPluginVersionIDFlag(cmd *cobra.Command) {
	cmd.Flags().Uint32("version-id", 0, "ID of the target plugin version")
}

// GetPluginVersionID reads the --version-id flag.
// Returns nil if the flag was not set (zero value means unset).
func GetPluginVersionID(cmd *cobra.Command) (*uint32, error) {
	v, err := cmd.Flags().GetUint32("version-id")
	if err != nil {
		return nil, err
	}
	if v == 0 {
		return nil, nil
	}
	return &v, nil
}

// ---------------------------------------------------------------------------
// Docker image flag (for version publish)
// ---------------------------------------------------------------------------

// AddDockerImageFlag registers a --docker-image string flag.
func AddDockerImageFlag(cmd *cobra.Command) {
	cmd.Flags().String("docker-image", "", "Docker image reference for the plugin version (e.g. registry:5000/org/plugin:tag)")
}

// GetDockerImage returns the value of the --docker-image flag.
func GetDockerImage(cmd *cobra.Command) (string, error) {
	return cmd.Flags().GetString("docker-image")
}

// ---------------------------------------------------------------------------
// Plugin config flags (for install / upgrade)
// ---------------------------------------------------------------------------

// AddPluginConfigFlags registers --config key=val and --config-file path.json.
func AddPluginConfigFlags(cmd *cobra.Command) {
	cmd.Flags().StringToString("config", nil, "plugin configuration as key=value pairs. Can be specified multiple times.")
	cmd.Flags().StringP("config-file", "f", "", "path to a JSON file containing plugin configuration (map of string to string)")
	cmd.MarkFlagFilename("config-file", "json")
}

// GetPluginConfig reads and merges plugin configuration from --config-file (base)
// and --config key=val overrides.
func GetPluginConfig(cmd *cobra.Command) (map[string]string, error) {
	result := make(map[string]string)

	filename, err := cmd.Flags().GetString("config-file")
	if err != nil {
		return nil, err
	}
	if filename != "" {
		data, err := os.ReadFile(filename)
		if err != nil {
			return nil, fmt.Errorf("reading --config-file %q: %w", filename, err)
		}
		if err := json.Unmarshal(data, &result); err != nil {
			return nil, fmt.Errorf("parsing --config-file %q as JSON object: %w", filename, err)
		}
	}

	kvPairs, err := cmd.Flags().GetStringToString("config")
	if err != nil {
		return nil, err
	}
	for k, v := range kvPairs {
		result[k] = v
	}

	return result, nil
}

// ---------------------------------------------------------------------------
// Plugin install ID resolution + completion
// ---------------------------------------------------------------------------

// CompletePluginInstallID offers both numeric-ID and name-based completions.
// The shell inserts the selected value; RunE resolves names to IDs via ResolvePluginInstallID.
func CompletePluginInstallID(cmd *cobra.Command, args []string, toComplete string) ([]cobra.Completion, cobra.ShellCompDirective) {
	org, err := GetOrg(cmd)
	if err != nil {
		return cobra.AppendActiveHelp(nil, "missing org: "+err.Error()), cobra.ShellCompDirectiveNoFileComp
	}

	m := apiclient.NewAPIClient(common.NewAPI())
	plugins, _, err := m.ListPlugins(org)
	if err != nil {
		return cobra.AppendActiveHelp(nil, "failed to list plugins: "+err.Error()), cobra.ShellCompDirectiveNoFileComp
	}

	completions := make([]cobra.Completion, 0, len(plugins)*2)
	for _, p := range plugins {
		if p.Install == nil || p.Install.ID == nil {
			continue
		}
		idStr := strconv.Itoa(int(*p.Install.ID))
		status := ""
		if p.Install.Status != nil {
			status = *p.Install.Status
		}
		name := idStr
		if p.Name != nil {
			name = *p.Name
		}
		completions = append(completions,
			cobra.CompletionWithDesc(idStr, fmt.Sprintf("%s (%s)", name, status)),
			cobra.CompletionWithDesc(name, fmt.Sprintf("id:%s (%s)", idStr, status)),
		)
	}
	return completions, cobra.ShellCompDirectiveNoFileComp
}

// ResolvePluginInstallID resolves a numeric ID string or name to a uint32 plugin install ID.
// The install ID lives at Plugin.Install.ID (not Plugin.ID which is the registry plugin ID).
func ResolvePluginInstallID(org, nameOrID string, m apiclient.APIClient) (uint32, error) {
	if id, err := strconv.ParseUint(nameOrID, 10, 32); err == nil {
		return uint32(id), nil
	}
	plugins, _, err := m.ListPlugins(org)
	if err != nil {
		return 0, fmt.Errorf("listing plugins to resolve %q: %w", nameOrID, err)
	}
	var matches []uint32
	for _, p := range plugins {
		if p.Install == nil || p.Install.ID == nil {
			continue
		}
		if p.Name != nil && *p.Name == nameOrID {
			matches = append(matches, *p.Install.ID)
		}
	}
	return resolveUnique("plugin install", nameOrID, matches)
}

// ---------------------------------------------------------------------------
// Plugin manager ID resolution + completion
// ---------------------------------------------------------------------------

// CompletePluginManagerID offers both ID and name completions for plugin managers.
func CompletePluginManagerID(cmd *cobra.Command, args []string, toComplete string) ([]cobra.Completion, cobra.ShellCompDirective) {
	org, err := GetOrg(cmd)
	if err != nil {
		return cobra.AppendActiveHelp(nil, "missing org: "+err.Error()), cobra.ShellCompDirectiveNoFileComp
	}

	m := apiclient.NewAPIClient(common.NewAPI())
	managers, _, err := m.ListPluginManagers(org)
	if err != nil {
		return cobra.AppendActiveHelp(nil, "failed to list plugin managers: "+err.Error()), cobra.ShellCompDirectiveNoFileComp
	}

	completions := make([]cobra.Completion, 0, len(managers)*2)
	for _, pm := range managers {
		if pm.ID == nil || pm.Name == nil || pm.Status == nil {
			continue
		}
		idStr := strconv.Itoa(int(*pm.ID))
		name := *pm.Name
		status := *pm.Status
		completions = append(completions,
			cobra.CompletionWithDesc(idStr, fmt.Sprintf("%s (%s)", name, status)),
			cobra.CompletionWithDesc(name, fmt.Sprintf("id:%s (%s)", idStr, status)),
		)
	}
	return completions, cobra.ShellCompDirectiveNoFileComp
}

// ResolvePluginManagerID resolves a numeric ID, name, or URL to a plugin manager ID.
func ResolvePluginManagerID(org, nameOrID string, m apiclient.APIClient) (uint32, error) {
	if id, err := strconv.ParseUint(nameOrID, 10, 32); err == nil {
		return uint32(id), nil
	}
	managers, _, err := m.ListPluginManagers(org)
	if err != nil {
		return 0, fmt.Errorf("listing plugin managers to resolve %q: %w", nameOrID, err)
	}
	var matches []uint32
	for _, pm := range managers {
		if pm.ID == nil {
			continue
		}
		if (pm.Name != nil && *pm.Name == nameOrID) ||
			(pm.URL != nil && pm.URL.String() == nameOrID) {
			matches = append(matches, *pm.ID)
		}
	}
	return resolveUnique("plugin manager", nameOrID, matches)
}

// ---------------------------------------------------------------------------
// Plugin registry ID resolution + completion
// ---------------------------------------------------------------------------

// CompletePluginRegistryID offers ID, name, and URL completions for plugin registries.
func CompletePluginRegistryID(cmd *cobra.Command, args []string, toComplete string) ([]cobra.Completion, cobra.ShellCompDirective) {
	org, err := GetOrg(cmd)
	if err != nil {
		return cobra.AppendActiveHelp(nil, "missing org: "+err.Error()), cobra.ShellCompDirectiveNoFileComp
	}

	m := apiclient.NewAPIClient(common.NewAPI())
	registries, _, err := m.ListPluginRegistries(org)
	if err != nil {
		return cobra.AppendActiveHelp(nil, "failed to list plugin registries: "+err.Error()), cobra.ShellCompDirectiveNoFileComp
	}

	completions := make([]cobra.Completion, 0, len(registries)*3)
	for _, r := range registries {
		if r.ID == nil || r.Name == nil || r.Status == nil {
			continue
		}
		idStr := strconv.Itoa(int(*r.ID))
		name := *r.Name
		status := *r.Status
		urlStr := ""
		if r.URL != nil {
			urlStr = r.URL.String()
		}
		completions = append(completions,
			cobra.CompletionWithDesc(idStr, fmt.Sprintf("%s (%s)", name, status)),
			cobra.CompletionWithDesc(name, fmt.Sprintf("id:%s (%s)", idStr, status)),
		)
		if urlStr != "" {
			completions = append(completions,
				cobra.CompletionWithDesc(urlStr, fmt.Sprintf("%s id:%s (%s)", name, idStr, status)),
			)
		}
	}
	return completions, cobra.ShellCompDirectiveNoFileComp
}

// ResolvePluginRegistryID resolves a numeric ID, name, or URL to a plugin registry ID.
// URL matching is tried when the value starts with "http".
func ResolvePluginRegistryID(org, nameOrID string, m apiclient.APIClient) (uint32, error) {
	if id, err := strconv.ParseUint(nameOrID, 10, 32); err == nil {
		return uint32(id), nil
	}
	registries, _, err := m.ListPluginRegistries(org)
	if err != nil {
		return 0, fmt.Errorf("listing plugin registries to resolve %q: %w", nameOrID, err)
	}
	isURL := strings.HasPrefix(nameOrID, "http")
	var matches []uint32
	for _, r := range registries {
		if r.ID == nil {
			continue
		}
		if isURL {
			if r.URL != nil && r.URL.String() == nameOrID {
				matches = append(matches, *r.ID)
			}
		} else {
			if r.Name != nil && *r.Name == nameOrID {
				matches = append(matches, *r.ID)
			}
		}
	}
	return resolveUnique("plugin registry", nameOrID, matches)
}

// CompletePluginIDFromRegistryFlag completes plugin names/IDs by reading the registry
// from the --registry flag rather than positional args.
func CompletePluginIDFromRegistryFlag(cmd *cobra.Command, args []string, toComplete string) ([]cobra.Completion, cobra.ShellCompDirective) {
	registryStr, _ := cmd.Flags().GetString("registry")
	if registryStr == "" {
		return cobra.AppendActiveHelp(nil, "provide --registry first"), cobra.ShellCompDirectiveNoFileComp
	}
	org, err := GetOrg(cmd)
	if err != nil {
		return cobra.AppendActiveHelp(nil, "missing org: "+err.Error()), cobra.ShellCompDirectiveNoFileComp
	}
	m := apiclient.NewAPIClient(common.NewAPI())
	registryID, err := ResolvePluginRegistryID(org, registryStr, m)
	if err != nil {
		return cobra.AppendActiveHelp(nil, "failed to resolve registry: "+err.Error()), cobra.ShellCompDirectiveNoFileComp
	}
	plugins, _, err := m.ListRegistryPlugins(org, registryID)
	if err != nil {
		return cobra.AppendActiveHelp(nil, "failed to list plugins: "+err.Error()), cobra.ShellCompDirectiveNoFileComp
	}
	completions := make([]cobra.Completion, 0, len(plugins)*2)
	for _, p := range plugins {
		if p.ID == nil || p.Name == nil {
			continue
		}
		idStr := strconv.Itoa(int(*p.ID))
		completions = append(completions,
			cobra.CompletionWithDesc(idStr, *p.Name),
			cobra.CompletionWithDesc(*p.Name, "id:"+idStr),
		)
	}
	return completions, cobra.ShellCompDirectiveNoFileComp
}

// CompletePluginVersionID completes version IDs by reading registry and plugin from flags.
func CompletePluginVersionID(cmd *cobra.Command, args []string, toComplete string) ([]cobra.Completion, cobra.ShellCompDirective) {
	registryStr, _ := cmd.Flags().GetString("registry")
	pluginStr, _ := cmd.Flags().GetString("plugin")
	if registryStr == "" || pluginStr == "" {
		return cobra.AppendActiveHelp(nil, "provide --registry and --plugin first"), cobra.ShellCompDirectiveNoFileComp
	}
	org, err := GetOrg(cmd)
	if err != nil {
		return cobra.AppendActiveHelp(nil, "missing org: "+err.Error()), cobra.ShellCompDirectiveNoFileComp
	}
	m := apiclient.NewAPIClient(common.NewAPI())
	registryID, err := ResolvePluginRegistryID(org, registryStr, m)
	if err != nil {
		return cobra.AppendActiveHelp(nil, "failed to resolve registry: "+err.Error()), cobra.ShellCompDirectiveNoFileComp
	}
	pluginID, err := ResolveRegistryPluginID(org, registryID, pluginStr, m)
	if err != nil {
		return cobra.AppendActiveHelp(nil, "failed to resolve plugin: "+err.Error()), cobra.ShellCompDirectiveNoFileComp
	}
	versions, _, err := m.ListPluginVersions(org, registryID, pluginID)
	if err != nil {
		return cobra.AppendActiveHelp(nil, "failed to list versions: "+err.Error()), cobra.ShellCompDirectiveNoFileComp
	}
	completions := make([]cobra.Completion, 0, len(versions))
	for _, v := range versions {
		if v.ID == nil {
			continue
		}
		idStr := strconv.Itoa(int(*v.ID))
		desc := ""
		if v.Status != nil {
			desc = *v.Status
		}
		completions = append(completions, cobra.CompletionWithDesc(idStr, desc))
	}
	return completions, cobra.ShellCompDirectiveNoFileComp
}

// CompleteRegistryPluginID offers ID and name completions for plugins within a registry.
// The registry must already be resolved from args[0].
func CompleteRegistryPluginID(cmd *cobra.Command, args []string, toComplete string) ([]cobra.Completion, cobra.ShellCompDirective) {
	if len(args) < 1 {
		return cobra.AppendActiveHelp(nil, "registry ID or name required first"), cobra.ShellCompDirectiveNoFileComp
	}
	org, err := GetOrg(cmd)
	if err != nil {
		return cobra.AppendActiveHelp(nil, "missing org: "+err.Error()), cobra.ShellCompDirectiveNoFileComp
	}

	m := apiclient.NewAPIClient(common.NewAPI())
	registryID, err := ResolvePluginRegistryID(org, args[0], m)
	if err != nil {
		return cobra.AppendActiveHelp(nil, "failed to resolve registry: "+err.Error()), cobra.ShellCompDirectiveNoFileComp
	}

	plugins, _, err := m.ListRegistryPlugins(org, registryID)
	if err != nil {
		return cobra.AppendActiveHelp(nil, "failed to list registry plugins: "+err.Error()), cobra.ShellCompDirectiveNoFileComp
	}

	completions := make([]cobra.Completion, 0, len(plugins)*2)
	for _, p := range plugins {
		if p.ID == nil || p.Name == nil {
			continue
		}
		idStr := strconv.Itoa(int(*p.ID))
		name := *p.Name
		completions = append(completions,
			cobra.CompletionWithDesc(idStr, name),
			cobra.CompletionWithDesc(name, "id:"+idStr),
		)
	}
	return completions, cobra.ShellCompDirectiveNoFileComp
}

// ResolveRegistryPluginID resolves a numeric ID or name to a plugin ID within a registry.
func ResolveRegistryPluginID(org string, registryID uint32, nameOrID string, m apiclient.APIClient) (uint32, error) {
	if id, err := strconv.ParseUint(nameOrID, 10, 32); err == nil {
		return uint32(id), nil
	}
	plugins, _, err := m.ListRegistryPlugins(org, registryID)
	if err != nil {
		return 0, fmt.Errorf("listing registry plugins to resolve %q: %w", nameOrID, err)
	}
	var matches []uint32
	for _, p := range plugins {
		if p.ID == nil {
			continue
		}
		if p.Name != nil && *p.Name == nameOrID {
			matches = append(matches, *p.ID)
		}
	}
	return resolveUnique("registry plugin", nameOrID, matches)
}

// ---------------------------------------------------------------------------
// Internal helpers
// ---------------------------------------------------------------------------

// noMatchPrefix is the sentinel prefix used by resolveUnique for not-found errors.
const noMatchPrefix = "no "

// resolveUnique returns the single ID from matches, or a descriptive error.
// Not-found errors are detectable via IsNoMatchError.
func resolveUnique(kind, nameOrID string, matches []uint32) (uint32, error) {
	switch len(matches) {
	case 0:
		return 0, fmt.Errorf(noMatchPrefix+"%s found matching %q; use a numeric ID or check spelling", kind, nameOrID)
	case 1:
		return matches[0], nil
	default:
		ids := make([]string, len(matches))
		for i, id := range matches {
			ids[i] = strconv.Itoa(int(id))
		}
		return 0, fmt.Errorf(
			"%q matches %d %ss (IDs: %s); use a numeric ID to disambiguate",
			nameOrID, len(matches), kind, strings.Join(ids, ", "),
		)
	}
}

// IsNoMatchError reports whether err is a "not found" error from resolveUnique.
func IsNoMatchError(err error) bool {
	if err == nil {
		return false
	}
	return strings.HasPrefix(err.Error(), noMatchPrefix)
}
