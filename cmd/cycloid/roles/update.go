package roles

import (
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/spf13/cobra"
)

func NewUpdateCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:  "update",
		Args: cobra.NoArgs,
		Example: `  cy --org my-org roles update --role my-role --name "My role" \
    --rule-json '{"action": "organization:list", "effect": "allow", "resources": []}'`,
		Short: "Create or update a role (same as roles create --update).",
		Long:  `Updates an existing role or creates it if missing. Same behavior as cy roles create --update. See API: https://docs.cycloid.io/api/#tag/Organization-Roles/operation/updateRole`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runRoleWrite(cmd, true)
		},
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

	cmd.Flags().String("canonical", "", "the role canonical")
	_ = cmd.Flags().MarkDeprecated("canonical", "use --role or pass the canonical as argument directly")

	return cmd
}
