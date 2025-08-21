package cyargs

import (
	"strings"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/spf13/cobra"
)

func AddOrgNameFlag(cmd *cobra.Command) string {
	flagName := "name"
	cmd.Flags().StringP(flagName, "n", "", "the organization's name")
	return flagName
}

func GetOrgName(cmd *cobra.Command) (string, error) {
	return cmd.Flags().GetString("name")
}

func CompleteOrg(cmd *cobra.Command, args []string, toComplete string) ([]cobra.Completion, cobra.ShellCompDirective) {
	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	org, err := GetOrg(cmd)
	if err != nil {
		return []string{}, cobra.ShellCompDirectiveNoFileComp
	}

	orgs, err := m.ListOrganizationChildrens(org)
	if err != nil {
		return cobra.AppendActiveHelp(nil, "failed to list orgs for completion: "+err.Error()), cobra.ShellCompDirectiveNoFileComp
	}

	completions := make([]cobra.Completion, len(orgs))
	for index, org := range orgs {
		if org.Canonical != nil && strings.HasPrefix(*org.Canonical, toComplete) {
			completions[index] = cobra.CompletionWithDesc(*org.Canonical, *org.Name)
		}
	}

	return completions, cobra.ShellCompDirectiveNoFileComp
}

func AddOrgChildOfFlag(cmd *cobra.Command) string {
	flagName := "parent-canonical"
	cmd.Flags().String(flagName, "p", "the parent organization canonical")
	cmd.RegisterFlagCompletionFunc(flagName, CompleteOrg)
	return flagName
}

func GetOrgParentCanonical(cmd *cobra.Command) (string, error) {
	return cmd.Flags().GetString("parent-canonical")
}
