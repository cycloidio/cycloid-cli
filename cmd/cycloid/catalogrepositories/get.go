package catalogrepositories

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
		ValidArgsFunction: cyargs.CompleteCatalogRepository,
		Short:             "get a catalog repository",
		Example: `
	# get a catalog repository by canonical
	cy --org my-org catalog-repo get my-catalog-repository

	# get a catalog repository using the deprecated --canonical flag
	cy --org my-org catalog-repo get --canonical my-catalog-repository

	# get multiple catalog repositories at once
	cy --org my-org catalog-repo get repo-one repo-two
`,
		RunE: getCatalogRepository,
	}

	// Keep --canonical for backward compatibility
	can := cyargs.AddCatalogRepoCanonicalFlag(cmd)
	_ = cmd.Flags().MarkHidden(can)

	return cmd
}

func getCatalogRepository(cmd *cobra.Command, args []string) error {
	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}

	// Support --canonical for backward compat; positional args take precedence
	if len(args) == 0 {
		can, err := cyargs.GetCatalogRepoCanonical(cmd)
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
		cr, _, err := m.GetCatalogRepository(org, args[0])
		if err != nil {
			return printer.SmartPrint(p, nil, err, "unable to get catalog repository", printer.Options{}, cmd.OutOrStderr())
		}
		return printer.SmartPrint(p, cr, nil, "", printer.Options{}, cmd.OutOrStdout())
	}

	results := make([]interface{}, 0, len(args))
	for _, can := range args {
		cr, _, err := m.GetCatalogRepository(org, can)
		if err != nil {
			return printer.SmartPrint(p, nil, err, "unable to get catalog repository "+can, printer.Options{}, cmd.OutOrStderr())
		}
		results = append(results, cr)
	}
	return printer.SmartPrint(p, results, nil, "", printer.Options{}, cmd.OutOrStdout())
}
