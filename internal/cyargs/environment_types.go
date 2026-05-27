package cyargs

import (
	"strings"

	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
)

func AddEnvironmentTypeNameFlag(cmd *cobra.Command) string {
	flagName := "environment-type-name"
	cmd.Flags().String(flagName, "", "display name of the environment type")
	return flagName
}

func GetEnvironmentTypeName(cmd *cobra.Command) (string, error) {
	return cmd.Flags().GetString("environment-type-name")
}

func CompleteEnvironmentTypeCanonical(cmd *cobra.Command, args []string, toComplete string) ([]cobra.Completion, cobra.ShellCompDirective) {
	org, err := GetOrg(cmd)
	if err != nil {
		return cobra.AppendActiveHelp(nil, "missing org parameter for completion"),
			cobra.ShellCompDirectiveNoFileComp
	}

	api := common.NewAPI()
	m := middleware.NewMiddleware(api)
	types, _, err := m.ListEnvironmentTypes(org)
	if err != nil {
		return cobra.AppendActiveHelp(nil, "failed to list environment types: "+err.Error()),
			cobra.ShellCompDirectiveNoFileComp
	}

	completions := make([]cobra.Completion, 0, len(types))
	for _, envType := range types {
		if envType.Canonical != nil && strings.HasPrefix(*envType.Canonical, toComplete) {
			name := ""
			if envType.Name != nil {
				name = *envType.Name
			}
			completions = append(completions, cobra.CompletionWithDesc(*envType.Canonical, name))
		}
	}
	return completions, cobra.ShellCompDirectiveNoFileComp
}
