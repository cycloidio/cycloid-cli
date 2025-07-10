package credentials

import (
	"fmt"

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
		Use:   "delete",
		Args:  cobra.NoArgs,
		Short: "delete a credential",
		Example: `
	# delete a credential with canonical my-cred
	cy --org my-org credential delete --canonical my-cred
`,
		RunE: del,
	}

	cyargs.AddCredentialCanonicalFlag(cmd)
	cyargs.AddCredentialPathFlag(cmd)
	return cmd
}

func del(cmd *cobra.Command, args []string) error {
	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}

	credential, err := cmd.Flags().GetString("canonical")
	if err != nil {
		return err
	}

	path, err := cyargs.GetCredentialPath(cmd)
	if err != nil {
		return err
	}

	if path == "" && credential == "" {
		return fmt.Errorf("please fill --canonical or --path argument.")
	}

	output, err := cmd.Flags().GetString("output")
	if err != nil {
		return errors.Wrap(err, "unable to get output flag")
	}

	// fetch the printer from the factory
	p, err := factory.GetPrinter(output)
	if err != nil {
		return errors.Wrap(err, "unable to get printer")
	}

	err = m.DeleteCredential(org, credential)
	if err
	// TODO: support getting cred by path
	return printer.SmartPrint(p, nil, err, "unable to delete credential", printer.Options{}, cmd.OutOrStdout())
}
