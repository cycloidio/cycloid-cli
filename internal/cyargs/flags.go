package cyargs

import "github.com/spf13/cobra"

// IsSet reports whether the user explicitly set a flag on the command.
func IsSet(cmd *cobra.Command, name string) bool {
	return cmd.Flags().Changed(name)
}
