package cloudaccounts

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/client/models"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
	"github.com/cycloidio/cycloid-cli/internal/ptr"
	"github.com/cycloidio/cycloid-cli/internal/utils"
)

func NewUpdateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update",
		Short: "Update a cloud account",
		RunE:  updateCloudAccount,
		Args:  cobra.NoArgs,
	}

	cmd.MarkFlagRequired(cyargs.AddCloudAccountFlag(cmd))
	cyargs.AddNameFlag(cmd)
	cyargs.AddExistingCredentialFlag(cmd)
	cyargs.AddNewCredentialTypeFlag(cmd)
	cyargs.AddDescriptionFlag(cmd)
	cyargs.AddOwnerFlag(cmd)
	addCredentialFlags(cmd)
	return cmd
}

func updateCloudAccount(cmd *cobra.Command, args []string) error {
	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}

	canonical, err := cyargs.GetCloudAccount(cmd)
	if err != nil {
		return err
	}

	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	current, _, err := m.GetCloudAccount(org, canonical)
	if err != nil {
		return cyout.PrintWithOptions(cmd, nil, err, "cloud account not found", cloudAccountTableOptions)
	}

	// Flag readers below ignore their err return on purpose: cobra guarantees
	// these flags are registered as the right type, so GetString/GetBool can
	// only fail on a programmer mistake (wrong type lookup), not user input.
	// Treating them as best-effort lookups keeps the PATCH-style merge readable.
	name, _ := cyargs.GetName(cmd)
	description, _ := cyargs.GetDescription(cmd)
	owner, _ := cyargs.GetOwner(cmd)
	existingCredential, _ := cyargs.GetExistingCredential(cmd)
	newCredentialType, _ := cyargs.GetNewCredentialType(cmd)
	if existingCredential != "" && newCredentialType != "" {
		return cyout.PrintWithOptions(cmd, nil, fmt.Errorf("cannot use --credential and --new-credential at the same time"), "", cloudAccountTableOptions)
	}

	body := &models.UpdateCloudAccount{
		Name: ptr.Ptr(utils.CoalesceNonZero(name, ptrValue(current.Name))),
	}
	if cyargs.IsSet(cmd, "description") {
		body.Description = description
	} else {
		body.Description = current.Description
	}
	if cyargs.IsSet(cmd, "owner") {
		body.Owner = owner
	} else if current.Owner != nil && current.Owner.Username != nil {
		body.Owner = *current.Owner.Username
	}

	switch {
	case newCredentialType != "":
		credName := utils.CoalesceNonZero(name, ptrValue(current.Name)) + " access"
		credential, err := createInlineCredential(cmd, m, org, newCredentialType, credName)
		if err != nil {
			return cyout.PrintWithOptions(cmd, nil, err, "failed to create inline credential", cloudAccountTableOptions)
		}
		body.CredentialCanonical = credential.Canonical
	case cyargs.IsSet(cmd, "credential"):
		body.CredentialCanonical = ptr.Ptr(existingCredential)
	case current.Credential != nil && current.Credential.Canonical != nil:
		body.CredentialCanonical = current.Credential.Canonical
	}

	result, _, err := m.UpdateCloudAccount(org, canonical, body)
	return cyout.PrintWithOptions(cmd, result, err, "failed to update cloud account", cloudAccountTableOptions)
}

func ptrValue(value *string) string {
	if value == nil {
		return ""
	}
	return *value
}
