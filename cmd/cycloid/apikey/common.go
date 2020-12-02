package apikey

import "github.com/spf13/cobra"

var (
	canonical   string
	name        string
	description string
	roleID      uint32
)

func WithFlagName(cmd *cobra.Command) {
	cmd.Flags().StringVar(&name, "name", "", "name of the API key")
}

func WithFlagDescription(cmd *cobra.Command) {
	cmd.Flags().StringVar(&description, "description", "", "description of the API key")
}

func WithFlagRoleID(cmd *cobra.Command) {
	cmd.Flags().Uint32Var(&roleID, "role-id", 0, "ID of the role")
}

func WithFlagCanonical(cmd *cobra.Command) {
	cmd.Flags().StringVar(&canonical, "canonical", "", "canonical of the API key")
}
