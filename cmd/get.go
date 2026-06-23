package cmd

import (
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/uri"
)

func NewGetCommand() *cobra.Command {
	return uri.NewGetCommand()
}
