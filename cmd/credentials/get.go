package credentials

import (
	"fmt"

	"github.com/spf13/cobra"
	"golang.org/x/exp/slices"

	"github.com/cycloidio/cycloid-cli/cmd/apiclient"
	"github.com/cycloidio/cycloid-cli/cmd/common"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/gen/models"
)

// credentialGetTableOptions excludes the Raw field (sensitive) but shows Keys.
var credentialGetTableOptions = printer.Options{
	Columns:    []string{"Canonical", "Name", "Type", "Path", "Keys"},
	Identifier: "Canonical",
}

func NewGetCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "get [canonical...]",
		Args:              cobra.ArbitraryArgs,
		ValidArgsFunction: cyargs.CompleteCredentialCanonical,
		Short:             "get a credential",
		Example: `
	# get a credential by canonical
	cy --org my-org credential get credential-canonical

	# get multiple credentials
	cy --org my-org credential get cred-a cred-b

	# get a credential by path
	cy --org my-org credential get --path /my/secret/path
`,
		RunE: get,
	}

	cyargs.AddCredentialCanonicalFlag(cmd)
	cyargs.AddCredentialPathFlag(cmd)
	cyout.RegisterModel(cmd, models.Credential{})
	return cmd
}

func get(cmd *cobra.Command, args []string) error {
	api := common.NewAPI()
	m := apiclient.NewMiddleware(api)

	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}

	credentialFlag, _ := cyargs.GetCredentialCanonical(cmd)
	credentialPath, _ := cyargs.GetCredentialPath(cmd)

	if credentialFlag != "" {
		found := false
		for _, a := range args {
			if a == credentialFlag {
				found = true
				break
			}
		}
		if !found {
			args = append(args, credentialFlag)
		}
	}

	// Multi-arg mode
	if len(args) > 1 && credentialPath == "" {
		results := make([]*models.Credential, 0, len(args))
		for _, canonical := range args {
			c, _, err := m.GetCredential(org, canonical)
			if err != nil {
				return cyout.PrintWithOptions(cmd, nil, err, "unable to get credential "+canonical, credentialGetTableOptions)
			}
			results = append(results, c)
		}
		return cyout.PrintWithOptions(cmd, results, nil, "", credentialGetTableOptions)
	}

	// Single credential: path > positional arg
	var credential string
	if credentialPath != "" {
		credList, _, err := m.ListCredentials(org, "")
		if err != nil {
			return fmt.Errorf("failed to fetch cred list to match credential by path %q: %w", credentialPath, err)
		}

		index := slices.IndexFunc(credList, func(c *models.CredentialSimple) bool {
			if c.Path != nil {
				return *c.Path == credentialPath
			}
			return false
		})
		if index == -1 || credList[index].Canonical == nil {
			return fmt.Errorf("credential with path %q not found in org %q", credentialPath, org)
		}

		credential = *credList[index].Canonical
	} else if len(args) == 1 {
		credential = args[0]
	} else {
		return fmt.Errorf("please fill --canonical or --path flags or as argument")
	}

	c, _, err := m.GetCredential(org, credential)
	return cyout.PrintWithOptions(cmd, c, err, "unable to get credential from API", credentialGetTableOptions)
}
