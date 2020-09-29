package projects

import (
	"github.com/spf13/cobra"
)

var (
	nameFlag             string
	canonicalFlag        string
	descriptionFlag      string
	CloudProviderFlag    string
	stackRefFlag         string
	usecaseFlag          string
	configRepositoryFlag uint32
	configsFlag          map[string]string
	pipelineFlag         string
	varsFlag             string
)

func WithFlagConfig(cmd *cobra.Command) string {
	flagName := "config"
	cmd.Flags().StringToStringVar(&configsFlag, flagName, nil, "<file_path>=<git_dest>")
	return flagName
}
func WithFlagPipeline(cmd *cobra.Command) string {
	flagName := "pipeline"
	cmd.Flags().StringVar(&pipelineFlag, flagName, "", "")
	return flagName
}
func WithFlagVars(cmd *cobra.Command) string {
	flagName := "vars"
	cmd.Flags().StringVar(&varsFlag, flagName, "", "")
	return flagName
}
func WithFlagName(cmd *cobra.Command) string {
	flagName := "name"
	cmd.Flags().StringVar(&nameFlag, flagName, "", "")
	return flagName
}
func WithFlagUsecase(cmd *cobra.Command) string {
	flagName := "usecase"
	cmd.Flags().StringVar(&usecaseFlag, flagName, "default", "")
	return flagName
}
func WithFlagStackRef(cmd *cobra.Command) string {
	flagName := "stack-ref"
	cmd.Flags().StringVar(&stackRefFlag, flagName, "", "")
	return flagName
}
func WithFlagCanonical(cmd *cobra.Command) string {
	flagName := "canonical"
	cmd.Flags().StringVar(&canonicalFlag, flagName, "", "")
	return flagName
}
func WithFlagDescription(cmd *cobra.Command) string {
	flagName := "description"
	cmd.Flags().StringVar(&descriptionFlag, flagName, "", "")
	return flagName
}
func WithFlagCloudProvider(cmd *cobra.Command) string {
	flagName := "cloud-provider"
	cmd.Flags().StringVar(&CloudProviderFlag, flagName, "", "")
	return flagName
}

func WithFlagConfigRepository(cmd *cobra.Command) string {
	flagName := "config-repo"
	cmd.Flags().Uint32Var(&configRepositoryFlag, flagName, 0, "")
	return flagName
}
