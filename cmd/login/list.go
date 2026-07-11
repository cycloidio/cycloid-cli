package login

import (
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/config"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
	"github.com/cycloidio/cycloid-cli/printer"
)

// NewListCommand returns the cobra command holding
// the list login subcommand
func NewListCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Args:  cobra.NoArgs,
		Short: "list the current logged organizations",
		Example: `
	# list the organizations where the user is logged in
	cy login list
`,
		RunE: listLogins,
	}
}

func listLogins(cmd *cobra.Command, args []string) error {
	// fetch any existing config
	// we skip the error in case it's the first usage and the config
	// file does not exist
	conf, _ := config.Read()

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
			Name: name,
			// Special formatting to display only the 7 last chars of a token
			Token: o.Token[len(o.Token)-7:],
		})
	}

	return cyout.PrintWithOptions(cmd, orgs, nil, "", printer.Options{})
}
