package printer

// Options carries per-command display hints for the table printer.
// Other printers (JSON, YAML, JQ, field) ignore these fields.
type Options struct {
	// Columns lists the default columns for table output.
	// Empty = reflection fallback (backwards compat during migration).
	Columns []string

	// Identifier is the primary identifier column name.
	// Always shown even when terminal width forces truncation.
	Identifier string

	// Transform flattens/computes display values for table output.
	// Called per-item with the raw model object.
	// If nil, the table printer uses reflection directly.
	Transform func(obj interface{}) map[string]string
}

// TableOptions holds parsed options from the --output table grammar.
// Grammar: table[=col1,col2][:key=val|:flag]*
// Set by the factory when parsing the --output flag value.
type TableOptions struct {
	// Columns overrides per-command column defaults (from --output table=col1,col2).
	Columns []string
	// NoHeader suppresses the header row (from --output table:noheader).
	NoHeader bool
	// Border enables the rounded-border grid style (from --output table:border).
	Border bool
	// Future: Sort string, Limit int, Wide bool, Compact bool
}
