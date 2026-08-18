package terracost

import "github.com/spf13/cobra"

var planPath string

// WithFlagPlanPath binds the `planPath` variable to its flag
func WithFlagPlanPath(cmd *cobra.Command) string {
	cmd.PersistentFlags().StringVar(&planPath, "plan-path", "", "Path to the terraform plan file")
	return "plan-path"
}
