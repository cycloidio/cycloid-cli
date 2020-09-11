package members

import (
	"github.com/spf13/cobra"

	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/common"
	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/middleware"
)

func NewDeleteCommand() *cobra.Command {
	var (
		example = `
	# Remove a member from my-org organization
	cy --org my-org members delete --name my-username
	`
		short = "Remove a user from the organization"
		long  = short
	)

	var cmd = &cobra.Command{
		Use:     "delete",
		Example: example,
		Short:   short,
		Long:    long,
		RunE:    deleteMember,
	}

	common.RequiredFlag(WithFlagName, cmd)

	return cmd
}

func deleteMember(cmd *cobra.Command, args []string) error {
	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	org, err := cmd.Flags().GetString("org")
	if err != nil {
		return err
	}

	name, err := cmd.Flags().GetString("name")
	if err != nil {
		return err
	}

	err = m.DeleteMember(org, name)

	return err
}
