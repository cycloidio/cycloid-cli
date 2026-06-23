package cyargs

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/apiclient"
	"github.com/cycloidio/cycloid-cli/cmd/common"
)

// AddMemberIDFlag registers --id for member get/delete/update commands.
func AddMemberIDFlag(cmd *cobra.Command) string {
	const flagName = "id"
	cmd.Flags().Uint32(flagName, 0, "member id")
	cmd.RegisterFlagCompletionFunc(flagName, CompleteMemberID)
	return flagName
}

func GetMemberID(cmd *cobra.Command) (uint32, error) {
	return cmd.Flags().GetUint32("id")
}

// AddMemberEmailFlag registers --email for member invite commands.
func AddMemberEmailFlag(cmd *cobra.Command) string {
	const flagName = "email"
	cmd.Flags().String(flagName, "", "member email address")
	cmd.RegisterFlagCompletionFunc(flagName, CompleteMemberEmail)
	return flagName
}

func GetMemberEmail(cmd *cobra.Command) (string, error) {
	return cmd.Flags().GetString("email")
}

// AddMemberRoleFlag registers --role for member invite/update commands.
// Completion lists available role canonicals.
func AddMemberRoleFlag(cmd *cobra.Command) string {
	const flagName = "role"
	cmd.Flags().String(flagName, "", "role canonical to assign to the member")
	cmd.RegisterFlagCompletionFunc(flagName, CompleteRoleCanonical)
	return flagName
}

func GetMemberRole(cmd *cobra.Command) (string, error) {
	return cmd.Flags().GetString("role")
}

func CompleteMemberID(cmd *cobra.Command, args []string, toComplete string) ([]cobra.Completion, cobra.ShellCompDirective) {
	org, err := GetOrg(cmd)
	if err != nil {
		return cobra.AppendActiveHelp(nil, "missing org for completion: "+err.Error()),
			cobra.ShellCompDirectiveNoFileComp
	}

	api := common.NewAPI()
	m := apiclient.NewMiddleware(api)

	members, _, err := m.ListMembers(org)
	if err != nil {
		return cobra.AppendActiveHelp(nil, "failed to list members for completion: "+err.Error()),
			cobra.ShellCompDirectiveNoFileComp
	}

	completions := make([]cobra.Completion, 0, len(members))
	for _, member := range members {
		if member.ID == nil {
			continue
		}
		idStr := strconv.Itoa(int(*member.ID))
		if strings.HasPrefix(idStr, toComplete) || toComplete == "" {
			completions = append(completions, cobra.CompletionWithDesc(idStr,
				fmt.Sprintf("(%s)", member.Email.String()),
			))
		}
	}

	return completions, cobra.ShellCompDirectiveNoFileComp
}

func CompleteMemberEmail(cmd *cobra.Command, args []string, toComplete string) ([]cobra.Completion, cobra.ShellCompDirective) {
	org, err := GetOrg(cmd)
	if err != nil {
		return cobra.AppendActiveHelp(nil, "missing org for completion: "+err.Error()),
			cobra.ShellCompDirectiveNoFileComp
	}

	api := common.NewAPI()
	m := apiclient.NewMiddleware(api)

	members, _, err := m.ListMembers(org)
	if err != nil {
		return cobra.AppendActiveHelp(nil, "failed to list members for completion: "+err.Error()),
			cobra.ShellCompDirectiveNoFileComp
	}

	completions := make([]cobra.Completion, 0, len(members))
	for _, member := range members {
		emailStr := member.Email.String()
		if strings.HasPrefix(emailStr, toComplete) || toComplete == "" {
			completions = append(completions, cobra.CompletionWithDesc(emailStr,
				fmt.Sprintf("(%s)", member.Username),
			))
		}
	}

	return completions, cobra.ShellCompDirectiveNoFileComp
}
