package roles

import (
	"fmt"
	"strings"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/printer/factory"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func NewCreateCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:  "create",
		Args: cobra.NoArgs,
		Example: strings.Join([]string{
			"cy", "roles", "create",
			"--name", `"New role"`, "--description", `"My cool role."`,
			"--rule-json", `'{"action": "organization:delete", "effect": "allow", "resources": []}'`,
			"--rule-file", "my_rule_file.json",
		}, " "),
		Short: "Create an new role.",
		Long:  `Please check the API for rules specifications: https://docs.cycloid.io/api/#tag/Organization-Roles/operation/createRole `,
		RunE:  createRole,
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
