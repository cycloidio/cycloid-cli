package cyargs

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
)

// AddPluginIDFlag registers --plugin-id for plugin commands.
func AddPluginIDFlag(cmd *cobra.Command) string {
	const flagName = "plugin-id"
	cmd.Flags().Uint32(flagName, 0, "plugin id")
	return flagName
}

func GetPluginID(cmd *cobra.Command) (uint32, error) {
	return cmd.Flags().GetUint32("plugin-id")
}

// AddPluginVersionIDFlag registers --version-id for plugin version commands.
func AddPluginVersionIDFlag(cmd *cobra.Command) string {
	const flagName = "version-id"
	cmd.Flags().Uint32(flagName, 0, "plugin version id")
	return flagName
}

func GetPluginVersionID(cmd *cobra.Command) (uint32, error) {
	return cmd.Flags().GetUint32("version-id")
}

// AddPluginInstallIDFlag registers --install-id for organization plugin commands.
func AddPluginInstallIDFlag(cmd *cobra.Command) string {
	const flagName = "install-id"
	cmd.Flags().Uint32(flagName, 0, "plugin install id")
	cmd.RegisterFlagCompletionFunc(flagName, CompletePluginInstallID)
	return flagName
}

func GetPluginInstallID(cmd *cobra.Command) (uint32, error) {
	return cmd.Flags().GetUint32("install-id")
}

func CompletePluginInstallID(cmd *cobra.Command, args []string, toComplete string) ([]cobra.Completion, cobra.ShellCompDirective) {
	org, err := GetOrg(cmd)
	if err != nil {
		return cobra.AppendActiveHelp(nil, "missing org for completion: "+err.Error()),
			cobra.ShellCompDirectiveNoFileComp
	}

	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	plugins, _, err := m.ListPlugins(org)
	if err != nil {
		return cobra.AppendActiveHelp(nil, "failed to list plugins for completion: "+err.Error()),
			cobra.ShellCompDirectiveNoFileComp
	}

	completions := make([]cobra.Completion, 0, len(plugins))
	for _, p := range plugins {
		if p.Install == nil || p.Install.ID == nil {
			continue
		}
		idStr := strconv.Itoa(int(*p.Install.ID))
		if strings.HasPrefix(idStr, toComplete) || toComplete == "" {
			completions = append(completions, cobra.CompletionWithDesc(idStr,
				fmt.Sprintf("%s", *p.Name),
			))
		}
	}

	return completions, cobra.ShellCompDirectiveNoFileComp
}
