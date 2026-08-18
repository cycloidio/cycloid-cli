package catalogrepositories

import (
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/apiclient"
	"github.com/cycloidio/cycloid-cli/cmd/common"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
	"github.com/cycloidio/cycloid-cli/printer"
)

func NewRefreshCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "refresh [canonical...]",
		Args:              cyargs.RequireArgsOrFlag("canonical"),
		ValidArgsFunction: cyargs.CompleteCatalogRepository,
		Short:             "refresh one or more catalog repositories",
		Long:              "refresh action can be used if the .cycloid.yml definition has been updated",
		Example: `
	# refresh a catalog repository by canonical
	cy --org my-org catalog-repo refresh my-catalog-repository

	# refresh multiple catalog repositories at once
	cy --org my-org catalog-repo refresh repo-one repo-two

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
	m := apiclient.NewAPIClient(api)

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

	if len(args) == 1 {
		cr, _, err := m.RefreshCatalogRepository(org, args[0])
		return cyout.PrintWithOptions(cmd, cr, err, "unable to refresh catalog repository", printer.Options{})
	}

	results := make([]interface{}, 0, len(args))
	for _, can := range args {
		cr, _, err := m.RefreshCatalogRepository(org, can)
		if err != nil {
			return cyout.PrintWithOptions(cmd, nil, err, "unable to refresh catalog repository "+can, printer.Options{})
		}
		results = append(results, cr)
	}
	return cyout.PrintWithOptions(cmd, results, nil, "", printer.Options{})
}
