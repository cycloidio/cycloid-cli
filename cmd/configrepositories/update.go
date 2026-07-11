package configrepositories

import (
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/apiclient"
	"github.com/cycloidio/cycloid-cli/cmd/common"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
	"github.com/cycloidio/cycloid-cli/printer"
)

func NewUpdateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update",
		Args:  cobra.NoArgs,
		Short: "update a config repository",
		Example: `
	# update a config repository
	cy --org my-org config-repo update --branch my-branch --cred my-cred --url "git@github.com:my/repo.git" --name my-config-name --canonical my-config-repo
`,
		RunE: updateConfigRepository,
	}

	cmd.MarkFlagRequired(cyargs.AddConfigRepoCanonicalFlag(cmd))
	cyargs.AddRepoCredFlag(cmd)
	cmd.MarkFlagRequired(cyargs.AddNameFlag(cmd))
	cmd.MarkFlagRequired(cyargs.AddRepoBranchFlag(cmd))
	cmd.MarkFlagRequired(cyargs.AddRepoURLFlag(cmd))
	cyargs.AddRepoDefaultFlag(cmd)

	return cmd
}

func updateConfigRepository(cmd *cobra.Command, args []string) error {
	api := common.NewAPI()
	m := apiclient.NewAPIClient(api)

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

	setDefault, err := cyargs.GetRepoDefault(cmd)
	if err != nil {
		return err
	}

	cred, err := cyargs.GetRepoCred(cmd)
	if err != nil {
		return err
	}

	cr, _, err := m.UpdateConfigRepository(org, can, cred, name, url, branch, setDefault)
	return cyout.PrintWithOptions(cmd, cr, err, "unable to update config repository", printer.Options{})
}
