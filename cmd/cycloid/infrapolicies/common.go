package infrapolicies

import "github.com/spf13/cobra"

var planPath string

func WithFlagPlanPath(cmd *cobra.Command) string {
	cmd.PersistentFlags().StringVar(&planPath, "plan-path", "", "Path to the terraform plan result")
	return "plan-path"
}
