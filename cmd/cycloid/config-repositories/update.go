package configRepositories

import (
	"fmt"
	"time"

	root "github.com/cycloidio/youdeploy-cli/cmd/cycloid"
	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/common"
	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/middleware"
	"github.com/spf13/cobra"
)

func NewUpdateCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "update",
		Short: "...",
		Long:  `........ . . .... .. .. ....`,
		RunE:  updateConfigRepository,
	}

	common.RequiredFlag(common.WithFlagID, cmd)
	common.RequiredFlag(common.WithFlagCred, cmd)
	common.RequiredFlag(WithFlagName, cmd)
	common.RequiredFlag(WithFlagBranch, cmd)
	common.RequiredFlag(WithFlagURL, cmd)
	WithFlagDefault(cmd)

	//TODO : dont Required flags and if not set, use value from the getConfigRepository

	return cmd
}

func updateConfigRepository(cmd *cobra.Command, args []string) error {
	api := root.NewAPI()
	m := middleware.NewMiddleware(api)

	org, err := cmd.Flags().GetString("org")
	if err != nil {
		return err
	}

	id, err := cmd.Flags().GetUint32("id")
	if err != nil {
		return err
	}

	name, err := cmd.Flags().GetString("name")
	if err != nil {
		return err
	}

	url, err := cmd.Flags().GetString("url")
	if err != nil {
		return err
	}

	branch, err := cmd.Flags().GetString("branch")
	if err != nil {
		return err
	}

	setDefault, err := cmd.Flags().GetBool("default")
	if err != nil {
		return err
	}

	cred, err := cmd.Flags().GetUint32("cred")
	if err != nil {
		return err
	}

	cr, err := m.UpdateConfigRepository(org, id, name, url, branch, setDefault, cred)
	if err != nil {
		return err
	}

	fmt.Printf("id: %d    name: %s    url: %s    branch: %s    default: %t    credential_id: %d\n", *cr.ID, *cr.Name, *cr.URL, cr.Branch, *cr.Default, cr.CredentialID)
	fmt.Printf("created_at: %v    updated_at: %v\n", time.Unix(*cr.CreatedAt, 0), time.Unix(*cr.UpdatedAt, 0))

	fmt.Println(cr)
	fmt.Printf("%+v\n", err)

	return nil
}
