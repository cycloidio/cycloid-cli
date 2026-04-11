package configrepositories

import (
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
	"github.com/cycloidio/cycloid-cli/printer"
)

func NewDeleteCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:               "delete [canonical...]",
		Aliases:           []string{"rm"},
		Args:              cyargs.RequireArgsOrFlag("canonical"),
		ValidArgsFunction: cyargs.CompleteConfigRepository,
		Short:             "delete a config repository",
		Example: `
	# delete a config repository by canonical
	cy --org my-org config-repo delete my-config-repo

	# delete multiple config repositories at once
	cy --org my-org config-repo delete repo-one repo-two

	# delete using the deprecated --canonical flag
	cy --org my-org config-repo delete --canonical my-config-repo
`,
		RunE: deleteConfigRepository,
	}

	// Keep --canonical for backward compatibility
	can := cyargs.AddConfigRepoCanonicalFlag(cmd)
	_ = cmd.Flags().MarkHidden(can)

	return cmd
}

func deleteConfigRepository(cmd *cobra.Command, args []string) error {
	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}

	// Support --canonical for backward compat; positional args take precedence
	if len(args) == 0 {
		can, err := cyargs.GetConfigRepoCanonical(cmd)
		if err != nil {
			return err
		}
		args = []string{can}
	}

	deleted := make([]string, 0, len(args))
	for _, can := range args {
		_, err = m.DeleteConfigRepository(org, can)
		if err != nil {
			return cyout.PrintWithOptions(cmd, nil, err, "unable to delete config repository "+can, printer.Options{})
		}
		deleted = append(deleted, can)
	}
	return cyout.PrintWithOptions(cmd, deleted, nil, "", printer.Options{})
}
