package members

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/printer/factory"
)

func NewUpdateCommand() *cobra.Command {
	var cmd = &cobra.Command{
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
	m := middleware.NewMiddleware(api)

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

	output, err := cyargs.GetOutput(cmd)
	if err != nil {
		return errors.Wrap(err, "unable to get output flag")
	}

	p, err := factory.GetPrinter(output)
	if err != nil {
		return errors.Wrap(err, "unable to get printer")
	}

	mb, _, err := m.UpdateMember(org, id, role)
	if err != nil {
		return printer.SmartPrint(p, nil, err, "unable to update member", printer.Options{}, cmd.OutOrStderr())
	}
	return printer.SmartPrint(p, mb, nil, "", printer.Options{}, cmd.OutOrStdout())
}
