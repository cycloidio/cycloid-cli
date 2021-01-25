package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/cycloidio/cycloid-cli/cmd"
	"github.com/cycloidio/cycloid-cli/internal/version"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// Used for flags.
	cfgFile    string
	userOutput string
	verbosity  string

	versionString = fmt.Sprintf("%s, revision %s, branch %s, date %s; go %s", version.Version, version.Revision, version.Branch, version.BuildDate, version.GoVersion)

	rootCmd = &cobra.Command{
		Version: versionString,
		Use:     "cy",
		Short:   "Cycloid CLI",
		Long:    `Cy is a CLI for Cycloid framework. Learn more at https://www.cycloid.io/.`,
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	// cobra.OnInitialize(initConfig)
	viper.SetEnvPrefix("CY")
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	viper.AutomaticEnv()

	rootCmd.PersistentFlags().StringVarP(&userOutput, "output", "o", "table", "The formatting style for command output: json|yaml|table")
	viper.BindPFlag("output", rootCmd.PersistentFlags().Lookup("output"))

	rootCmd.PersistentFlags().StringP("verbosity", "v", "warning", "Override the default verbosity for this command. VERBOSITY must be one of: debug, info, warning, error, critical, none.")
	viper.BindPFlag("verbosity", rootCmd.PersistentFlags().Lookup("verbosity"))
	viper.SetDefault("verbosity", "warning")

	rootCmd.PersistentFlags().String("api-url", "https://http-api.cycloid.io", "Specify the HTTP url of Cycloid API to use eg https://http-api.cycloid.io. This can also be given by CY_API_URL environment variable.")
	viper.BindPFlag("api-url", rootCmd.PersistentFlags().Lookup("api-url"))

	rootCmd.PersistentFlags().Bool("insecure", false, "Decide to skip or not TLS verification")
	viper.BindPFlag("insecure", rootCmd.PersistentFlags().Lookup("insecure"))
}

func main() {
	cmd.AttachCommands(rootCmd)

	Execute()
}
