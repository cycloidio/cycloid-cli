package organizations

import (
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/apiclient"
	"github.com/cycloidio/cycloid-cli/cmd/common"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
)

func NewDeleteCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete [canonical...]",
		Args:  cobra.ArbitraryArgs,
		Short: "delete an organization (require root API_KEY)",
		Example: `
	# delete the organization set via --org flag
	# The API_KEY must be obtained from the root organization
	cy organization delete --org my-org

	# delete multiple organizations by canonical
	cy organization delete my-org-a my-org-b
`,
		RunE: del,
	}

	return cmd
}

func del(cmd *cobra.Command, args []string) error {
	api := common.NewAPI()
	m := apiclient.NewAPIClient(api)

	var orgs []string
	if len(args) > 0 {
		orgs = args
	} else {
		org, err := cyargs.GetOrg(cmd)
		if err != nil {
			return err
		}
		orgs = []string{org}
	}

	for _, org := range orgs {
		_, err := m.DeleteOrganization(org)
		if err != nil {
			return cyout.Print(cmd, nil, err, "unable to delete organization "+org)
		}
	}
	return nil
}
