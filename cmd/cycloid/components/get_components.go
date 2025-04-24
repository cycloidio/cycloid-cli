package components

import (
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/printer/factory"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func NewGetComponentsCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list -p project -e env",
		Short: "List components in a project",
		RunE:  getComponents,
	}
	flagSet := cmd.Flags()
	common.AddProjectFlag(flagSet)
	common.AddEnvFlag(flagSet)
	return cmd
}

func getComponents(cmd *cobra.Command, args []string) error {
	org, err := common.GetOrg(cmd)
	if err != nil {
		return err
	}

	project, err := common.GetProject(cmd)
	if err != nil {
		return err
	}

	env, err := common.GetEnv(cmd)
	if err != nil {
		return err
	}

	output, err := cmd.Flags().GetString("output")
	if err != nil {
		return err
	}

	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	p, err := factory.GetPrinter(output)
	if err != nil {
		return errors.Wrap(err, "unable to get printer")
	}

	components, err := m.GetComponents(org, project, env)
	if err != nil {
		return printer.SmartPrint(p, nil, err, "failed to fetch list of components in  '"+project+"', '"+env+"'", printer.Options{}, cmd.OutOrStderr())
	}

	return printer.SmartPrint(p, components, nil, "", printer.Options{}, cmd.OutOrStderr())
}
