package infrapolicies

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/printer/factory"
)

// NewCreateCommand returns the cobra command
// to create a new infrapolicy using a file
// Note! For boolean flags it is required var=bool
//
//	https://github.com/spf13/cobra/issues/613
func NewCreateCommand() *cobra.Command {

	var cmd = &cobra.Command{
		Use:   "create",
		Args:  cobra.NoArgs,
		Short: "create a infrapolicy",
		Example: `
	# create a infrapolicy my-policy
	cy --org my-org ip create \
	   --policy-path /path/to/file/file.rego \
	   --name my-policy
	   --description "an awesome infrapolicy" \
	   --owner user_canonical \
	   --severity "advisory" \
	   --enabled=true
		`,
		RunE: create,
	}
	common.RequiredFlag(WithFlagPolicyPath, cmd)
	common.RequiredFlag(WithFlagName, cmd)
	common.RequiredFlag(WithFlagOwner, cmd)
	common.RequiredFlag(WithFlagSeverity, cmd)

	WithFlagcanonical(cmd)
	WithFlagDescription(cmd)
	WithFlagEnabled(cmd)

	return cmd

}

func create(cmd *cobra.Command, args []string) error {
	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}

	policyPath, err := cmd.Flags().GetString("policy-path")
	if err != nil {
		return err
	}

	name, err := cmd.Flags().GetString("name")
	if err != nil {
		return err
	}

	owner, err := cmd.Flags().GetString("owner")
	if err != nil {
		return err
	}

	severity, err := cmd.Flags().GetString("severity")
	if err != nil {
		return err
	}

	canonical, err := cmd.Flags().GetString("canonical")
	if err != nil {
		return err
	}

	description, err := cmd.Flags().GetString("description")
	if err != nil {
		return err
	}

	enabled, err := cmd.Flags().GetBool("enabled")
	if err != nil {
		return err
	}

	//to allow to specify the output flag as specified in cmd/cycloid.go
	output, err := cmd.Flags().GetString("output")
	if err != nil {
		return errors.Wrap(err, "unable to get output flag")
	}

	// fetch the printer from the factory
	p, err := factory.GetPrinter(output)
	if err != nil {
		return errors.Wrap(err, "unable to get printer")
	}

	res, err := m.CreateInfraPolicy(org, policyPath, canonical, description, name, owner, severity, enabled)
	return printer.SmartPrint(p, res, err, "unable to create infrapolicy", printer.Options{}, cmd.OutOrStdout())

}
