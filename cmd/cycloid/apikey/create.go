package apikey

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/printer/factory"
)

// NewCreateCommand returns the cobra command holding
// the create API key subcommand
func NewCreateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create",
		Short: "create an API key",
		Example: `
	# create an API key in the org my-org
	cy api-key create --role-id 2 --name "CI API key" --description "Cycloid API key to be used in a CI context"
`,
		RunE: func(cmd *cobra.Command, args []string) error {
			output, err := cmd.Flags().GetString("output")
			if err != nil {
				return fmt.Errorf("unable to get output flag: %w", err)
			}
			name, err := cmd.Flags().GetString("name")
			if err != nil {
				return fmt.Errorf("unable to get name flag: %w", err)
			}
			description, err := cmd.Flags().GetString("description")
			if err != nil {
				return fmt.Errorf("unable to get description flag: %w", err)
			}
			roleID, err := cmd.Flags().GetUint32("role-id")
			if err != nil {
				return fmt.Errorf("unable to get role ID flag: %w", err)
			}
			canonical, err := cmd.Flags().GetString("canonical")
			if err != nil {
				return fmt.Errorf("unable to get canonical flag: %w", err)
			}
			org, err := cmd.Flags().GetString("org")
			if err != nil {
				return fmt.Errorf("unable to get org flag: %w", err)
			}
			return create(org, name, canonical, description, output, roleID)
		},
	}

	WithFlagName(cmd)
	WithFlagDescription(cmd)
	WithFlagRoleID(cmd)
	WithFlagCanonical(cmd)

	return cmd
}

// create will send the POST request to the API in order to
// create an API token which will be displayed on the screen
func create(org, name, canonical, description, output string, roleID uint32) error {
	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	if len(canonical) == 0 {
		canonical = common.GenerateCanonical(name)
	}

	key, err := m.CreateAPIKey(org, name, canonical, description, roleID)
	if err != nil {
		return fmt.Errorf("unable to create API key: %w", err)
	}

	// fetch the printer from the factory
	p, err := factory.GetPrinter(output)
	if err != nil {
		return fmt.Errorf("unable to get printer: %w", err)
	}

	// print the result on the standard output
	if err := p.Print(key, printer.Options{}, os.Stdout); err != nil {
		return fmt.Errorf("unable to print result: %w", err)
	}
	return nil
}
