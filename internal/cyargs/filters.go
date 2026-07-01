package cyargs

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
)

const lhsFilterFlagName = "filter"

// AddLHSFilterFlag registers a repeatable --filter flag for LHS bracket filters.
// The expected syntax is: attribute[condition]=value (e.g. output_type[eq]=string).
// Shared across any List command that forwards filters to a middleware method.
func AddLHSFilterFlag(cmd *cobra.Command) string {
	cmd.Flags().StringArray(lhsFilterFlagName, nil,
		`LHS bracket filter in the form attribute[condition]=value (repeatable).
Conditions: eq, ne, gt, gte, lt, lte, like, rlike.
Example: --filter 'output_type[eq]=string' --filter 'project_canonical[eq]=my-project'`)
	_ = cmd.RegisterFlagCompletionFunc(lhsFilterFlagName, cobra.NoFileCompletions)
	return lhsFilterFlagName
}

// GetLHSFilters parses the --filter flag values into middleware.LHSFilter structs.
// Each token must follow the form: attribute[condition]=value
func GetLHSFilters(cmd *cobra.Command) ([]middleware.LHSFilter, error) {
	tokens, err := cmd.Flags().GetStringArray(lhsFilterFlagName)
	if err != nil {
		return nil, err
	}

	filters := make([]middleware.LHSFilter, 0, len(tokens))
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
func parseLHSFilter(tok string) (middleware.LHSFilter, error) {
	bracketOpen := strings.IndexByte(tok, '[')
	if bracketOpen < 1 {
		return middleware.LHSFilter{}, fmt.Errorf("invalid --filter %q: expected format attribute[condition]=value", tok)
	}

	closeEq := strings.Index(tok[bracketOpen:], "]=")
	if closeEq < 0 {
		return middleware.LHSFilter{}, fmt.Errorf("invalid --filter %q: expected format attribute[condition]=value", tok)
	}
	closeEq += bracketOpen // absolute offset of ']'

	attribute := tok[:bracketOpen]
	condition := tok[bracketOpen+1 : closeEq]
	value := tok[closeEq+2:] // skip "]="

	if condition == "" {
		return middleware.LHSFilter{}, fmt.Errorf("invalid --filter %q: condition must not be empty", tok)
	}

	return middleware.LHSFilter{
		Attribute: attribute,
		Condition: condition,
		Value:     value,
	}, nil
}
