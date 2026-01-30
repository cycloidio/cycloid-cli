package cyargs

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/cycloidio/cycloid-cli/client/models"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/spf13/cobra"
	"go.yaml.in/yaml/v4"
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

func AddRoleCanonicalFlag(cmd *cobra.Command) string {
	flagName := "role"
	cmd.Flags().StringP(flagName, "r", "", "the role canonical")
	cmd.RegisterFlagCompletionFunc(flagName, CompleteRoleCanonicals)
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

func CompleteRoleCanonicals(cmd *cobra.Command, args []string, toComplete string) ([]cobra.Completion, cobra.ShellCompDirective) {
	org, err := GetOrg(cmd)
	if err != nil {
		return cobra.AppendActiveHelp(nil, "failed to list role, missing org for completion "+err.Error()), cobra.ShellCompDirectiveNoFileComp
	}

	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	roles, err := m.ListRoles(org)
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

	var rules = make([]*models.NewRule, len(rulesJSON))
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

	var rules = []*models.NewRule{}
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
