package apikey

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
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
	cy api-key create --role my-role --name "CI API key" --description "Cycloid API key to be used in a CI context"
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
			role, err := cmd.Flags().GetString("role")
			if err != nil {
				return fmt.Errorf("unable to get role flag: %w", err)
			}
			org, err := cmd.Flags().GetString("org")
			if err != nil {
				return fmt.Errorf("unable to get org flag: %w", err)
			}
			return create(org, name, description, output, role)
		},
	}

	WithFlagName(cmd)
	WithFlagDescription(cmd)
	WithFlagRole(cmd)

	return cmd
}

// create will send the POST request to the API in order to
// create an API token which will be displayed on the screen
func create(org, name, description, output, role string) error {
	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	// fetch the printer from the factory
	p, err := factory.GetPrinter(output)
	if err != nil {
		return errors.Wrap(err, "unable to get printer")
	}
	key, err := m.CreateAPIKey(org, name, description, role)
	return printer.SmartPrint(p, key, err, "unable to create API key", printer.Options{}, os.Stdout)
}
