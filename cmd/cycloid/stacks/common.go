package stacks

import (
	"github.com/spf13/cobra"
)

var formsFlag string

func WithFlagForms(cmd *cobra.Command) string {
	flagName := "forms"
	cmd.Flags().StringVar(&formsFlag, flagName, "./.forms.yml", "Path to the Cycloid stackforms file")
	return flagName
}
