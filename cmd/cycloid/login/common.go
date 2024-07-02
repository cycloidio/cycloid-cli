package login

import "github.com/spf13/cobra"

var (
	apiKeyFlag string
	orgFlag    string
)

func WithFlagAPIKey(cmd *cobra.Command) string {
	flagName := "api-key"
	cmd.Flags().StringVar(&apiKeyFlag, flagName, "", "API key")
	return flagName
}
