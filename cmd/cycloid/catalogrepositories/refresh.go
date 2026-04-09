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

func NewRefreshCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:               "refresh [canonical]",
		Args:              cobra.MaximumNArgs(1),
		ValidArgsFunction: cyargs.CompleteCatalogRepository,
		Short:             "refresh a catalog repository",
		Long:              "refresh action can be used if the .cycloid.yml definition has been updated",
		Example: `
	# refresh a catalog repository by canonical
	cy --org my-org catalog-repo refresh my-catalog-repository

	# refresh using the deprecated --canonical flag
	cy --org my-org catalog-repo refresh --canonical my-catalog-repository
`,
		RunE: refreshCatalogRepository,
	}

	// Keep --canonical for backward compatibility
	can := cyargs.AddCatalogRepoCanonicalFlag(cmd)
	_ = cmd.Flags().MarkHidden(can)

	return cmd
}

func refreshCatalogRepository(cmd *cobra.Command, args []string) error {
	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}

	var can string
	if len(args) == 1 {
		can = args[0]
	} else {
		can, err = cyargs.GetCatalogRepoCanonical(cmd)
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

	cr, _, err := m.RefreshCatalogRepository(org, can)
	if err != nil {
		return printer.SmartPrint(p, nil, err, "unable to refresh catalog repository", printer.Options{}, cmd.OutOrStderr())
	}
	return printer.SmartPrint(p, cr, nil, "", printer.Options{}, cmd.OutOrStdout())
}
