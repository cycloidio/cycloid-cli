package cmd

import (
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/apiclient"
	"github.com/cycloidio/cycloid-cli/cmd/common"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
	"github.com/cycloidio/cycloid-cli/printer"
)

func NewStatusCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "status",
		Short: "Get the status of the Cycloid services",
		Example: `
	# show the status of Cycloid services
	cy status
`,
		RunE: getStatus,
	}
}

func getStatus(cmd *cobra.Command, args []string) error {
	api := common.NewAPI()
	m := apiclient.NewMiddleware(api)

	services, _, err := m.GetStatus()
	if err != nil {
		return cyout.PrintWithOptions(cmd, nil, err, "unable to get status", printer.Options{})
	}

	output, _ := cyargs.GetOutput(cmd)

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
			errColor = "\033[0;31mError\033[0m"
		)
		for _, service := range services.Checks {
			switch *service.Status {
			case "Success":
				service.Status = &success
			case "Unknown":
				service.Status = &unknown
			case "Error":
				service.Status = &errColor
			}
		}
	}

	return cyout.PrintWithOptions(cmd, services.Checks, nil, "", printer.Options{})
}
