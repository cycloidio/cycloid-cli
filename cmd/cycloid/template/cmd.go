package template

import (
	"github.com/spf13/cobra"
)

// NewCommands builds the `cy template` command group: local templating and
// interpolation tooling. Today it ships the offline `render` verb; the
// backend-backed `context` verb (pull real context from an existing component)
// lands in a follow-up once the templating endpoint exists.
func NewCommands() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "template",
		Aliases: []string{"tpl", "tmpl"},
		Short:   "Test Cycloid templating and interpolation locally",
	}

	cmd.AddCommand(
		NewRenderCommand(),
	)

	return cmd
}
