package members

import (
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/apiclient"
	"github.com/cycloidio/cycloid-cli/cmd/common"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
	"github.com/cycloidio/cycloid-cli/printer"
)

func NewUpdateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update",
		Args:  cobra.NoArgs,
		Short: "Update a member",
		Example: `
	# Update a member within my-org organization
	cy --org my-org members update --id 50 --role my-role
`,
		RunE: updateMember,
	}

	cmd.MarkFlagRequired(cyargs.AddMemberIDFlag(cmd))
	cmd.MarkFlagRequired(cyargs.AddMemberRoleFlag(cmd))

	return cmd
}

func updateMember(cmd *cobra.Command, args []string) error {
	api := common.NewAPI()
	m := apiclient.NewMiddleware(api)

	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}

	id, err := cyargs.GetMemberID(cmd)
	if err != nil {
		return err
	}

	role, err := cyargs.GetMemberRole(cmd)
	if err != nil {
		return err
	}

	mb, _, err := m.UpdateMember(org, id, role)
	return cyout.PrintWithOptions(cmd, mb, err, "unable to update member", printer.Options{})
}
