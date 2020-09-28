package creds

import (
	"fmt"

	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/common"
	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/middleware"
	"github.com/spf13/cobra"
)

func NewDeleteCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "delete",
		Short: "delete a credential",
		Example: `
	# delete a credential with ID 123
	cy --org my-org creds delete --id 123
`,
		RunE: del,
	}

	common.RequiredFlag(common.WithFlagID, cmd)
	return cmd
}

func del(cmd *cobra.Command, args []string) error {
	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	org, err := cmd.Flags().GetString("org")
	if err != nil {
		return err
	}
	id, err := cmd.Flags().GetUint32("id")
	if err != nil {
		return err
	}

	err = m.DeleteCredential(org, id)

	fmt.Printf("%+v\n", err)
	return nil
}
