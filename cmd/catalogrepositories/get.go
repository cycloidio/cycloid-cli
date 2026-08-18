package catalogrepositories

import (
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/apiclient"
	"github.com/cycloidio/cycloid-cli/cmd/common"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
	"github.com/cycloidio/cycloid-cli/gen/models"
)

func NewGetCommand() *cobra.Command {
	cmd := &cobra.Command{
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
	cyout.RegisterModel(cmd, models.ServiceCatalogSource{})

	return cmd
}

func getCatalogRepository(cmd *cobra.Command, args []string) error {
	api := common.NewAPI()
	m := apiclient.NewAPIClient(api)

	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}

	if can, err := cyargs.GetCatalogRepoCanonical(cmd); err != nil {
		return err
	} else if can != "" {
		found := false
		for _, a := range args {
			if a == can {
				found = true
				break
			}
		}
		if !found {
			args = append(args, can)
		}
	}

	if len(args) == 1 {
		cr, _, err := m.GetCatalogRepository(org, args[0])
		return cyout.PrintWithOptions(cmd, cr, err, "unable to get catalog repository", catalogSourceTableOptions)
	}

	results := make([]*models.ServiceCatalogSource, 0, len(args))
	for _, can := range args {
		cr, _, err := m.GetCatalogRepository(org, can)
		if err != nil {
			return cyout.PrintWithOptions(cmd, nil, err, "unable to get catalog repository "+can, catalogSourceTableOptions)
		}
		results = append(results, cr)
	}
	return cyout.PrintWithOptions(cmd, results, nil, "", catalogSourceTableOptions)
}
