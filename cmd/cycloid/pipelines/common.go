package pipelines

import (
	"github.com/spf13/cobra"
)

var (
	jobFlag      string
	taskFlag     string
	pipelineFlag string
	varsFlag     string
	configsFlag  map[string]string
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

func WithFlagJob(cmd *cobra.Command) string {
	flagName := "job"
	cmd.Flags().StringVar(&jobFlag, flagName, "", "")
	return flagName
}
func WithFlagTask(cmd *cobra.Command) string {
	flagName := "task"
	cmd.Flags().StringVar(&taskFlag, flagName, "", "")
	return flagName
}
