package terracost

import (
	"fmt"
	"io/ioutil"
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
		Use: "estimate",
		RunE: func(cmd *cobra.Command, args []string) error {
			org, err := cmd.Flags().GetString("org")
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
			return estimate(org, planPath, output)
		},
	}
	common.RequiredFlag(WithFlagPlanPath, cmd)
	return cmd
}

func estimate(org, planPath, output string) error {
	plan, err := ioutil.ReadFile(planPath)
	if err != nil {
		return fmt.Errorf("unable to read terraform plan file: %w", err)
	}

	api := common.NewAPI()
	m := middleware.NewMiddleware(api)
	res, err := m.CostEstimation(org, plan)
	if err != nil {
		return fmt.Errorf("unable to estimate terraform plan file: %w", err)
	}

	// fetch the printer from the factory
	p, err := factory.GetPrinter(output)
	if err != nil {
		return fmt.Errorf("unable to get printer: %w", err)
	}

	// print the result on the standard output
	if err := p.Print(res, printer.Options{}, os.Stdout); err != nil {
		return fmt.Errorf("unable to print result: %w", err)
	}
	return nil
}