package roles

import (
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/spf13/cobra"
)

func NewCreateCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:  "create",
		Args: cobra.NoArgs,
		Example: `  cy --org my-org roles create --name "New role" --description "My cool role." \
    --rule-json '{"action": "organization:delete", "effect": "allow", "resources": []}' --rule-file my_rule_file.json

  cy --org my-org roles create --update --role my-role --name "My role" \
    --rule-json '{"action": "organization:list", "effect": "allow", "resources": []}'`,
		Short: "Create a new role.",
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
	cmd.Flags().Bool("update", false, "If set, update the role when it already exists.")

	// keep legacy flag just in case
	// TODO: deprecate in next update
	cmd.Flags().String("canonical", "", "the role canonical")
	_ = cmd.Flags().MarkDeprecated("canonical", "use --role or pass the canonical as argument directly")

	return cmd
}

func createRole(cmd *cobra.Command, args []string) error {
	update, err := cmd.Flags().GetBool("update")
	if err != nil {
		return err
	}
	return runRoleWrite(cmd, update)
}
