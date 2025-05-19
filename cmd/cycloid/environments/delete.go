package environments

import (
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/internal"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cy_args"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/printer/factory"
)

func NewDeleteCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:     "delete",
		Args:    cobra.NoArgs,
		Aliases: []string{"del", "rm"},
		Short:   "delete a environment",
		Example: `cy --org my-org environment delete --env my-environment`,
		RunE:    deleteEnvironment,
		PreRunE: internal.CheckAPIAndCLIVersion,
	}

	cy_args.AddProjectFlag(cmd)
	cy_args.AddEnvFlag(cmd)
	return cmd
}

func deleteEnvironment(cmd *cobra.Command, args []string) error {
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

	env, err := cy_args.GetEnv(cmd)
	if err != nil {
		return err
	}

	output, err := cy_args.GetOutput(cmd)
	if err != nil {
		return err
	}

	// fetch the printer from the factory
	p, err := factory.GetPrinter(output)
	if err != nil {
		return err
	}

	err = m.DeleteEnv(org, project, env)
	return printer.SmartPrint(p, nil, err, "unable to delete environment", printer.Options{}, cmd.OutOrStdout())
}
