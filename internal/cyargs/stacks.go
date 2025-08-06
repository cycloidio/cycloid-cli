package cyargs

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/spf13/cobra"
)

func ValidateForms(cmd *cobra.Command, args []string, toComplete string) ([]cobra.Completion, cobra.ShellCompDirective) {
	var (
		err error
		dir string
	)
	if toComplete != "" {
		absPath, err := filepath.Abs(toComplete)
		if err != nil {
			return []cobra.Completion{}, cobra.ShellCompDirectiveDefault
		}

		dir = filepath.Base(absPath)
	} else {
		cwd, err := os.Getwd()
		if err != nil {
			return []cobra.Completion{}, cobra.ShellCompDirectiveDefault
		}

		dir = filepath.Base(cwd)
	}

	dirEntries, err := os.ReadDir(dir)
	if err != nil {
		return []cobra.Completion{}, cobra.ShellCompDirectiveDefault
	}

	var comp = make([]cobra.Completion, len(dirEntries))
	for index, dirEntry := range dirEntries {

		if dirEntry.IsDir() {
			continue
		}

		name := dirEntry.Name()
		if (strings.HasSuffix(name, "yml") ||
			strings.HasSuffix(name, "yaml") ||
			name == ".forms.yaml" ||
			name == ".forms.yml") && strings.HasPrefix(name, toComplete) {
			comp[index] = name
		}
	}

	return comp, cobra.ShellCompDirectiveDefault
}

func AddStackRefFlag(cmd *cobra.Command) string {
	flagName := "stack-ref"
	cmd.Flags().StringP(flagName, "s", "", "referential of the stack, format 'org:canonical'")
	cmd.RegisterFlagCompletionFunc(flagName, CompleteStackRef)
	return flagName
}

func CompleteStackRef(cmd *cobra.Command, args []string, toComplete string) ([]cobra.Completion, cobra.ShellCompDirective) {
	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	org, err := GetOrg(cmd)
	if err != nil {
		return cobra.AppendActiveHelp(nil, "completion failed: "+err.Error()),
			cobra.ShellCompDirectiveNoFileComp
	}

	stacks, err := m.ListStacks(org)
	if err != nil {
		return cobra.AppendActiveHelp(nil, "failed to list stacks in org '"+org+"': "+err.Error()),
			cobra.ShellCompDirectiveNoFileComp
	}

	var stackRefs = make([]string, len(stacks))
	for index, stack := range stacks {
		if stack.Ref != nil && strings.HasPrefix(*stack.Ref, toComplete) {
			desc := *stack.Name
			if stack.Description != "" {
				desc = desc + " - " + stack.Description
			}
			stackRefs[index] = cobra.CompletionWithDesc(*stack.Ref, desc)
		}
	}

	return stackRefs, cobra.ShellCompDirectiveNoFileComp
}

func GetStackRef(cmd *cobra.Command) (string, error) {
	return cmd.Flags().GetString("stack-ref")
}

func AddVisibilityFlag(cmd *cobra.Command) string {
	flagName := "visibility"
	cmd.Flags().StringP(flagName, "V", "", "set the visibility of a stack")
	cmd.RegisterFlagCompletionFunc(flagName, CompleteVisibility)
	return flagName
}

func GetVisibility(cmd *cobra.Command) (string, error) {
	return cmd.Flags().GetString("visibility")
}

func CompleteVisibility(cmd *cobra.Command, args []string, toComplete string) ([]cobra.Completion, cobra.ShellCompDirective) {
	return []cobra.Completion{
		"hidden",
		"shared",
		"local",
	}, cobra.ShellCompDirectiveNoFileComp
}

func AddTeamFlag(cmd *cobra.Command) string {
	flagName := "team"
	cmd.Flags().StringP(flagName, "t", "", "designate the maintainers of a stack")
	cmd.RegisterFlagCompletionFunc(flagName, CompleteTeam)
	return flagName
}

func GetTeam(cmd *cobra.Command) (string, error) {
	return cmd.Flags().GetString("team")
}

// TODO: We don't care that much of that
// This may disappear
func CompleteTeam(cmd *cobra.Command, args []string, toComplete string) ([]cobra.Completion, cobra.ShellCompDirective) {
	return []cobra.Completion{}, cobra.ShellCompDirectiveNoFileComp
}

// Blueprint-related functions
func AddBlueprintRefFlag(cmd *cobra.Command) string {
	flagName := "blueprint-ref"
	cmd.Flags().StringP(flagName, "", "", "Blueprint reference to use")
	cmd.RegisterFlagCompletionFunc(flagName, CompleteBlueprint)
	return flagName
}

func CompleteBlueprint(cmd *cobra.Command, args []string, toComplete string) ([]cobra.Completion, cobra.ShellCompDirective) {
	org, err := GetOrg(cmd)
	if err != nil {
		return cobra.AppendActiveHelp(nil, "completion failed: "+err.Error()),
			cobra.ShellCompDirectiveNoFileComp
	}

	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	blueprints, err := m.ListBlueprints(org)
	if err != nil {
		return cobra.AppendActiveHelp(nil, "failed to list stacks in org '"+org+"': "+err.Error()),
			cobra.ShellCompDirectiveNoFileComp
	}

	var completions = make([]cobra.Completion, len(blueprints))
	for index, blueprint := range blueprints {
		if blueprint.Ref != nil && strings.HasPrefix(*blueprint.Ref, toComplete) {
			completions[index] = cobra.CompletionWithDesc(*blueprint.Ref, *blueprint.Name+" - "+blueprint.Description)
		}
	}

	return completions, cobra.ShellCompDirectiveNoFileComp
}

func GetBlueprintRef(cmd *cobra.Command) (string, error) {
	return cmd.Flags().GetString("blueprint-ref")
}

func AddStackNameFlag(cmd *cobra.Command) string {
	flagName := "name"
	cmd.Flags().StringP(flagName, "n", "", "name of the stack")
	return flagName
}

func GetStackName(cmd *cobra.Command) (string, error) {
	return cmd.Flags().GetString("name")
}

func AddStackFlag(cmd *cobra.Command) string {
	flagStack := "stack"
	cmd.Flags().StringP(flagStack, "s", "", "canonical of the stack")
	return flagStack
}

func GetStack(cmd *cobra.Command) (string, error) {
	return cmd.Flags().GetString("stack")
}
