package api_key

import "github.com/spf13/cobra"

var (
	canonical   string
	name        string
	description string
	role        string
)

func WithFlagName(cmd *cobra.Command) {
	cmd.Flags().StringVar(&name, "name", "", "name of the API key")
}

func WithFlagDescription(cmd *cobra.Command) {
	cmd.Flags().StringVar(&description, "description", "", "description of the API key")
}

func WithFlagRole(cmd *cobra.Command) {
	cmd.Flags().StringVar(&role, "role", "", "Canonical of the role")
}

func WithFlagCanonical(cmd *cobra.Command) {
	cmd.Flags().StringVar(&canonical, "canonical", "", "canonical of the API key")
}
