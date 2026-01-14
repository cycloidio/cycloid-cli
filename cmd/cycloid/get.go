package cycloid

import (
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/uri"
)

func NewGetCommand() *cobra.Command {
	return uri.NewGetCommand()
}
