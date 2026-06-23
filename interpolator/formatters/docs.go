package formatters

const (
	// Docs contains the documentation of Formatting parameters for the end user
	// must be indented at 2.
	Docs = `  Output parameter:

  To specify the output format only the key is required. e.g. cy://path?output=yaml will output in yaml
  each format key can be paired with options as values e.g. cy://path?output=yaml&indent_size=2 set the indent to 2
  here is all the format and their options

  json
    Format the API output in JSON.

  yaml
    Format the API output in YAML.

	No matter which output format, any simple value (string, number, bool) will be output
	as is.

	Extra output related parameters:

		json_compact (only for json output)
			Will output a compact json on one line like {"hello":"world"}

		json_escape  (only for json output)
			Will escape the output json like "{\"hello\":\"world\"}"

    indent_size=<int>
			Set the indentation level for the output (for YAML or JSON)
`
)
