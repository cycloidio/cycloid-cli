package cyargs

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func AddFSRecurseFlag(cmd *cobra.Command) string {
	flagName := "recurse"
	cmd.Flags().BoolP(flagName, "r", false, "will walk directories recursively if set")
	return flagName
}

func GetFSRecurseFlag(cmd *cobra.Command) (bool, error) {
	return cmd.Flags().GetBool("recurse")
}

func AddFSIgnoreFlag(cmd *cobra.Command) string {
	flagName := "ignore"
	cmd.Flags().StringArray(
		flagName, []string{".git"},
		strings.Join([]string{
			"add glob pattern to ignore when using --recurse argument",
			"can be specified multiple time",
		}, "\n"),
	)
	return flagName
}

func GetFSIgnoreFlag(cmd *cobra.Command) ([]string, error) {
	return cmd.Flags().GetStringArray("ignore")
}

func ValidateFSArguments(cmd *cobra.Command, args []string) error {
	recurse, _ := GetFSRecurseFlag(cmd)
	for _, path := range args {
		file, err := os.Stat(path)
		if err != nil {
			return fmt.Errorf("no file was found at path %q: %w", path, err)
		}

		if file.IsDir() && !recurse {
			return fmt.Errorf("path %q is a directory, add --recurse argument for walking directories", path)
		}
	}

	return nil
}

// ValidFileArgs return an error if a file argument is invalid
// will validate the use of --recurse
func ValidFileArgs() cobra.PositionalArgs {
	return ValidateFSArguments
}

func AddInPlaceFlag(cmd *cobra.Command) string {
	flagName := "in-place"
	cmd.Flags().BoolP(flagName, "i", false, "inject values directly in the file in-place")
	return flagName
}

func GetInPlaceFlag(cmd *cobra.Command) (bool, error) {
	return cmd.Flags().GetBool("in-place")
}

func AddOutputDirectoryFlag(cmd *cobra.Command) string {
	flagName := "directory"
	cmd.Flags().StringP(flagName, "d", "", "add an output directory path")
	cmd.MarkFlagDirname(flagName)
	cmd.MarkFlagsMutuallyExclusive("in-place")
	return flagName
}

func GetOutputDirectoryFlag(cmd *cobra.Command) (string, error) {
	return cmd.Flags().GetString("directory")
}
