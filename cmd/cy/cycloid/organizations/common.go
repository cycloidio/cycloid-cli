package organizations

import "github.com/spf13/cobra"

var (
	nameFlag               string
	canonicalFlag          string
	parentOrganizationFlag string
)

func WithFlagName(cmd *cobra.Command) string {
	flagName := "name"
	cmd.Flags().StringVar(&nameFlag, flagName, "", "name")
	return flagName
}

func WithPersistentFlagCanonical(cmd *cobra.Command) string {
	flagName := "canonical"
	cmd.PersistentFlags().StringVar(&canonicalFlag, flagName, "", "Canonical")
	return flagName
}

func WithFlagParentOrganization(cmd *cobra.Command) string {
	flagName := "parent-org"
	cmd.Flags().StringVar(&parentOrganizationFlag, flagName, "", "parent organization")
	return flagName
}
