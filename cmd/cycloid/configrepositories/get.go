package configrepositories

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
		Args:              cyargs.RequireArgsOrFlag("canonical"),
		ValidArgsFunction: cyargs.CompleteConfigRepository,
		Short:             "get a config repository",
		Example: `
	# get a config repository by canonical
	cy --org my-org config-repo get my-config-repo

	# get multiple config repositories at once
	cy --org my-org config-repo get repo-one repo-two

	# get using the deprecated --canonical flag
	cy --org my-org config-repo get --canonical my-config-repo -o yaml
`,
		RunE: getConfigRepository,
	}

	// Keep --canonical for backward compatibility
	can := cyargs.AddConfigRepoCanonicalFlag(cmd)
	_ = cmd.Flags().MarkHidden(can)

	return cmd
}

func getConfigRepository(cmd *cobra.Command, args []string) error {
	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}

	// Support --canonical for backward compat; positional args take precedence
	if len(args) == 0 {
		can, err := cyargs.GetConfigRepoCanonical(cmd)
		if err != nil {
			return err
		}
		args = []string{can}
	}

	output, err := cyargs.GetOutput(cmd)
	if err != nil {
		return errors.Wrap(err, "unable to get output flag")
	}

	p, err := factory.GetPrinter(output)
	if err != nil {
		return errors.Wrap(err, "unable to get printer")
	}

	if len(args) == 1 {
		cr, _, err := m.GetConfigRepository(org, args[0])
		if err != nil {
			return printer.SmartPrint(p, nil, err, "unable to get config repository", printer.Options{}, cmd.OutOrStderr())
		}
		return printer.SmartPrint(p, cr, nil, "", printer.Options{}, cmd.OutOrStdout())
	}

	results := make([]interface{}, 0, len(args))
	for _, can := range args {
		cr, _, err := m.GetConfigRepository(org, can)
		if err != nil {
			return printer.SmartPrint(p, nil, err, "unable to get config repository "+can, printer.Options{}, cmd.OutOrStderr())
		}
		results = append(results, cr)
	}
	if output == "table" {
		p, _ = factory.GetPrinter("json")
	}
	return printer.SmartPrint(p, results, nil, "", printer.Options{}, cmd.OutOrStdout())
}
