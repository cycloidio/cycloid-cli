package cyargs

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/cycloidio/cycloid-cli/client/models"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/ptr"
	"github.com/spf13/cobra"
)

const (
	teamFlagName               = "team"
	teamNameFlagName           = "name"
	teamOrderByFlagName        = "order-by"
	teamMemberIDFlagName       = "member-id"
	teamMemberFlagName         = "member"
	teamMemberEmailFlagName    = "email"
	teamMemberUsernameFlagName = "username"
	teamCreatedAtFlagName      = "created-at"
	teamOwnerFlagName          = "owner"
	teamRolesFlagName          = "role"
)

func AddTeamFlag(cmd *cobra.Command) string {
	_ = cmd.Flags().StringP(teamFlagName, "t", "", "the team canonical")
	cmd.RegisterFlagCompletionFunc(teamFlagName, CompleteTeam)
	return teamFlagName
}

func CompleteTeam(cmd *cobra.Command, args []string, toComplete string) ([]cobra.Completion, cobra.ShellCompDirective) {
	org, err := GetOrg(cmd)
	if err != nil {
		return cobra.AppendActiveHelp(nil, "missing org for completion: "+err.Error()),
			cobra.ShellCompDirectiveNoFileComp
	}

	name, _ := GetTeamName(cmd)
	memberID, _ := GetTeamMemberID(cmd)
	createdAt, _ := GetTeamCreatedAt(cmd)
	orderBy, _ := GetTeamOrderBy(cmd)

	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	teams, err := m.ListTeams(org, &name, createdAt, &memberID, orderBy)
	if err != nil {
		return cobra.AppendActiveHelp(nil, "Failed to list team for completion: "+err.Error()),
			cobra.ShellCompDirectiveNoFileComp
	}

	completions := make([]cobra.Completion, len(teams))
	for i, t := range teams {
		if t.Canonical != nil {
			completions[i] = cobra.CompletionWithDesc(*t.Canonical, ptr.Value(t.Name))
		}
	}

	return completions, cobra.ShellCompDirectiveNoFileComp
}

func GetTeam(cmd *cobra.Command) (string, error) {
	return cmd.Flags().GetString(teamFlagName)
}

func AddTeamNameFlag(cmd *cobra.Command) string {
	cmd.Flags().StringP(teamNameFlagName, "n", "", "the name of the team")
	return teamNameFlagName
}

func GetTeamName(cmd *cobra.Command) (string, error) {
	return cmd.Flags().GetString(teamNameFlagName)
}

// TeamMembersCompletionHelper is meant to be consumed by other members related
// complete functions.
func TeamMembersCompletionHelper(cmd *cobra.Command) ([]*models.MemberTeam, error) {
	org, err := GetOrg(cmd)
	if err != nil {
		return nil, fmt.Errorf("missing org for completion: %w", err)
	}

	team, err := GetOrg(cmd)
	if err != nil {
		return nil, fmt.Errorf("missing team canoncal for completion: %w", err)
	}

	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	teams, err := m.ListTeamMembers(org, team)
	if err != nil {
		return nil, fmt.Errorf("cannot list members from team %q for completion: %w", team, err)
	}

	return teams, nil
}

func CompleteTeamMemberID(cmd *cobra.Command, args []string, toComplete string) ([]cobra.Completion, cobra.ShellCompDirective) {
	members, err := TeamMembersCompletionHelper(cmd)
	if err != nil {
		return cobra.AppendActiveHelp([]cobra.Completion{}, err.Error()),
			cobra.ShellCompDirectiveNoFileComp
	}

	completions := make([]cobra.Completion, len(members))
	for i, member := range members {
		if member.ID == nil {
			continue
		}

		idStr := strconv.Itoa(int(*member.ID))
		if strings.HasPrefix(idStr, toComplete) || toComplete == "" {
			completions[i] = cobra.CompletionWithDesc(idStr, member.Email.String()+": "+member.FullName)
		}
	}

	return completions, cobra.ShellCompDirectiveNoFileComp
}

func CompleteTeamMemberUsername(cmd *cobra.Command, args []string, toComplete string) ([]cobra.Completion, cobra.ShellCompDirective) {
	members, err := TeamMembersCompletionHelper(cmd)
	if err != nil {
		return cobra.AppendActiveHelp([]cobra.Completion{}, err.Error()),
			cobra.ShellCompDirectiveNoFileComp
	}

	completions := make([]cobra.Completion, len(members))
	for i, member := range members {
		if strings.HasPrefix(member.Username, toComplete) || toComplete == "" {
			completions[i] = cobra.CompletionWithDesc(member.Username, member.Email.String()+": "+member.Username)
		}
	}

	return completions, cobra.ShellCompDirectiveNoFileComp
}

func CompleteTeamMemberEmail(cmd *cobra.Command, args []string, toComplete string) ([]cobra.Completion, cobra.ShellCompDirective) {
	members, err := TeamMembersCompletionHelper(cmd)
	if err != nil {
		return cobra.AppendActiveHelp([]cobra.Completion{}, err.Error()),
			cobra.ShellCompDirectiveNoFileComp
	}

	completions := make([]cobra.Completion, len(members))
	for i, member := range members {
		if strings.HasPrefix(member.Email.String(), toComplete) || toComplete == "" {
			completions[i] = cobra.CompletionWithDesc(member.Email.String(), member.Username+": "+member.FullName)
		}
	}

	return completions, cobra.ShellCompDirectiveNoFileComp
}

func CompleteTeamMemberEmailOrUsername(cmd *cobra.Command, args []string, toComplete string) ([]cobra.Completion, cobra.ShellCompDirective) {
	members, err := TeamMembersCompletionHelper(cmd)
	if err != nil {
		return cobra.AppendActiveHelp([]cobra.Completion{}, err.Error()),
			cobra.ShellCompDirectiveNoFileComp
	}

	completions := make([]cobra.Completion, len(members))
	for i, member := range members {
		if strings.HasPrefix(member.Email.String(), toComplete) ||
			strings.HasPrefix(member.Username, toComplete) ||
			toComplete == "" {
			completions[i] = cobra.CompletionWithDesc(member.Email.String(), member.Username+": "+member.FullName)
		}
	}

	return completions, cobra.ShellCompDirectiveNoFileComp
}

func CompleteTeamMemberAny(cmd *cobra.Command, args []string, toComplete string) ([]cobra.Completion, cobra.ShellCompDirective) {
	members, err := TeamMembersCompletionHelper(cmd)
	if err != nil {
		return cobra.AppendActiveHelp([]cobra.Completion{}, err.Error()),
			cobra.ShellCompDirectiveNoFileComp
	}

	completions := make([]cobra.Completion, len(members))
	for i, member := range members {
		// Try to complete any of those team member attributes, fullname, email, id
		idStr := strconv.Itoa(int(ptr.Value(member.ID)))
		if strings.HasPrefix(idStr, toComplete) || toComplete == "" {
			completions[i] = cobra.CompletionWithDesc(idStr, member.Username+": "+member.FullName)
			continue
		}

		if strings.HasPrefix(member.FullName, toComplete) ||
			strings.HasPrefix(member.Username, toComplete) || toComplete == "" {
			completions[i] = cobra.CompletionWithDesc(member.FullName, member.Username+": "+member.FullName)
			continue
		}

		if strings.HasPrefix(member.Email.String(), toComplete) || toComplete == "" {
			completions[i] = cobra.CompletionWithDesc(member.Email.String(), member.Username+": "+member.FullName)
		}
	}

	return completions, cobra.ShellCompDirectiveNoFileComp
}

func AddTeamMemberIDFlag(cmd *cobra.Command) string {
	cmd.Flags().Uint32P(teamMemberIDFlagName, "i", 0, "specify a team member id")
	cmd.RegisterFlagCompletionFunc(teamMemberIDFlagName, CompleteTeamMemberID)
	return teamMemberIDFlagName
}

func GetTeamMemberID(cmd *cobra.Command) (uint32, error) {
	return cmd.Flags().GetUint32(teamMemberIDFlagName)
}

func AddTeamMemberUsernameFlag(cmd *cobra.Command) string {
	cmd.Flags().StringP(teamMemberFlagName, "m", "", "specify a team member by its username")
	cmd.RegisterFlagCompletionFunc(teamMemberFlagName, CompleteTeamMemberUsername)
	return teamMemberIDFlagName
}

func GetTeamMember(cmd *cobra.Command) (string, error) {
	return cmd.Flags().GetString(teamMemberFlagName)
}

func AddTeamMemberEmailFlagName(cmd *cobra.Command) string {
	cmd.Flags().StringP(teamMemberEmailFlagName, "e", "", "team member email")
	cmd.RegisterFlagCompletionFunc(teamMemberEmailFlagName, CompleteTeamMemberEmail)
	return teamMemberEmailFlagName
}

func GetTeamMemberUsername(cmd *cobra.Command) (string, error) {
	return cmd.Flags().GetString(teamMemberUsernameFlagName)
}

func AddTeamCreatedAt(cmd *cobra.Command) string {
	cmd.Flags().Time(teamCreatedAtFlagName, time.UnixMilli(0), []string{}, "team member email")
	return teamMemberEmailFlagName
}

func GetTeamCreatedAt(cmd *cobra.Command) (*uint64, error) {
	time, err := cmd.Flags().GetTime(teamCreatedAtFlagName)
	if err != nil {
		return nil, err
	}
	return ptr.Ptr(uint64(time.UnixMilli())), nil
}

func AddTeamOrderBy(cmd *cobra.Command) string {
	cmd.Flags().String(teamOrderByFlagName, "asc", "specify the output order, asc for ascending and desc for descending")
	cmd.RegisterFlagCompletionFunc(teamOrderByFlagName, cobra.FixedCompletions([]string{"asc", "desc"}, cobra.ShellCompDirectiveNoFileComp))
	return teamCreatedAtFlagName
}

func GetTeamOrderBy(cmd *cobra.Command) (*middleware.TeamOrderByParam, error) {
	orderByStr, err := cmd.Flags().GetString(teamOrderByFlagName)
	if err != nil {
		return &middleware.Ascending, err
	}

	switch orderByStr {
	case "desc", "descending":
		return &middleware.Descending, nil
	default:
		return &middleware.Ascending, nil
	}
}

func AddTeamOwnerFlag(cmd *cobra.Command) string {
	cmd.Flags().String(teamOwnerFlagName, "", "specify a team owner canonical")
	cmd.RegisterFlagCompletionFunc(teamOwnerFlagName, CompleteTeamOwnerCanonical)
	return teamOwnerFlagName
}

func CompleteTeamOwnerCanonical(cmd *cobra.Command, args []string, toComplete string) ([]cobra.Completion, cobra.ShellCompDirective) {
	org, err := GetOrg(cmd)
	if err != nil {
		return cobra.AppendActiveHelp([]string{}, "missing org for completion: "+err.Error()), cobra.ShellCompDirectiveNoFileComp
	}

	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	users, err := m.ListMembers(org)
	if err != nil {
		return cobra.AppendActiveHelp([]string{}, "failed to list current org members: "+err.Error()), cobra.ShellCompDirectiveNoFileComp
	}

	completions := make([]cobra.Completion, len(users))
	for i, user := range users {
		if strings.HasPrefix(user.Username, toComplete) {
			completions[i] = cobra.CompletionWithDesc(user.Username, user.Email.String()+": "+user.FullName)
		}
	}

	return completions, cobra.ShellCompDirectiveNoFileComp
}

func GetTeamOwner(cmd *cobra.Command) (string, error) {
	return cmd.Flags().GetString(teamOwnerFlagName)
}

func AddTeamRolesFlag(cmd *cobra.Command) string {
	cmd.Flags().StringArrayP(teamRolesFlagName, "r", []string{}, "list of role canonical to associate to the team")
	cmd.RegisterFlagCompletionFunc(teamRolesFlagName, CompleteRoleCanonical)
	return teamRolesFlagName
}

func GetTeamRoles(cmd *cobra.Command) ([]string, error) {
	return cmd.Flags().GetStringArray(teamRolesFlagName)
}
