package catalogrepositories

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/printer/factory"
)

func NewDeleteCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:               "delete [canonical...]",
		Aliases:           []string{"rm"},
		Args:              cyargs.RequireArgsOrFlag("canonical"),
		ValidArgsFunction: cyargs.CompleteCatalogRepository,
		Short:             "delete a catalog repository",
		Example: `
	# delete a catalog repository by canonical
	cy --org my-org catalog-repo delete my-catalog-repository

	# delete multiple catalog repositories at once
	cy --org my-org catalog-repo delete repo-one repo-two

	# delete using the deprecated --canonical flag
	cy --org my-org catalog-repo delete --canonical my-catalog-repository
`,
		RunE: deleteCatalogRepository,
	}

	// Keep --canonical for backward compatibility
	can := cyargs.AddCatalogRepoCanonicalFlag(cmd)
	_ = cmd.Flags().MarkHidden(can)

	return cmd
}

func deleteCatalogRepository(cmd *cobra.Command, args []string) error {
	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}

	// Support --canonical for backward compat; positional args take precedence
	if len(args) == 0 {
		can, err := cyargs.GetCatalogRepoCanonical(cmd)
		if err != nil {
			return err
		}
		args = []string{can}
	}

	output, err := cyargs.GetOutput(cmd)
	if err != nil {
		return errors.Wrap(err, "unable to get output flag")
	}

	p, err := factory.GetPrinter(output)
	if err != nil {
		return errors.Wrap(err, "unable to get printer")
	}

	deleted := make([]string, 0, len(args))
	for _, canonical := range args {
		_, err = m.DeleteCatalogRepository(org, canonical)
		if err != nil {
			return printer.SmartPrint(p, nil, err, "unable to delete catalog repository "+canonical, printer.Options{}, cmd.OutOrStderr())
		}
		deleted = append(deleted, canonical)
	}
	return printer.SmartPrint(p, deleted, nil, "", printer.Options{}, cmd.OutOrStdout())
}
