package bootstrapfirstorg

import (
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/internal/cyargs"
)

func NewCommands() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "bootstrap-first-org",
		Short: "Bootstrap the first organization on a fresh Cycloid install",
		Long: `Bootstrap the very first organization on a fresh Cycloid install.

Chains:
  1. POST /user            — signup (ignores 409 if user already exists)
  2. POST /user/login      — login with email + password
  3. POST /organizations   — create the org (ignores 409 if it already exists)
  4. GET  /user/refresh_token?organization_canonical=<org>
                            — refresh token to get org scope
  5. POST /organizations/<org>/licence — activate the licence

Optionally, when --api-key-canonical is set:
  6. Creates an admin api-key under that canonical
  7. Stores it in a custom credential under the same canonical

Intended for first-install bootstrap or fully scripted environment recreation.
All steps tolerate "already exists" responses, so re-running is safe.

This command is BETA — output shape and flag names may change.`,
		Args: cobra.NoArgs,
		RunE: bootstrap,
	}

	cyargs.AddBootstrapUserFlags(cmd)
	cyargs.AddLicenceFlag(cmd)
	cyargs.AddLicenceFileFlag(cmd)
	cyargs.AddBootstrapAPIKeyCanonicalFlag(cmd)
	cmd.MarkFlagsMutuallyExclusive("licence", "licence-file")

	return cmd
}
