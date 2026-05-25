package cyargs

import "github.com/spf13/cobra"

// UpdateFlag is the canonical name for the upsert flag used by create/update commands.
const UpdateFlag = "update"

// IsSet reports whether the user explicitly set a flag on the command.
func IsSet(cmd *cobra.Command, name string) bool {
	return cmd.Flags().Changed(name)
}

// AddUpdateFlag registers the boolean upsert flag and returns its name so
// callers can MarkFlag* without hardcoding the literal.
func AddUpdateFlag(cmd *cobra.Command, usage string) string {
	if usage == "" {
		usage = "update the resource if it already exists"
	}
	cmd.Flags().Bool(UpdateFlag, false, usage)
	return UpdateFlag
}

// GetUpdate returns the value of the upsert flag (false when not set).
func GetUpdate(cmd *cobra.Command) bool {
	v, _ := cmd.Flags().GetBool(UpdateFlag)
	return v
}
