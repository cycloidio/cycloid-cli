package login

import "github.com/spf13/cobra"

var (
	apiKeyFlag   string
	orgFlag      string
)

func WithFlagAPIKey(cmd *cobra.Command) string {
	flagName := "api-key"
	cmd.Flags().StringVar(&apiKeyFlag, flagName, "", "API key")
	return flagName
}

func WithFlagOrg(cmd *cobra.Command) string {
	flagName := "org"
	cmd.Flags().StringVar(&orgFlag, flagName, "", "Org cannonical name")
	return flagName
}
