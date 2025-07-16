package credentials

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"golang.org/x/exp/slices"

	"github.com/cycloidio/cycloid-cli/client/models"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/internal"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/printer/factory"
)

func NewGetCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use: "get [credential]",
		Args: cobra.MatchAll(
			cobra.OnlyValidArgs,
			cobra.RangeArgs(0, 1),
		),
		ValidArgsFunction: cyargs.CompleteCredentialCanonical,
		Short:             "get a credential",
		Example: `# get a credential by its canonical
	cy --org my-org credential get my-cred`,
		RunE:    get,
		PreRunE: internal.CheckAPIAndCLIVersion,
	}

	cyargs.AddCredentialCanonicalFlag(cmd)
	return cmd
}

func get(cmd *cobra.Command, args []string) error {
	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}

	var credential string
	credentialFlag, _ := cyargs.GetCredentialCanonical(cmd)
	credentialPath, _ := cyargs.GetCredentialPath(cmd)
	// Fill credential with precedence Canflag > PathFlag > Args
	if credentialFlag != "" {
		credential = credentialFlag
	} else if credentialPath != "" && credentialFlag == "" {
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

		credential = *credList[index].Canonical
	} else if credentialFlag == "" && credentialPath == "" && len(args) == 1 {
		credential = args[0]
	} else {
		return errors.New("please fill --canonical or --path flags or as argument")
	}

	output, err := cyargs.GetOutput(cmd)
	if err != nil {
		return errors.Wrap(err, "unable to get output flag")
	}

	// fetch the printer from the factory
	p, err := factory.GetPrinter(output)
	if err != nil {
		return errors.Wrap(err, "unable to get printer")
	}

	c, err := m.GetCredential(org, credential)
	if err != nil {
		return printer.SmartPrint(p, nil, err, "unable to get credential from API", printer.Options{}, cmd.OutOrStderr())
	}

	return printer.SmartPrint(p, c, nil, "", printer.Options{}, cmd.OutOrStdout())
}
