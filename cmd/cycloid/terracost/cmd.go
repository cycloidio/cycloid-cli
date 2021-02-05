package terracost

import (
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
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

	common.RequiredFlag(common.WithFlagOrg, cmd)

	cmd.AddCommand(NewEstimateCommand())
	return cmd
}
