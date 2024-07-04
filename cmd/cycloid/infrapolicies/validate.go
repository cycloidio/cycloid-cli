package infrapolicies

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/internal"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/printer/factory"
)

// NewValidateCommand returns the cobra command holding
// the validate Terraform plan against infrapolicy
func NewValidateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "validate",
		Short: "validate Terraform plan against Cycloid Infrapolicy",
		Example: `
	# validate saved terraform plan against the infra policies rule in my-org/my-project/my-env
	cy --org my-org --project my-project --env my-env ip validate --plan-path terraform-plan.json
`,
		RunE:    validate,
		PreRunE: internal.CheckAPIAndCLIVersion,
	}
	common.RequiredFlag(WithFlagPlanPath, cmd)
	common.RequiredFlag(common.WithFlagProject, cmd)
	common.RequiredFlag(common.WithFlagEnv, cmd)
	return cmd
}

// validate will send the GET request to the API in order to
// validate the terraform Plan located in planPath
func validate(cmd *cobra.Command, args []string) error {
	org, err := common.GetOrg(cmd)
	if err != nil {
		return fmt.Errorf("unable to validate org flag: %w", err)
	}
	project, err := cmd.Flags().GetString("project")
	if err != nil {
		return fmt.Errorf("unable to get project flag: %w", err)
	}
	env, err := cmd.Flags().GetString("env")
	if err != nil {
		return fmt.Errorf("unable to get env flag: %w", err)
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
	res, err := m.ValidateInfraPolicies(org, project, env, plan)
	return printer.SmartPrint(p, res, err, "unable to validate terraform plan file", printer.Options{}, cmd.OutOrStdout())
}
