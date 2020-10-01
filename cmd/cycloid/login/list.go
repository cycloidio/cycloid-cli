package login

import (
	"os"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/config"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/printer/factory"
)

var (
	list = &cobra.Command{
		Use:   "list",
		Short: "list the current logged organizations",
		Example: `
	# list the organizations where the user is logged in
	cy login list
`,
		RunE: func(cmd *cobra.Command, args []string) error {
			output, err := cmd.Flags().GetString("output")
			if err != nil {
				return errors.Wrap(err, "unable to get output flag")
			}
			// fetch any existing config
			// we skip the error in case it's the first usage and the config
			// file does not exist
			conf, _ := config.ReadConfig()

			// we need to peform this hack because the printer is waiting for
			// a struct or a slice of structs. Not a map, since the header of the table
			// is the name of the field, we need to pass to the printer a slice of anonymous
			// struct
			var orgs []*struct {
				Name  string
				Token string
			}
			for name, o := range conf.Organizations {
				orgs = append(orgs, &struct {
					Name  string
					Token string
				}{
					Name:  name,
					Token: o.Token,
				})
			}

			// fetch the printer from the factory
			p, err := factory.GetPrinter(output)
			if err != nil {
				return errors.Wrap(err, "unable to get printer")
			}

			// print the result on the standard output
			if err := p.Print(orgs, printer.Options{}, os.Stdout); err != nil {
				return errors.Wrap(err, "unable to print result")
			}
			return nil
		},
	}
)
