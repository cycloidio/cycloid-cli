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

func NewDeleteComponentCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete -p project -e env -c component",
		Short: "Delete a component",
		RunE:  deleteComponent,
	}
	cy_args.AddCyContext(cmd)
	return cmd
}

func deleteComponent(cmd *cobra.Command, args []string) error {
	org, project, env, component, err := cy_args.GetCyContext(cmd)
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

	err = m.DeleteComponent(org, project, env, component)
	return printer.SmartPrint(p, nil, err, "failed to delete component '"+component+"'", printer.Options{}, cmd.OutOrStdout())
}
