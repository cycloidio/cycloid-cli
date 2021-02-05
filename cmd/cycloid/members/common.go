package members

import (
	"github.com/spf13/cobra"
)

var (
	nameFlag, emailFlag, roleFlag string
)

func WithFlagName(cmd *cobra.Command) string {
	flagName := "name"
	cmd.Flags().StringVar(&nameFlag, flagName, "", "name")
	return flagName
}

func WithFlagEmail(cmd *cobra.Command) string {
	flagName := "email"
	cmd.Flags().StringVar(&emailFlag, flagName, "", "email")
	return flagName
}

func WithFlagRoleCanonical(cmd *cobra.Command) string {
	flagName := "role"
	cmd.Flags().StringVar(&roleFlag, flagName, "", "role")
	return flagName
}
