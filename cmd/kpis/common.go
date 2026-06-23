package kpis

import (
	"github.com/spf13/cobra"
)

var (
	nameFlag   string
	typeFlag   string
	widgetFlag string
	jobFlag    string
	configFlag string
)

func WithFlagName(cmd *cobra.Command) string {
	flagName := "name"
	cmd.Flags().StringVar(&nameFlag, flagName, "", "")
	return flagName
}

func WithFlagConfig(cmd *cobra.Command) string {
	flagName := "config"
	cmd.Flags().StringVar(&configFlag, flagName, "", "")
	return flagName
}

func WithFlagType(cmd *cobra.Command) string {
	flagName := "type"
	cmd.Flags().StringVar(&typeFlag, flagName, "", "")
	return flagName
}

func WithFlagWidget(cmd *cobra.Command) string {
	flagName := "widget"
	cmd.Flags().StringVar(&widgetFlag, flagName, "", "")
	return flagName
}

func WithFlagJob(cmd *cobra.Command) string {
	flagName := "job"
	cmd.Flags().StringVar(&jobFlag, flagName, "", "")
	return flagName
}
