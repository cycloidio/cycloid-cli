package main

import (
	"fmt"
	"os"

	"github.com/cycloidio/cycloid-cli/cmd"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// Used for flags.
	rootCmd *cobra.Command
)

func inRed(msg string) string {
	return fmt.Sprintf("\033[1;31m%s\033[0m", msg)
}

// Execute runs the CLI root command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		rootCmd.PrintErrln(inRed("Error:"), err.Error())
		// rootCmd.PrintErrf("Run '%v --help' for usage.\n", rootCmd.CommandPath())
		os.Exit(1)
	}
}

func main() {
	rootCmd = cmd.NewRootCommand()
	viper.BindPFlag("api-url", rootCmd.PersistentFlags().Lookup("api-url"))

	Execute()
}
