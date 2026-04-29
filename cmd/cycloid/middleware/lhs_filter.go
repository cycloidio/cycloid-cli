package middleware

import (
	"fmt"
	"strings"
)

// LHSFilter represents a single LHS bracket filter: attribute[condition]=value.
// Encoded as a raw query parameter with literal brackets so the API receives
// exactly attribute[condition]=value (not attribute%5Bcondition%5D=value).
type LHSFilter struct {
	Attribute string
	Condition string
	Value     string
}

// buildLHSFilterQuery builds a raw query string for LHS bracket filter params.
// Brackets are kept literal in keys; values use lhsEscapeValue to preserve
// regex metacharacters while encoding structurally significant chars.
func buildLHSFilterQuery(filters []LHSFilter) string {
	parts := make([]string, 0, len(filters))
	for _, f := range filters {
		key := f.Attribute + "[" + f.Condition + "]"
		parts = append(parts, key+"="+lhsEscapeValue(f.Value))
	}
	return strings.Join(parts, "&")
}

// lhsEscapeValue percent-encodes only characters that are structurally
// significant in query strings (&, =, #, space, control chars), preserving
// regex metacharacters such as ?, *, +, [, ], (, ), {, }, |, ^, $, \.
func lhsEscapeValue(s string) string {
	var buf strings.Builder
	for i := 0; i < len(s); i++ {
		c := s[i]
		switch {
		case c == '&' || c == '=' || c == '#' || c == ' ':
			fmt.Fprintf(&buf, "%%%02X", c)
		case c < 0x20 || c >= 0x7F:
			fmt.Fprintf(&buf, "%%%02X", c)
		default:
			buf.WriteByte(c)
		}
	}
	return buf.String()
}
