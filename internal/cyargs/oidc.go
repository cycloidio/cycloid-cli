package cyargs

import (
	"github.com/spf13/cobra"
)

// OIDC flag names. Exported so commands can test cmd.Flags().Changed(...) for
// read-merge-write semantics on the full-replace settings endpoint.
const (
	OIDCGroupNameFlagName     = "group-name"
	OIDCMappingIDFlagName     = "mapping-id"
	OIDCManagedFlagName       = "oidc-managed"
	OIDCNoMatchPolicyFlagName = "no-match-policy"
	OIDCDefaultRoleFlagName   = "default-role"
)

// AddOIDCGroupNameFlag registers --group-name for OIDC group mapping commands.
func AddOIDCGroupNameFlag(cmd *cobra.Command) string {
	cmd.Flags().String(OIDCGroupNameFlagName, "", "OIDC group claim value to map")
	return OIDCGroupNameFlagName
}

// GetOIDCGroupName reads the --group-name flag.
func GetOIDCGroupName(cmd *cobra.Command) (string, error) {
	return cmd.Flags().GetString(OIDCGroupNameFlagName)
}

// AddOIDCMappingIDFlag registers --mapping-id for OIDC group mapping delete.
func AddOIDCMappingIDFlag(cmd *cobra.Command) string {
	cmd.Flags().Uint32(OIDCMappingIDFlagName, 0, "OIDC group mapping ID")
	return OIDCMappingIDFlagName
}

// GetOIDCMappingID reads the --mapping-id flag.
func GetOIDCMappingID(cmd *cobra.Command) (uint32, error) {
	return cmd.Flags().GetUint32(OIDCMappingIDFlagName)
}

// AddOIDCManagedFlag registers --oidc-managed for OIDC settings commands.
func AddOIDCManagedFlag(cmd *cobra.Command) string {
	cmd.Flags().Bool(OIDCManagedFlagName, false, "make OIDC the authoritative source for org membership")
	return OIDCManagedFlagName
}

// GetOIDCManaged reads the --oidc-managed flag.
func GetOIDCManaged(cmd *cobra.Command) (bool, error) {
	return cmd.Flags().GetBool(OIDCManagedFlagName)
}

// AddOIDCNoMatchPolicyFlag registers --no-match-policy with fixed completions.
func AddOIDCNoMatchPolicyFlag(cmd *cobra.Command) string {
	cmd.Flags().String(OIDCNoMatchPolicyFlagName, "", "action for members whose OIDC groups match no team (keep_membership|eject)")
	_ = cmd.RegisterFlagCompletionFunc(OIDCNoMatchPolicyFlagName,
		cobra.FixedCompletions(
			[]string{"keep_membership", "eject"},
			cobra.ShellCompDirectiveNoFileComp,
		),
	)
	return OIDCNoMatchPolicyFlagName
}

// GetOIDCNoMatchPolicy reads the --no-match-policy flag.
func GetOIDCNoMatchPolicy(cmd *cobra.Command) (string, error) {
	return cmd.Flags().GetString(OIDCNoMatchPolicyFlagName)
}

// AddOIDCDefaultRoleFlag registers --default-role for OIDC settings.
// This is distinct from the per-member --role flag (AddMemberRoleFlag / AddRoleCanonicalFlag)
// because it configures the org-level default role granted to OIDC-provisioned members, and
// is only used on the settings subcommand.
func AddOIDCDefaultRoleFlag(cmd *cobra.Command) string {
	cmd.Flags().String(OIDCDefaultRoleFlagName, "", "role canonical granted by default to OIDC-provisioned members (empty clears it)")
	_ = cmd.RegisterFlagCompletionFunc(OIDCDefaultRoleFlagName, CompleteRoleCanonical)
	return OIDCDefaultRoleFlagName
}

// GetOIDCDefaultRole reads the --default-role flag.
func GetOIDCDefaultRole(cmd *cobra.Command) (string, error) {
	return cmd.Flags().GetString(OIDCDefaultRoleFlagName)
}
