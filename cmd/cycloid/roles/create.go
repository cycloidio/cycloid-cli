package roles

import (
	"fmt"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/printer/factory"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func NewCreateCommand() *cobra.Command {
	var (
		example = `cy roles create --name "New role" --`
		short   = "Get a role specification."
		long    = short
	)

	var cmd = &cobra.Command{
		Use: "create",
		Args: cobra.MatchAll(
			cobra.MaximumNArgs(1),
		),
		Example:           example,
		Short:             short,
		Long:              long,
		RunE:              getRole,
		ValidArgsFunction: cyargs.CompleteRoleCanonicals,
	}

	cmd.MarkFlagsOneRequired(
		cyargs.AddRoleNameFlag(cmd),
		cyargs.AddRoleCanonicalFlag(cmd),
	)
	cmd.MarkFlagsOneRequired(
		cyargs.AddRoleRulesJSONFlag(cmd),
		cyargs.AddRoleRulesFilesFlag(cmd),
	)
	cyargs.AddDescriptionFlag(cmd)

	// keep legacy flag just in case
	// TODO: deprecate in next update
	cmd.Flags().String("canonical", "", "the role canonical")
	cmd.Flags().MarkDeprecated("canonical", "use --role or pass the canonical as argument directly")

	return cmd
}

func createRole(cmd *cobra.Command, args []string) error {
	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}

	roleName, err := cyargs.GetRoleName(cmd)
	if err != nil {
		return err
	}

	roleCan, err := cyargs.GetRoleCanonical(cmd)
	if err != nil {
		return err
	}

	deprecatedCan, err := cmd.Flags().GetString("canonical")
	if err != nil {
		return err
	}

	if roleCan == "" && deprecatedCan != "" {
		roleCan = deprecatedCan
	}

	name, role, err := middleware.NameOrCanonical(&roleName, &roleCan)
	if err != nil {
		return err
	}

	rulesJSON, err := cyargs.GetRoleRulesJSON(cmd)
	if err != nil {
		return err
	}

	rulesFiles, err := cyargs.GetRoleRulesFiles(cmd)
	if err != nil {
		return err
	}

	var rules = append(rulesJSON, rulesFiles...)

	description, err := cyargs.GetDescription(cmd)
	if err != nil {
		return err
	}

	output, err := cyargs.GetOutput(cmd)
	if err != nil {
		return err
	}

	p, err := factory.GetPrinter(output)
	if err != nil {
		return errors.Wrap(err, "unable to get printer")
	}

	api := common.NewAPI()
	m := middleware.NewMiddleware(api)
	outRole, err := m.CreateRole(org, &name, &role, &description, rules)
	if err != nil {
		return fmt.Errorf("failed to create the role: %w", err)
	}

	return printer.SmartPrint(p, outRole, nil, "", printer.Options{}, cmd.OutOrStdout())
}
