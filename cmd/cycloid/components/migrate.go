package components

import (
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cy_args"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/printer/factory"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func NewMigrateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "migrate --project old-project --env old-env --component component --new-project new-project --new-env new-env --new-canonical new-comp-canonical",
		Short: "move a component from one project / env to another.",
		RunE:  migrate,
	}

	cy_args.AddCyContext(cmd)
	cmd.Flags().String("new-project", "", "specify a new project, if unset, will use the previous one")
	cmd.Flags().String("new-env", "", "specify a new env, if unset, will use the previous one")
	cmd.Flags().String("new-canonical", "", "specify a new component canonical, if unset, will use the previous one")
	cmd.Flags().String("new-name", "", "specify a new component name, if unset, will use the new canonical")
	return cmd
}

func migrate(cmd *cobra.Command, args []string) error {
	org, project, env, component, err := cy_args.GetCyContext(cmd)
	if err != nil {
		return err
	}

	newProject, err := cmd.Flags().GetString("new-project")
	if err != nil {
		return err
	}
	if newProject == "" {
		newProject = project
	}

	newEnv, err := cmd.Flags().GetString("new-env")
	if err != nil {
		return err
	}
	if newEnv == "" {
		newEnv = env
	}

	newComponent, err := cmd.Flags().GetString("new-canonical")
	if err != nil {
		return err
	}
	if newComponent == "" {
		newComponent = component
	}

	newComponentName, err := cmd.Flags().GetString("new-name")
	if err != nil {
		return err
	}
	if newComponentName == "" {
		newComponentName = newComponent
	}

	output, err := cmd.Flags().GetString("output")
	if err != nil {
		return err
	}

	p, err := factory.GetPrinter(output)
	if err != nil {
		return errors.Wrap(err, "unable to get printer")
	}

	api := common.NewAPI()
	m := middleware.NewMiddleware(api)
	compResponse, err := m.MigrateComponent(org, project, env, component, newProject, newEnv, newComponent, newComponentName)
	if err != nil {
		return printer.SmartPrint(p, nil, err, "failed to migrate component", printer.Options{}, cmd.OutOrStderr())
	}

	return printer.SmartPrint(p, compResponse, nil, "", printer.Options{}, cmd.OutOrStdout())
}
