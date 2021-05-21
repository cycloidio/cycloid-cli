package kpis

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/internal"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/printer/factory"
)

func NewCreateCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "create",
		Short: "create a kpi",
		Example: `
	# create a build_history kpi
	cy --org my-org kpi create \
		--name "Build history" \
		--type build_history \
		--widget history \
		--project foo \
		--env environment-name \
		--job hello
`,
		RunE:    create,
		PreRunE: internal.CheckAPIAndCLIVersion,
	}
	common.RequiredFlag(WithFlagName, cmd)
	common.RequiredFlag(WithFlagType, cmd)
	common.RequiredFlag(WithFlagWidget, cmd)

	common.WithFlagProject(cmd)
	common.WithFlagEnv(cmd)
	WithFlagJob(cmd)
	WithFlagConfig(cmd)

	return cmd
}

func create(cmd *cobra.Command, args []string) error {
	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	org, err := cmd.Flags().GetString("org")
	if err != nil {
		return err
	}
	name, err := cmd.Flags().GetString("name")
	if err != nil {
		return err
	}
	env, err := cmd.Flags().GetString("env")
	if err != nil {
		return err
	}
	project, err := cmd.Flags().GetString("project")
	if err != nil {
		return err
	}
	job, err := cmd.Flags().GetString("job")
	if err != nil {
		return err
	}
	widget, err := cmd.Flags().GetString("widget")
	if err != nil {
		return err
	}
	kpiType, err := cmd.Flags().GetString("type")
	if err != nil {
		return err
	}
	config, err := cmd.Flags().GetString("config")
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

	kpi, err := m.CreateKpi(name, kpiType, widget, org, project, job, env, config)
	return printer.SmartPrint(p, kpi, err, "Unable to create the KPI", printer.Options{}, cmd.OutOrStdout())
}
