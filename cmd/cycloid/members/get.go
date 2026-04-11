package members

import (
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/client/models"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
)

func NewGetCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "get",
		Args:  cobra.NoArgs,
		Short: "Get the organization member",
		Example: `
	# Get a member within my-org organization
	cy --org my-org members get --id 50
`,
		RunE: getMember,
	}

	cmd.MarkFlagRequired(cyargs.AddMemberIDFlag(cmd))
	cyout.RegisterModel(cmd, models.User{})

	return cmd
}

func getMember(cmd *cobra.Command, args []string) error {
	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}

	id, err := cyargs.GetMemberID(cmd)
	if err != nil {
		return err
	}

	mb, _, err := m.GetMember(org, id)
	return cyout.PrintWithOptions(cmd, mb, err, "unable to get member", memberTableOptions)
}
