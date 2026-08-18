package configrepositories

import (
	stderrors "errors"
	"fmt"
	"net/http"

	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/apiclient"
	"github.com/cycloidio/cycloid-cli/cmd/common"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
	"github.com/cycloidio/cycloid-cli/printer"
)

func NewCreateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create",
		Args:  cobra.NoArgs,
		Short: "create a config repository",
		Example: `
	# create a config repository and set up as default
	cy --org my-org config-repo create --branch config --cred my-cred --url "git@github.com:my/repo.git" --name default-config --default
`,
		RunE: createConfigRepository,
	}

	cmd.MarkFlagRequired(cyargs.AddRepoCredFlag(cmd))
	cmd.MarkFlagRequired(cyargs.AddNameFlag(cmd))
	cmd.MarkFlagRequired(cyargs.AddRepoBranchFlag(cmd))
	cmd.MarkFlagRequired(cyargs.AddRepoURLFlag(cmd))
	cyargs.AddConfigRepositoryFlag(cmd)
	cyargs.AddRepoDefaultFlag(cmd)
	cmd.Flags().Bool("update", false, "update the config repository if it already exists")

	return cmd
}

func createConfigRepository(cmd *cobra.Command, args []string) error {
	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}

	canonical, err := cyargs.GetConfigRepository(cmd)
	if err != nil {
		return err
	}

	name, err := cyargs.GetName(cmd)
	if err != nil {
		return err
	}

	displayName, repoCanonical, err := apiclient.NameOrCanonical(&name, &canonical)
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

	setDefault, err := cyargs.GetRepoDefault(cmd)
	if err != nil {
		return err
	}

	cred, err := cyargs.GetRepoCred(cmd)
	if err != nil {
		return err
	}

	update, err := cmd.Flags().GetBool("update")
	if err != nil {
		return err
	}

	api := common.NewAPI()
	m := apiclient.NewAPIClient(api)

	_, _, getErr := m.GetConfigRepository(org, repoCanonical)
	exists := getErr == nil
	if getErr != nil {
		var apiErr *apiclient.APIResponseError
		if !stderrors.As(getErr, &apiErr) || apiErr.StatusCode != http.StatusNotFound {
			return cyout.PrintWithOptions(cmd, nil, getErr, "failed to check if config repository exists", printer.Options{})
		}
	}

	if exists && !update {
		return cyout.PrintWithOptions(cmd, nil,
			fmt.Errorf("config repository %q already exists; use --update or `cy config-repo update`", repoCanonical),
			"unable to create config repository", printer.Options{})
	}

	if exists {
		cr, _, err := m.UpdateConfigRepository(org, repoCanonical, cred, displayName, url, branch, setDefault)
		return cyout.PrintWithOptions(cmd, cr, err, "unable to update config repository", printer.Options{})
	}

	cr, _, err := m.CreateConfigRepository(org, displayName, repoCanonical, url, branch, cred, setDefault)
	return cyout.PrintWithOptions(cmd, cr, err, "unable to create config repository", printer.Options{})
}
