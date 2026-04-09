package configrepositories

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

	_, _, getErr := m.GetConfigRepository(org, repoCanonical)
	exists := getErr == nil
	if getErr != nil {
		var apiErr *middleware.APIResponseError
		if !stderrors.As(getErr, &apiErr) || apiErr.StatusCode != http.StatusNotFound {
			return printer.SmartPrint(p, nil, getErr, "failed to check if config repository exists", printer.Options{}, cmd.OutOrStderr())
		}
	}

	if exists && !update {
		return printer.SmartPrint(p, nil,
			fmt.Errorf("config repository %q already exists; use --update or `cy config-repo update`", repoCanonical),
			"unable to create config repository", printer.Options{}, cmd.OutOrStderr())
	}

	if exists {
		cr, _, err := m.UpdateConfigRepository(org, repoCanonical, cred, displayName, url, branch, setDefault)
		if err != nil {
			return printer.SmartPrint(p, nil, err, "unable to update config repository", printer.Options{}, cmd.OutOrStderr())
		}
		return printer.SmartPrint(p, cr, nil, "", printer.Options{}, cmd.OutOrStdout())
	}

	cr, _, err := m.CreateConfigRepository(org, displayName, repoCanonical, url, branch, cred, setDefault)
	if err != nil {
		return printer.SmartPrint(p, nil, err, "unable to create config repository", printer.Options{}, cmd.OutOrStderr())
	}
	return printer.SmartPrint(p, cr, nil, "", printer.Options{}, cmd.OutOrStdout())
}
