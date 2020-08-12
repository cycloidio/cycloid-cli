package externalBackends

import (
	"fmt"

	"github.com/cycloidio/youdeploy-cli/CLIs/v1/cycloid/common"
	"github.com/cycloidio/youdeploy-cli/client/client/organization_external_backends"
	"github.com/spf13/cobra"

	root "github.com/cycloidio/youdeploy-cli/cmd/cycloid"
)

func NewDeleteCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "delete",
		Short: "...",
		Long:  `........ . . .... .. .. ....`,
		RunE:  delete,
	}
	common.RequiredFlag(common.WithFlagID, cmd)
	common.RequiredPersistentFlag(common.WithFlagOrg, cmd)

	return cmd
}

// /organizations/{organization_canonical}/external_backends/{external_backend_id}
// delete: deleteExternalBackend
// delete an External Backend
func delete(cmd *cobra.Command, args []string) error {
	api := root.NewAPI()

	org, err := cmd.Flags().GetString("org")
	if err != nil {
		return err
	}
	id, err := cmd.Flags().GetUint32("id")
	if err != nil {
		return err
	}

	ebP := organization_external_backends.NewDeleteExternalBackendParams()
	ebP.SetOrganizationCanonical(org)
	ebP.SetExternalBackendID(id)

	resp, err := api.OrganizationExternalBackends.DeleteExternalBackend(ebP, root.ClientCredentials())
	if err != nil {
		return err
	}

	fmt.Println(resp)
	fmt.Printf("%+v\n", err)
	return nil
}
