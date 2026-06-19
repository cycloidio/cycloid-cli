package integration

import (
	"errors"
	"net/http"

	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
)

// NewSetCommand returns the `integration set` command.
func NewSetCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set",
		Args:  cobra.NoArgs,
		Short: "Create or update the org's OIDC SSO integration",
		Long: `Create or update the AuthenticationOIDC SSO integration for the organization.

Only the flags you pass are changed; unset flags keep their current value
(read-merge-write). The backend always requires 'type' and 'enabled' — they
are always included in the update payload.

Secrets (--client-secret, --ca-cert) are write-only: the GET never returns
them. Omitting them preserves the stored value; passing an empty string is
equivalent to omitting them (the backend ignores blank secrets).

The client secret can also be supplied via the CY_OIDC_CLIENT_SECRET
environment variable; the CA certificate via CY_OIDC_CA_CERT.`,
		Example: `
  # Recommended: pass the secret via env var to keep it out of shell history
  export CY_OIDC_CLIENT_SECRET="$(cat ./oidc-secret)"
  cy --org my-org oidc integration set \
    --enabled \
    --issuer https://idp.example.com \
    --client-id my-client-id

  # Flag form (secret is visible in shell history / process list — avoid on shared hosts)
  cy --org my-org oidc integration set --client-secret "$MY_SECRET"

  # Update only the groups claim name (all other settings unchanged)
  cy --org my-org oidc integration set --groups-claim-name groups

  # Use the alias 'config' instead of 'integration'
  cy --org my-org oidc config set --enabled=false
`,
		RunE: setIntegration,
	}

	cyargs.AddOIDCEnabledFlag(cmd)
	cyargs.AddOIDCIssuerFlag(cmd)
	cyargs.AddOIDCClientIDFlag(cmd)
	cyargs.AddOIDCClientSecretFlag(cmd)
	cyargs.AddOIDCDisplayNameFlag(cmd)
	cyargs.AddOIDCIconFlag(cmd)
	cyargs.AddOIDCDiscoveryURLFlag(cmd)
	cyargs.AddOIDCGroupsClaimNameFlag(cmd)
	cyargs.AddOIDCSessionTTLSecondsFlag(cmd)
	cyargs.AddOIDCClientSecretJwtFlag(cmd)
	cyargs.AddOIDCUseCaCertFlag(cmd)
	cyargs.AddOIDCCaCertFlag(cmd)
	cyargs.AddOIDCSkipTLSVerifyFlag(cmd)
	cyargs.AddOIDCAllowInsecureDiscoveryFlag(cmd)
	cyargs.AddOIDCAdoptManualMembersFlag(cmd)

	cyout.RegisterModel(cmd, middleware.OIDCIntegration{})
	return cmd
}

func setIntegration(cmd *cobra.Command, args []string) error {
	// Hard Rule 4: ALL flags parsed before NewAPI / NewMiddleware.
	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}

	enabled, err := cyargs.GetOIDCEnabled(cmd)
	if err != nil {
		return err
	}
	issuer, err := cyargs.GetOIDCIssuer(cmd)
	if err != nil {
		return err
	}
	clientID, err := cyargs.GetOIDCClientID(cmd)
	if err != nil {
		return err
	}
	clientSecret, err := cyargs.GetOIDCClientSecret(cmd)
	if err != nil {
		return err
	}
	displayName, err := cyargs.GetOIDCDisplayName(cmd)
	if err != nil {
		return err
	}
	icon, err := cyargs.GetOIDCIcon(cmd)
	if err != nil {
		return err
	}
	discoveryURL, err := cyargs.GetOIDCDiscoveryURL(cmd)
	if err != nil {
		return err
	}
	groupsClaimName, err := cyargs.GetOIDCGroupsClaimName(cmd)
	if err != nil {
		return err
	}
	sessionTTL, err := cyargs.GetOIDCSessionTTLSeconds(cmd)
	if err != nil {
		return err
	}
	clientSecretJwt, err := cyargs.GetOIDCClientSecretJwt(cmd)
	if err != nil {
		return err
	}
	useCaCert, err := cyargs.GetOIDCUseCaCert(cmd)
	if err != nil {
		return err
	}
	caCert, err := cyargs.GetOIDCCaCert(cmd)
	if err != nil {
		return err
	}
	skipTLSVerify, err := cyargs.GetOIDCSkipTLSVerify(cmd)
	if err != nil {
		return err
	}
	allowInsecureDiscovery, err := cyargs.GetOIDCAllowInsecureDiscovery(cmd)
	if err != nil {
		return err
	}
	adoptManualMembers, err := cyargs.GetOIDCAdoptManualMembers(cmd)
	if err != nil {
		return err
	}

	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	// Read-merge-write: fetch current config so we can preserve 'enabled' when
	// the user does not explicitly pass --enabled. A 404 means the integration
	// does not exist yet — treat as a fresh create (current == nil).
	var current *middleware.OIDCIntegration
	current, _, err = m.GetOIDCIntegration(org)
	if err != nil {
		var apiErr *middleware.APIResponseError
		if !errors.As(err, &apiErr) || apiErr.StatusCode != http.StatusNotFound {
			return cyout.Print(cmd, nil, err, "unable to read current OIDC integration")
		}
		// 404 → create path; current stays nil
	}

	// Always required by the backend.
	config := map[string]interface{}{
		"type": "AuthenticationOIDC",
	}

	// enabled: use the flag value if explicitly changed; otherwise carry forward
	// the current value; default to true for a fresh create.
	if cmd.Flags().Changed(cyargs.OIDCEnabledFlagName) {
		config["enabled"] = enabled
	} else if current != nil {
		config["enabled"] = current.Enabled
	} else {
		config["enabled"] = true
	}

	// For each optional field: only include in the payload when the user
	// explicitly passed the flag, so the backend merge keeps untouched values.
	if cmd.Flags().Changed(cyargs.OIDCIssuerFlagName) {
		config["oidc_issuer"] = issuer
	}
	if cmd.Flags().Changed(cyargs.OIDCClientIDFlagName) {
		config["oidc_client_id"] = clientID
	}
	if cmd.Flags().Changed(cyargs.OIDCDisplayNameFlagName) {
		config["oidc_display_name"] = displayName
	}
	if cmd.Flags().Changed(cyargs.OIDCIconFlagName) {
		config["oidc_icon"] = icon
	}
	if cmd.Flags().Changed(cyargs.OIDCDiscoveryURLFlagName) {
		config["oidc_discovery_url"] = discoveryURL
	}
	if cmd.Flags().Changed(cyargs.OIDCGroupsClaimNameFlagName) {
		config["oidc_groups_claim_name"] = groupsClaimName
	}
	if cmd.Flags().Changed(cyargs.OIDCSessionTTLSecondsFlagName) {
		config["oidc_session_ttl_seconds"] = sessionTTL
	}
	if cmd.Flags().Changed(cyargs.OIDCClientSecretJwtFlagName) {
		config["oidc_client_secret_jwt"] = clientSecretJwt
	}
	if cmd.Flags().Changed(cyargs.OIDCUseCaCertFlagName) {
		config["oidc_use_ca_cert"] = useCaCert
	}
	if cmd.Flags().Changed(cyargs.OIDCSkipTLSVerifyFlagName) {
		config["oidc_skip_tls_verify"] = skipTLSVerify
	}
	if cmd.Flags().Changed(cyargs.OIDCAllowInsecureDiscoveryFlagName) {
		config["oidc_allow_insecure_discovery"] = allowInsecureDiscovery
	}
	if cmd.Flags().Changed(cyargs.OIDCAdoptManualMembersFlagName) {
		config["oidc_adopt_manual_members"] = adoptManualMembers
	}

	// Secrets: send only when a non-empty value was supplied (via --client-secret
	// flag or CY_OIDC_CLIENT_SECRET env var). The backend preserves the stored
	// secret when the key is absent or the value is empty. Note: if the env var
	// is set in the shell, the secret will be sent on every invocation — this is
	// intentional (the user opted in by exporting the var).
	if clientSecret != "" {
		config["oidc_client_secret"] = clientSecret
	}
	if caCert != "" {
		config["oidc_ca_cert"] = caCert
	}

	result, _, err := m.UpdateOIDCIntegration(org, config)
	return cyout.PrintWithOptions(cmd, result, err, "unable to update OIDC integration", integrationTableOptions)
}
