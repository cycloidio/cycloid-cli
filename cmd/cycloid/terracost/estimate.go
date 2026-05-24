package terracost

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
	"github.com/cycloidio/cycloid-cli/printer"
)

// NewEstimateCommands returns a cobra implementation
// of the cost estimation
func NewEstimateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "estimate",
		Args: cobra.NoArgs,
		RunE: estimate,
	}
	common.RequiredFlag(WithFlagPlanPath, cmd)
	return cmd
}

func estimate(cmd *cobra.Command, args []string) error {
	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return fmt.Errorf("unable to validate org flag: %w", err)
	}
	planPath, err := cmd.Flags().GetString("plan-path")
	if err != nil {
		return fmt.Errorf("unable to get plan path flag: %w", err)
	}
	plan, err := os.ReadFile(planPath)
	if err != nil {
		return fmt.Errorf("unable to read terraform plan file: %w", err)
	}

	api := common.NewAPI()
	m := middleware.NewMiddleware(api)
	res, _, err := m.CostEstimation(org, plan)
	return cyout.PrintWithOptions(cmd, res, err, "unable to estimate terraform plan file", printer.Options{})
}
