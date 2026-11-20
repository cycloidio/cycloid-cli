package cyargs

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"go.yaml.in/yaml/v4"

	"github.com/cycloidio/cycloid-cli/cmd/apiclient"
	"github.com/cycloidio/cycloid-cli/cmd/common"
	"github.com/cycloidio/cycloid-cli/gen/models"
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
	m := apiclient.NewAPIClient(api)

	org, err := GetOrg(cmd)
	if err != nil {
		return []string{}, cobra.ShellCompDirectiveNoFileComp
	}

	orgs, _, err := m.ListOrganizationChildrens(org)
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
	cmd.Flags().StringP(flagName, "p", "", "the parent organization canonical")
	cmd.RegisterFlagCompletionFunc(flagName, CompleteOrg)
	return flagName
}

func GetOrgParentCanonical(cmd *cobra.Command) (string, error) {
	return cmd.Flags().GetString("parent-canonical")
}

func AddRoleCanonicalFlag(cmd *cobra.Command) string {
	flagName := "role"
	cmd.Flags().StringP(flagName, "r", "", "the role canonical")
	cmd.RegisterFlagCompletionFunc(flagName, CompleteRoleCanonical)
	return flagName
}

func GetRoleCanonical(cmd *cobra.Command) (string, error) {
	roleFlag, err := cmd.Flags().GetString("role")
	if err != nil {
		return "", err
	}
	if roleFlag != "" {
		return roleFlag, nil
	}

	// TODO: remove next release
	// Support for the old flag, deprecated
	can, err := cmd.Flags().GetString("canonical")
	if err != nil {
		return "", err
	}

	return can, nil
}

func CompleteRoleCanonical(cmd *cobra.Command, args []string, toComplete string) ([]cobra.Completion, cobra.ShellCompDirective) {
	org, err := GetOrg(cmd)
	if err != nil {
		return cobra.AppendActiveHelp(nil, "failed to list role, missing org for completion "+err.Error()), cobra.ShellCompDirectiveNoFileComp
	}

	api := common.NewAPI()
	m := apiclient.NewAPIClient(api)

	roles, _, err := m.ListRoles(org)
	if err != nil {
		return cobra.AppendActiveHelp(nil, "failed to list roles, "+err.Error()), cobra.ShellCompDirectiveNoFileComp
	}

	completions := make([]cobra.Completion, len(roles))
	for i, role := range roles {
		completions[i] = *role.Canonical
	}

	return completions, cobra.ShellCompDirectiveNoFileComp
}

func AddRoleNameFlag(cmd *cobra.Command) string {
	flagName := "name"
	cmd.Flags().String(flagName, "", "The name of the role")
	return flagName
}

func AddRoleRulesJSONFlag(cmd *cobra.Command) string {
	flagName := "rule-json"
	cmd.Flags().StringArrayP(flagName, "j", []string{},
		`a JSON string representing a list of rules, spec: {"action": "string", "effect": "string", "resources": ["string"]}`,
	)
	return flagName
}

func AddRoleRulesFilesFlag(cmd *cobra.Command) string {
	flagName := "rule-file"
	cmd.Flags().StringArrayP(flagName, "f", []string{},
		`path to a YAML or JSON file (use "-" to read from stdin) containing a list of rules, spec: '[{"action": "string", "effect": "string", "resources": ["strings"]}]'`,
	)
	cmd.MarkFlagFilename(flagName, "json", "yaml")
	return flagName
}

func GetRoleName(cmd *cobra.Command) (string, error) {
	return cmd.Flags().GetString("name")
}

func GetRoleRulesJSON(cmd *cobra.Command) ([]*models.NewRule, error) {
	rulesJSON, err := cmd.Flags().GetStringArray("rule-json")
	if err != nil {
		return nil, err
	}

	rules := make([]*models.NewRule, len(rulesJSON))
	for i, ruleJSON := range rulesJSON {
		var rule models.NewRule
		err := json.Unmarshal([]byte(ruleJSON), &rule)
		if err != nil {
			return nil, fmt.Errorf("failed to parse JSON: %s, as rule: %w", ruleJSON, err)
		}

		rules[i] = &rule
	}

	return rules, nil
}

func GetRoleRulesFiles(cmd *cobra.Command) ([]*models.NewRule, error) {
	rulesFiles, err := cmd.Flags().GetStringArray("rule-file")
	if err != nil {
		return nil, err
	}

	rules := []*models.NewRule{}
	for _, ruleFile := range rulesFiles {
		var content []byte
		if ruleFile == "-" && common.DetectStdinInput() {
			ruleFile = "stdin"
			content, err = io.ReadAll(cmd.InOrStdin())
			if err != nil {
				return nil, fmt.Errorf("failed to read rule from stdin %q: %w", ruleFile, err)
			}
		} else {
			content, err = os.ReadFile(ruleFile)
			if err != nil {
				return nil, fmt.Errorf("failed to read rule file %q: %w", ruleFile, err)
			}
		}

		var fileRules []*models.NewRule
		err = yaml.Unmarshal(content, &fileRules)
		if err != nil {
			return nil, fmt.Errorf("failed to decode file content from %q: %w", ruleFile, err)
		}

		rules = append(rules, fileRules...)
	}

	return rules, nil
}

// AddStartFromFlags adds --from-role and --from-team, the CLI side of the console's "Start from"
func AddStartFromFlags(cmd *cobra.Command) (string, string) {
	cmd.Flags().String("from-role", "", "start from the rules of this role; rules given with the other rule flags are added to them")
	cmd.Flags().String("from-team", "", "start from the rules of every role of this team; rules given with the other rule flags are added to them")
	cmd.MarkFlagsMutuallyExclusive("from-role", "from-team")
	return "from-role", "from-team"
}

// GetStartFromRules reads the rules of the role or team named by --from-role or --from-team, ids
// and duplicates dropped. The API still refuses any rule the caller does not hold
func GetStartFromRules(cmd *cobra.Command, m apiclient.APIClient, org string) ([]*models.NewRule, error) {
	if cmd.Flags().Lookup("from-role") == nil {
		return nil, nil
	}

	fromRole, err := cmd.Flags().GetString("from-role")
	if err != nil {
		return nil, err
	}
	fromTeam, err := cmd.Flags().GetString("from-team")
	if err != nil {
		return nil, err
	}

	var roles []*models.Role
	switch {
	case fromRole != "":
		r, _, err := m.GetRole(org, fromRole)
		if err != nil {
			return nil, fmt.Errorf("failed to get role %q: %w", fromRole, err)
		}
		roles = []*models.Role{r}
	case fromTeam != "":
		t, _, err := m.GetTeam(org, fromTeam)
		if err != nil {
			return nil, fmt.Errorf("failed to get team %q: %w", fromTeam, err)
		}
		roles = t.Roles
	}

	rules := []*models.NewRule{}
	seen := map[string]bool{}
	for _, r := range roles {
		for _, ru := range r.Rules {
			nr := &models.NewRule{Action: ru.Action, Effect: ru.Effect, Resources: ru.Resources}
			key, err := json.Marshal(nr)
			if err != nil {
				return nil, err
			}
			if seen[string(key)] {
				continue
			}
			seen[string(key)] = true
			rules = append(rules, nr)
		}
	}

	return rules, nil
}
