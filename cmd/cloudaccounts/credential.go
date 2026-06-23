package cloudaccounts

import (
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/apiclient"
	"github.com/cycloidio/cycloid-cli/cmd/credentials"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/gen/models"
	"github.com/cycloidio/cycloid-cli/utils/ptr"
)

// addCredentialFlags wires every flag a credential type might need on a
// cloud-account command. It delegates to the shared credentials package so
// the supported flag set stays in sync with `cy credential` subcommands.
func addCredentialFlags(cmd *cobra.Command) {
	credentials.AddAllRawFlags(cmd)
}

// createInlineCredential creates a standalone Credential for an existing
// cloud account (used by `cy cloud-account update --new-credential …`).
func createInlineCredential(cmd *cobra.Command, m apiclient.Middleware, org, credType, name string) (*models.Credential, error) {
	raw, err := credentials.BuildCredentialRaw(cmd, credType)
	if err != nil {
		return nil, err
	}

	// path/canonical/description are user-friendly hints, missing values are
	// fine: GetCredentialPath/Canonical only error on type mismatch (cobra
	// guarantees the flags we registered are strings), so the discarded err
	// is unreachable and we keep the local code path readable.
	description, _ := cyargs.GetDescription(cmd)
	path, _ := cyargs.GetCredentialPath(cmd)
	canonical, _ := cyargs.GetCredentialCanonical(cmd)

	cred, _, err := m.CreateCredential(org, name, credType, raw, path, canonical, description)
	if err != nil {
		return nil, err
	}

	return cred, nil
}

// buildAccessCredential builds an inline credential payload to send alongside
// a cloud-account create call (atomic POST /cloud_accounts/with_credentials).
func buildAccessCredential(cmd *cobra.Command, credType, name, canonical string) (*models.NewCloudAccountCredential, error) {
	raw, err := credentials.BuildCredentialRaw(cmd, credType)
	if err != nil {
		return nil, err
	}

	description, err := cyargs.GetDescription(cmd)
	if err != nil {
		return nil, err
	}

	path, _ := cyargs.GetCredentialPath(cmd)

	credential := &models.NewCloudAccountCredential{
		Name:        ptr.Ptr(name),
		Type:        ptr.Ptr(credType),
		Raw:         raw,
		Description: description,
	}
	if canonical != "" {
		credential.Canonical = canonical
	}
	if path != "" {
		credential.Path = path
	}

	return credential, nil
}

func cloudAccountName(cmd *cobra.Command) (string, error) {
	name, err := cyargs.GetName(cmd)
	if err != nil {
		return "", err
	}

	if name == "" {
		canonical, err := cyargs.GetCloudAccount(cmd)
		if err != nil {
			return "", err
		}
		name = canonical
	}

	return name, nil
}
