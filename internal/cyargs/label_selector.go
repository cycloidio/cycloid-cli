package cyargs

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/gen/models"
	"github.com/cycloidio/cycloid-cli/utils/ptr"
)

const (
	labelSelectorEnforcementFlag = "label-selector-enforcement"
	labelSelectorRequirementFlag = "label-selector-requirement"
	deleteLabelSelectorFlag      = "delete-label-selector"
)

// AddLabelSelectorFlags registers --label-selector-enforcement,
// --label-selector-requirement (repeatable), and --delete-label-selector
// on the given command.
func AddLabelSelectorFlags(cmd *cobra.Command) {
	cmd.Flags().String(labelSelectorEnforcementFlag, "", "label selector enforcement mode: 'soft' or 'hard'")
	_ = cmd.RegisterFlagCompletionFunc(labelSelectorEnforcementFlag, func(cmd *cobra.Command, args []cobra.Completion, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{"soft", "hard"}, cobra.ShellCompDirectiveNoFileComp
	})
	cmd.Flags().StringArray(labelSelectorRequirementFlag, nil,
		"label selector requirement as key:operator:value1,value2 (repeatable). "+
			"Operator must be 'eq' or 'in'. Example: env-type:in:production,staging")
	cmd.Flags().Bool(deleteLabelSelectorFlag, false, "remove the label selector from the environment type")

	cmd.MarkFlagsMutuallyExclusive(deleteLabelSelectorFlag, labelSelectorEnforcementFlag)
	cmd.MarkFlagsMutuallyExclusive(deleteLabelSelectorFlag, labelSelectorRequirementFlag)
}

// AddLabelSelectorCreateFlags registers label selector flags for the create
// command (no --delete-label-selector since there's nothing to delete yet).
func AddLabelSelectorCreateFlags(cmd *cobra.Command) {
	cmd.Flags().String(labelSelectorEnforcementFlag, "", "label selector enforcement mode: 'soft' or 'hard'")
	_ = cmd.RegisterFlagCompletionFunc(labelSelectorEnforcementFlag, func(cmd *cobra.Command, args []cobra.Completion, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{"soft", "hard"}, cobra.ShellCompDirectiveNoFileComp
	})
	cmd.Flags().StringArray(labelSelectorRequirementFlag, nil,
		"label selector requirement as key:operator:value1,value2 (repeatable). "+
			"Operator must be 'eq' or 'in'. Example: env-type:in:production,staging")
}

// GetDeleteLabelSelector returns true if --delete-label-selector was set.
func GetDeleteLabelSelector(cmd *cobra.Command) bool {
	v, _ := cmd.Flags().GetBool(deleteLabelSelectorFlag)
	return v
}

// HasLabelSelectorFlags returns true if any label selector flag was set.
func HasLabelSelectorFlags(cmd *cobra.Command) bool {
	return IsSet(cmd, labelSelectorEnforcementFlag) || IsSet(cmd, labelSelectorRequirementFlag)
}

// GetLabelSelector parses the label selector flags and returns a LabelSelector model.
// Returns an error if the flags are inconsistent (e.g. requirements without enforcement).
func GetLabelSelector(cmd *cobra.Command) (*models.LabelSelector, error) {
	if !HasLabelSelectorFlags(cmd) {
		return nil, nil
	}

	enforcement, err := cmd.Flags().GetString(labelSelectorEnforcementFlag)
	if err != nil {
		return nil, err
	}

	requirements, err := cmd.Flags().GetStringArray(labelSelectorRequirementFlag)
	if err != nil {
		return nil, err
	}

	if len(requirements) == 0 {
		return nil, fmt.Errorf("--%s requires at least one --%s", labelSelectorEnforcementFlag, labelSelectorRequirementFlag)
	}

	if enforcement == "" {
		return nil, fmt.Errorf("--%s is required when --%s is set", labelSelectorEnforcementFlag, labelSelectorRequirementFlag)
	}

	if enforcement != "soft" && enforcement != "hard" {
		return nil, fmt.Errorf("--%s must be 'soft' or 'hard', got %q", labelSelectorEnforcementFlag, enforcement)
	}

	parsedReqs, err := parseLabelSelectorRequirements(requirements)
	if err != nil {
		return nil, err
	}

	return &models.LabelSelector{
		Enforcement:  ptr.Ptr(enforcement),
		Requirements: parsedReqs,
	}, nil
}

// parseLabelSelectorRequirements parses requirement strings in the format
// "key:operator:value1,value2" into model objects.
func parseLabelSelectorRequirements(raw []string) ([]*models.LabelSelectorRequirement, error) {
	reqs := make([]*models.LabelSelectorRequirement, 0, len(raw))
	for _, s := range raw {
		req, err := parseSingleRequirement(s)
		if err != nil {
			return nil, err
		}
		reqs = append(reqs, req)
	}
	return reqs, nil
}

// parseSingleRequirement parses "key:operator:value1,value2".
func parseSingleRequirement(s string) (*models.LabelSelectorRequirement, error) {
	// Split into exactly 3 parts: key, operator, values
	parts := strings.SplitN(s, ":", 3)
	if len(parts) != 3 {
		return nil, fmt.Errorf("invalid requirement format %q: expected key:operator:value1,value2", s)
	}

	key := strings.TrimSpace(parts[0])
	operator := strings.TrimSpace(parts[1])
	valuesRaw := strings.TrimSpace(parts[2])

	if key == "" {
		return nil, fmt.Errorf("invalid requirement %q: key cannot be empty", s)
	}

	if operator != "eq" && operator != "in" {
		return nil, fmt.Errorf("invalid requirement %q: operator must be 'eq' or 'in', got %q", s, operator)
	}

	if valuesRaw == "" {
		return nil, fmt.Errorf("invalid requirement %q: at least one value is required", s)
	}

	values := strings.Split(valuesRaw, ",")
	for i, v := range values {
		values[i] = strings.TrimSpace(v)
		if values[i] == "" {
			return nil, fmt.Errorf("invalid requirement %q: empty value at position %d", s, i+1)
		}
	}

	if operator == "eq" && len(values) != 1 {
		return nil, fmt.Errorf("invalid requirement %q: operator 'eq' requires exactly one value, got %d", s, len(values))
	}

	return &models.LabelSelectorRequirement{
		Key:      &key,
		Operator: &operator,
		Values:   values,
	}, nil
}
