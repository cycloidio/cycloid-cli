package cyargs

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/cycloidio/cycloid-cli/config"
)

// GetOutput resolves the effective output format using this priority order:
//  1. --jq flag (sugar for jq=<expr>)
//  2. --output / -o flag (if explicitly set by the user on the CLI)
//  3. CY_OUTPUT environment variable (via viper AutomaticEnv)
//  4. Output field in the config file (~/.config/cycloid-cli/config.yaml)
//  5. "table" (built-in default)
func GetOutput(cmd *cobra.Command) (string, error) {
	// 1. --jq takes precedence
	if jqExpr, err := cmd.Flags().GetString("jq"); err == nil && jqExpr != "" {
		return "jq=" + jqExpr, nil
	}

	// 2. --output / -o explicitly set on CLI
	if cmd.Flags().Changed("output") {
		return cmd.Flags().GetString("output")
	}

	// 3. CY_OUTPUT env var (viper AutomaticEnv maps CY_OUTPUT → "output")
	if envOutput := viper.GetString("output"); envOutput != "" {
		return envOutput, nil
	}

	// 4. Config file
	if conf, err := config.Read(); err == nil && conf.Output != "" {
		return conf.Output, nil
	}

	// 5. Default
	return "table", nil
}
