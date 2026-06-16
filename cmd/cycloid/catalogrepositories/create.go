package catalogrepositories

import (
	stderrors "errors"
	"fmt"
	"net/http"

	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
	"github.com/cycloidio/cycloid-cli/printer"
)

func NewCreateCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "create",
		Args:  cobra.NoArgs,
		Short: "create a catalog repository",
		Example: `
	# create a catalog repository using credential canonical 123, branch 'stacks' and git URL
	cy --org my-org catalog-repo create --branch stacks --cred my-cred --url "git@github.com:my/repo.git" --name my-catalog-name

	# create a catalog repository using public git repository
	cy --org my-org catalog-repo create --branch stacks --url "https://github.com:my/repo.git" --name my-catalog-name

	# create and immediately re-index all branches/tags so they are resolvable (fixes branch-stack presence race)
	cy --org my-org catalog-repo create --branch stacks --url "git@github.com:my/repo.git" --name my-catalog-name --refresh
`,
		RunE: createCatalogRepository,
	}

	cyargs.AddRepoCredFlag(cmd)
	cmd.MarkFlagRequired(cyargs.AddNameFlag(cmd))
	cmd.MarkFlagRequired(cyargs.AddRepoBranchFlag(cmd))
	cmd.MarkFlagRequired(cyargs.AddRepoURLFlag(cmd))
	cyargs.AddCatalogRepoCanonicalFlag(cmd)

	cmd.Flags().String("visibility", "", "set the stacks base visibility in the catalog. accepted values are 'local', 'shared' or 'hidden' (default: local)")
	cmd.Flags().String("team", "", "set the team canonical to be set as maintener of the stacks")
	cmd.Flags().Bool("update", false, "update the catalog repository if it already exists")
	cmd.Flags().Bool("refresh", false, "trigger a synchronous version re-index after create (or update) to make all branches and tags immediately resolvable")

	return cmd
}

// /organizations/{organization_canonical}/service_catalog_sources
// post: createServiceCatalogSource
// Creates a Service catalog source
func createCatalogRepository(cmd *cobra.Command, args []string) error {
	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}

	name, err := cyargs.GetName(cmd)
	if err != nil {
		return err
	}

	canonical, err := cyargs.GetCatalogRepoCanonical(cmd)
	if err != nil {
		return err
	}

	displayName, repoCanonical, err := middleware.NameOrCanonical(&name, &canonical)
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

	visibility, err := cmd.Flags().GetString("visibility")
	if err != nil {
		return err
	}

	teamCanonical, err := cmd.Flags().GetString("team")
	if err != nil {
		return err
	}

	update, err := cmd.Flags().GetBool("update")
	if err != nil {
		return err
	}

	refresh, err := cmd.Flags().GetBool("refresh")
	if err != nil {
		return err
	}

	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	_, _, getErr := m.GetCatalogRepository(org, repoCanonical)
	exists := getErr == nil
	if getErr != nil {
		var apiErr *middleware.APIResponseError
		if !stderrors.As(getErr, &apiErr) || apiErr.StatusCode != http.StatusNotFound {
			return cyout.PrintWithOptions(cmd, nil, getErr, "failed to check if catalog repository exists", printer.Options{})
		}
	}

	if exists && !update {
		return cyout.PrintWithOptions(cmd, nil,
			fmt.Errorf("catalog repository %q already exists; use --update or `cy catalog-repo update`", repoCanonical),
			"unable to create catalog repository", printer.Options{})
	}

	var cr interface{}
	if exists {
		cr, _, err = m.UpdateCatalogRepository(org, repoCanonical, displayName, url, branch, cred, nil)
	} else {
		cr, _, err = m.CreateCatalogRepository(org, displayName, url, branch, cred, visibility, teamCanonical)
	}
	if err != nil {
		return cyout.PrintWithOptions(cmd, nil, err, "unable to create catalog repository", printer.Options{})
	}

	if refresh {
		if _, _, refreshErr := m.RefreshCatalogRepositoryVersions(org, repoCanonical); refreshErr != nil {
			_ = cyout.PrintWithOptions(cmd, cr, nil, "", printer.Options{})
			return cyout.PrintWithOptions(cmd, nil, refreshErr, "unable to refresh catalog repository versions", printer.Options{})
		}
	}

	return cyout.PrintWithOptions(cmd, cr, nil, "", printer.Options{})
}
