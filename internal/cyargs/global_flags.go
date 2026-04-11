package cyargs

import "github.com/spf13/cobra"

func GetOutput(cmd *cobra.Command) (string, error) {
	// --jq <expr> is sugar for --output jq=<expr>
	if jqExpr, err := cmd.Flags().GetString("jq"); err == nil && jqExpr != "" {
		return "jq=" + jqExpr, nil
	}

	output, err := cmd.Flags().GetString("output")
	if err != nil {
		return "", err
	}

	return output, nil
}
