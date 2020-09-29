package members

import (
	"github.com/spf13/cobra"
)

var (
	nameFlag, emailFlag string
	roleIDFlag          uint32
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

func WithFlagRoleID(cmd *cobra.Command) string {
	flagRoleID := "role-id"
	cmd.Flags().Uint32Var(&roleIDFlag, flagRoleID, 0, "role-id")
	return flagRoleID
}
