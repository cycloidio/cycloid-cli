package catalogrepositories

import (
	stderrors "errors"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/printer/factory"
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
`,
		RunE: createCatalogRepository,
	}

	common.WithFlagCred(cmd)
	common.RequiredFlag(WithFlagName, cmd)
	common.RequiredFlag(WithFlagBranch, cmd)
	common.RequiredFlag(WithFlagURL, cmd)
	cyargs.AddCatalogRepoCanonicalFlag(cmd)

	cmd.Flags().String("visibility", "", "set the stacks base visibility in the catalog. accepted values are 'local', 'shared' or 'hidden' (default: local)")
	cmd.Flags().String("team", "", "set the team canonical to be set as maintener of the stacks")
	cmd.Flags().Bool("update", false, "update the catalog repository if it already exists")

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

	name, err := cmd.Flags().GetString("name")
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

	url, err := cmd.Flags().GetString("url")
	if err != nil {
		return err
	}

	branch, err := cmd.Flags().GetString("branch")
	if err != nil {
		return err
	}

	cred, err := cmd.Flags().GetString("cred")
	if err != nil {
		return err
	}

	visibility, err := cmd.Flags().GetString("visibility")
	if err != nil {
		return errors.Wrap(err, "unable to get visibility flag")
	}

	teamCanonical, err := cmd.Flags().GetString("team")
	if err != nil {
		return errors.Wrap(err, "unable to get team flag")
	}

	update, err := cmd.Flags().GetBool("update")
	if err != nil {
		return err
	}

	output, err := cyargs.GetOutput(cmd)
	if err != nil {
		return errors.Wrap(err, "unable to get output flag")
	}

	p, err := factory.GetPrinter(output)
	if err != nil {
		return errors.Wrap(err, "unable to get printer")
	}

	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	_, _, getErr := m.GetCatalogRepository(org, repoCanonical)
	exists := getErr == nil
	if getErr != nil {
		var apiErr *middleware.APIResponseError
		if !stderrors.As(getErr, &apiErr) || apiErr.StatusCode != http.StatusNotFound {
			return printer.SmartPrint(p, nil, getErr, "failed to check if catalog repository exists", printer.Options{}, cmd.OutOrStderr())
		}
	}

	if exists && !update {
		return printer.SmartPrint(p, nil,
			fmt.Errorf("catalog repository %q already exists; use --update or `cy catalog-repo update`", repoCanonical),
			"unable to create catalog repository", printer.Options{}, cmd.OutOrStderr())
	}

	if exists {
		cr, _, err := m.UpdateCatalogRepository(org, repoCanonical, displayName, url, branch, cred, nil)
		if err != nil {
			return printer.SmartPrint(p, nil, err, "unable to update catalog repository", printer.Options{}, cmd.OutOrStderr())
		}
		return printer.SmartPrint(p, cr, nil, "", printer.Options{}, cmd.OutOrStdout())
	}

	cr, _, err := m.CreateCatalogRepository(org, displayName, url, branch, cred, visibility, teamCanonical)
	if err != nil {
		return printer.SmartPrint(p, nil, err, "unable to create catalog repository", printer.Options{}, cmd.OutOrStderr())
	}
	return printer.SmartPrint(p, cr, nil, "", printer.Options{}, cmd.OutOrStdout())
}
