package cyargs

import (
	"os"

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

// OIDC integration flag names. Exported so commands can test
// cmd.Flags().Changed(<name>) for read-merge-write semantics.
const (
	OIDCIssuerFlagName                 = "issuer"
	OIDCClientIDFlagName               = "client-id"
	OIDCClientSecretFlagName           = "client-secret"
	OIDCDisplayNameFlagName            = "display-name"
	OIDCIconFlagName                   = "icon"
	OIDCDiscoveryURLFlagName           = "discovery-url"
	OIDCGroupsClaimNameFlagName        = "groups-claim-name"
	OIDCSessionTTLSecondsFlagName      = "session-ttl-seconds"
	OIDCClientSecretJwtFlagName        = "client-secret-jwt"
	OIDCUseCaCertFlagName              = "use-ca-cert"
	OIDCCaCertFlagName                 = "ca-cert"
	OIDCSkipTLSVerifyFlagName          = "skip-tls-verify"
	OIDCAllowInsecureDiscoveryFlagName = "allow-insecure-discovery"
	OIDCEnabledFlagName                = "enabled"
)

// AddOIDCIssuerFlag registers --issuer for OIDC integration commands.
func AddOIDCIssuerFlag(cmd *cobra.Command) {
	cmd.Flags().String(OIDCIssuerFlagName, "", "OIDC issuer URL")
}

// GetOIDCIssuer reads the --issuer flag.
func GetOIDCIssuer(cmd *cobra.Command) (string, error) {
	return cmd.Flags().GetString(OIDCIssuerFlagName)
}

// AddOIDCClientIDFlag registers --client-id for OIDC integration commands.
func AddOIDCClientIDFlag(cmd *cobra.Command) {
	cmd.Flags().String(OIDCClientIDFlagName, "", "OIDC client ID")
}

// GetOIDCClientID reads the --client-id flag.
func GetOIDCClientID(cmd *cobra.Command) (string, error) {
	return cmd.Flags().GetString(OIDCClientIDFlagName)
}

// AddOIDCClientSecretFlag registers --client-secret for OIDC integration commands.
// The secret can also be supplied via the CY_OIDC_CLIENT_SECRET environment variable.
func AddOIDCClientSecretFlag(cmd *cobra.Command) {
	cmd.Flags().String(OIDCClientSecretFlagName, "", "OIDC client secret (also read from CY_OIDC_CLIENT_SECRET env var)")
}

// GetOIDCClientSecret returns the OIDC client secret: flag value if the flag was
// explicitly set, otherwise the CY_OIDC_CLIENT_SECRET environment variable.
// An empty string means "no secret supplied" — the backend then preserves the
// stored secret unchanged.
func GetOIDCClientSecret(cmd *cobra.Command) (string, error) {
	if cmd.Flags().Changed(OIDCClientSecretFlagName) {
		return cmd.Flags().GetString(OIDCClientSecretFlagName)
	}
	return os.Getenv("CY_OIDC_CLIENT_SECRET"), nil
}

// AddOIDCDisplayNameFlag registers --display-name for OIDC integration commands.
func AddOIDCDisplayNameFlag(cmd *cobra.Command) {
	cmd.Flags().String(OIDCDisplayNameFlagName, "", "Display name shown on the OIDC login button")
}

// GetOIDCDisplayName reads the --display-name flag.
func GetOIDCDisplayName(cmd *cobra.Command) (string, error) {
	return cmd.Flags().GetString(OIDCDisplayNameFlagName)
}

// AddOIDCIconFlag registers --icon for OIDC integration commands.
func AddOIDCIconFlag(cmd *cobra.Command) {
	cmd.Flags().String(OIDCIconFlagName, "", "Icon URL for the OIDC login button")
}

// GetOIDCIcon reads the --icon flag.
func GetOIDCIcon(cmd *cobra.Command) (string, error) {
	return cmd.Flags().GetString(OIDCIconFlagName)
}

// AddOIDCDiscoveryURLFlag registers --discovery-url for OIDC integration commands.
func AddOIDCDiscoveryURLFlag(cmd *cobra.Command) {
	cmd.Flags().String(OIDCDiscoveryURLFlagName, "", "OIDC discovery (well-known) URL (overrides issuer-derived URL when set)")
}

// GetOIDCDiscoveryURL reads the --discovery-url flag.
func GetOIDCDiscoveryURL(cmd *cobra.Command) (string, error) {
	return cmd.Flags().GetString(OIDCDiscoveryURLFlagName)
}

// AddOIDCGroupsClaimNameFlag registers --groups-claim-name for OIDC integration commands.
func AddOIDCGroupsClaimNameFlag(cmd *cobra.Command) {
	cmd.Flags().String(OIDCGroupsClaimNameFlagName, "", "JWT claim name that carries the user's group memberships")
}

// GetOIDCGroupsClaimName reads the --groups-claim-name flag.
func GetOIDCGroupsClaimName(cmd *cobra.Command) (string, error) {
	return cmd.Flags().GetString(OIDCGroupsClaimNameFlagName)
}

// AddOIDCSessionTTLSecondsFlag registers --session-ttl-seconds for OIDC integration commands.
func AddOIDCSessionTTLSecondsFlag(cmd *cobra.Command) {
	cmd.Flags().Int64(OIDCSessionTTLSecondsFlagName, 0, "OIDC session lifetime in seconds (0 uses provider default)")
}

// GetOIDCSessionTTLSeconds reads the --session-ttl-seconds flag.
func GetOIDCSessionTTLSeconds(cmd *cobra.Command) (int64, error) {
	return cmd.Flags().GetInt64(OIDCSessionTTLSecondsFlagName)
}

// AddOIDCClientSecretJwtFlag registers --client-secret-jwt for OIDC integration commands.
func AddOIDCClientSecretJwtFlag(cmd *cobra.Command) {
	cmd.Flags().Bool(OIDCClientSecretJwtFlagName, false, "Use client_secret_jwt authentication method")
}

// GetOIDCClientSecretJwt reads the --client-secret-jwt flag.
func GetOIDCClientSecretJwt(cmd *cobra.Command) (bool, error) {
	return cmd.Flags().GetBool(OIDCClientSecretJwtFlagName)
}

// AddOIDCUseCaCertFlag registers --use-ca-cert for OIDC integration commands.
func AddOIDCUseCaCertFlag(cmd *cobra.Command) {
	cmd.Flags().Bool(OIDCUseCaCertFlagName, false, "Validate the OIDC provider's TLS certificate against the configured CA cert")
}

// GetOIDCUseCaCert reads the --use-ca-cert flag.
func GetOIDCUseCaCert(cmd *cobra.Command) (bool, error) {
	return cmd.Flags().GetBool(OIDCUseCaCertFlagName)
}

// AddOIDCCaCertFlag registers --ca-cert for OIDC integration commands.
// The CA cert can also be supplied via the CY_OIDC_CA_CERT environment variable.
func AddOIDCCaCertFlag(cmd *cobra.Command) {
	cmd.Flags().String(OIDCCaCertFlagName, "", "PEM-encoded CA certificate for the OIDC provider (also read from CY_OIDC_CA_CERT env var)")
}

// GetOIDCCaCert returns the CA certificate PEM: flag value if explicitly set,
// otherwise the CY_OIDC_CA_CERT environment variable.
func GetOIDCCaCert(cmd *cobra.Command) (string, error) {
	if cmd.Flags().Changed(OIDCCaCertFlagName) {
		return cmd.Flags().GetString(OIDCCaCertFlagName)
	}
	return os.Getenv("CY_OIDC_CA_CERT"), nil
}

// AddOIDCSkipTLSVerifyFlag registers --skip-tls-verify for OIDC integration commands.
func AddOIDCSkipTLSVerifyFlag(cmd *cobra.Command) {
	cmd.Flags().Bool(OIDCSkipTLSVerifyFlagName, false, "Skip TLS certificate verification for the OIDC provider (insecure)")
}

// GetOIDCSkipTLSVerify reads the --skip-tls-verify flag.
func GetOIDCSkipTLSVerify(cmd *cobra.Command) (bool, error) {
	return cmd.Flags().GetBool(OIDCSkipTLSVerifyFlagName)
}

// AddOIDCAllowInsecureDiscoveryFlag registers --allow-insecure-discovery for OIDC integration commands.
func AddOIDCAllowInsecureDiscoveryFlag(cmd *cobra.Command) {
	cmd.Flags().Bool(OIDCAllowInsecureDiscoveryFlagName, false, "Allow fetching the OIDC discovery document over HTTP (insecure)")
}

// GetOIDCAllowInsecureDiscovery reads the --allow-insecure-discovery flag.
func GetOIDCAllowInsecureDiscovery(cmd *cobra.Command) (bool, error) {
	return cmd.Flags().GetBool(OIDCAllowInsecureDiscoveryFlagName)
}

// AddOIDCEnabledFlag registers --enabled for OIDC integration commands.
func AddOIDCEnabledFlag(cmd *cobra.Command) {
	cmd.Flags().Bool(OIDCEnabledFlagName, false, "Enable the OIDC SSO integration")
}

// GetOIDCEnabled reads the --enabled flag.
func GetOIDCEnabled(cmd *cobra.Command) (bool, error) {
	return cmd.Flags().GetBool(OIDCEnabledFlagName)
}

const OIDCAdoptManualMembersFlagName = "adopt-manual-members"

// AddOIDCAdoptManualMembersFlag registers --adopt-manual-members for OIDC integration commands.
func AddOIDCAdoptManualMembersFlag(cmd *cobra.Command) {
	cmd.Flags().Bool(OIDCAdoptManualMembersFlagName, false, "Adopt manually-invited members on OIDC login (flips their source to 'oidc' so group mapping manages them)")
}

// GetOIDCAdoptManualMembers reads the --adopt-manual-members flag.
func GetOIDCAdoptManualMembers(cmd *cobra.Command) (bool, error) {
	return cmd.Flags().GetBool(OIDCAdoptManualMembersFlagName)
}
