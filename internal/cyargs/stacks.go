package cyargs

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/apiclient"
	"github.com/cycloidio/cycloid-cli/cmd/common"
	"github.com/cycloidio/cycloid-cli/utils/ptr"
)

func ValidateForms(cmd *cobra.Command, args []string, toComplete string) ([]cobra.Completion, cobra.ShellCompDirective) {
	return []cobra.Completion{"yml", "yaml"}, cobra.ShellCompDirectiveFilterFileExt
}

func AddStackRefFlag(cmd *cobra.Command) string {
	flagName := "stack-ref"
	cmd.Flags().StringP(flagName, "s", "", "referential of the stack, format 'org:canonical'")
	cmd.RegisterFlagCompletionFunc(flagName, CompleteStackRef)
	return flagName
}

func CompleteStackRef(cmd *cobra.Command, args []string, toComplete string) ([]cobra.Completion, cobra.ShellCompDirective) {
	api := common.NewAPI()
	m := apiclient.NewAPIClient(api)

	org, err := GetOrg(cmd)
	if err != nil {
		return cobra.AppendActiveHelp(nil, "completion failed: "+err.Error()),
			cobra.ShellCompDirectiveNoFileComp
	}

	stacks, _, err := m.ListStacks(org)
	if err != nil {
		return cobra.AppendActiveHelp(nil, "failed to list stacks in org '"+org+"': "+err.Error()),
			cobra.ShellCompDirectiveNoFileComp
	}

	stackRefs := make([]string, len(stacks))
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
	m := apiclient.NewAPIClient(api)

	blueprints, _, err := m.ListBlueprints(org)
	if err != nil {
		return cobra.AppendActiveHelp(nil, "failed to list stacks in org '"+org+"': "+err.Error()),
			cobra.ShellCompDirectiveNoFileComp
	}

	completions := make([]cobra.Completion, len(blueprints))
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

// AddStackVersionFlags adds flags for specifying stack versions.
// The primary flag is --stack-version, which accepts a bare name (resolved via API)
// or a typed prefix: tag:<name>, branch:<name>, sha:<hash>.
// Legacy flags --stack-tag, --stack-branch, --stack-commit-hash are kept for
// backward compatibility but are deprecated and will be removed in a future major release.
func AddStackVersionFlags(cmd *cobra.Command) {
	cmd.Flags().String("stack-version", "", "stack version (tag, branch, or commit hash). "+
		"Use type prefixes to avoid ambiguity: tag:<name>, branch:<name>, sha:<hash>, version:<id>")
	cmd.RegisterFlagCompletionFunc("stack-version", CompleteStackVersionUnified)

	// Legacy flags — kept for backward compatibility.
	cmd.Flags().String("stack-tag", "", "the stack version tag to use (e.g., v1.0.0)")
	cmd.Flags().String("stack-branch", "", "the stack version branch to use (e.g., main)")
	cmd.Flags().String("stack-commit-hash", "", "the stack version commit hash to use")

	//nolint:errcheck
	cmd.Flags().MarkDeprecated("stack-tag", "use --stack-version=tag:<value> instead (will be removed in a future major release)")
	//nolint:errcheck
	cmd.Flags().MarkDeprecated("stack-branch", "use --stack-version=branch:<value> instead (will be removed in a future major release)")
	//nolint:errcheck
	cmd.Flags().MarkDeprecated("stack-commit-hash", "use --stack-version=sha:<value> instead (will be removed in a future major release)")

	cmd.RegisterFlagCompletionFunc("stack-tag", CompleteStackVersionTag)
	cmd.RegisterFlagCompletionFunc("stack-branch", CompleteStackVersionBranch)
	cmd.RegisterFlagCompletionFunc("stack-commit-hash", CompleteCatalogRepoCommitHash)

	// Mutual exclusion: --stack-version vs each legacy flag
	cmd.MarkFlagsMutuallyExclusive("stack-version", "stack-tag")
	cmd.MarkFlagsMutuallyExclusive("stack-version", "stack-branch")
	cmd.MarkFlagsMutuallyExclusive("stack-version", "stack-commit-hash")
	// Legacy flags remain mutually exclusive with each other
	cmd.MarkFlagsMutuallyExclusive("stack-tag", "stack-branch", "stack-commit-hash")
}

// GetStackVersionFlags reads the legacy stack version flags from the command.
// Prefer ResolveStackVersionArg for new code.
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

// ResolveStackVersionArg resolves the --stack-version flag (or legacy flags) to
// the (tag, branch, hash) triple expected by apiclient methods.
//
// Resolution order:
//  1. Legacy flags (--stack-tag / --stack-branch / --stack-commit-hash) — returned as-is.
//  2. --stack-version is empty — returns ("", "", "", nil); callers should preserve the current version.
//  3. --stack-version has a type prefix (tag: / branch: / sha: / commit:) — parsed client-side, no API call.
//  4. Bare value — calls m.ListStackVersions to resolve; requires a non-empty stackRef.
//     Precedence: tag > branch > commit-hash-prefix (min 7 chars).
//     Collision (same name matches both tag and branch) → error with disambiguation hint.
func ResolveStackVersionArg(cmd *cobra.Command, m apiclient.APIClient, org, stackRef string) (tag, branch, hash string, err error) {
	// Legacy flags take priority; Cobra already emits the deprecation notice.
	tag, branch, hash, err = GetStackVersionFlags(cmd)
	if err != nil || tag != "" || branch != "" || hash != "" {
		return tag, branch, hash, err
	}

	v, err := cmd.Flags().GetString("stack-version")
	if err != nil {
		return "", "", "", err
	}
	if v == "" {
		return "", "", "", nil
	}

	// Typed prefix form — no API call needed.
	typeHint, value, hasPrefix := strings.Cut(v, ":")
	if hasPrefix {
		switch strings.ToLower(typeHint) {
		case "tag":
			return value, "", "", nil
		case "branch":
			return "", value, "", nil
		case "sha", "commit":
			return "", "", value, nil
		case "version":
			return "", "", "", fmt.Errorf(
				"--stack-version=version:<id> is only supported by 'cy component config get'; " +
					"use tag:<name>, branch:<name>, or sha:<hash> for other commands")
		}
		// Unknown prefix: fall through and attempt bare resolution on the full value.
	}

	// Bare form requires the stack ref to query available versions.
	if stackRef == "" {
		return "", "", "", fmt.Errorf(
			"cannot auto-detect version type for %q without a stack reference; "+
				"use --stack-version=tag:<value>, --stack-version=branch:<value>, or --stack-version=sha:<value>", v)
	}

	versions, _, err := m.ListStackVersions(org, stackRef)
	if err != nil {
		return "", "", "", fmt.Errorf("failed to resolve --stack-version=%q: %w", v, err)
	}

	var matchTag, matchBranch, matchCommit string
	for _, ver := range versions {
		name := ptr.Value(ver.Name)
		if name == v {
			if ptr.Value(ver.Type) == "tag" {
				matchTag = name
			} else {
				matchBranch = name
			}
		}
		if commitHash := ptr.Value(ver.CommitHash); commitHash != "" && len(v) >= 7 && strings.HasPrefix(commitHash, v) {
			matchCommit = commitHash
		}
	}

	switch {
	case matchTag != "" && matchBranch != "":
		return "", "", "", fmt.Errorf(
			"ambiguous version %q (matches both a tag and a branch); "+
				"use --stack-version=tag:%s or --stack-version=branch:%s", v, v, v)
	case matchTag != "":
		return matchTag, "", "", nil
	case matchBranch != "":
		return "", matchBranch, "", nil
	case matchCommit != "":
		return "", "", matchCommit, nil
	default:
		return "", "", "", fmt.Errorf("version %q not found in stack %q", v, stackRef)
	}
}

// GetStackVersionID parses --stack-version=version:<id> and returns the numeric catalog
// version ID. Returns (id, true, nil) when that form is used, (0, false, nil) otherwise.
func GetStackVersionID(cmd *cobra.Command) (uint32, bool, error) {
	v, err := cmd.Flags().GetString("stack-version")
	if err != nil || v == "" {
		return 0, false, err
	}
	idStr, ok := strings.CutPrefix(v, "version:")
	if !ok {
		return 0, false, nil
	}
	id64, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		return 0, false, fmt.Errorf("--stack-version=version:<id>: expected a numeric ID, got %q", idStr)
	}
	return uint32(id64), true, nil
}

// CompleteStackVersionUnified provides prefix-aware completion for --stack-version.
//
// - No prefix typed (e.g. "v1"): lists all versions with type in description.
// - "tag:<TAB>": lists only tags, formatted as "tag:<name>".
// - "branch:<TAB>": lists only branches, formatted as "branch:<name>".
// - "sha:<TAB>": lists only commit hashes (short), formatted as "sha:<hash>".
func CompleteStackVersionUnified(cmd *cobra.Command, args []string, toComplete string) ([]cobra.Completion, cobra.ShellCompDirective) {
	org, err := GetOrg(cmd)
	if err != nil {
		return cobra.AppendActiveHelp(nil, "missing org for completion: "+err.Error()), cobra.ShellCompDirectiveError
	}

	stackRef := resolveStackRefForCompletion(cmd, org)
	if stackRef == "" {
		return cobra.AppendActiveHelp(nil, "missing stack-ref for completion: please provide --stack-ref flag"),
			cobra.ShellCompDirectiveNoFileComp
	}

	api := common.NewAPI()
	m := apiclient.NewAPIClient(api)

	versions, _, err := m.ListStackVersions(org, stackRef)
	if err != nil {
		return cobra.AppendActiveHelp(nil, "cannot find versions for stack: "+stackRef+", err: "+err.Error()),
			cobra.ShellCompDirectiveNoFileComp
	}

	typeHint, valuePrefix, hasPrefix := strings.Cut(toComplete, ":")

	var completions []cobra.Completion
	for _, ver := range versions {
		if ver.Name == nil || ver.Type == nil {
			continue
		}
		name := *ver.Name
		vType := *ver.Type

		if hasPrefix {
			switch strings.ToLower(typeHint) {
			case "tag":
				if vType == "tag" && strings.HasPrefix(name, valuePrefix) {
					completions = append(completions, cobra.CompletionWithDesc("tag:"+name, versionDesc(ver)))
				}
			case "branch":
				if vType == "branch" && strings.HasPrefix(name, valuePrefix) {
					completions = append(completions, cobra.CompletionWithDesc("branch:"+name, versionDesc(ver)))
				}
			case "sha", "commit":
				if ver.CommitHash != nil {
					h := shortHash(*ver.CommitHash)
					if strings.HasPrefix(h, valuePrefix) {
						completions = append(completions, cobra.CompletionWithDesc("sha:"+h, versionDesc(ver)))
					}
				}
			}
		} else {
			if strings.HasPrefix(name, toComplete) {
				completions = append(completions, cobra.CompletionWithDesc(name, "("+vType+") "+versionDesc(ver)))
			}
		}
	}

	return completions, cobra.ShellCompDirectiveNoFileComp
}

// resolveStackRefForCompletion tries to determine the stack ref for completion functions.
// It checks --stack-ref first, then falls back to fetching it from the existing component.
func resolveStackRefForCompletion(cmd *cobra.Command, org string) string {
	stackRef, _ := GetStackRef(cmd)
	if stackRef != "" {
		return stackRef
	}

	project, errProj := GetProject(cmd)
	env, errEnv := GetEnv(cmd)
	component, errComp := GetComponent(cmd)
	if errProj != nil || errEnv != nil || errComp != nil {
		return ""
	}

	api := common.NewAPI()
	m := apiclient.NewAPIClient(api)
	currentComponent, _, errGet := m.GetComponent(org, project, env, component)
	if errGet == nil && currentComponent.ServiceCatalog != nil && currentComponent.ServiceCatalog.Ref != nil {
		return *currentComponent.ServiceCatalog.Ref
	}
	return ""
}

func versionDesc(ver *apiclient.StackVersion) string {
	desc := ""
	if ver.CommitHash != nil {
		desc = shortHash(*ver.CommitHash)
	}
	if ptr.Value(ver.IsLatest) {
		if desc != "" {
			desc += " "
		}
		desc += "[latest]"
	}
	return desc
}

func shortHash(h string) string {
	if len(h) > 8 {
		return h[:8]
	}
	return h
}

// CompleteCatalogRepoCommitHash provides completion for the --stack-commit-hash flag.
func CompleteCatalogRepoCommitHash(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	org, err := GetOrg(cmd)
	if err != nil {
		return cobra.AppendActiveHelp(nil, "missing org for completion: "+err.Error()), cobra.ShellCompDirectiveError
	}

	stackRef := resolveStackRefForCompletion(cmd, org)
	if stackRef == "" {
		return cobra.AppendActiveHelp(nil, "missing stack-ref for completion: please provide --stack-ref flag"),
			cobra.ShellCompDirectiveNoFileComp
	}

	api := common.NewAPI()
	m := apiclient.NewAPIClient(api)

	versions, _, err := m.ListStackVersions(org, stackRef)
	if err != nil {
		return cobra.AppendActiveHelp(nil, "cannot find versions for stack: "+stackRef+", err: "+err.Error()),
			cobra.ShellCompDirectiveNoFileComp
	}

	var commitHashes []string
	for _, version := range versions {
		if version.CommitHash != nil && strings.HasPrefix(*version.CommitHash, toComplete) {
			desc := ptr.Value(version.Name)
			if t := ptr.Value(version.Type); t != "" {
				if desc != "" {
					desc = desc + " (" + t + ")"
				} else {
					desc = t
				}
			}
			if ptr.Value(version.IsLatest) {
				desc = desc + " [latest]"
			}
			commitHashes = append(commitHashes, cobra.CompletionWithDesc(*version.CommitHash, desc))
		}
	}

	return commitHashes, cobra.ShellCompDirectiveNoFileComp
}

// CompleteStackVersionTag provides completion for the --stack-tag flag.
func CompleteStackVersionTag(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	org, err := GetOrg(cmd)
	if err != nil {
		return cobra.AppendActiveHelp(nil, "missing org for completion: "+err.Error()), cobra.ShellCompDirectiveError
	}

	stackRef := resolveStackRefForCompletion(cmd, org)
	if stackRef == "" {
		return cobra.AppendActiveHelp(nil, "missing stack-ref for completion: please provide --stack-ref flag"),
			cobra.ShellCompDirectiveNoFileComp
	}

	api := common.NewAPI()
	m := apiclient.NewAPIClient(api)

	versions, _, err := m.ListStackVersions(org, stackRef)
	if err != nil {
		return cobra.AppendActiveHelp(nil, "cannot find versions for stack: "+stackRef+", err: "+err.Error()),
			cobra.ShellCompDirectiveNoFileComp
	}

	var tags []string
	for _, version := range versions {
		if ptr.Value(version.Type) == "tag" && version.Name != nil && strings.HasPrefix(*version.Name, toComplete) {
			tags = append(tags, cobra.CompletionWithDesc(*version.Name, versionDesc(version)))
		}
	}

	return tags, cobra.ShellCompDirectiveNoFileComp
}

// CompleteStackVersionBranch provides completion for the --stack-branch flag.
func CompleteStackVersionBranch(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	org, err := GetOrg(cmd)
	if err != nil {
		return cobra.AppendActiveHelp(nil, "missing org for completion: "+err.Error()), cobra.ShellCompDirectiveError
	}

	stackRef := resolveStackRefForCompletion(cmd, org)
	if stackRef == "" {
		return cobra.AppendActiveHelp(nil, "missing stack-ref for completion: please provide --stack-ref flag"),
			cobra.ShellCompDirectiveNoFileComp
	}

	api := common.NewAPI()
	m := apiclient.NewAPIClient(api)

	versions, _, err := m.ListStackVersions(org, stackRef)
	if err != nil {
		return cobra.AppendActiveHelp(nil, "cannot find versions for stack: "+stackRef+", err: "+err.Error()),
			cobra.ShellCompDirectiveNoFileComp
	}

	var branches []string
	for _, version := range versions {
		if ptr.Value(version.Type) == "branch" && version.Name != nil && strings.HasPrefix(*version.Name, toComplete) {
			branches = append(branches, cobra.CompletionWithDesc(*version.Name, versionDesc(version)))
		}
	}

	return branches, cobra.ShellCompDirectiveNoFileComp
}
