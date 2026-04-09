package projects

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
		Use:               "get [canonical...]",
		Args:              cobra.MaximumNArgs(1),
		ValidArgsFunction: cyargs.CompleteProject,
		Short:             "get a project",
		Example: `
	# get a project by canonical
	cy --org my-org project get my-project

	# get a project using the --project flag or CY_PROJECT env var
	cy --org my-org project get --project my-project -o yaml
`,
		RunE: get,
	}

	cyargs.AddProjectFlag(cmd)
	return cmd
}

func get(cmd *cobra.Command, args []string) error {
	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}

	var project string
	if len(args) == 1 {
		project = args[0]
	} else {
		project, err = cyargs.GetProject(cmd)
		if err != nil {
			return err
		}
	}

	output, err := cyargs.GetOutput(cmd)
	if err != nil {
		return errors.Wrap(err, "unable to get output flag")
	}

	p, err := factory.GetPrinter(output)
	if err != nil {
		return errors.Wrap(err, "unable to get printer")
	}

	proj, _, err := m.GetProject(org, project)
	if err != nil {
		return printer.SmartPrint(p, nil, err, "unable to get project", printer.Options{}, cmd.OutOrStderr())
	}
	return printer.SmartPrint(p, proj, nil, "", printer.Options{}, cmd.OutOrStdout())
}
