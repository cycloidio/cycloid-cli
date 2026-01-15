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

func AddStackFlag(cmd *cobra.Command) string {
	flagStack := "stack"
	cmd.Flags().StringP(flagStack, "s", "", "canonical of the stack")
	return flagStack
}

func GetStack(cmd *cobra.Command) (string, error) {
	return cmd.Flags().GetString("stack")
}

// AddStackVersionFlags adds mutually exclusive flags for specifying stack versions.
// Similar to AddStackFormsInputFlags, this groups related version flags together.
// At least one of these flags must be provided.
func AddStackVersionFlags(cmd *cobra.Command) {
	cmd.Flags().String("stack-tag", "", "the stack version tag to use (e.g., v1.0.0)")
	cmd.Flags().String("stack-branch", "", "the stack version branch to use (e.g., main)")
	cmd.Flags().String("stack-commit-hash", "", "the stack version commit hash to use")

	// Register completion functions
	cmd.RegisterFlagCompletionFunc("stack-tag", CompleteStackVersionTag)
	cmd.RegisterFlagCompletionFunc("stack-branch", CompleteStackVersionBranch)
	cmd.RegisterFlagCompletionFunc("stack-commit-hash", CompleteCatalogRepoCommitHash)

	// Make flags mutually exclusive
	cmd.MarkFlagsMutuallyExclusive("stack-tag", "stack-branch", "stack-commit-hash")
}

// GetStackVersionFlags reads the stack version flags from the command.
func GetStackVersionFlags(cmd *cobra.Command) (tag, branch, hash string, err error) {
	tag, err = cmd.Flags().GetString("stack-tag")
	if err != nil {
		return "", "", "", err
	}

	branch, err = cmd.Flags().GetString("stack-branch")
	if err != nil {
		return "", "", "", err
	}

	hash, err = cmd.Flags().GetString("stack-commit-hash")
	if err != nil {
		return "", "", "", err
	}

	return tag, branch, hash, nil
}

// CompleteCatalogRepoCommitHash provides completion for the catalog-repo-commit-hash flag.
// It fetches available commit hashes from the stack's catalog repository versions.
// For component create: uses the required stack-ref flag.
// For component update: uses stack-ref flag if provided, otherwise fetches it from the existing component.
func CompleteCatalogRepoCommitHash(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	org, err := GetOrg(cmd)
	if err != nil {
		return cobra.AppendActiveHelp(nil, "missing org for completion: "+err.Error()), cobra.ShellCompDirectiveError
	}

	var stackRef string
	stackRef, _ = GetStackRef(cmd)

	// If stack-ref is not provided, try to get it from the component (for update command)
	if stackRef == "" {
		project, errProj := GetProject(cmd)
		env, errEnv := GetEnv(cmd)
		component, errComp := GetComponent(cmd)

		if errProj == nil && errEnv == nil && errComp == nil {
			api := common.NewAPI()
			m := middleware.NewMiddleware(api)

			currentComponent, errGet := m.GetComponent(org, project, env, component)
			if errGet == nil && currentComponent.ServiceCatalog != nil && currentComponent.ServiceCatalog.Ref != nil {
				stackRef = *currentComponent.ServiceCatalog.Ref
			}
		}
	}

	if stackRef == "" {
		return cobra.AppendActiveHelp(nil, "missing stack-ref for completion: please provide --stack-ref flag"),
			cobra.ShellCompDirectiveNoFileComp
	}

	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	versions, err := m.ListStackVersions(org, stackRef)
	if err != nil {
		return cobra.AppendActiveHelp(nil, "cannot find versions for stack: "+stackRef+", err: "+err.Error()),
			cobra.ShellCompDirectiveNoFileComp
	}

	var commitHashes []string
	for _, version := range versions {
		if version.CommitHash != nil && strings.HasPrefix(*version.CommitHash, toComplete) {
			// Build description with version name and type
			desc := ""
			if version.Name != nil {
				desc = *version.Name
			}
			if version.Type != nil {
				if desc != "" {
					desc = desc + " (" + *version.Type + ")"
				} else {
					desc = *version.Type
				}
			}
			if version.IsLatest != nil && *version.IsLatest {
				desc = desc + " [latest]"
			}

			commitHashes = append(commitHashes, cobra.CompletionWithDesc(*version.CommitHash, desc))
		}
	}

	return commitHashes, cobra.ShellCompDirectiveNoFileComp
}

// CompleteStackVersionTag provides completion for the stack-tag flag.
// It fetches available tags from the stack's catalog repository versions.
func CompleteStackVersionTag(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	org, err := GetOrg(cmd)
	if err != nil {
		return cobra.AppendActiveHelp(nil, "missing org for completion: "+err.Error()), cobra.ShellCompDirectiveError
	}

	var stackRef string
	stackRef, _ = GetStackRef(cmd)

	// If stack-ref is not provided, try to get it from the component (for update command)
	if stackRef == "" {
		project, errProj := GetProject(cmd)
		env, errEnv := GetEnv(cmd)
		component, errComp := GetComponent(cmd)

		if errProj == nil && errEnv == nil && errComp == nil {
			api := common.NewAPI()
			m := middleware.NewMiddleware(api)

			currentComponent, errGet := m.GetComponent(org, project, env, component)
			if errGet == nil && currentComponent.ServiceCatalog != nil && currentComponent.ServiceCatalog.Ref != nil {
				stackRef = *currentComponent.ServiceCatalog.Ref
			}
		}
	}

	if stackRef == "" {
		return cobra.AppendActiveHelp(nil, "missing stack-ref for completion: please provide --stack-ref flag"),
			cobra.ShellCompDirectiveNoFileComp
	}

	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	versions, err := m.ListStackVersions(org, stackRef)
	if err != nil {
		return cobra.AppendActiveHelp(nil, "cannot find versions for stack: "+stackRef+", err: "+err.Error()),
			cobra.ShellCompDirectiveNoFileComp
	}

	var tags []string
	for _, version := range versions {
		if version.Type != nil && *version.Type == "tag" && version.Name != nil && strings.HasPrefix(*version.Name, toComplete) {
			// Build description with commit hash
			desc := ""
			if version.CommitHash != nil {
				commitHashShort := *version.CommitHash
				if len(commitHashShort) > 8 {
					commitHashShort = commitHashShort[:8]
				}
				desc = commitHashShort
			}
			if version.IsLatest != nil && *version.IsLatest {
				desc = desc + " [latest]"
			}

			tags = append(tags, cobra.CompletionWithDesc(*version.Name, desc))
		}
	}

	return tags, cobra.ShellCompDirectiveNoFileComp
}

// CompleteStackVersionBranch provides completion for the stack-branch flag.
// It fetches available branches from the stack's catalog repository versions.
func CompleteStackVersionBranch(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	org, err := GetOrg(cmd)
	if err != nil {
		return cobra.AppendActiveHelp(nil, "missing org for completion: "+err.Error()), cobra.ShellCompDirectiveError
	}

	var stackRef string
	stackRef, _ = GetStackRef(cmd)

	// If stack-ref is not provided, try to get it from the component (for update command)
	if stackRef == "" {
		project, errProj := GetProject(cmd)
		env, errEnv := GetEnv(cmd)
		component, errComp := GetComponent(cmd)

		if errProj == nil && errEnv == nil && errComp == nil {
			api := common.NewAPI()
			m := middleware.NewMiddleware(api)

			currentComponent, errGet := m.GetComponent(org, project, env, component)
			if errGet == nil && currentComponent.ServiceCatalog != nil && currentComponent.ServiceCatalog.Ref != nil {
				stackRef = *currentComponent.ServiceCatalog.Ref
			}
		}
	}

	if stackRef == "" {
		return cobra.AppendActiveHelp(nil, "missing stack-ref for completion: please provide --stack-ref flag"),
			cobra.ShellCompDirectiveNoFileComp
	}

	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	versions, err := m.ListStackVersions(org, stackRef)
	if err != nil {
		return cobra.AppendActiveHelp(nil, "cannot find versions for stack: "+stackRef+", err: "+err.Error()),
			cobra.ShellCompDirectiveNoFileComp
	}

	var branches []string
	for _, version := range versions {
		if version.Type != nil && *version.Type == "branch" && version.Name != nil && strings.HasPrefix(*version.Name, toComplete) {
			// Build description with commit hash
			desc := ""
			if version.CommitHash != nil {
				commitHashShort := *version.CommitHash
				if len(commitHashShort) > 8 {
					commitHashShort = commitHashShort[:8]
				}
				desc = commitHashShort
			}

			branches = append(branches, cobra.CompletionWithDesc(*version.Name, desc))
		}
	}

	return branches, cobra.ShellCompDirectiveNoFileComp
}
