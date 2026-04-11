package stacks

import (
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
	"github.com/cycloidio/cycloid-cli/printer"
)

func NewGetCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:     "get",
		Args:    cobra.NoArgs,
		Short:   "get information on a stack",
		Example: `cy --org my-org stacks get --ref my:stack-ref`,
		RunE:    get,
	}

	cmd.MarkFlagRequired(cyargs.AddStackRefFlag(cmd))
	return cmd
}

func get(cmd *cobra.Command, args []string) error {
	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}

	ref, err := cyargs.GetStackRef(cmd)
	if err != nil {
		return err
	}

	// Initialize middleware after all arguments are collected
	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	s, _, err := m.GetStack(org, ref)
	return cyout.PrintWithOptions(cmd, s, err, "failed to get stack from API", printer.Options{})
}
