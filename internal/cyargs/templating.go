package cyargs

import "github.com/spf13/cobra"

// AddTemplateRenderFlags registers the flags for `cy template render`: the
// template inputs (--file/--dir) and the layered context sources
// (--context-file, --context, --set). Output format is the global --output.
func AddTemplateRenderFlags(cmd *cobra.Command) {
	cmd.Flags().StringArrayP("file", "f", nil, "Template file to render (repeatable). Use - for stdin.")
	cmd.Flags().String("dir", "", "Directory of templates to render recursively.")
	cmd.Flags().String("context-file", "", "Path to a JSON or YAML file of context variables.")
	cmd.Flags().String("context", "", "Raw JSON object of context variables (highest-fidelity floor; merged below --set).")
	cmd.Flags().StringArray("set", nil, "Context variable as key=value (repeatable). Dotted keys nest, e.g. env_vars.region=eu-west-1. Highest precedence.")
}

// GetTemplateFiles returns the --file values.
func GetTemplateFiles(cmd *cobra.Command) ([]string, error) {
	return cmd.Flags().GetStringArray("file")
}

// GetTemplateDir returns the --dir value.
func GetTemplateDir(cmd *cobra.Command) (string, error) {
	return cmd.Flags().GetString("dir")
}

// GetTemplateContextFile returns the --context-file value.
func GetTemplateContextFile(cmd *cobra.Command) (string, error) {
	return cmd.Flags().GetString("context-file")
}

// GetTemplateContextString returns the --context value.
func GetTemplateContextString(cmd *cobra.Command) (string, error) {
	return cmd.Flags().GetString("context")
}

// GetTemplateSet returns the --set key=value pairs.
func GetTemplateSet(cmd *cobra.Command) ([]string, error) {
	return cmd.Flags().GetStringArray("set")
}
