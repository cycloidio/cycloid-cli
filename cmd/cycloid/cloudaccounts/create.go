package cloudaccounts

import (
	stderrors "errors"
	"fmt"
	"net/http"

	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/client/models"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
	"github.com/cycloidio/cycloid-cli/internal/ptr"
	"github.com/cycloidio/cycloid-cli/printer"
)

func NewCreateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "create",
		Short:   "Create a cloud account",
		Example: `cy --org my-org cloud-account create --cloud-account my-aws --name "My AWS" --cloud-provider aws --credential my-aws-cred`,
		RunE:    create,
		Args:    cobra.NoArgs,
	}

	cmd.MarkFlagsOneRequired(
		cyargs.AddNameFlag(cmd),
		cyargs.AddCloudAccountFlag(cmd),
	)
	cmd.MarkFlagRequired(cyargs.AddCloudProviderFlag(cmd))
	cyargs.AddExistingCredentialFlag(cmd)
	cyargs.AddNewCredentialTypeFlag(cmd)
	cyargs.AddDescriptionFlag(cmd)
	cyargs.AddOwnerFlag(cmd)
	addCredentialFlags(cmd)
	cyargs.AddUpdateFlag(cmd, "update the cloud account if it already exists")
	return cmd
}

func create(cmd *cobra.Command, args []string) error {
	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}

	name, err := cloudAccountName(cmd)
	if err != nil {
		return err
	}

	canonical, err := cyargs.GetCloudAccount(cmd)
	if err != nil {
		return err
	}
	name, canonical, err = middleware.NameOrCanonical(&name, &canonical)
	if err != nil {
		return err
	}

	cloudProviderPtr, err := cyargs.GetCloudProvider(cmd)
	if err != nil {
		return err
	}
	if cloudProviderPtr == nil || *cloudProviderPtr == "" {
		return fmt.Errorf("cloud-provider is required")
	}
	cloudProvider := *cloudProviderPtr

	existingCredential, err := cyargs.GetExistingCredential(cmd)
	if err != nil {
		return err
	}
	newCredentialType, err := cyargs.GetNewCredentialType(cmd)
	if err != nil {
		return err
	}
	if existingCredential != "" && newCredentialType != "" {
		return fmt.Errorf("cannot use --credential and --new-credential at the same time")
	}
	if existingCredential == "" && newCredentialType == "" {
		return fmt.Errorf("either --credential or --new-credential must be set")
	}

	description, err := cyargs.GetDescription(cmd)
	if err != nil {
		return err
	}
	owner, err := cyargs.GetOwner(cmd)
	if err != nil {
		return err
	}

	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	if cyargs.GetUpdate(cmd) {
		_, _, getErr := m.GetCloudAccount(org, canonical)
		if getErr == nil {
			return updateCloudAccount(cmd, args)
		}
		var apiErr *middleware.APIResponseError
		if !stderrors.As(getErr, &apiErr) || apiErr.StatusCode != http.StatusNotFound {
			return cyout.PrintWithOptions(cmd, nil, getErr, "failed to check existing cloud account", printer.Options{})
		}
	}

	var result *models.CloudAccount
	if newCredentialType != "" {
		accessCredential, buildErr := buildAccessCredential(cmd, newCredentialType, name+" access", "")
		if buildErr != nil {
			return buildErr
		}
		body := &models.NewCloudAccountWithCredentials{
			Canonical:        canonical,
			Name:             ptr.Ptr(name),
			CloudProvider:    ptr.Ptr(cloudProvider),
			Description:      description,
			Owner:            owner,
			AccessCredential: accessCredential,
		}
		result, _, err = m.CreateCloudAccountWithCredentials(org, body)
	} else {
		body := &models.NewCloudAccount{
			Canonical:           canonical,
			Name:                ptr.Ptr(name),
			CloudProvider:       ptr.Ptr(cloudProvider),
			CredentialCanonical: ptr.Ptr(existingCredential),
			Description:         description,
			Owner:               owner,
		}
		result, _, err = m.CreateCloudAccount(org, body)
	}

	return cyout.PrintWithOptions(cmd, result, err, "failed to create cloud account", cloudAccountTableOptions)
}
