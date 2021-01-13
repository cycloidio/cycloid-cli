package infrapolicies

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
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
		RunE: func(cmd *cobra.Command, args []string) error {
			org, err := cmd.Flags().GetString("org")
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
			return validate(org, project, env, planPath, output)
		},
	}
	common.RequiredFlag(WithFlagPlanPath, cmd)
	return cmd
}

// validate will send the GET request to the API in order to
// validate the terraform Plan located in planPath
func validate(org, project, env, planPath, output string) error {
	plan, err := ioutil.ReadFile(planPath)
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
	if err != nil {
		// print the result on the standard output
		if err := p.Print(err, printer.Options{}, os.Stdout); err != nil {
			return errors.Wrap(err, "unable to print result")
		}
		return fmt.Errorf("unable to validate terraform plan file: %w", err)
	}

	// print the result on the standard output
	if err := p.Print(res, printer.Options{}, os.Stdout); err != nil {
		return fmt.Errorf("unable to print result: %w", err)
	}
	return nil
}
