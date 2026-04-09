package members

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/printer/factory"
)

func NewDeleteCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:     "delete",
		Aliases: []string{"rm"},
		Args:    cobra.NoArgs,
		Short:   "Remove a user from the organization",
		Example: `
	# Remove a member from my-org organization
	cy --org my-org members delete --id 50
`,
		RunE: deleteMember,
	}

	cmd.MarkFlagRequired(cyargs.AddMemberIDFlag(cmd))

	return cmd
}

func deleteMember(cmd *cobra.Command, args []string) error {
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

	if output == "table" {
		output = "json"
	}
	p, err := factory.GetPrinter(output)
	if err != nil {
		return errors.Wrap(err, "unable to get printer")
	}

	_, err = m.DeleteMember(org, id)
	if err != nil {
		return printer.SmartPrint(p, nil, err, "unable to remove member", printer.Options{}, cmd.OutOrStderr())
	}
	return printer.SmartPrint(p, []string{fmt.Sprintf("%d", id)}, nil, "", printer.Options{}, cmd.OutOrStdout())
}
