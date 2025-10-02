package cycloid

import (
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/uri"
	"github.com/spf13/cobra"
)

func NewGetCommand() *cobra.Command {
	return uri.NewGetCommand()
}
