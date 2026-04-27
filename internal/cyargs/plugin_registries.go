package cyargs

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
)

// AddPluginRegistryIDFlag registers --registry-id for plugin registry commands.
func AddPluginRegistryIDFlag(cmd *cobra.Command) string {
	const flagName = "registry-id"
	cmd.Flags().Uint32(flagName, 0, "plugin registry id")
	cmd.RegisterFlagCompletionFunc(flagName, CompletePluginRegistryID)
	return flagName
}

func GetPluginRegistryID(cmd *cobra.Command) (uint32, error) {
	return cmd.Flags().GetUint32("registry-id")
}

func CompletePluginRegistryID(cmd *cobra.Command, args []string, toComplete string) ([]cobra.Completion, cobra.ShellCompDirective) {
	org, err := GetOrg(cmd)
	if err != nil {
		return cobra.AppendActiveHelp(nil, "missing org for completion: "+err.Error()),
			cobra.ShellCompDirectiveNoFileComp
	}

	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	registries, _, err := m.ListPluginRegistries(org)
	if err != nil {
		return cobra.AppendActiveHelp(nil, "failed to list plugin registries for completion: "+err.Error()),
			cobra.ShellCompDirectiveNoFileComp
	}

	completions := make([]cobra.Completion, 0, len(registries))
	for _, r := range registries {
		if r.ID == nil {
			continue
		}
		idStr := strconv.Itoa(int(*r.ID))
		if strings.HasPrefix(idStr, toComplete) || toComplete == "" {
			completions = append(completions, cobra.CompletionWithDesc(idStr,
				fmt.Sprintf("%s (%s)", *r.Name, *r.Status),
			))
		}
	}

	return completions, cobra.ShellCompDirectiveNoFileComp
}
