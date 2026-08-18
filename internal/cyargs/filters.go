package cyargs

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/apiclient"
)

const lhsFilterFlagName = "filter"

// AddLHSFilterFlag registers a repeatable --filter flag for LHS bracket filters.
// The expected syntax is: attribute[condition]=value (e.g. output_type[eq]=string).
// Shared across any List command that forwards filters to a apiclient method.
func AddLHSFilterFlag(cmd *cobra.Command) string {
	cmd.Flags().StringArray(lhsFilterFlagName, nil,
		`LHS bracket filter in the form attribute[condition]=value (repeatable).
Conditions: eq, ne, gt, gte, lt, lte, like, rlike.
Example: --filter 'output_type[eq]=string' --filter 'project_canonical[eq]=my-project'`)
	_ = cmd.RegisterFlagCompletionFunc(lhsFilterFlagName, cobra.NoFileCompletions)
	return lhsFilterFlagName
}

// GetLHSFilters parses the --filter flag values into apiclient.LHSFilter structs.
// Each token must follow the form: attribute[condition]=value
func GetLHSFilters(cmd *cobra.Command) ([]apiclient.LHSFilter, error) {
	tokens, err := cmd.Flags().GetStringArray(lhsFilterFlagName)
	if err != nil {
		return nil, err
	}

	filters := make([]apiclient.LHSFilter, 0, len(tokens))
	for _, tok := range tokens {
		f, err := parseLHSFilter(tok)
		if err != nil {
			return nil, err
		}
		filters = append(filters, f)
	}
	return filters, nil
}

// parseLHSFilter parses a single "attribute[condition]=value" token.
func parseLHSFilter(tok string) (apiclient.LHSFilter, error) {
	bracketOpen := strings.IndexByte(tok, '[')
	if bracketOpen < 1 {
		return apiclient.LHSFilter{}, fmt.Errorf("invalid --filter %q: expected format attribute[condition]=value", tok)
	}

	closeEq := strings.Index(tok[bracketOpen:], "]=")
	if closeEq < 0 {
		return apiclient.LHSFilter{}, fmt.Errorf("invalid --filter %q: expected format attribute[condition]=value", tok)
	}
	closeEq += bracketOpen // absolute offset of ']'

	attribute := tok[:bracketOpen]
	condition := tok[bracketOpen+1 : closeEq]
	value := tok[closeEq+2:] // skip "]="

	if condition == "" {
		return apiclient.LHSFilter{}, fmt.Errorf("invalid --filter %q: condition must not be empty", tok)
	}

	return apiclient.LHSFilter{
		Attribute: attribute,
		Condition: condition,
		Value:     value,
	}, nil
}
