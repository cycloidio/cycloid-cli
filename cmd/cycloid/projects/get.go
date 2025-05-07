package projects

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/internal"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cy_args"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/printer/factory"
)

func NewGetCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "get",
		Short: "get a project",
		Example: `
	# get a project in YAML format
	cy --org my-org project get --project my-project -o yaml
`,
		RunE:    get,
		PreRunE: internal.CheckAPIAndCLIVersion,
	}

	common.RequiredPersistentFlag(common.WithFlagProject, cmd)

	return cmd
}

func get(cmd *cobra.Command, args []string) error {
	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	org, err := cy_args.GetOrg(cmd)
	if err != nil {
		return err
	}
	project, err := cy_args.GetProject(cmd)
	if err != nil {
		return err
	}
	output, err := cy_args.GetOutput(cmd)
	if err != nil {
		return errors.Wrap(err, "unable to get output flag")
	}

	// fetch the printer from the factory
	p, err := factory.GetPrinter(output)
	if err != nil {
		return errors.Wrap(err, "unable to get printer")
	}

	proj, err := m.GetProject(org, project)
	return printer.SmartPrint(p, proj, err, "unable to get project", printer.Options{}, cmd.OutOrStdout())
}
