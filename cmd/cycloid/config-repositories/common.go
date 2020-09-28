package configRepositories

import (
	"github.com/spf13/cobra"
)

var (
	nameFlag, branchFlag, urlFlag string
	defaultFlag                   bool
)

func WithFlagName(cmd *cobra.Command) string {
	flagName := "name"
	cmd.Flags().StringVar(&nameFlag, flagName, "", "name")
	return flagName
}

func WithFlagBranch(cmd *cobra.Command) string {
	flagName := "branch"
	cmd.Flags().StringVar(&branchFlag, flagName, "", "git branch")
	return flagName
}

func WithFlagURL(cmd *cobra.Command) string {
	flagName := "url"
	cmd.Flags().StringVar(&urlFlag, flagName, "", "git url")
	return flagName
}

func WithFlagDefault(cmd *cobra.Command) string {
	flagName := "default"
	cmd.Flags().BoolVar(&defaultFlag, flagName, false, "use it as default config repo")
	return flagName
}
