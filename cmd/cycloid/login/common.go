package login

import "github.com/spf13/cobra"

var (
	org      string
	email    string
	password string
	child    string
)

func WithFlagOrg(cmd *cobra.Command) {
	cmd.Flags().StringVar(&org, "org", "", "organization name")
}

func WithFlagEmail(cmd *cobra.Command) {
	cmd.Flags().StringVar(&email, "email", "", "email")
}

func WithFlagPassword(cmd *cobra.Command) {
	cmd.Flags().StringVar(&password, "password", "", "password")
}

func WithFlagChild(cmd *cobra.Command) {
	cmd.Flags().StringVar(&child, "child", "", "child organization")
}
