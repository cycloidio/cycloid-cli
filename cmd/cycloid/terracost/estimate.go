package terracost

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/printer/factory"
)

// NewEstimateCommands returns a cobra implementation
// of the cost estimation
func NewEstimateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "estimate",
		RunE: estimate,
	}
	common.RequiredFlag(WithFlagPlanPath, cmd)
	return cmd
}

func estimate(cmd *cobra.Command, args []string) error {
	org, err := common.GetOrg(cmd)
	if err != nil {
		return fmt.Errorf("unable to validate org flag: %w", err)
	}
	planPath, err := cmd.Flags().GetString("plan-path")
	if err != nil {
		return fmt.Errorf("unable to get plan path flag: %w", err)
	}
	output, err := cmd.Flags().GetString("output")
	if err != nil {
		return fmt.Errorf("unable to get output flag: %w", err)
	}
	plan, err := os.ReadFile(planPath)
	if err != nil {
		return fmt.Errorf("unable to read terraform plan file: %w", err)
	}

	// fetch the printer from the factory
	p, err := factory.GetPrinter(output)
	if err != nil {
		return fmt.Errorf("unable to get printer: %w", err)
	}

	api := common.NewAPI()
	m := middleware.NewMiddleware(api)
	res, err := m.CostEstimation(org, plan)
	return printer.SmartPrint(p, res, err, "unable to estimate terraform plan file", printer.Options{}, cmd.OutOrStdout())
}
