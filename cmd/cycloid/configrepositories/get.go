package configrepositories

import (
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/client/models"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
)

func NewGetCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:               "get [canonical...]",
		Args:              cyargs.RequireArgsOrFlag("canonical"),
		ValidArgsFunction: cyargs.CompleteConfigRepository,
		Short:             "get a config repository",
		Example: `
	# get a config repository by canonical
	cy --org my-org config-repo get my-config-repo

	# get multiple config repositories at once
	cy --org my-org config-repo get repo-one repo-two

	# get using the deprecated --canonical flag
	cy --org my-org config-repo get --canonical my-config-repo -o yaml
`,
		RunE: getConfigRepository,
	}

	// Keep --canonical for backward compatibility
	can := cyargs.AddConfigRepoCanonicalFlag(cmd)
	_ = cmd.Flags().MarkHidden(can)
	cyout.RegisterModel(cmd, models.ConfigRepository{})

	return cmd
}

func getConfigRepository(cmd *cobra.Command, args []string) error {
	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}

	if can, err := cyargs.GetConfigRepoCanonical(cmd); err != nil {
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
		cr, _, err := m.GetConfigRepository(org, args[0])
		return cyout.PrintWithOptions(cmd, cr, err, "unable to get config repository", configRepoTableOptions)
	}

	results := make([]*models.ConfigRepository, 0, len(args))
	for _, can := range args {
		cr, _, err := m.GetConfigRepository(org, can)
		if err != nil {
			return cyout.PrintWithOptions(cmd, nil, err, "unable to get config repository "+can, configRepoTableOptions)
		}
		results = append(results, cr)
	}
	return cyout.PrintWithOptions(cmd, results, nil, "", configRepoTableOptions)
}
