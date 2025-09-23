package cycloid

import (
	"os"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/printer/factory"
)

func NewStatusCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "status",
		Short: "Get the status of the Cycloid services",
		Example: `
	# show the status of Cycloid services
	cy status
`,
		RunE: func(cmd *cobra.Command, args []string) error {
			output, err := cmd.Flags().GetString("output")
			if err != nil {
				return errors.Wrap(err, "unable to get output flag")
			}

			return status(output)
		},
	}

}

func status(output string) error {
	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	services, err := m.GetStatus()
	if err != nil {
		return errors.Wrap(err, "unable to get status")
	}

	// for a better table display, we process the result
	// in order to add some colors following the service
	// state
	if output == "table" {
		var (
			// green
			success = "\033[0;32mSuccess\033[0m"
			// orange
			unknown = "\033[0;33mUnknown\033[0m"
			// red
			err = "\033[0;31mError\033[0m"
		)
		for _, service := range services.Checks {
			switch *service.Status {
			case "Success":
				service.Status = &success
			case "Unknown":
				service.Status = &unknown
			case "Error":
				service.Status = &err
			}
		}
	}

	// fetch the printer from the factory
	p, err := factory.GetPrinter(output)
	if err != nil {
		return errors.Wrap(err, "unable to get printer")
	}

	// print the result on the standard output
	if err := p.Print(services.Checks, printer.Options{}, os.Stdout); err != nil {
		return errors.Wrap(err, "unable to print result")
	}
	return nil
}
