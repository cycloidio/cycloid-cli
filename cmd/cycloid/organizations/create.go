package organizations

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/client/models"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/printer/factory"
)

// This command have been Hidden because it is not compatible with API key login.
// Advanced user still can use it passing a user token in CY_API_KEY env var during a login.
func NewCreateCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:    "create",
		Args:   cobra.NoArgs,
		Short:  "create an organization",
		Hidden: true,
		Example: `# create an organization foo
cy organization create --name foo

# create a child organization
cy organization create --name bar --child-of foo
`,
		RunE: create,
	}

	cmd.MarkFlagRequired(cyargs.AddOrgNameFlag(cmd))
	cyargs.AddOrgChildOfFlag(cmd)
	return cmd
}

func create(cmd *cobra.Command, args []string) error {
	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	name, err := cyargs.GetOrgName(cmd)
	if err != nil {
		return err
	}
	org := common.GenerateCanonical(name)

	parentOrg, err := cyargs.GetOrgChildOf(cmd)
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

	var outOrg *models.Organization
	if parentOrg != "" {
		outOrg, err = m.CreateOrganizationChild(parentOrg, org, &name)
	} else {
		outOrg, err = m.CreateOrganization(name)
	}
	if err != nil {
		return printer.SmartPrint(p, nil, err, "failed to create org named"+name, printer.Options{}, cmd.OutOrStderr())
	}

	return printer.SmartPrint(p, outOrg, nil, "", printer.Options{}, cmd.OutOrStdout())
}
