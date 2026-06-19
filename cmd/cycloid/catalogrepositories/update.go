package catalogrepositories

import (
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
	"github.com/cycloidio/cycloid-cli/printer"
)

func NewUpdateCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "update",
		Args:  cobra.NoArgs,
		Short: "update a catalog repository",
		Example: `
	# update a catalog repository
	cy --org my-org cr update --branch my-branch --cred my-cred --url "git@github.com:my/repo.git" --name my-catalog-name --canonical my-catalog-repository

	# update and immediately re-index all branches/tags so they are resolvable
	cy --org my-org cr update --branch my-branch --url "git@github.com:my/repo.git" --name my-catalog-name --canonical my-catalog-repository --refresh
`,
		RunE: updateCatalogRepository,
	}

	cmd.MarkFlagRequired(cyargs.AddCatalogRepoCanonicalFlag(cmd))
	cyargs.AddRepoCredFlag(cmd)
	cmd.MarkFlagRequired(cyargs.AddNameFlag(cmd))
	cmd.MarkFlagRequired(cyargs.AddRepoBranchFlag(cmd))
	cmd.MarkFlagRequired(cyargs.AddRepoURLFlag(cmd))
	cmd.Flags().Bool("refresh", true, "trigger a synchronous version re-index after update to make all branches and tags immediately resolvable")

	return cmd
}

func updateCatalogRepository(cmd *cobra.Command, args []string) error {
	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}

	can, err := cyargs.GetCatalogRepoCanonical(cmd)
	if err != nil {
		return err
	}

	name, err := cyargs.GetName(cmd)
	if err != nil {
		return err
	}

	url, err := cyargs.GetRepoURL(cmd)
	if err != nil {
		return err
	}

	branch, err := cyargs.GetRepoBranch(cmd)
	if err != nil {
		return err
	}

	cred, err := cyargs.GetRepoCred(cmd)
	if err != nil {
		return err
	}

	refresh, err := cmd.Flags().GetBool("refresh")
	if err != nil {
		return err
	}

	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	cr, _, err := m.UpdateCatalogRepository(org, can, name, url, branch, cred, nil)
	if err != nil {
		return cyout.PrintWithOptions(cmd, nil, err, "unable to update catalog repository", printer.Options{})
	}

	if refresh {
		if _, _, refreshErr := m.RefreshCatalogRepositoryVersions(org, can); refreshErr != nil {
			// Print the successful update result before returning the refresh error
			_ = cyout.PrintWithOptions(cmd, cr, nil, "", printer.Options{})
			return cyout.PrintWithOptions(cmd, nil, refreshErr, "unable to refresh catalog repository versions", printer.Options{})
		}
	}

	return cyout.PrintWithOptions(cmd, cr, nil, "", printer.Options{})
}
