package organizations

import (
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/apiclient"
	"github.com/cycloidio/cycloid-cli/cmd/common"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/gen/models"
)

var orgTableOptions = printer.Options{
	Columns:    []string{"Canonical", "Name"},
	Identifier: "Canonical",
}

func NewGetCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get [canonical...]",
		Args:  cobra.ArbitraryArgs,
		Short: "get an organization",
		Example: `
	# get the organization set via --org flag
	cy organization get --org my-org -o yaml

	# get multiple organizations by canonical
	cy organization get my-org-a my-org-b
`,
		RunE: get,
	}

	cyout.RegisterModel(cmd, models.Organization{})
	return cmd
}

func get(cmd *cobra.Command, args []string) error {
	api := common.NewAPI()
	m := apiclient.NewMiddleware(api)

	// Multi-arg
	if len(args) > 1 {
		results := make([]*models.Organization, 0, len(args))
		for _, canonical := range args {
			o, _, err := m.GetOrganization(canonical)
			if err != nil {
				return cyout.PrintWithOptions(cmd, nil, err, "unable to get organization "+canonical, orgTableOptions)
			}
			results = append(results, o)
		}
		return cyout.PrintWithOptions(cmd, results, nil, "", orgTableOptions)
	}

	var org string
	if len(args) == 1 {
		org = args[0]
	} else {
		var err error
		org, err = cyargs.GetOrg(cmd)
		if err != nil {
			return err
		}
	}

	o, _, err := m.GetOrganization(org)
	return cyout.PrintWithOptions(cmd, o, err, "unable to get organization", orgTableOptions)
}
