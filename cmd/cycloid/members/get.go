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

	output, err := cyargs.GetOutput(cmd)
	if err != nil {
		return errors.Wrap(err, "unable to get output flag")
	}

	p, err := factory.GetPrinter(output)
	if err != nil {
		return errors.Wrap(err, "unable to get printer")
	}

	mb, _, err := m.GetMember(org, id)
	if err != nil {
		return printer.SmartPrint(p, nil, err, "unable to get member", printer.Options{}, cmd.OutOrStderr())
	}
	return printer.SmartPrint(p, mb, nil, "", printer.Options{}, cmd.OutOrStdout())
}
