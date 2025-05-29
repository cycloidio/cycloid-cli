package terracost

import (
	"github.com/spf13/cobra"
)

// NewCommands returns an implementation of Terracost commands
func NewCommands() *cobra.Command {
	cmd := &cobra.Command{
		Use: "terracost",
		Example: `
	# estimate the cost of a Terraform plan
	cy --org my-org terracost estimate --plan-path ./plan.json
`,
		Short: "Use terracost feature",
	}

	cmd.AddCommand(NewEstimateCommand())
	return cmd
}
