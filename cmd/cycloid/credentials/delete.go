package credentials

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"golang.org/x/exp/slices"

	"github.com/cycloidio/cycloid-cli/client/models"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/printer/factory"
)

func NewDeleteCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:               "delete [canonicals...]",
		Args:              cobra.OnlyValidArgs,
		ValidArgsFunction: cyargs.CompleteCredentialCanonical,
		Short:             "delete a credential",
		Example: `# Delete 3 credentials with canonical: cred1, cred2, cred3
cy --org my-org credential delete cred1 cred2 --canonical cred3`,
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

	credentialFlag, _ := cyargs.GetCredentialCanonical(cmd)
	credentialPath, _ := cyargs.GetCredentialPath(cmd)
	if credentialPath == "" && credentialFlag == "" && len(args) == 0 {
		return errors.New("please fill --canonical or --path flags or pass canonicals as arguments")
	}

	if credentialPath != "" && credentialFlag == "" {
		credList, err := m.ListCredentials(org, "")
		if err != nil {
			return fmt.Errorf("failed to fetch cred list to match credential by path '%s': %w", credentialPath, err)
		}

		index := slices.IndexFunc(credList, func(c *models.CredentialSimple) bool {
			if c.Path != nil {
				return *c.Path == credentialPath
			} else {
				return false
			}
		})
		if index == -1 || credList[index].Canonical == nil {
			return fmt.Errorf("credential with path '%s' not found in org '%s'", credentialPath, org)
		}

		credentialFlag = *credList[index].Canonical
	}

	credList := append(args, credentialFlag)

	output, err := cyargs.GetOutput(cmd)
	if err != nil {
		return errors.Wrap(err, "unable to get output flag")
	}

	// fetch the printer from the factory
	p, err := factory.GetPrinter(output)
	if err != nil {
		return errors.Wrap(err, "unable to get printer")
	}

	for _, credential := range credList {
		if credential == "" {
			continue
		}

		err := m.DeleteCredential(org, credential)
		if err != nil {
			return printer.SmartPrint(p, nil, err, fmt.Sprintf("unable to delete credential '%s'", credential), printer.Options{}, cmd.OutOrStderr())
		}

		fmt.Fprintf(cmd.OutOrStderr(), "successfully deleted credential '%s'\n", credential)
	}

	return printer.SmartPrint(p, nil, nil, "", printer.Options{}, cmd.OutOrStdout())
}
