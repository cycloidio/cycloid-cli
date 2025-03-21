package projects

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/internal"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/printer/factory"
)

func NewDeleteEnvCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "delete-env",
		Short: "delete an environment within a project",
		Example: `
	# delete env 'my-env' in 'my-project'
	cy --org my-org project delete-env --project my-project --env my-env
`,
		RunE:    deleteEnv,
		PreRunE: internal.CheckAPIAndCLIVersion,
	}
	common.RequiredPersistentFlag(common.WithFlagEnv, cmd)
	common.RequiredPersistentFlag(common.WithFlagProject, cmd)

	return cmd
}

func deleteEnv(cmd *cobra.Command, args []string) error {
	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	org, err := common.GetOrg(cmd)
	if err != nil {
		return err
	}
	project, err := cmd.Flags().GetString("project")
	if err != nil {
		return err
	}
	env, err := cmd.Flags().GetString("env")
	if err != nil {
		return err
	}
	output, err := cmd.Flags().GetString("output")
	if err != nil {
		return errors.Wrap(err, "unable to get output flag")
	}

	// fetch the printer from the factory
	p, err := factory.GetPrinter(output)
	if err != nil {
		return errors.Wrap(err, "unable to get printer")
	}

	err = m.DeleteEnv(org, project, env)
	return printer.SmartPrint(p, nil, err, "unable to delete environment", printer.Options{}, cmd.OutOrStdout())
}
