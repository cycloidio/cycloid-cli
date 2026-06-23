package cyargs

import (
	"github.com/spf13/cobra"
)

// AddRepoBranchFlag registers --branch for catalog/config repository commands.
func AddRepoBranchFlag(cmd *cobra.Command) string {
	const flagName = "branch"
	cmd.Flags().String(flagName, "", "git branch")
	return flagName
}

func GetRepoBranch(cmd *cobra.Command) (string, error) {
	return cmd.Flags().GetString("branch")
}

// AddRepoURLFlag registers --url for catalog/config repository commands.
func AddRepoURLFlag(cmd *cobra.Command) string {
	const flagName = "url"
	cmd.Flags().String(flagName, "", "git url")
	return flagName
}

func GetRepoURL(cmd *cobra.Command) (string, error) {
	return cmd.Flags().GetString("url")
}

// AddRepoCredFlag registers --cred for catalog/config repository commands.
// Completion lists credential canonicals.
func AddRepoCredFlag(cmd *cobra.Command) string {
	const flagName = "cred"
	cmd.Flags().String(flagName, "", "credential canonical used to access the repository")
	cmd.RegisterFlagCompletionFunc(flagName, CompleteCredentialCanonical)
	return flagName
}

func GetRepoCred(cmd *cobra.Command) (string, error) {
	return cmd.Flags().GetString("cred")
}

// AddRepoDefaultFlag registers --default for config repository commands.
func AddRepoDefaultFlag(cmd *cobra.Command) string {
	const flagName = "default"
	cmd.Flags().Bool(flagName, false, "set as default config repository for the org")
	return flagName
}

func GetRepoDefault(cmd *cobra.Command) (bool, error) {
	return cmd.Flags().GetBool("default")
}
